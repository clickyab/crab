#!/usr/bin/env bash
set -eo pipefail
## This is the only script here. its going to be on the final image.

ROOT="$(readlink -f $(dirname ${BASH_SOURCE[0]})/)"
${ROOT}/migration --action=up
${ROOT}/webserver
