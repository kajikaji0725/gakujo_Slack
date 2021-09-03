#!/bin/bash

filename=".env"               # .env の path
appname="frozen-sierra-65437" # appname

while read line; do
    echo "line: $line"
    heroku config:add $line --app $appname
done < $filename