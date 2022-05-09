package main

import (
	"fmt"
	"image/png"
	"image/color"
	"math"
	// "os"
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
	colors := [9]Color{Color{"\033[31m", [3]byte{255, 255, 255}},
		Color{"\033[32m", [3]byte{128, 0, 0}},
		Color{"\033[33m", [3]byte{0, 128, 0}},
		Color{"\033[34m", [3]byte{128, 128, 0}},
		Color{"\033[35m", [3]byte{0, 0, 128}},
		Color{"\033[36m", [3]byte{128, 0, 128}},
		Color{"\033[37m", [3]byte{0, 128, 128}},
		Color{"", [3]byte{0, 0, 0}},
		Color{"\033[97m", [3]byte{192, 192, 192}}}

	// fmt.Println(Red + "Red" + Reset)
	fmt.Println(closestColor(colors[:], [3]byte{0, 255, 255}))

	file, err := ioutil.ReadFile("test.png")
	if err != nil {
		panic(err)
	}

	image, err := png.Decode(bytes.NewReader(file))
	if err != nil {
		panic(err)
	}

	cm := image.ColorModel()
	if cm != color.RGBAModel && cm != color.NRGBAModel && cm != color.AlphaModel && cm != color.GrayModel {
		panic("Only PNG 8 bit per pixel are supported")
	}

	fmt.Println(image)
	x0, y0 := image.Bounds().Min.X, image.Bounds().Min.Y 
	x1, y1 := image.Bounds().Max.X, image.Bounds().Max.Y
	x := x1 - x0
	y := y1 - y0

	unit := '#'
	for vy := 0; vy < y; vy++ {
		for vx := 0; vx < x; vx++ {
			r, g, b, _ := color.RGBAModel.Convert(image.At(vx, vy)).RGBA()
			color := [3]byte { byte(r), byte(g), byte(b) }
			selColor := closestColor(colors[:], color)

			if selColor.rgb[0] == 0 && selColor.rgb[1] == 0 && selColor.rgb[2] == 0 {
				fmt.Print(" ")
				continue
			}
			fmt.Print(selColor.code + string(unit) + Reset)			
		}
		fmt.Printf("\n")
	}
}
