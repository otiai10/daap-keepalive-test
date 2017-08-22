#!/bin/bash

# t=60 # negative
# t=70 # negative
# t=80 # negative
# t=90 # negative
# t=100 # POSITIVE
# t=96 # netative
# t=98 # negative
# t=100 # negative ????
# t=120 # negative !!!!
# t=600 # POSITIVE
# t=140 # POSITIVE
t=122

echo "Let's sleep for $t seconds"

NOW=`date +"%D %H:%M"`
echo "STARTED: ${NOW}"
sleep $t
NOW=`date +"%D %H:%M"`
echo "FINISHED: ${NOW}"
