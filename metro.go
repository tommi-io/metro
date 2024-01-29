package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Stop struct {
    nome string
    indice int //numero della fermata, per stampare nell'output dopo quante fermate deve scendere/cambiare linea
	linea string //nome della linea
    innesto *[]Stop //puntatore all'array della linea d'innesto (se esiste)
}

var (
	Gialla = []Stop {
	Stop{nome: "S. Donato", indice: 1, linea: "Gialla", innesto: nil},
	Stop{nome: "", indice: 2, linea: "Gialla", innesto: nil},
	
	}
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Da quale fermata vuoi partire? ")
	partenza, _ := reader.ReadString('\n')
	fmt.Print("Dove vuoi arrivare? ")
	arrivo, _ := reader.ReadString('\n')
	partenza = strings.TrimSpace(partenza)
	arrivo = strings.TrimSpace(arrivo)
	
}

//dato un nome di fermata, restituisce la linea in cui si trova e la struct della fermata stessa
func getLine(fermata string) ([]Stop, Stop) {
	
}

//data un nome di fermata e una linea, returna true se la fermata è nella linea/false se non lo è.
func isOnSameLine(fermata Stop, linea map[Stop][]string) bool {
	
}