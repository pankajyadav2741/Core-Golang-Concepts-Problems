package src

import (
	"fmt"
	"os"
	"strings"
)

const (
	celciusTemp    = "C"
	fahrenheitTemp = "F"
)

var (
	temp  float64
	again string
)

func StartApp() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}

	for {
		tempType := strings.ToUpper(os.Args[1])
		if tempType != celciusTemp && tempType != fahrenheitTemp {
			fmt.Println("Invalid Temperature Type")
			os.Exit(1)
		}

		if tempType == celciusTemp {
			fmt.Println("Celcius to Fahrenheit Temperature Converter")
		} else if tempType == fahrenheitTemp {
			fmt.Println("Fahrenheit to Celcius Temperature Converter")
		}

		fmt.Printf("Enter current temperature in %s ? : ", tempType)
		fmt.Scanf("%f", &temp)
		if tempType == celciusTemp {
			convertTofahrenheit(temp)
		} else if tempType == fahrenheitTemp {
			convertToCelcius(temp)
		}
		fmt.Print("Do you want to convert another temperature? (y/n)")
		fmt.Scanf("%s", &again)
		again = strings.ToUpper(again)
		if again != "Y" {
			fmt.Println("Good Bye!")
			break
		}
	}
}

func convertToCelcius(temp float64) {
	fmt.Printf("%v F = %.02f C\n", temp, (temp-32)*5/9)
}

func convertTofahrenheit(temp float64) {
	fmt.Printf("%v C = %.02f F\n", temp, (temp*9/5)+32)
}
