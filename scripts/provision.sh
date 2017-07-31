#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR=$(readlink -f $(dirname ${BASH_SOURCE[0]}))

echo "export GOPATH=/home/develop/go" >> /home/develop/.zshrc
echo "export GOPATH=/home/develop/go" >> /etc/environment
echo "export PATH=$PATH:/usr/local/go/bin:/home/develop/go/bin" >> /home/develop/.zshrc
echo "alias cdp=\"cd /home/develop/go/src/clickyab.com/crab\"" >> /home/develop/.zshrc

chown develop. /home/develop/go
chown develop. /home/develop/go/src
chown develop. /home/develop/go/src/clickyab.com

cd /home/develop/go/src/clickyab.com/crab
make -f /home/develop/go/src/clickyab.com/crab/Makefile database-setup
make -f /home/develop/go/src/clickyab.com/crab/Makefile broker-setup

sudo -u develop /home/develop/go/src/clickyab.com/crab/scripts/provision_user.sh