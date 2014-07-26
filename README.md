gift
====

Command line interface to the Go Image Filtering toolkit

`Usage of gift:
  -blur=0: blur value
  -brightness=-200: brightness value (-100, 100)
  -contrast=-200: contrast value (-100, 100)
  -crop="": crop x1,y1,x2,y2
  -edge=false: edge
  -emboss=false: emboss
  -fliph=false: flip horizontal
  -flipv=false: flip vertical
  -gamma=0: gamma value
  -gray=false: grayscale
  -hue=-200: hue value (-180, 180)
  -invert=false: invert
  -max=0: local maximum (kernel size)
  -mean=0: local mean filter (kernel size)
  -median=0: local median filter (kernel size)
  -min=0: local minimum (kernel size)
  -resize="": resize w,h
  -rotate=0: rotate 90, 180, 270 degrees counter-clockwise
  -saturation=-200: saturation value (-100, 500)
  -sepia=-1: sepia percentage (0-100)
  -sigmoid="": sigmoid contrast (midpoint,factor)
  -transpose=false: flip horizontally and rotate 90° counter-clockwise
  -transverse=false:  flips vertically and rotate 90° counter-clockwise
  -unsharp="": unsharp mask (sigma,amount,threshold)
`

See testgift.sh for typical usage
