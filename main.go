package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chonla/rnd"
)

// VERSION of rand
const VERSION = "0.1.0"

func main() {
	var dataSet string
	var entries string
	var ver bool
	var help bool
	var length int

	flag.Usage = usage

	flag.StringVar(&dataSet, "d", "alphanum", "set data set")
	flag.StringVar(&entries, "e", "", "set data entries, used only when data set is custom")
	flag.IntVar(&length, "l", 10, "result length")
	flag.BoolVar(&ver, "v", false, "show rand version")
	flag.BoolVar(&help, "h", false, "show this help")
	flag.Parse()

	if ver {
		fmt.Printf("rand %s\n", VERSION)
		os.Exit(0)
	}

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if length <= 0 {
		length = 10
	}

	switch dataSet {
	case "alphanum":
		entries = rnd.ALPHANUM
	case "alphanumcase":
		entries = rnd.ALPHANUMCASE
	case "num":
		entries = rnd.NUM
	case "alpha":
		entries = rnd.ALPHA
	case "alphacase":
		entries = rnd.ALPHACASE
	case "hex":
		entries = rnd.HEX
	case "hexcase":
		entries = rnd.HEXCASE
	case "custom":
		if entries == "" {
			entries = rnd.ALPHANUM
		}
	default:
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(rnd.Of(entries, length))
}

func usage() {
	fmt.Println("Usage of rand:")
	fmt.Println()
	fmt.Println("  rand [-d <dataset>] [-e <dataset-entries>]")
	fmt.Println()
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("  possible dataset:")
	fmt.Println("    alpha : alphabetic (a-z, lowercase)")
	fmt.Println("    alphacase : alphabetic with case-sensitive (a-z + A-Z)")
	fmt.Println("    num : numeric (0-9)")
	fmt.Println("    alphanum : alphanumeric (a-z + 0-9, lowercase)")
	fmt.Println("    alphanumcase : alphanumeric with case-sensitive (a-z + A-Z + 0-9)")
	fmt.Println("    hex : hexeric (0-9 + a-f, lowercase)")
	fmt.Println("    hexcase : hexeric with case-sensitive (0-9 + a-f + A-F)")
	fmt.Println("    custom : custom data set, use -e to specify entries in data set")
}
