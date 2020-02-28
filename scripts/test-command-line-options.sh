#!/usr/bin/env bash

COMMAND='./bin/dnsmonitor -version'
echo $COMMAND
eval $COMMAND
if [ $? -ne 0 ]; then
    echo "-version should exit with 0"
    exit 1
fi
echo "OK"

FILENAME=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
COMMAND='./bin/dnsmonitor -configfile=/tmp/'$FILENAME
echo $COMMAND
eval $COMMAND
if [ $? -ne 1 ]; then
  echo "-configfile should exit with 1 if the file doesn't exist"
  exit 1
fi
echo "OK"

COMMAND='./bin/dnsmonitor -configfile=./test/config.yml -domain www.google.com'
echo $COMMAND
eval $COMMAND
if [ $? -ne 2 ]; then
  echo "exit code 2 expected if -configfile and -domain is provided"
  exit 1
fi
echo "OK"
