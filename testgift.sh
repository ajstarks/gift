#!/bin/sh
go build || exit 1
if test $# -lt 1
then
  echo "specify a file" 1>&2
  exit 2
fi
input=$1
enc=`echo $1 | cut -d . -f2`
./gift -blur=3 $input              > blur.$enc
./gift -brightness=20 $input       > bright30.$enc
./gift -brightness=-20 $input      > bright-30.$enc
./gift -contrast=30 $input         > contrast30.$enc
./gift -contrast=-30 $input        > contrast-30.$enc
./gift -crop=90,90,250,250 $input  > crop.$enc
./gift -edge $input                > edge.$enc
./gift -emboss $input              > emboss.$enc
./gift -fliph $input               > fliph.$enc
./gift -flipv $input               > flipv.$enc
./gift -gamma=1.5 $input           > gamma.$enc
./gift -gray $input                > gray.$enc
./gift -hue=45 $input              > hue45.$enc
./gift -hue=-45 $input             > hue-45.$enc
./gift -invert $input              > invert.$enc
./gift -max=5 $input               > max.$enc
./gift -mean=5 $input              > mean.$enc
./gift -median=5 $input            > median.$enc
./gift -min=5 $input               > min.$enc
./gift -resize=200,0 $input        > resize.$enc
./gift -rotate=60 $input           > rotate60.$enc
./gift -rotate=90 $input           > rotate90.$enc
./gift -rotate=180 $input          > rotate180.$enc
./gift -rotate=270 $input          > rotate270.$enc
./gift -saturation=50 $input       > sat50.$enc
./gift -saturation=-50 $input      > sat-50.$enc
./gift -sepia=100 $input           > sepia.$enc
./gift -sigmoid=0.5,5.0 $input     > sigmoid.$enc
./gift -transpose $input           > transpose.$enc
./gift -transverse $input          > transverse.$enc
./gift -unsharp=1.0,1.5,0.0 $input > unsharp.$enc
