package main

import "fmt"

const tempKelvin = 373

func main() {

	tempCelsius := tempKelvin - 273
	fmt.Printf("A temperatura de ebulição da água em °C é: %v\n", tempCelsius)

}
