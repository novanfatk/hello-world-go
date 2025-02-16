package main

import "fmt"

func main() {
	fmt.Print("masukan nama anda: ")
	var name string
	fmt.Scanln(&name)

	fmt.Printf("halo, %s! selamat datang di dunia Go\n", name)
}
