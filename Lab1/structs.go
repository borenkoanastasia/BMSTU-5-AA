package main

import "fmt"

// Структуры

// Обычная матрица
type Matrix struct {
    matrix [][]int
    rows int
    columns int
}
// Матрица для алгоритма Левенштейна
type MatrixLeventein struct {
    matrix Matrix
    x string // horisontal
    y string // vertical
}

// Функции для поиска минимума
func min3(x1 int, x2 int, x3 int) int{
	if (x1 <= x2 && x1 <= x3){
		return x1
	} else if (x2 <= x3) {
		return x2
	}
	return x3
}
func min4(x1 int, x2 int, x3 int, x4 int) int{
	if (x1 <= x2 && x1 <= x3 && x1 <= x4){
		return x1
	} else if (x2 <= x3 && x2 <= x4) {
		return x2
	} else if (x3 <= x4) {
		return x3
	}
	return x4
}

// Ввод строки
func input_string()string{
	var str1 string
    fmt.Scanf("%s\n", &str1)
	return str1
}

// Создание пустой обычной матрицы
func make_empty_matrix(rows int, columns int)  Matrix{
    var m1 Matrix
    m1.matrix = make([][]int, rows)
    m1.rows = rows
    m1.columns = columns
    for row:= 0; row < m1.rows; row++{
        m1.matrix[row] = make([]int, columns)
    }
    return m1
}

// Вывод обычной матрицы
func print_matrix(matrix Matrix) {
    for row:= 0; row < matrix.rows; row++{
        for column:=0; column < matrix.columns; column++{
            fmt.Printf("%3d ", matrix.matrix[row][column]);
        }
        fmt.Println()
    }
}

// Вывод матрицы Левенштейн
func print_levenshtein_matrix(str1 string, str2 string, m Matrix){
	for i := 0; i < len(str1) + 1; i++{
		if i == 0{
			fmt.Printf("%3c ", '\\')
			for j:=0;j < len(str2); j++{
				fmt.Printf("%3c", str2[j])
			}
		} else {
			fmt.Printf("%3c ", str1[i - 1])
			for j:=0;j < len(str2); j++{
				fmt.Printf("%3d", m.matrix[i - 1][j])
			}
		}
		fmt.Println()
	}
}

type test_result_t struct{
	len int
	times [4]float64
}

func print_result_auto(res test_result_t){
	fmt.Printf("|%6d", res.len)
	for i := range res.times{
		fmt.Printf("|%24f", res.times[i])
	}
	fmt.Printf("|\n")
}

func print_result_manual(res [2]test_result_t){
	fmt.Printf("|%6s|%24s|%24s|%24s|%24s|\n", " ", "Левенштейн мат.", "Левенштейн рек.", 
											  "Левенштейн рек. с кешем", "Дамерау-Левенштейн рек.")
	line(6+24*4+6)
	for j,cur := range res{
		if j==0{
			fmt.Printf("|%6s", "Время")
			for i := range cur.times{
				fmt.Printf("|%24f", cur.times[i])
			}
		} else{
			fmt.Printf("|%6s", "Ответ")
			for i := range cur.times{
				fmt.Printf("|%24d", int(cur.times[i]))
			}
		}
		fmt.Printf("|\n")
	}
}