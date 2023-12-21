package main

import "github.com/labstack/echo"

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func GenerateCars() {
	cars = append(cars, Car{Name: "Audi", Price: 1000000})
	cars = append(cars, Car{Name: "Bmw", Price: 1000002})
	cars = append(cars, Car{Name: "Fusca", Price: 10000})
}

func main() {
	GenerateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	return c.JSON(200, cars)
}
