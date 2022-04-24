package main

import (
"fmt"
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

func defaultCell() Cell{
	 return Cell{0, []int{1,2,3,4,5,6,7,8,9} }
}
func main(){
	var sudoku = Sudoku{}
	sudoku.cells = make([][]Cell,9)
	for i:= 0; i < 9; i++ {
		sudoku.cells[i] = make([]Cell,9)
		for j:= 0; j < 9 ; j++ {
			sudoku.cells[i][j] =  defaultCell()
		}
	}
	display(sudoku)

}
