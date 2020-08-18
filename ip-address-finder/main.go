package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "192125213112"
	fmt.Println("convert", s)
	convertIps(s)

	s = "12121212"
	fmt.Println("convert", s)
	convertIps(s)

	s = "127001"
	fmt.Println("convert", s)
	convertIps(s)

	s = "911111111"
	fmt.Println("convert", s)
	convertIps(s)
}

func convertIps(s string) {
	length := len(s)
	if length > 4 && length <= 12 {
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				for k := 1; k <= 3; k++ {
					for m := 1; m <= 3; m++ {
						sum := i + j + k + m
						if sum > length {
							break
						}
						position1, remainder1 := s[:i], s[i:]
						val, err := strconv.Atoi(position1)
						if err != nil || val > 255 {
							break
						}
						position2, remainder2 := remainder1[:j], remainder1[j:]
						val, err = strconv.Atoi(position2)
						if err != nil || val > 255 {
							break
						}
						position3, remainder3 := remainder2[:k], remainder2[k:]
						val, err = strconv.Atoi(position3)
						if err != nil || val > 255 {
							break
						}
						position4, remainder4 := remainder3[:m], remainder3[m:]
						val, err = strconv.Atoi(position4)
						if (err != nil) || (val > 255) || len(remainder4) > 0 {
							continue
						}
						fmt.Println(position1 + "." + position2 + "." + position3 + "." + position4)
					}
				}
			}
		}
	}
}
