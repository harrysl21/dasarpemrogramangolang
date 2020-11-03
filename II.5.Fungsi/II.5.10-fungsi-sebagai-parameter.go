package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
    var result []string
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}


func main() {
	var data = []string{"bruce", "clack", "peter"}
	var ContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var numLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)	

	fmt.Println("filter ada huruf \"o\"\t:", ContainsO)

	fmt.Println("filter jumlah huruf \"5\"\t:", numLenght5)
	
}
