package main

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fmtutil"
	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/options"
	"github.com/essentialkaos/ek/v12/strutil"
	"github.com/essentialkaos/ek/v12/usage"
	"github.com/essentialkaos/ek/v12/usage/completion/bash"
	"github.com/essentialkaos/ek/v12/usage/completion/fish"
	"github.com/essentialkaos/ek/v12/usage/completion/zsh"
	"github.com/essentialkaos/ek/v12/usage/man"
	"github.com/essentialkaos/ek/v12/usage/update"

	"github.com/disintegration/imaging"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Basic utility info
const (
	APP  = "rsz"
	VER  = "0.0.3"
	DESC = "Simple utility for image resizing"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Options
const (
	OPT_FILTER       = "f:filter"
	OPT_LIST_FILTERS = "F:list-filters"
	OPT_NO_COLOR     = "nc:no-color"
	OPT_HELP         = "h:help"
	OPT_VER          = "v:version"

	OPT_COMPLETION   = "completion"
	OPT_GENERATE_MAN = "generate-man"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// optMap contains information about all supported options
var optMap = options.Map{
	OPT_FILTER:       {Value: "CatmullRom"},
	OPT_LIST_FILTERS: {Type: options.BOOL},
	OPT_NO_COLOR:     {Type: options.BOOL},
	OPT_HELP:         {Type: options.BOOL, Alias: "u:usage"},
	OPT_VER:          {Type: options.BOOL, Alias: "ver"},

	OPT_COMPLETION:   {},
	OPT_GENERATE_MAN: {Type: options.BOOL},
}

// supportedFilters is a slice with supported filters
var supportedFilters = []string{
	"BSpline",
	"Bartlett",
	"Blackman",
	"Box",
	"CatmullRom",
	"Cosine",
	"Gaussian",
	"Hamming",
	"Hann",
	"Hermite",
	"Lanczos",
	"Linear",
	"MitchellNetravali",
	"NearestNeighbor",
	"Welch",
}

// ////////////////////////////////////////////////////////////////////////////////// //

// main is main function
func main() {
	args, errs := options.Parse(optMap)

	if len(errs) != 0 {
		for _, err := range errs {
			printError(err.Error())
		}

		os.Exit(1)
	}

	preConfigureUI()

	if options.Has(OPT_COMPLETION) {
		os.Exit(genCompletion())
	}

	if options.Has(OPT_GENERATE_MAN) {
		os.Exit(genMan())
	}

	configureUI()

	if options.GetB(OPT_LIST_FILTERS) {
		listFilters()
		os.Exit(0)
	}

	if options.GetB(OPT_VER) {
		os.Exit(showAbout())
	}

	if options.GetB(OPT_HELP) || len(args) < 3 {
		os.Exit(showUsage())
	}

	process(args)
}

// preConfigureUI preconfigures UI based on information about user terminal
func preConfigureUI() {
	term := os.Getenv("TERM")

	fmtc.DisableColors = true

	if term != "" {
		switch {
		case strings.Contains(term, "xterm"),
			strings.Contains(term, "color"),
			term == "screen":
			fmtc.DisableColors = false
		}
	}

	if !fsutil.IsCharacterDevice("/dev/stdout") && os.Getenv("FAKETTY") == "" {
		fmtc.DisableColors = true
	}

	if os.Getenv("NO_COLOR") != "" {
		fmtc.DisableColors = true
	}
}

// configureUI configures user interface
func configureUI() {
	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}
}

// listFilters prints names of supported resampling filters
func listFilters() {
	current := options.GetS(OPT_FILTER)

	fmtc.NewLine()

	for _, filter := range supportedFilters {
		if filter == current {
			fmtc.Printf(" {s}•{!} %s {s-}(default){!}\n", filter)
		} else {
			fmtc.Printf(" {s}•{!} %s\n", filter)
		}
	}

	fmtc.NewLine()
}

// process starts image processing
func process(args []string) {
	srcImage, size, outImage := args[0], args[1], args[2]

	err := checkSrcImage(srcImage)

	if err != nil {
		printErrorAndExit(err.Error())
	}

	err = resizeImage(srcImage, outImage, size)

	if err != nil {
		printErrorAndExit(err.Error())
	}

	fmtc.Printf(
		"{g}Image successfully resized and saved as {g*}%s{!} {s-}(%s){!}\n",
		outImage, fmtutil.PrettySize(fsutil.GetSize(outImage)),
	)
}

// resizeImage resizes image
func resizeImage(srcImage, outImage, size string) error {
	filter, err := getResampleFilter()

	if err != nil {
		return err
	}

	img, err := imaging.Open(srcImage)

	if err != nil {
		return fmt.Errorf("Can't open image: %v", err.Error())
	}

	w, h, err := parseSize(size, img.Bounds())

	img = imaging.Resize(img, w, h, filter)
	err = imaging.Save(img, outImage)

	if err != nil {
		return fmt.Errorf("Can't save image: %v", err.Error())
	}

	return nil
}

// getResampleFilter returns resampling filter config
func getResampleFilter() (imaging.ResampleFilter, error) {
	filter := options.GetS(OPT_FILTER)

	switch strings.ToLower(filter) {
	case strings.ToLower("BSpline"):
		return imaging.BSpline, nil
	case strings.ToLower("Bartlett"):
		return imaging.Bartlett, nil
	case strings.ToLower("Blackman"):
		return imaging.Blackman, nil
	case strings.ToLower("Box"):
		return imaging.Box, nil
	case strings.ToLower("CatmullRom"):
		return imaging.CatmullRom, nil
	case strings.ToLower("Cosine"):
		return imaging.Cosine, nil
	case strings.ToLower("Gaussian"):
		return imaging.Gaussian, nil
	case strings.ToLower("Hamming"):
		return imaging.Hamming, nil
	case strings.ToLower("Hann"):
		return imaging.Hann, nil
	case strings.ToLower("Hermite"):
		return imaging.Hermite, nil
	case strings.ToLower("Lanczos"):
		return imaging.Lanczos, nil
	case strings.ToLower("Linear"):
		return imaging.Linear, nil
	case strings.ToLower("MitchellNetravali"):
		return imaging.MitchellNetravali, nil
	case strings.ToLower("NearestNeighbor"):
		return imaging.NearestNeighbor, nil
	case strings.ToLower("Welch"):
		return imaging.Welch, nil
	}

	return imaging.ResampleFilter{}, fmt.Errorf("Unknown resampling filter \"%s\"", filter)
}

// checkSrcImage checks source image before processing
func checkSrcImage(srcImage string) error {
	if !fsutil.IsExist(srcImage) {
		return fmt.Errorf("Image %s doesn't exist", srcImage)
	}

	if !fsutil.IsReadable(srcImage) {
		return fmt.Errorf("Image file %s is not readable", srcImage)
	}

	if !fsutil.IsNonEmpty(srcImage) {
		return fmt.Errorf("Image file %s is empty", srcImage)
	}

	return nil
}

// parseSize parses new image size
func parseSize(size string, bounds image.Rectangle) (int, int, error) {
	switch {
	case strings.Contains(size, "x"):
		return parseExactSize(size)
	case strings.Contains(size, "%"), strings.Contains(size, "."):
		return parseRelativeSize(size, bounds)
	}

	return 0, 0, fmt.Errorf("Unsupported size definition \"%s\"", size)
}

// parseExactSize parses exact image size
func parseExactSize(size string) (int, int, error) {
	ws := strutil.ReadField(size, 0, false, "x")
	hs := strutil.ReadField(size, 1, false, "x")

	w, err := strconv.Atoi(ws)

	if err != nil {
		return 0, 0, fmt.Errorf("Can't parse width value: %v", err)
	}

	h, err := strconv.Atoi(hs)

	if err != nil {
		return 0, 0, fmt.Errorf("Can't parse height value: %v", err)
	}

	return w, h, nil
}

// parseRelativeSize parses relative image size
func parseRelativeSize(size string, bounds image.Rectangle) (int, int, error) {
	var err error
	var mod float64

	if strings.Contains(size, "%") {
		mod, err = strconv.ParseFloat(strings.Trim(size, "%"), 64)

		if err != nil {
			return 0, 0, fmt.Errorf("Can't parse size: %v", err)
		}

		mod /= 100.0
	} else {
		mod, err = strconv.ParseFloat(size, 64)

		if err != nil {
			return 0, 0, fmt.Errorf("Can't parse size: %v", err)
		}
	}

	return int(float64(bounds.Max.X) * mod),
		int(float64(bounds.Max.Y) * mod), nil
}

// printError prints error message to console
func printError(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{r}"+f+"{!}\n", a...)
}

// printError prints warning message to console
func printWarn(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{y}"+f+"{!}\n", a...)
}

// printErrorAndExit print error mesage and exit with exit code 1
func printErrorAndExit(f string, a ...interface{}) {
	printError(f, a...)
	os.Exit(1)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// showUsage prints usage info
func showUsage() int {
	genUsage().Render()
	return 0
}

// showAbout prints info about version
func showAbout() int {
	genAbout().Render()
	return 0
}

// genCompletion generates completion for different shells
func genCompletion() int {
	info := genUsage()

	switch options.GetS(OPT_COMPLETION) {
	case "bash":
		fmt.Printf(bash.Generate(info, "rsz"))
	case "fish":
		fmt.Printf(fish.Generate(info, "rsz"))
	case "zsh":
		fmt.Printf(zsh.Generate(info, optMap, "rsz"))
	default:
		return 1
	}

	return 0
}

// genMan generates man page
func genMan() int {
	fmt.Println(
		man.Generate(
			genUsage(),
			genAbout(),
		),
	)

	return 0
}

// genUsage generates usage info
func genUsage() *usage.Info {
	info := usage.NewInfo("", "src-image", "size", "output-image")

	info.AddOption(OPT_FILTER, "Resampling filter name", "name")
	info.AddOption(OPT_LIST_FILTERS, "Print list of supported resampling filters")
	info.AddOption(OPT_NO_COLOR, "Disable colors in output")
	info.AddOption(OPT_HELP, "Show this help message")
	info.AddOption(OPT_VER, "Show version")

	info.AddExample(
		"image.png 256x256 thumbnail.png",
		"Convert image to exact size",
	)

	info.AddExample(
		"-f Lanczos image.png 256x256 thumbnail.png",
		"Convert image to exact size using Lanczos resampling filter",
	)

	info.AddExample(
		"image.png 25% thumbnail.png",
		"Convert image to relative size (25% of original)",
	)

	info.AddExample(
		"image.png 0.55 thumbnail.png",
		"Convert image to relative size (55% of original)",
	)

	return info
}

// genAbout generates info about version
func genAbout() *usage.About {
	return &usage.About{
		App:           APP,
		Version:       VER,
		Desc:          DESC,
		Year:          2009,
		Owner:         "ESSENTIAL KAOS",
		License:       "Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>",
		UpdateChecker: usage.UpdateChecker{"essentialkaos/rsz", update.GitHubChecker},
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //
