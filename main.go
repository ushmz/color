package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Hex string

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (h *Hex) toRGB() (RGB, error) {
	return Hex2RGB(*h)
}

func (r *RGB) toHex() (Hex, error) {
	hexR := strconv.FormatUint(uint64(r.Red), 16)
	hexG := strconv.FormatUint(uint64(r.Green), 16)
	hexB := strconv.FormatUint(uint64(r.Blue), 16)
	return Hex(fmt.Sprintf("%02s%02s%02s", hexR, hexG, hexB)), nil
}

func Hex2RGB(hex Hex) (RGB, error) {
	var rgb RGB
	v, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return rgb, err
	}

	rgb = RGB{
		Red:   uint8(v >> 16),
		Green: uint8((v >> 8) & 0xFF),
		Blue:  uint8(v & 0xFF),
	}

	return rgb, nil
}

func main() {
	// var rgb RGB
	// var err error

	flag.Parse()
	args := flag.Args()

	if len(args) == 1 {
		hex := Hex(args[0])
		rgb, err := hex.toRGB()
		if err != nil {
			fmt.Println("Cannot convert")
			os.Exit(1)
		}
		fmt.Printf("#%s = RGB(%d, %d, %d)", hex, rgb.Red, rgb.Green, rgb.Blue)
	} else if len(args) == 3 {
		var rgb RGB
		if v, err := strconv.ParseUint(args[0], 10, 8); err != nil {
			os.Exit(1)
		} else {
			rgb.Red = uint8(v)
		}
		if v, err := strconv.ParseUint(args[1], 10, 8); err != nil {
			os.Exit(1)
		} else {
			rgb.Green = uint8(v)
		}
		if v, err := strconv.ParseUint(args[2], 10, 8); err != nil {
			os.Exit(1)
		} else {
			rgb.Blue = uint8(v)
		}

		hex, err := rgb.toHex()
		if err != nil {
			fmt.Println("Cannot convert")
			os.Exit(1)
		}
		fmt.Printf("RGB(%d, %d, %d) = #%s", rgb.Red, rgb.Green, rgb.Blue, hex)
	}
}
