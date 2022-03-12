package units

import "fmt"

// temperature units
type Celsius float64
type Fahrenheit float64
type Kelvin float64

// distance units
type Feet float64
type Meter float64

// weight units
type Pound float64
type Kilo float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g\u00B0C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g\u00B0F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g'", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (lb Pound) String() string {
	return fmt.Sprintf("%g lb", lb)
}

func (kg Kilo) String() string {
	return fmt.Sprintf("%g kg", kg)
}
