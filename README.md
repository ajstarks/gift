gift
====

Command line interface to the [Go Image Filtering Toolkit](https://github.com/disintegration/gift)

Install
===

    go get github.com/ajstarks/gift
    
Usage
===

gift reads either from a single file (PNG or JPEG) or standard input and always writes to standard output:

    gift -contrast 80 foo.jpg > contrast.jpg
    gift -flipv < foo.png > flipped.png
    
Filters can be combined.  For example to blur and invert:

    gift -invert -blur 4 foo.jpg > inverted-blurred.jpg
		
Here are the command flags:

    -blur=0: blur value
    -brightness=-200: brightness value (-100, 100)
    -colorbalance="": color balance (%red, %green, %blue)
    -colorize="": colorize (hue, saturation, percentage)
    -contrast=-200: contrast value (-100, 100)
    -crop="": crop x1,y1,x2,y2
    -cropsize="": crop w h
    -edge=false: edge filter
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
    -resizefill="": resizefill w,h
    -resizefit="": resizefit w,h
    -rotate=0: rotate specified degrees counter-clockwise
    -saturation=-200: saturation value (-100, 500)
    -sepia=-1: sepia percentage (0-100)
    -sigmoid="": sigmoid contrast (midpoint,factor)
    -sobel=false: sobel filter
    -transpose=false: flip horizontally and rotate 90° counter-clockwise
    -transverse=false:  flips vertically and rotate 90° counter-clockwise
    -unsharp="": unsharp mask (sigma,amount,threshold)


See testgift.sh for typical usage
