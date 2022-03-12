package main

import (
	"fmt"
	"os"
	"strconv"
	"units"
)

func main() {
	var temperature, weight, distance float64
	var err error
	switch len(os.Args[1:]) {
	case 0:
		break
	case 1:
		temperature, err = strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
	case 2:
		temperature, err = strconv.ParseFloat(os.Args[1], 64)
		weight, err = strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
	default: // 3 or more
		temperature, err = strconv.ParseFloat(os.Args[1], 64)
		weight, err = strconv.ParseFloat(os.Args[2], 64)
		distance, err = strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("%v = %v\n", units.Celsius(temperature), units.CToF(units.Celsius(temperature)))
	fmt.Printf("%v = %v\n", units.Fahrenheit(temperature), units.FToC(units.Fahrenheit(temperature)))
	fmt.Printf("%v = %v\n", units.Celsius(temperature), units.CToK(units.Celsius(temperature)))
	fmt.Printf("%v = %v\n", units.Kelvin(temperature), units.KToC(units.Kelvin(temperature)))
	fmt.Printf("%v = %v\n", units.Fahrenheit(temperature), units.FToK(units.Fahrenheit(temperature)))
	fmt.Printf("%v = %v\n", units.Kelvin(temperature), units.KToF(units.Kelvin(temperature)))
	fmt.Printf("%v = %v\n", units.Pound(weight), units.PToK(units.Pound(weight)))
	fmt.Printf("%v = %v\n", units.Kilo(weight), units.KToP(units.Kilo(weight)))
	fmt.Printf("%v = %v\n", units.Feet(distance), units.FToM(units.Feet(distance)))
	fmt.Printf("%v = %v\n", units.Meter(distance), units.MToF(units.Meter(distance)))
}
