// gift: command line interface to Go image filtering toolkit
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/disintegration/gift"
)

var (
	blurvalue, brvalue, contvalue, hvalue, satvalue, gammavalue, sepiavalue, threshvalue, opacityvalue float64
	gray, neg, xpose, xverse, fliph, flipv, emboss, edge, sobel                                        bool
	res, resfit, resfill, cropspec, cropsize, sigspec, unsharp, colorize, colorbal, colorspace         string
	rotvalue, minvalue, maxvalue, meanvalue, medvalue, pixelatevalue                                   int
)

func digits(s string, index int) (int, int) {
	var x, y int
	var err error

	x, err = strconv.Atoi(s[0:index])
	if err != nil {
		return 0, 0
	}
	y, err = strconv.Atoi(s[index+1:])
	if err != nil {
		return 0, 0
	}
	return x, y
}

func dimen(s string) (int, int) {
	ci := strings.Index(s, ",")
	cx := strings.Index(s, "x")
	if ci > 0 && len(s) > ci+1 {
		return digits(s, ci)
	}
	if cx > 0 && len(s) > cx+1 {
		return digits(s, cx)
	}
	return 0, 0
}

func main() {
	flag.Float64Var(&blurvalue, "blur", 0, "blur value")
	flag.Float64Var(&brvalue, "brightness", -200, "brightness value (-100, 100)")
	flag.Float64Var(&hvalue, "hue", -200, "hue value (-180, 180)")
	flag.Float64Var(&contvalue, "contrast", -200, "contrast value (-100, 100)")
	flag.Float64Var(&satvalue, "saturation", -200, "saturation value (-100, 500)")
	flag.Float64Var(&gammavalue, "gamma", 0, "gamma value")
	flag.Float64Var(&sepiavalue, "sepia", -1, "sepia percentage (0-100)")
	flag.Float64Var(&opacityvalue, "opacity", -1, "opacity percentage (0-100)")
	flag.Float64Var(&threshvalue, "threshold", -1, "color threshold percentage (0-100)")
	flag.IntVar(&rotvalue, "rotate", 0, "rotate specified degrees counter-clockwise")
	flag.IntVar(&maxvalue, "max", 0, "local maximum (kernel size)")
	flag.IntVar(&minvalue, "min", 0, "local minimum (kernel size)")
	flag.IntVar(&medvalue, "median", 0, "local median filter (kernel size)")
	flag.IntVar(&meanvalue, "mean", 0, "local mean filter (kernel size)")
	flag.IntVar(&pixelatevalue, "pixelate", 0, "pixelate")
	flag.BoolVar(&flipv, "flipv", false, "flip vertical")
	flag.BoolVar(&fliph, "fliph", false, "flip horizontal")
	flag.BoolVar(&gray, "gray", false, "grayscale")
	flag.BoolVar(&neg, "invert", false, "invert")
	flag.BoolVar(&xpose, "transpose", false, "flip horizontally and rotate 90° counter-clockwise")
	flag.BoolVar(&xverse, "transverse", false, " flips vertically and rotate 90° counter-clockwise")
	flag.BoolVar(&emboss, "emboss", false, "emboss")
	flag.BoolVar(&edge, "edge", false, "edge filter")
	flag.BoolVar(&sobel, "sobel", false, "sobel filter")
	flag.StringVar(&colorspace, "colorspace", "", "colorspace; s: linear->sRGB, l: sRGB->linear")
	flag.StringVar(&res, "resize", "", "resize WxH")
	flag.StringVar(&resfit, "resizefill", "", "resizefill WxH")
	flag.StringVar(&resfill, "resizefit", "", "resizefit WxH")
	flag.StringVar(&cropspec, "crop", "", "crop x1,y1,x2,y2")
	flag.StringVar(&cropsize, "cropsize", "", "crop WxH")
	flag.StringVar(&sigspec, "sigmoid", "", "sigmoid contrast (midpoint,factor)")
	flag.StringVar(&unsharp, "unsharp", "", "unsharp mask (sigma,amount,threshold)")
	flag.StringVar(&colorize, "colorize", "", "colorize (hue, saturation, percentage)")
	flag.StringVar(&colorbal, "colorbalance", "", "color balance (%red, %green, %blue)")

	flag.Parse()

	var f io.Reader = os.Stdin
	var ferr error
	var fname string
	if len(flag.Args()) > 0 {
		fname = flag.Args()[0]
		f, ferr = os.Open(fname)
		if ferr != nil {
			fmt.Fprintf(os.Stderr, "%v\n", ferr)
			os.Exit(1)
		}
	}

	src, format, ierr := image.Decode(f)
	if ierr != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fname, ierr)
		os.Exit(2)
	}

	g := gift.New() // initial state

	// stack filters when command flags are set (not default values)
	if blurvalue > 0 {
		g.Add(gift.GaussianBlur(float32(blurvalue)))
	}

	// brightness
	if brvalue >= -100 && brvalue <= 100 {
		g.Add(gift.Brightness(float32(brvalue)))
	}

	// opacity
	if opacityvalue >= 0 && opacityvalue <= 100 {
		g.Add(gift.ColorFunc(func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
			return r0, g0, b0, float32(opacityvalue) / 100
		}))
	}

	// Linear to sRGB colorspace
	if colorspace == "s" {
		g.Add(gift.ColorspaceLinearToSRGB())
	}

	// sRGB to linear colorspace
	if colorspace == "l" {
		g.Add(gift.ColorspaceSRGBToLinear())
	}

	// hue
	if hvalue >= -180 && hvalue <= 180 {
		g.Add(gift.Hue(float32(hvalue)))
	}

	// contrast
	if contvalue >= -100 && contvalue <= 100 {
		g.Add(gift.Contrast(float32(contvalue)))
	}

	// saturation
	if satvalue >= -100 && satvalue <= 500 {
		g.Add(gift.Saturation(float32(satvalue)))
	}

	// gamma
	if gammavalue > 0 {
		g.Add(gift.Gamma(float32(gammavalue)))
	}

	// sepia
	if sepiavalue >= 0 && sepiavalue <= 100 {
		g.Add(gift.Sepia(float32(sepiavalue)))
	}

	// median
	if medvalue > 0 && medvalue%1 == 0 {
		g.Add(gift.Median(medvalue, true))
	}

	// mean
	if meanvalue > 0 && meanvalue%1 == 0 {
		g.Add(gift.Mean(meanvalue, true))
	}

	// minimum
	if minvalue > 0 && minvalue%1 == 0 {
		g.Add(gift.Minimum(minvalue, true))
	}

	// maximum
	if maxvalue > 0 && maxvalue%1 == 0 {
		g.Add(gift.Maximum(maxvalue, true))
	}

	// pixelate
	if pixelatevalue > 0 {
		g.Add(gift.Pixelate(pixelatevalue))
	}

	// threshold
	if threshvalue >= 0 && threshvalue <= 100 {
		g.Add(gift.Threshold(float32(threshvalue)))
	}

	// rotate
	if rotvalue > 0 && rotvalue <= 360 {
		switch rotvalue {
		case 90:
			g.Add(gift.Rotate90())
		case 180:
			g.Add(gift.Rotate180())
		case 270:
			g.Add(gift.Rotate270())
		default:
			g.Add(gift.Rotate(float32(rotvalue), color.White, gift.LinearInterpolation))
		}
	}

	// grayscale
	if gray {
		g.Add(gift.Grayscale())
	}

	// invert
	if neg {
		g.Add(gift.Invert())
	}

	// transpose
	if xpose {
		g.Add(gift.Transpose())
	}

	// transverse
	if xverse {
		g.Add(gift.Transverse())
	}

	// flip horizontal
	if fliph {
		g.Add(gift.FlipHorizontal())
	}

	// flip vertical
	if flipv {
		g.Add(gift.FlipVertical())
	}

	// emboss
	if emboss {
		g.Add(gift.Convolution(
			[]float32{-1, -1, 0, -1, 1, 1, 0, 1, 1},
			false, false, false, 0.0))
	}

	// edge detections
	if edge {
		g.Add(gift.Convolution(
			[]float32{-1, -1, -1, -1, 8, -1, -1, -1, -1},
			false, false, false, 0.0))
	}

	// Sobel filter
	if sobel {
		g.Add(gift.Sobel())
	}

	// resize
	if len(res) > 0 {
		w, h := dimen(res)
		if w == 0 && h == 0 {
			fmt.Fprintln(os.Stderr, "use: -resize WxH")
			os.Exit(3)
		}
		g.Add(gift.Resize(w, h, gift.LanczosResampling))
	}

	// resize fit
	if len(resfit) > 0 {
		w, h := dimen(resfit)
		if w == 0 && h == 0 {
			fmt.Fprintln(os.Stderr, "use: -resizefit WxH")
			os.Exit(3)
		}
		g.Add(gift.ResizeToFit(w, h, gift.LanczosResampling))
	}

	// resize fill
	if len(resfill) > 0 {
		w, h := dimen(resfill)
		if w == 0 && h == 0 {
			fmt.Fprintln(os.Stderr, "use: -resizefill WxH")
			os.Exit(3)
		}
		g.Add(gift.ResizeToFill(w, h, gift.LanczosResampling, gift.CenterAnchor))
	}

	// crop
	if len(cropspec) > 0 {
		var x1, y1, x2, y2 int
		nr, err := fmt.Sscanf(cropspec, "%d,%d,%d,%d", &x1, &y1, &x2, &y2)
		if nr != 4 || err != nil {
			fmt.Fprintln(os.Stderr, "use: -crop x1,y1,x2,y2")
			os.Exit(4)
		}
		g.Add(gift.Crop(image.Rect(x1, y1, x2, y2)))
	}

	// cropsize
	if len(cropsize) > 0 {
		w, h := dimen(cropsize)
		if w == 0 && h == 0 {
			fmt.Fprintln(os.Stderr, "use: -cropsize WxH")
			os.Exit(4)
		}
		g.Add(gift.CropToSize(w, h, gift.CenterAnchor))
	}

	// unsharp
	if len(unsharp) > 0 {
		var sigma, amount, threshold float32
		nr, err := fmt.Sscanf(unsharp, "%g,%g,%g", &sigma, &amount, &threshold)
		if nr != 3 || err != nil {
			fmt.Fprintln(os.Stderr, "use: -unsharp sigma,amount,threshold")
			os.Exit(5)
		}
		g.Add(gift.UnsharpMask(sigma, amount, threshold))
	}

	// sigmoid
	if len(sigspec) > 0 {
		var midpoint, factor float32
		nr, err := fmt.Sscanf(sigspec, "%g,%g", &midpoint, &factor)
		if nr != 2 || err != nil {
			fmt.Fprintln(os.Stderr, "use: -sigma midpoint,factor")
			os.Exit(6)
		}
		g.Add(gift.Sigmoid(midpoint, factor))
	}

	// colorize
	if len(colorize) > 0 {
		var chue, csaturation, cpercent float32
		nr, err := fmt.Sscanf(colorize, "%g,%g,%g", &chue, &csaturation, &cpercent)
		if nr != 3 || err != nil {
			fmt.Fprintln(os.Stderr, "use: -colorize hue,saturation,percent")
			os.Exit(7)
		}
		g.Add(gift.Colorize(chue, csaturation, cpercent))
	}

	// color balance
	if len(colorbal) > 0 {
		var pctred, pctgreen, pctblue float32
		nr, err := fmt.Sscanf(colorbal, "%g,%g,%g", &pctred, &pctgreen, &pctblue)
		if nr != 3 || err != nil {
			fmt.Fprintln(os.Stderr, "use: -colorbalance %red,%green,%blue")
			os.Exit(8)
		}
		g.Add(gift.ColorBalance(pctred, pctgreen, pctblue))
	}

	// make the filtered image, writing to stdout
	dst := image.NewRGBA(g.Bounds(src.Bounds()))
	g.Draw(dst, src)
	switch format {
	case "png":
		png.Encode(os.Stdout, dst)
	case "jpeg":
		jpeg.Encode(os.Stdout, dst, nil)
	}
}
