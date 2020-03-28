#!/usr/bin/env bash
echo 'Runing migrations...'
/goapp/migrate up > /dev/null 2>&1 &
echo 'Deleting mysql-client...'
apk del mysql-client
echo 'Start application...'
/goapp/app