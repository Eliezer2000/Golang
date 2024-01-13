package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 42
		close(canal)
	}()

	v, ok := <-canal
	fmt.Println(v, ok)



	Npar := make(chan int)
	Nimpar := make(chan int)
	Nquit := make(chan bool)

	go mandaNumeros(Npar, Nimpar, Nquit)
	receive(Npar, Nimpar, Nquit)
}

func mandaNumeros (Npar, Nimpar chan int, Nquit chan bool) {
	total := 500
	for i := 0; i < total; i++{
		if i % 2 == 0{
			Npar <- i
		} else {
			Nimpar <- i
		}
	}
	close(Npar)
	close(Nimpar)
	Nquit<-true
}


func receive (par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número ", v, "é par.")
		case v := <-impar:
			fmt.Println("O número", v, "é impar.")
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Bad :(", v)
			} else {
				fmt.Println("Encerrando :)", v)
			}
			return
		}
	}
}