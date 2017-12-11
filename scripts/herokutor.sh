#!/bin/bash
#set -x
set -eo pipefail

# This job is from jenkins. so kill it if it is a pull request
exit_message() {
    echo ${1:-'exiting...'}
    code=${2:-1}
    if [ "${code}" == "0" ]
    then
        echo "Build was OK, but its not the correct branch. ignore this" >> ${OUT_LOG}
        echo "green" > ${OUT_LOG_COLOR}
    else
        echo "Build was NOT OK. verify with dev team" >> ${OUT_LOG}
        echo "red" > ${OUT_LOG_COLOR}
    fi;
    exit ${code}
}

OUT_LOG=${OUT_LOG:-/dev/null}
OUT_LOG_COLOR=${OUT_LOG_COLOR:-/dev/null}
echo "red" > ${OUT_LOG_COLOR}
echo ""> ${OUT_LOG}

APP=${APP:-}
PROJECT=${PROJECT:-ccc}
BRANCH=${BRANCH_NAME:-master}
BRANCH=${CHANGE_TARGET:-${BRANCH}}

PUSH="--push"
[ -z ${CHANGE_AUTHOR} ] || PUSH=""
[ -z ${APP} ] && exit_message "The APP is not defined." # WTF, the APP NAME is important

SCRIPT_DIR=$(readlink -f $(dirname ${BASH_SOURCE[0]}))

SOURCE_DIR=${1:-}
[ -z ${SOURCE_DIR} ] && exit_message "Must pass the source directory as the first parameter" 1
SOURCE_DIR=$(cd "${SOURCE_DIR}/" && pwd)

BUILD_DIR=${2:-$(mktemp -d)}
CACHE_DIR=${3:-${SOURCE_DIR}-cache}
ENV_DIR=$(mktemp -d)

mkdir -p "${BUILD_DIR}" "${CACHE_DIR}" "${ENV_DIR}"
BUILD=$(cd "${BUILD_DIR}/" && pwd)
CACHE=$(cd "${CACHE_DIR}/" && pwd)
VARS=$(cd "${ENV_DIR}/" && pwd)

BUILD_PACKS_DIR=$(mktemp -d)

# Extract build data
pushd ${SOURCE_DIR}
GIT_WORK_TREE=${BUILD} git checkout -f HEAD

export LONG_HASH=$(git log -n1 --pretty="format:%H" | cat)
export SHORT_HASH=$(git log -n1 --pretty="format:%h"| cat)
export COMMIT_DATE=$(git log -n1 --date="format:%D-%H-%I-%S" --pretty="format:%cd"| sed -e "s/\//-/g")
export IMP_DATE=$(date +%Y%m%d)
export COMMIT_COUNT=$(git rev-list HEAD --count| cat)
export BUILD_DATE=$(date "+%D/%H/%I/%S"| sed -e "s/\//-/g")
popd

# Populate env for herokuish
env -0 | while IFS='=' read -r -d '' n v; do
    echo "${v}">"${VARS}/${n}";
done< <(env -0)

TEMPORARY=$(mktemp -d)

# Create Rockerfile to build with rocker (the Dockerfile enhancer tool)
cat > ${TEMPORARY}/Rockerfile <<EOF
FROM gliderlabs/herokuish

MOUNT {{ .Build }}:/tmp/app
MOUNT {{ .EnvDir }}:/tmp/env
MOUNT {{ .Target }}:/tmp/build
MOUNT {{ .Cache }}:/tmp/cache

ENV TZ=Asia/Tehran
RUN ln -snf /usr/share/zoneinfo/\$TZ /etc/localtime && echo \$TZ > /etc/timezone

RUN /bin/herokuish buildpack build
EXPORT /app/bin /app


FROM ubuntu:16.04
IMPORT /app

ENV TZ=Asia/Tehran
RUN ln -snf /usr/share/zoneinfo/\$TZ /etc/localtime && echo \$TZ > /etc/timezone

RUN apt-get update && apt-get install -y tzdata ca-certificates && apt-get clean

CMD ["bash", "/app/bin/run_order.sh"]

TAG registry.clickyab.ae/clickyab/{{ .App }}:{{ .Version }}
PUSH registry.clickyab.ae/clickyab/{{ .App }}:{{ .Version }}
EOF

TARGET=$(mktemp -d)
pushd ${TEMPORARY}
# Actual build
rocker build ${PUSH} -var Build=${BUILD} -var EnvDir=${VARS} -var Cache=${CACHE} -var Target=${TARGET} -var Version=${BRANCH}.${COMMIT_COUNT} -var App=${APP}
popd

echo "${VARS}" >> /tmp/kill-me
echo "${TARGET}" >> /tmp/kill-me
echo "${TEMPORARY}" >> /tmp/kill-me
echo "${BUILD_DIR}" >> /tmp/kill-me
echo "${BUILD_PACKS_DIR}" >> /tmp/kill-me

[ -z ${CHANGE_AUTHOR} ] || exit_message "It's a PR, bail out" 0
if [[ ( "${BRANCH}" != "master" ) && ( "${BRANCH}" != "deploy" ) ]]; then
    exit_message "Its not on correct branch, bail out" 0
fi
echo "The branch ${BRANCH} build finished, try to deploy it" >> ${OUT_LOG}
echo "If there is no report after this for successful deploy, it means the deploy failed. report it please." >> ${OUT_LOG}
kubectl -n ${PROJECT} set image deployment  ${APP}-${BRANCH} ${APP}-${BRANCH}=registry.clickyab.ae/clickyab/${APP}:${BRANCH}.${COMMIT_COUNT} --record
echo "Deploy done successfully to image registry.clickyab.ae/clickyab/${APP}:${BRANCH}.${COMMIT_COUNT}" >> ${OUT_LOG}
echo "green" > ${OUT_LOG_COLOR}
