#!/bin/sh
if test $# -ne 1
then
	echo "specify an image file" 2>&1
	exit 1
fi
t=/tmp/perfchart.$$
sh -x timegift.sh $1 2> $t
awk '$0  ~ /gift/ {printf "%s\t", $3} $1 == "real" {print substr($2, 3, 5)}' $t | 
(echo "# `ims $1`"; sort -k2 -nr) | 
dchart -left 20 -top 90 -hbar -datafmt %0.3f -textsize 1.4 -ls 1.6 -color lightgray | pdfdeck -stdout - > giftime.pdf
open giftime.pdf
rm $t
