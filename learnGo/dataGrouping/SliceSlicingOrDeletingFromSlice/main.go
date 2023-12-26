package main

import "fmt"

func main() {

	sabores := []string{"queijo", "cheedar", "alface", "bacon", "frango"}
	sabores = append(sabores[:2], sabores[4:]...)
	fmt.Println(sabores)
}
