package main

import (
	"fmt"
	"image/png"
	"image/jpeg"
	"image"
	"image/color"
	"math"
	"os"
	"bytes"
	"io/ioutil"
)

var Reset = "\033[0m"

// var Red = "\033[31m"
// var Green = "\033[32m"
// var Yellow = "\033[33m"
// var Blue = "\033[34m"
// var Purple = "\033[35m"
// var Cyan = "\033[36m"
// var Gray = "\033[37m"
// var White = "\033[97m"

type Color struct {
	code string
	rgb  [3]byte
}

func hypot3d(p1 [3]int, p2 [3]int) int {
	dx := float64(p2[0] - p1[0])
	dy := float64(p2[1] - p1[1])
	dz := float64(p2[2] - p1[2])
	return int(math.Sqrt(dx*dx + dy*dy + dz*dz))
}

func closestColor(colors []Color, color [3]byte) Color {
	selColor := colors[0]
	minDistance := math.MaxInt64
	for _, availColor := range colors {
		distance := hypot3d([3]int{int(color[0]), int(color[1]), int(color[2])},
			[3]int{int(availColor.rgb[0]), int(availColor.rgb[1]), int(availColor.rgb[2])})

		if distance < minDistance {
			minDistance = distance
			selColor = availColor
		}
	}

	return selColor
}

func main() {
	colors := [9]Color{
		Color{"\033[31m", [3]byte{128, 0, 0}}, // Red
		Color{"\033[32m", [3]byte{0, 128, 0}}, // Green
		Color{"\033[33m", [3]byte{128, 128, 0}}, // Yellow
		Color{"\033[34m", [3]byte{0, 0, 128}}, // Blue
		Color{"\033[35m", [3]byte{128, 0, 128}}, // Purple
		Color{"\033[36m", [3]byte{0, 128, 128}}, // Cyan
		Color{"\033[37m", [3]byte{192, 192, 192}}, // Gray
		Color{"", [3]byte{0, 0, 0}}, // Black
		Color{"\033[97m", [3]byte{255, 255, 255}}} // White

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("cli-image <path> <draw-character>")
		return
	}

	path := args[0]
	unit := "#"
	if len(args) >= 2 {
		unit = args[1]
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	pngSignature := []byte { 0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A }
	jpgSignature := []byte { 0xFF, 0xD8 }
	var img image.Image
	if bytes.Compare(file[0:8], pngSignature) == 0 {
		img, err = png.Decode(bytes.NewReader(file))
		if err != nil {
			panic(err)
		}
	} else if bytes.Compare(file[0:2], jpgSignature) == 0 {
		img, err = jpeg.Decode(bytes.NewReader(file))
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("Only PNG and JPG are supported")
		return
	}

	x0, y0 := img.Bounds().Min.X, img.Bounds().Min.Y 
	x1, y1 := img.Bounds().Max.X, img.Bounds().Max.Y
	x := x1 - x0
	y := y1 - y0

	for vy := 0; vy < y; vy++ {
		for vx := 0; vx < x; vx++ {
			r, g, b, a := color.RGBAModel.Convert(img.At(vx, vy)).RGBA()
			color := [3]byte { byte(r), byte(g), byte(b) }
			alpha := float64(a)/65535.0
			color[0] = byte(float64(color[0]) * alpha)
			color[1] = byte(float64(color[1]) * alpha)
			color[2] = byte(float64(color[2]) * alpha)
			selColor := closestColor(colors[:], color)

			if selColor.rgb[0] == 0 && selColor.rgb[1] == 0 && selColor.rgb[2] == 0 {
				fmt.Print(" ")
				continue
			}
			fmt.Print(selColor.code + unit + Reset)			
		}
		fmt.Printf("\n")
	}
}
