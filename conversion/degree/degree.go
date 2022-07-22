package degree

import "fmt"

type DegreeFunctions interface {
	Calculate(arg float64) string
	GetName() string
}

type Celsius struct {
	Name string
}

func (c Celsius) GetName() string { return c.Name }

func (c Celsius) Calculate(arg float64) string {
	return fmt.Sprintf("%v %v", FString(CToF(arg)), KString(CToK(arg)))
}

type Fahrenheit struct {
	Name string
}

func (f Fahrenheit) GetName() string { return f.Name }

func (f Fahrenheit) Calculate(arg float64) string {
	return fmt.Sprintf("%v %v", CString(FToC(arg)), KString(FToK(arg)))
}

type Kelvin struct {
	Name string
}

func (k Kelvin) GetName() string { return k.Name }

func (k Kelvin) Calculate(arg float64) string {
	return fmt.Sprintf("%v %v", CString(KToC(arg)), FString(KToF(arg)))
}

// Functions to convert from various units to each other
func CToF(arg float64) float64 { return (arg * 9 / 5) + 32 }
func CToK(arg float64) float64 { return arg + 273.15 }

func FToC(arg float64) float64 { return (arg - 32) * 5 / 9 }
func FToK(arg float64) float64 { return CToK(FToC(arg)) }

func KToC(arg float64) float64 { return (arg - 273.15) }
func KToF(arg float64) float64 { return KToC(arg)*9/5 + 32 }

// Functions to convert to string and show degrees with their symbols.
func CString(arg float64) string { return fmt.Sprintf("%g°C", arg) }
func FString(arg float64) string { return fmt.Sprintf("%g°F", arg) }
func KString(arg float64) string { return fmt.Sprintf("%g°K", arg) }
