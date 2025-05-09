#!/bin/bash

source ~/.profile

cd "$HOME"/compile/mailer-go || {
    echo "Status: $?"
    exit 4
}

echo "------------------------------------"
env | grep PATH
env | grep LOADED
echo "------------------------------------"



echo '
### LINUX #########################################################
'

stages="
stp,testta3
lgkk,testta3
stp,prodta3
lgkk,prodta3
"

for UMG in ${stages}
do
    cd /tmp/ || exit 1
    "$HOME"/bin/vicecersa.sh ${UMG} mailer \$HOME/bin/ || {
        echo "Status: $?"
        exit 2
    }
done


echo '

### AIX #########################################################

'

stages="
stp,testta2
stp,prodta2
"

for UMG in ${stages}
do
    cd /tmp/ || exit 1
    "$HOME"/bin/vicecersa.sh ${UMG} mailer.aix \$HOME/bin/ mailer || {
        echo "Status: $?"
        exit 2
    }
done
