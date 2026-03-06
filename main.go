package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	_ "golang.org/x/image/webp"

	"golang.org/x/image/draw"
)

func pixelsToANSI(px1, px2 [4]uint8) string {
	res := ""

	if px1[3] != 255 || px2[3] != 255 {
		res += "\033[0m"
	}

	if px1[3] != 255 && px2[3] != 255 {
		res += " "
		return res
	}

	if px1[3] == 255 {
		res += fmt.Sprintf("\033[38;2;%d;%d;%dm", px1[0], px1[1], px1[2])
		if px2[3] == 255 {
			res += fmt.Sprintf("\033[48;2;%d;%d;%dm", px2[0], px2[1], px2[2])
		}
		res += "▀"
	} else if px2[3] == 255 {
		res += fmt.Sprintf("\033[38;2;%d;%d;%dm", px2[0], px2[1], px2[2])
		res += "▄"
	}

	return res
}

func imageToANSI(img image.Image) string {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	result := ""

	getPixel := func(x, y int) [4]uint8 {
		if x < bounds.Min.X || x >= bounds.Max.X || y < bounds.Min.Y || y >= bounds.Max.Y {
			return [4]uint8{0, 0, 0, 0}
		}
		r, g, b, a := img.At(x, y).RGBA()
		return [4]uint8{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
	}

	for i := 0; i < height/2; i++ {
		for j := 0; j < width; j++ {
			px1 := getPixel(bounds.Min.X+j, bounds.Min.Y+i*2)
			px2 := getPixel(bounds.Min.X+j, bounds.Min.Y+i*2+1)
			result += pixelsToANSI(px1, px2)
		}
		result += "\033[0m\n"
	}

	if height%2 != 0 {
		for j := 0; j < width; j++ {
			px1 := getPixel(bounds.Min.X+j, bounds.Min.Y+height-1)
			px2 := [4]uint8{0, 0, 0, 0}
			result += pixelsToANSI(px1, px2)
		}
	}

	return result
}

func loadImage(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	// Convert to RGBA (equivalent to image.convert("RGBA") in Pillow)
	rgba := image.NewNRGBA(src.Bounds())
	draw.Draw(rgba, rgba.Bounds(), src, src.Bounds().Min, draw.Src)

	return rgba, nil
}

func main() {
	output := flag.String("o", "", "output file")
	flag.StringVar(output, "output", "", "output file")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: img2ansi [-o output] filename")
		os.Exit(1)
	}

	filename := flag.Arg(0)

	img, err := loadImage(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading image: %v\n", err)
		os.Exit(1)
	}

	result := imageToANSI(img)

	if *output != "" {
		if err := os.WriteFile(*output, []byte(result), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "error writing output: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(result)
	}
}
