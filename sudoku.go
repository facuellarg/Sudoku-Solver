package main

import (
	//"github.com/facuellarg/problem"
	"fmt"
)

//Sudoku struc sudoku
type Sudoku struct {
	state                                  [9][9]Casilla
	xMaximaRestriccion, yMaximaRestriccion int
}

//Casilla estructura tipo casilla
type Casilla struct {
	valor, fila, columna, restricciones int
	posibilidades                       [9]bool
}

//NewCasilla retorna una nueva casilla
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

//CalcularRestriccion calcula las restricciones de una casilla dada por los indices i,j
func CalcularRestriccion(fila, columna int, sd *Sudoku) {
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

//CalcularRestricciones calculas las restricciones de todas las casillas del sudoku
func (sd *Sudoku) CalcularRestricciones() {
	var max int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sd.state[i][j].valor == 0 {
				CalcularRestriccion(i, j, sd)
				if sd.state[i][j].restricciones > max {
					sd.xMaximaRestriccion = i
					sd.yMaximaRestriccion = j
					max = sd.state[i][j].restricciones
				}
			}
		}
	}
}

func (sd *Sudoku) Solve() bool {

	sd.CalcularRestricciones()
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

//Imprimir imprime el estado actual del sudoku
func (sd *Sudoku) Imprimir() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(sd.state[i][j].valor, " ")
		}
		fmt.Println("")
	}

}
