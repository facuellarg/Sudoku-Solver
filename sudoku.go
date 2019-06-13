package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//Sudoku struc sudoku
type Sudoku struct {
	state                                  [9][9]Casilla
	xMaximaRestriccion, yMaximaRestriccion int
}

type Casilla struct {
	valor, fila, columna, restricciones int
	posibilidades                       [9]bool
}

func NewCasilla(fila, columna int) Casilla {
	this := Casilla{}
	this.fila = fila
	this.columna = columna
	this.posibilidades = [9]bool{true, true, true, true, true, true, true, true, true}
	this.restricciones = 0
	return this
}

//NewSudoku retorna un nuevo sudoku dado el estado inicial de valores
func NewSudoku(valores [9][9]int) (sudoku Sudoku) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku.state[i][j] = NewCasilla(i, j)
			sudoku.state[i][j].valor = valores[i][j]
		}
	}
	sudoku.xMaximaRestriccion = 0
	sudoku.yMaximaRestriccion = 0
	return
}

func CalculateConstrain(fila, columna int, sd *Sudoku) {
	inicioSeccionX := (fila / 3) * 3
	inicioSeccionY := (columna / 3) * 3
	var seccionX, seccionY int
	for k := 0; k < 9; k++ {
		seccionX = inicioSeccionX + (k / 3)
		seccionY = inicioSeccionY + (k % 3)
		//fmt.Println("valor x ", seccionX, "valor y ", seccionY)

		if val := (*sd).state[fila][k].valor; val != 0 && (*sd).state[fila][columna].posibilidades[val-1] {
			(*sd).state[fila][columna].posibilidades[val-1] = false
			(*sd).state[fila][columna].restricciones++
		}
		if val := (*sd).state[k][columna].valor; val != 0 && (*sd).state[fila][columna].posibilidades[val-1] {
			(*sd).state[fila][columna].posibilidades[val-1] = false
			(*sd).state[fila][columna].restricciones++
		}
		if val := (*sd).state[seccionX][seccionY].valor; val != 0 && (*sd).state[fila][columna].posibilidades[val-1] {
			(*sd).state[fila][columna].posibilidades[val-1] = false
			(*sd).state[fila][columna].restricciones++
		}

	}

}

func (sd *Sudoku) CalculateConstrains() {
	var max int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sd.state[i][j].valor == 0 {
				CalculateConstrain(i, j, sd)
				if sd.state[i][j].restricciones > max {
					sd.xMaximaRestriccion = i
					sd.yMaximaRestriccion = j
					max = sd.state[i][j].restricciones
				}
			}
		}
	}
}

//Solve solve the sudoku problem
func (sd *Sudoku) Solve() bool {
	sd.CalculateConstrains()
	if sd.isSolved() {
		return true
	}
	for i := 0; i < 9; i++ {
		if sd.state[sd.xMaximaRestriccion][sd.yMaximaRestriccion].posibilidades[i] {
			// sd.state[sd.xMaximaRestriccion][sd.yMaximaRestriccion].valor = i + 1
			var newSudoku Sudoku
			newSudoku.state = sd.state
			newSudoku.state[sd.xMaximaRestriccion][sd.yMaximaRestriccion].valor = i + 1
			if newSudoku.Solve() {
				sd.state = newSudoku.state
				return true
			}
		}
	}
	return false
}

func (sd *Sudoku) isSolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sd.state[i][j].valor == 0 {
				return false
			}
		}
	}
	return true
}

//Print imprime el estado actual del sudoku
func (sd *Sudoku) Print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(sd.state[i][j].valor, " ")
		}
		fmt.Println("")
	}

}

//GetState give the state of the problem in format [9][9]int
func (sd Sudoku) GetState() (sol [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sol[i][j] = sd.state[i][j].valor
		}
	}
	return
}

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
