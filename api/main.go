package main

import (
	"database/sql"

	"github.com/labstack/echo"
)

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
	saveCar(*car)
	return c.JSON(200, cars)
}

func saveCar(car Car) error {
	db, err := sql.Open("postgres", "host=localhost port=5050 user=root password=root dbname=cars.db sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, price) VALUES (s1, s2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}
