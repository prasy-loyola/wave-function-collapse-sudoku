package main

import (
	"fmt"
	"strings"
	"strconv"
)

func display(sudoku Sudoku){
	for i:=0; i< 9 ; i++ {
		if (i % 3 == 0) {
			fmt.Println(" =====================================")
		}else {
			fmt.Println(" -------------------------------------")
		}
		for j := 0; j < 9 ; j++ {
			if j % 3 == 0 {
				fmt.Print(" | ")
			}else{
				fmt.Print(" : ")
			}
		     fmt.Print(sudoku.cells[i][j].value)
		}
		fmt.Println(" |")	
	}
	fmt.Println(" =====================================")
}

type Cell struct {
	value int
    possibilities []int
}
type Sudoku struct {
	cells [][]Cell
}


func loadSudoku(gameState string) Sudoku {
	var sudoku = Sudoku{}
	sudoku.cells = make([][]Cell,9)
	rows := strings.Split(gameState,"\n")
	fmt.Println(len(rows))
	for i:= 0; i < 9; i++ {
		sudoku.cells[i] = make([]Cell,9)
		row := strings.Fields(rows[i])
		for j:= 0; j < 9 ; j++ {
			if num, err := strconv.Atoi(row[j]); err == nil && num > 0 {
				sudoku.cells[i][j] =  Cell{num, []int{}}
			}else{
				sudoku.cells[i][j] =  Cell{0, []int{1,2,3,4,5,6,7,8,9}}
			}
		}
	}
	return sudoku
}
func main(){
	var sudoku =loadSudoku(`0 4 0 0 0 0 1 7 9
0 0 2 0 0 8 0 5 4
0 0 6 0 0 5 0 0 8
0 8 0 0 7 0 9 1 0
0 5 0 0 9 0 0 3 0
0 1 9 0 6 0 0 4 0
3 0 0 4 0 0 7 0 0
5 7 0 1 0 0 2 0 0
9 2 8 0 0 0 0 6 0
`)
	display(sudoku)

}
