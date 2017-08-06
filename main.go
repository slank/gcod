package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var width, depth, z, toolWidth, feedrate float64
	var units, unitCode string
	flag.Float64Var(&width, "width", 0, "The width (X) of the surfaced area")
	flag.Float64Var(&depth, "depth", 0, "The depth (Y) of the surfaced area")
	flag.Float64Var(&z, "height", 0, "The height (Z) value for the surfacing run")
	flag.Float64Var(&toolWidth, "tool-width", 0, "The width of the tool being used to surface")
	flag.Float64Var(&feedrate, "feedrate", 0, "Tool feed rate")
	flag.StringVar(&units, "units", "mm", "The units to use (inches/[mm])")
	flag.Parse()

	if units == "mm" {
		unitCode = "G21"
	} else if units == "inch" || units == "in" {
		unitCode = "G20"
	} else {
		log.Fatalf("Invalid unit '%s'. Should be 'in' or 'mm'.", units)
	}

	if width == 0 || depth == 0 || toolWidth == 0 {
		log.Fatalf("Width and depth must be greater than zero.")
	}

	// preamble
	fmt.Println("% Surfacing program")
	fmt.Printf("(W=%.5f H=%.5f D=%.5f TW=%.5f FR=%.5f U=%s)\n", width, depth, z, toolWidth, feedrate, units)
	fmt.Println("O1000")
	fmt.Println("G00 X0Y0Z0")
	fmt.Printf("F%.5f %s\n", feedrate, unitCode)
	fmt.Printf("G1 Z%.5f\n", z)

	// surfacing loop
	var x, y float64
	for x = 0.0; x < width; x = x + toolWidth*3/4 {
		if y == 0 {
			y = depth
		} else {
			y = 0.0
		}
		if x > width {
			x = width
		}
		fmt.Printf("G01 X%.5fY%.5f\n", x, y)
	}

	// end
	fmt.Println("%")
}
