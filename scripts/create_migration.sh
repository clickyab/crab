#!/usr/bin/env bash
set -euo pipefail


if [ "$#" -ne 1 ]; then
    echo "Name for this module : "
    read NAME
    echo "Name for this migration : "
    read MIG
else
    NAME=${1}
    MIG=${2}
fi

SCRIPT_DIR=$(readlink -f $(dirname ${BASH_SOURCE[0]}))
cd ${SCRIPT_DIR}/..
mkdir -p modules/${NAME}/migrations/db
cd modules/${NAME}/migrations/db
DATE=`date +%Y%m%d%H%M%S`
FILE="${DATE}_${MIG}.sql"

cat >>$FILE <<-EOGO

-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


EOGO

echo ${FILE}
