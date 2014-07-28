#!/bin/sh
go build || exit 1
if test $# -lt 1
then
  echo "specify a file" 
  exit 2
fi
input=$1 
./gift -blur=3 $input              > blur.jpg
./gift -brightness=20 $input       > bright30.jpg
./gift -brightness=-20 $input      > bright-30.jpg
./gift -contrast=30 $input         > contrast30.jpg
./gift -contrast=-30 $input        > contrast-30.jpg
./gift -crop=90,90,250,250 $input  > crop.jpg
./gift -edge $input                > edge.jpg
./gift -emboss $input              > emboss.jpg
./gift -fliph $input               > fliph.jpg
./gift -flipv $input               > flipv.jpg
./gift -gamma=1.5 $input           > gamma.jpg
./gift -gray $input                > gray.jpg
./gift -hue=45 $input              > hue45.jpg
./gift -hue=-45 $input             > hue-45.jpg
./gift -invert $input              > invert.jpg
./gift -max=5 $input               > max.jpg
./gift -mean=5 $input              > mean.jpg
./gift -median=5 $input            > median.jpg
./gift -min=5 $input               > min.jpg
./gift -resize=200,0 $input        > resize.jpg
./gift -rotate=60 $input           > rotate60.jpg
./gift -rotate=90 $input           > rotate90.jpg
./gift -rotate=180 $input          > rotate180.jpg
./gift -rotate=270 $input          > rotate270.jpg
./gift -saturation=50 $input       > sat50.jpg
./gift -saturation=-50 $input      > sat-50.jpg
./gift -sepia=100 $input           > sepia.jpg
./gift -sigmoid=0.5,5.0 $input     > sigmoid.jpg
./gift -transpose $input           > transpose.jpg
./gift -transverse $input          > transverse.jpg
./gift -unsharp=1.0,1.5,0.0 $input > unsharp.jpg
