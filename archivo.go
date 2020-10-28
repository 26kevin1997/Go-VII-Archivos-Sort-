package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var archivo []string
	var a string
	var n int

	fmt.Println("¿Cuantos string quiere? ")
	fmt.Scan(&n)
	for x := 0; x < n; x++ {
		fmt.Println("Palabra: ")
		fmt.Scan(&a)
		archivo = append(archivo, a)
	}
	for x := 0; x < n; x++ {
		fmt.Print(archivo[x], " ")
	}
	fmt.Print("\n")
	//Ordeno de forma ascendente
	sort.Strings(archivo)
	fmt.Println(archivo)
	//Guardo en un txt
	file, err := os.Create("asecendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for x := 0; x < n; x++ {
		file.WriteString(archivo[x])
		if x < n-1 {
			file.WriteString("\n")
		}
	}
	//Ordeno de forma descendente
	sort.Sort(sort.Reverse(sort.StringSlice(archivo)))
	fmt.Println(archivo)

	//Guardo en un txt
	file, errr := os.Create("descendente.txt")
	if errr != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for x := 0; x < n; x++ {
		file.WriteString(archivo[x])
		if x < n-1 {
			file.WriteString("\n")
		}
	}

}
