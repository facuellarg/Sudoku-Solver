package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readSudoku(input string) [9][9]int {
	var matriz [9][9]int
	text, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Print(err)
	}

	values := strings.Fields(string(text))

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			matriz[i][j], _ = strconv.Atoi(values[i*9+j])
		}
	}
	return matriz

}

func imprimirMatriz(matriz [9][9]int) {
	for _, fila := range matriz {
		for _, s := range fila {
			fmt.Printf("%d ", s)
		}
		fmt.Println("")
	}
}
func main() {
	sd := NewSudoku(readSudoku("sudoku"))
	sd.Imprimir()
	fmt.Println(sd.isSolved())

	fmt.Println(sd.Solve())
	sd.Imprimir()

}
