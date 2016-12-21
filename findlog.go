package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	base, product float64
)

func main() {
	usage := fmt.Sprintf("usage: %s [base] product\n", os.Args[0])
	if len(os.Args) > 3 {
		fmt.Printf("%s: too many arguments\n", os.Args[0])
		fmt.Print(usage)
		return
	}

	if len(os.Args) == 1 {
		fmt.Printf("%s: not enough arguments\n", os.Args[0])
		fmt.Print(usage)
		return
	}

	if len(os.Args) == 2 {
		var err error
		base = 2.0
		product, err = strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Println(err)
		}
	}

	if len(os.Args) == 3 {
		var err error
		base, err = strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Println(err)
		}
		product, err = strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Println(err)
		}
	}

	for i := 0.0; i < product; i++ {
		if math.Pow(base, i) == product {
			fmt.Printf("%.f\n", i)
			return
		}
	}
	fmt.Printf("could not find %.f in exponents of %.f\n", product, base)

}
