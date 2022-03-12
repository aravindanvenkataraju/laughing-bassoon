package units

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

func FToM(f Feet) Meter { return Meter(f / 3.2808) }

func MToF(m Meter) Feet { return Feet(m * 3.2808) }

func PToK(lb Pound) Kilo { return Kilo(lb / 2.2046) }

func KToP(kg Kilo) Pound { return Pound(kg * 2.2046) }
