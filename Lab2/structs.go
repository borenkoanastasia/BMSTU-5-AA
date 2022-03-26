package main 

import "fmt"

type vector_t struct{
	elem[] float32
	size int
}
type matrix_t struct{
	elem[][] float32
	rows int 
	columns int
}
func make_empty_vector(size int)  vector_t{
    var v vector_t
    v.elem = make([]float32, size)
    v.size = size
    return v
}

func make_empty_matrix(rows int, columns int)  matrix_t{
    var m matrix_t
    m.elem = make([][]float32, rows)
    m.rows = rows
    m.columns = columns
    for row:= 0; row < m.rows; row++{
        m.elem[row] = make([]float32, columns)
    }
    return m
}

func input_matrix(m *matrix_t)bool{
	var rows, columns int
	var x float32
	var rc error
	fmt.Print("Введите кол-во строк: ")
	_, rc = fmt.Scanf("%d\n", &rows)
	if(rc != nil){
		return false
	}
	fmt.Print("Введите кол-во столбцов: ")
	_, rc = fmt.Scanf("%d\n", &columns)
	if(rc != nil){
		return false
	}
	*m = make_empty_matrix(rows, columns)

	for row:=0;row<rows;row++{
		for column:=0;column<columns;column++{
			_, rc = fmt.Scanf("%f", &x)
			//fmt.Scanf("%f ", &x)
			if(rc != nil){
				return false
			}
			(*m).elem[row][column] = x
		}
	}
	fmt.Print("\n")
	return true
}
func print_matrix(m matrix_t){
	fmt.Print("Матрица [", m.rows, " x ", m.columns, "]:\n")
	for i:=0;i<m.rows;i++{
		for j:=0;j<m.columns;j++{
			fmt.Print(m.elem[i][j], " ")
		}
		fmt.Print("\n")
	}
}

// Ввод строки
func input_string()string{
	var str1 string
    fmt.Scanf("%s\n", &str1)
	return str1
}