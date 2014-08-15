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
    -brightness=<value>: brightness value (-100, 100)
    -colorize=hue,saturation,percent
    -colorbalance=%red,%green,%blue
    -contrast=<value>: contrast value (-100, 100)
    -crop=x1,y1,x2,y2
    -edge: edge
    -emboss: emboss
    -fliph: flip horizontal
    -flipv: flip vertical
    -gamma=<value>: gamma value
    -gray: grayscale
    -hue=<value>: hue value (-180, 180)
    -invert: invert
    -max=<value>: local maximum (kernel size)
    -mean=<value>: local mean filter (kernel size)
    -median=<value>: local median filter (kernel size)
    -min=<value>: local minimum (kernel size)
    -resize=w,h: resize w,h
    -rotate=<value>: rotate specified degrees counter-clockwise
    -saturation=<value>: saturation value (-100, 500)
    -sepia=<value>: sepia percentage (0-100)
    -sigmoid=midpint,factor: sigmoid contrast
    -transpose: flip horizontally and rotate 90° counter-clockwise
    -transverse:  flips vertically and rotate 90° counter-clockwise
    -unsharp=sigma,amount,threshold: unsharp mask


See testgift.sh for typical usage
