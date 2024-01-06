package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"banana", "maça", "laranja"}

	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println(ss)
}