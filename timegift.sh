#!/bin/sh
go build || exit 1
if test $# -lt 1
then
  echo "specify a file" 1>&2
  exit 2
fi
input="$1"
if test ! -r "$input"
then 
	echo "cannot read $input" 1>&2
	exit 3
fi
enc=`echo "$input" | awk -F\. '{print $NF}'`
time ./gift -blur=10 "$input"               > blur.$enc
time ./gift -brightness=20 "$input"         > bright30.$enc
time ./gift -brightness=-20 "$input"        > bright-30.$enc
time ./gift -contrast=30 "$input"           > contrast30.$enc
time ./gift -contrast=-30 "$input"          > contrast-30.$enc
time ./gift -colorize=240,50,100 "$input"   > colorize.$enc
time ./gift -colorbalance=20,-20,0 "$input" > colorbalance.$enc
time ./gift -crop=90,90,250,250 "$input"    > crop.$enc
time ./gift -cropsize=100,100 "$input"      > cropsize.$enc
time ./gift -edge "$input"                  > edge.$enc
time ./gift -emboss "$input"                > emboss.$enc
time ./gift -fliph "$input"                 > fliph.$enc
time ./gift -flipv "$input"                 > flipv.$enc
time ./gift -gamma=1.5 "$input"             > gamma.$enc
time ./gift -gray "$input"                  > gray.$enc
time ./gift -hue=45 "$input"                > hue45.$enc
time ./gift -hue=-45 "$input"               > hue-45.$enc
time ./gift -invert "$input"                > invert.$enc
time ./gift -max=5 "$input"                 > max.$enc
time ./gift -mean=5 "$input"                > mean.$enc
time ./gift -median=5 "$input"              > median.$enc
time ./gift -min=5 "$input"                 > min.$enc
time ./gift -pixelate=50 "$input"           > pixelate.$enc
time ./gift -resize=200,0 "$input"          > resize.$enc
time ./gift -resizefit=100,100 "$input"     > resizefit.$enc
time ./gift -resizefill=100,100 "$input"    > resizefill.$enc
time ./gift -rotate=60 "$input"             > rotate60.$enc
time ./gift -rotate=90 "$input"             > rotate90.$enc
time ./gift -rotate=180 "$input"            > rotate180.$enc
time ./gift -rotate=270 "$input"            > rotate270.$enc
time ./gift -saturation=50 "$input"         > sat50.$enc
time ./gift -saturation=-50 "$input"        > sat-50.$enc
time ./gift -sepia=100 "$input"             > sepia.$enc
time ./gift -sigmoid=0.5,5.0 "$input"       > sigmoid.$enc
time ./gift -sobel "$input"                 > sobel.$enc
time ./gift -transpose "$input"             > transpose.$enc
time ./gift -transverse "$input"            > transverse.$enc
time ./gift -unsharp=1.0,1.5,0.0 "$input"   > unsharp.$enc
time ./gift -threshold=50 "$input"          > threshold.$enc
