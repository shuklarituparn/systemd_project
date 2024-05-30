#!/bin/bash

current_datetime=$(date +"%Y-%m-%d %H:%M:%S")
status_code=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8888/status)
if [ $status_code -eq 200 ]; then
    echo "$current_datetime - SUCCESS"
else
    echo "$current_datetime - ERROR"
fi