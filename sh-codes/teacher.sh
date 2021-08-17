#!/usr/bin/env bash

PWD=$(pwd)

for folder in $PWD
do
        cd "$folder" || exit
        INTERVIEWNUMBER=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
done  
