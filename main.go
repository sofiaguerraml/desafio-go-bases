package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	tickets.Copy()

	total, err := tickets.GetTotalTickets("Uruguay")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Cantidad de pasajeros:", total)

	total, err = tickets.GetPassengers("early")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Pasajeros que vuelan en la madrugada:", total)

	total, err = tickets.GetPassengers("morning")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Pasajeros que vuelan en la manana:", total)

	total, err = tickets.GetPassengers("afternoon")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Pasajeros que vuelan en la afternoon:", total)

	total, err = tickets.GetPassengers("night")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Pasajeros que vuelan en la noche:", total)

	var total2 float64
	total2, err2 := tickets.AverageDestination("Brazil", 50)
	if err2 != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Promedio de pasajeros:", total2)

}
