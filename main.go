package main

import (
	"fmt"
	"github.com/maritza05/golang-bootcamp-2020/infraestructure/repository"
)

func main() {
	data, err := repository.ReadCsv("satellites.csv")
	if err != nil {
		fmt.Println("An error ocurred while fetching data: ", err)
		return
	}
	fmt.Println(data)
}
