package main

import (
	"fmt"
	"math"
)

var a, b, c string
var d, e, f = "dog", "ear", "fan"

func main() {
	fmt.Printf("hello, world\n")
	//test(3, 2)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(Small)
}

func test(i, j int) {
	// k := i + j
	// fmt.Printf("test %d\n", k)
	// a = "foo"
	// fmt.Printf("%s\n", a)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
	v := z // change me!
	fmt.Printf("v is of type %T\n", v)
	const World = "世界"
	fmt.Printf("Hello %s\n", World)

}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
