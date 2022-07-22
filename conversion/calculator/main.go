package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Property-Finder-Patika/week-4-homework-1-furkanarsl/conversion/degree"
)

type Calculator struct {
	functions []degree.DegreeFunctions
}

var flag bool = true

func (c *Calculator) addDegreeFunc(d degree.DegreeFunctions) {
	c.functions = append(c.functions, d)
}

func (c Calculator) doCalculation(name string, arg float64) (string, error) {
	for _, f := range c.functions {
		if strings.EqualFold(name, f.GetName()) {
			result := f.Calculate(arg)
			return result, nil
		}

	}
	return "", errors.New("no function with name: " + name)
}

func StartCalculator() {
	calc := Calculator{}

	calc.addDegreeFunc(degree.Celsius{Name: "C"})
	calc.addDegreeFunc(degree.Fahrenheit{Name: "F"})
	calc.addDegreeFunc(degree.Kelvin{Name: "K"})

	fmt.Println("Choose the unit you want to convert from")

	for _, f := range calc.functions {
		fmt.Println(f.GetName())
	}

	for flag {
		var fName string
		var arg float64
		fmt.Println("> Enter name of the calculation or enter x to exit:")
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if fName == "x" {
			flag = false
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanf("%f", &arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			value, err := calc.doCalculation(fName, arg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Result of %v%s : %v\n", arg, fName, value)
			}
		}
	}
	println("Exiting ...")
}

func main() {
	StartCalculator()
}
