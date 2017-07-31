#!/bin/bash -x
set -euo pipefail

echo -e "\nexport ENV=development\n" >> /home/develop/.zshrc
echo -e "\nexport PATH=\${PATH}:/home/develop/go/src/clickyab.com/crab/scripts:/home/develop/go/src/clickyab.com/crab/bin" >> /home/develop/.zshrc

#make all
