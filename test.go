package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(1974, 12, 1, 0, 0, 0, 0, time.UTC)

	switch a := date.Year(); {
	case a >= 1946 && a <= 1964:
		fmt.Println("Привет, бумер!")
		fallthrough
	case a >= 1965 && a <= 1980:
		fmt.Println("Привет, представитель X!")
		fallthrough
	case a >= 1981 && a <= 1996:
		fmt.Println("Привет, миллениал!")
	case a >= 1997 && a <= 2012:
		fmt.Println("Привет, зумер!")
	case a >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Default value handling")
	}
	fmt.Println(time.Now())
	fmt.Printf("Go launched at %s\n", date.Local())
}
