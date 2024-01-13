package main

import "fmt"

func main() {
	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("Canal A: ", v)
		case v := <-b:
			fmt.Println("Canal B: ", v)
	
		}
	}

	canal := make(chan int)
	quit := make(chan int)
	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)


	Npar := make(chan int)
	Nimpar := make(chan int)
	Nquit := make(chan bool)

	go mandaNumeros(Npar, Nimpar, Nquit)
	receive(Npar, Nimpar, Nquit)
}


func recebeQuit (canal chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido: ", <-canal)
	}
	quit <- 0
}

func enviaPraCanal (canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <- quit:
			return
		}
	}
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
		case <-quit:
			return
		}
	}
}