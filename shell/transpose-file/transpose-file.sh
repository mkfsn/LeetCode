#!/bin/bash
# Read from the file file.txt and print its transposed content to stdout.

cat file.txt | awk 'BEGIN{x=1;y=1;} {for(i=1;i<=NF;i++){arr[x][i]=$i}x++;if(NF>y)y=NF} END{for(i=1;i<=y;i++){for(j=1;j<x;j++){printf "%s",arr[j][i];if(j!=x-1){printf "%s"," "}}printf "\n"}}'
