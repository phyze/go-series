package main

import (
	"fmt"
	"time"
)

type Celsius float32

type Fahrenheit float32

func ConvertCelToFah(c Celsius) Fahrenheit {
	return Fahrenheit(c*1.8 + 32)
}

type Date time.Time

func (d Date) String() string {
	layout := "2006-01-02"
	return time.Now().Format(layout)
}

func main() {
	
	//-------------------------
	fmt.Println("convert ")
	fah := ConvertCelToFah(25.4)
	fmt.Println("celsius to fahrenheit = ", fah)

	//-------------------------
	var date Date
	fmt.Println("get date = ", date.String())
}
