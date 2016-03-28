#!/bin/bash
# Read from the file file.txt and output the tenth line to stdout.

cat file.txt | awk 'BEGIN{x=1} {if(x==10) print $0; x++}'
