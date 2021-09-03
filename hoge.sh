#!/bin/bash

filename=".env"               # .env の path
appname="floating-crag-63270" # appname

while read line; do
    echo "line: $line"
    sudo heroku config:add $line --app $appname
done < $filename