package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Life struct {
	Grid [][]string
}

func (l *Life) Out(col *int, row *int) {
	colorRed := "\033[34m"
	colorReset := "\033[0m"
	for i := 0; i < *col; i++ {
		for j := 0; j < *row; j++ {
			if l.Grid[i][j] == "*" {
				fmt.Print(string(colorRed), l.Grid[i][j])
			} else {
				fmt.Print(string(colorReset), l.Grid[i][j])
			}
		}
		fmt.Println()
	}
}

func (l *Life) Make(col *int, row *int, number int) {
	count := 0
	for i := 0; i < *col; i++ {
		for j := 0; j < *row; j++ {
			l.Grid[i][j] = "."
		}
	}
	for {
		i := rand.Intn(*col)
		j := rand.Intn(*row)
		if l.Grid[i][j] == "." {
			l.Grid[i][j] = "*"
			count++
		}
		if count == number {
			break
		}
	}
}

func (l *Life) MakeFigure(col *int, row *int, view string) {
	for i := 0; i < *col; i++ {
		for j := 0; j < *row; j++ {
			l.Grid[i][j] = "."
		}
	}
	switch view {
	case "Glider Generator":
		for i := 5; i <= 13; i += 2 {
			for j := 11; j <= 22; j++ {
				if i == 5 && j >= 15 && j <= 18 {
					l.Grid[i][j] = "*"
				}
				if i == 7 && j >= 13 && j <= 20 {
					l.Grid[i][j] = "*"
				}
				if i == 9 && j >= 11 && j <= 22 {
					l.Grid[i][j] = "*"
				}
				if i == 11 && j >= 13 && j <= 20 {
					l.Grid[i][j] = "*"
				}
				if i == 13 && j >= 15 && j <= 18 {
					l.Grid[i][j] = "*"
				}
			}
		}
	case "Revolver":
		l.Grid[3][5] = "*"
		l.Grid[3][6] = "*"
		l.Grid[3][11] = "*"
		l.Grid[3][12] = "*"
		l.Grid[4][6] = "*"
		l.Grid[4][11] = "*"
		l.Grid[5][6] = "*"
		l.Grid[5][11] = "*"
		l.Grid[5][8] = "*"
		l.Grid[5][9] = "*"
		l.Grid[6][7] = "*"
		l.Grid[6][10] = "*"
		l.Grid[7][9] = "*"
		l.Grid[8][7] = "*"
		l.Grid[9][10] = "*"
		l.Grid[9][11] = "*"
		l.Grid[10][6] = "*"
		l.Grid[10][7] = "*"
		l.Grid[11][10] = "*"
		l.Grid[12][8] = "*"
		l.Grid[13][7] = "*"
		l.Grid[13][10] = "*"
		l.Grid[14][6] = "*"
		l.Grid[14][8] = "*"
		l.Grid[14][9] = "*"
		l.Grid[14][11] = "*"
		l.Grid[15][6] = "*"
		l.Grid[15][11] = "*"
		l.Grid[16][6] = "*"
		l.Grid[16][11] = "*"
		l.Grid[16][5] = "*"
		l.Grid[16][12] = "*"

	}

}

func (l *Life) Game(col *int, row *int) {
	for i := 0; i < *col; i++ {
		for j := 0; j < *row; j++ {
			if l.CheckToCreateLife(col, row, i, j) && l.Grid[i][j] == "." {
				l.Grid[i][j] = "#"
			}
		}
	}
	for i := 0; i < *col; i++ {
		for j := 0; j < *row; j++ {
			if l.CheckToCreateDeath(col, row, i, j) && (l.Grid[i][j] == "*" || l.Grid[i][j] == "%") {
				l.Grid[i][j] = "%"
			}
		}
	}
	for i := 0; i < *col; i++ {
		for j := 0; j < *row; j++ {
			if l.Grid[i][j] == "#" {
				l.Grid[i][j] = "*"
			}
			if l.Grid[i][j] == "%" {
				l.Grid[i][j] = "."
			}
		}
	}
}

func (l *Life) CheckToCreateLife(col *int, row *int, lcol int, lrow int) bool {
	count := 0
	for i := lcol - 1; i <= lcol+1; i++ {
		for j := lrow - 1; j <= lrow+1; j++ {
			if i >= 0 && i < *col && j >= 0 && j < *row {
				if l.Grid[i][j] == "*" || l.Grid[i][j] == "%" {
					count++
				}

			}
		}
	}
	if count == 3 {
		return true
	} else {
		return false
	}
}

func (l *Life) CheckToCreateDeath(col *int, row *int, lcol int, lrow int) bool {
	count := -1
	for i := lcol - 1; i <= lcol+1; i++ {
		for j := lrow - 1; j <= lrow+1; j++ {
			if i >= 0 && i < *col && j >= 0 && j < *row {
				if l.Grid[i][j] == "*" || l.Grid[i][j] == "%" {
					count++
				}
			}
		}
	}
	if count == 3 || count == 2 {
		return false
	}

	return true
}

func main() {
	col := 18
	row := 30
	Grid := make([][]string, col)
	for i := range Grid {
		Grid[i] = make([]string, row)
	}
	Field := Life{Grid}
	fmt.Println("Введите, что хотите увидить:")
	fmt.Println("1, если хотите задать своё значение живых клеток")
	fmt.Println("2, если хотите увидеть Glider Generator ")
	fmt.Println("3, если хотите увидеть Revolver ")
	Wish := ""
	fmt.Scan(&Wish)
	switch Wish {
	case "1":
		fmt.Println("Введите число живых клеток, не привыющих", col*row)
		Number := 0
		fmt.Scan(&Number)
		Field.Make(&col, &row, Number)
	case "2":
		Field.MakeFigure(&col, &row, "Glider Generator")
	case "3":
		Field.MakeFigure(&col, &row, "Revolver")

	}

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for {
		Field.Out(&col, &row)
		Field.Game(&col, &row)
		time.Sleep(1 * time.Second)
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}
