package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"
)

type Ticket struct {
	Id                 string
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              string
}

var tickets []Ticket

const (
	early     = "early"
	morning   = "morning"
	afternoon = "afternoon"
	night     = "night"
)

func Copy() {
	file, err := os.Open("/Users/sguerra/Documents/desafioGo/desafio-go-bases/tickets.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	read := csv.NewReader(file)
	read.Comma = ','
	read.FieldsPerRecord = 6
	for {
		data, err := read.Read()
		if err != nil {
			break
		}
		ticket := Ticket{
			Id:                 data[0],
			Name:               data[1],
			Email:              data[2],
			DestinationCountry: data[3],
			FlightTime:         data[4],
			Price:              data[5],
		}
		tickets = append(tickets, ticket)
	}
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	count := 0
	for _, tciket := range tickets {
		if tciket.DestinationCountry == destination {
			count++
		}
	}
	if count == 0 {
		return 0, errors.New("Error: No se encuentran tickets a este destino")
	}
	return count, nil
}

// ejemplo 2
func GetPassengers(times string) (int, error) {
	count := 0
	switch times {
	case early: // de 00:00 a 06:59
		minStr := "00:00"
		maxStr := "06:59"
		min, _ := time.Parse("15:04", minStr)
		max, _ := time.Parse("15:04", maxStr)
		for _, tciket := range tickets {
			actualStr := tciket.FlightTime
			actual, _ := time.Parse("15:04", actualStr)
			if actual.After(min) && actual.Before(max) {
				count++
			}
		}
	case morning: // de 07:00 a 12:59
		minStr := "07:00"
		maxStr := "12:59"
		min, _ := time.Parse("15:04", minStr)
		max, _ := time.Parse("15:04", maxStr)
		for _, tciket := range tickets {
			actualStr := tciket.FlightTime
			actual, _ := time.Parse("15:04", actualStr)
			if actual.After(min) && actual.Before(max) {
				count++
			}
		}
	case afternoon: // de 13:00 a 19:59
		minStr := "13:00"
		maxStr := "19:59"
		min, _ := time.Parse("15:04", minStr)
		max, _ := time.Parse("15:04", maxStr)
		for _, tciket := range tickets {
			actualStr := tciket.FlightTime
			actual, _ := time.Parse("15:04", actualStr)
			if actual.After(min) && actual.Before(max) {
				count++
			}
		}
	case night: // de 20:00 a 23:59
		minStr := "20:00"
		maxStr := "23:59"
		min, _ := time.Parse("15:04", minStr)
		max, _ := time.Parse("15:04", maxStr)
		for _, tciket := range tickets {
			actualStr := tciket.FlightTime
			actual, _ := time.Parse("15:04", actualStr)
			if actual.After(min) && actual.Before(max) {
				count++
			}
		}
	default:
		return 0, errors.New("Error: Se ingreso un turno incorrecto")
	}

	return count, nil
}

// ejemplo 3
func AverageDestination(destination string, total int) (float64, error) {
	totalToday, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	if total == 0 {
		return 0, errors.New("Error: el total de pasajeros es 0, no se puede realizar el promedio")
	}
	var resultado float64 = float64(totalToday) / float64(total)
	return resultado, nil
}
