package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	Copy()
	t.Run("Total de tickets a Uruguay", func(t *testing.T) {
		//Average
		var expectedResult int = 1
		pais := "Uruguay"
		//Act
		result, _ := GetTotalTickets(pais)
		//Assert
		assert.Equal(t, expectedResult, result, "El total no es el esperado")
	})

	t.Run("Total de tickets a Roma", func(t *testing.T) {
		//Average
		var expectedResult int = 45
		//Act
		result, _ := GetTotalTickets("Brazil")
		//Assert
		assert.Equal(t, expectedResult, result, "El total no es el esperado")
	})
}

func TestGetPassengers(t *testing.T) {
	t.Run("Total de tickets en la madrugada", func(t *testing.T) {
		//Average
		var expectedResult int = 303
		//Act
		result, _ := GetPassengers("early")
		//Assert
		assert.Equal(t, expectedResult, result, "El total no es el esperado")
	})
	t.Run("Total de tickets en la manana", func(t *testing.T) {
		//Average
		var expectedResult int = 255
		//Act
		result, _ := GetPassengers("morning")
		//Assert
		assert.Equal(t, expectedResult, result, "El total no es el esperado")
	})
	t.Run("Total de tickets en la tarde", func(t *testing.T) {
		//Average
		var expectedResult int = 289
		//Act
		result, _ := GetPassengers("afternoon")
		//Assert
		assert.Equal(t, expectedResult, result, "El total no es el esperado")
	})
	t.Run("Total de tickets en la noche", func(t *testing.T) {
		//Average
		var expectedResult int = 151
		//Act
		result, _ := GetPassengers("night")
		//Assert
		assert.Equal(t, expectedResult, result, "El total no es el esperado")
	})

}
func TestAverageDestination(t *testing.T) {
	t.Run("Promedio de pasajeros que viajaron a Brasil", func(t *testing.T) {
		//Average
		var expectedResult float64 = 1
		//Act
		result, _ := AverageDestination("Brazil", 45)
		//Assert
		assert.Equal(t, expectedResult, result, "El total no es el esperado")
	})

	t.Run("Promedio de pasajeros que viajaron a Uruguay", func(t *testing.T) {
		//Average
		var expectedResult float64 = 2.2
		//Act
		result, _ := AverageDestination("Mexico", 5)
		//Assert
		assert.Equal(t, expectedResult, result, "El total no es el esperado")
	})
}
