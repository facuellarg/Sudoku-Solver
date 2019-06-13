# Sudoku-Solver
Solver for sudoku game using bactracking and constraint satisfaction problem



## Install 
go  get github.com/facuellarg/Sudoku-Solver

## Usage

### Import  the file   
~~~
import "github.com/facuellarg/Sudoku-Solver"
~~~
### Created a new sudoku passing a [9][9]int

~~~
sd := NewSudoku([9][9]int)
~~~
Or use the function readSudoku(path of file)
~~~
sd := NewSudoku(readSudoku(path of file))
~~~
#### Format to Read Sudoku
Edit sudoku file, each squeare is separated by single space, empty value is represented by 0 and any other value with the number.

Example  

8 0 0 0 0 0 0 0 0  
0 0 3 6 0 0 0 0 0  
0 7 0 0 9 0 2 0 0  
0 5 0 0 0 7 0 0 0  
0 0 0 0 4 5 7 0 0  
0 0 0 1 0 0 0 3 0  
0 0 1 0 0 0 0 6 8  
0 0 8 5 0 0 0 1 0  
0 9 0 0 0 0 4 0 0  

### Solve the sudoku

Function Solve() return true if can resolve the sudoku and false in other case
~~~
sd.Solve()
~~~

To see the solution

~~~
sd.Print()
~~~

If you want to store the solution in [9][9]int
~~~
if sd.Solve() {
  myArray := sd.GetState()
}
~~~

