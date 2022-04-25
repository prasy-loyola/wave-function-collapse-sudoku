package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
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


func remove(slice []int, value int) []int {

	for i := 0; i < len(slice); i++ {
		if value == slice[i] {
			slice[i] = slice[len(slice)-1]
			return slice[:len(slice)-1]
		}
	}
	return slice
}

func waveCollapse(sudoku Sudoku, i int , j int ) Sudoku {
	cell := sudoku.cells[i][j]
	if cell.value > 0 {
		return sudoku
	}
	//checking rows
	for j2 := 0; j2 < 9; j2++ {
		if j != j2{
		cell.possibilities = remove(cell.possibilities, sudoku.cells[i][j2].value)
		}
	}
	//checking cols
	for i2 := 0; i2 < 9; i2++ {
		if i != i2 {
		cell.possibilities = remove(cell.possibilities, sudoku.cells[i2][j].value)
		}
	}

	box_i := i / 3
	box_j := j / 3

	for i3:= box_i * 3; i3 < (box_i + 1) * 3; i3++ {
		for j3 := box_j * 3; j3 < (box_j +1) * 3; j3++ {
			if i3 == i && j3 == j {
				continue
			}
		cell.possibilities = remove(cell.possibilities, sudoku.cells[i3][j3].value)
		}
	}

	if len(cell.possibilities) == 1 {

		cell.value = cell.possibilities[0]

	}
	sudoku.cells[i][j] = cell
    fmt.Printf("Number of possibilities for %d, %d: %d\n", i,j , len(cell.possibilities))
	return sudoku
}

func collapseWaveFunc(sudoku Sudoku) (Sudoku,bool) {
	all_collapsed := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku.cells[i][j].value > 0 {
				continue
			} else {
				all_collapsed = false
				sudoku = waveCollapse(sudoku,i,j)
			}
		}
	}
	return sudoku,all_collapsed
}

func waitForUser()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to collapse next step: ")
	reader.ReadString('\n')
}

func main(){
	var sudoku =loadSudoku(`0 0 0 0 0 0 0 0 7
7 2 0 3 0 9 0 0 1
0 0 8 7 0 5 0 6 0
5 0 2 8 9 0 0 0 0
0 4 0 5 0 1 0 9 0
0 0 0 0 6 3 7 0 5
0 3 0 9 0 6 1 0 0
2 0 0 1 0 7 0 5 3
9 0 0 0 0 0 0 0 0
`)
	noOfSteps := 0
	for true {
		display(sudoku)
		//waitForUser()
		sudoku, solved := collapseWaveFunc(sudoku)
		noOfSteps++
		if solved{
			fmt.Println("Sudoku solved. No. of Steps: " , noOfSteps)
			display(sudoku)
			break
		}
	}
}
