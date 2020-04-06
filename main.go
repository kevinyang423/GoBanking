package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	food := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: food}
	fmt.Println(nico.name)
}
