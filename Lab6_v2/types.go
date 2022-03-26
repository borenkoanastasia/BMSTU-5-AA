package main 

import "fmt"

// Объявление типов
type int_array_t []int 
type string_array_t []string 
type float_matrix_t [][]float64
type int_matrix_t [][]int

// Создание типов
func make_int_array(size int)int_array_t{
	return make(int_array_t, size)
}
func make_string_array(size int)string_array_t{
	return make(string_array_t, size)
}
func make_float_matrix(rows int, columns int)float_matrix_t{
	var m = make(float_matrix_t, rows)
	for i:=0;i<rows;i++{
		m[i] = make([]float64, columns)
	}
	return m
}


func make_int_matrix(rows int, columns int)int_matrix_t{
	var m = make(int_matrix_t, rows)
	for i:=0;i<rows;i++{
		m[i] = make([]int, columns)
	}
	return m
}

// Работа с типами

// Копирование 
/*func (arr int_array_t)copy_array(num int){
	a
}*/

// Ввод типов

func input_int_matrix(m *int_matrix_t)bool{
	var rows, columns int
	var x int
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
	*m = make_int_matrix(rows, columns)

	for row:=0;row<rows;row++{
		for column:=0;column<columns;column++{
			_, rc = fmt.Scanf("%d", &x)
			//fmt.Scanf("%f ", &x)
			if(rc != nil){
				return false
			}
			(*m)[row][column] = x
		}
	}
	fmt.Print("\n")
	return true
}
func input_float_matrix(m *float_matrix_t)bool{
	var rows, columns int
	var x float64
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
	*m = make_float_matrix(rows, columns)

	for row:=0;row<rows;row++{
		for column:=0;column<columns;column++{
			_, rc = fmt.Scanf("%f", &x)
			//fmt.Scanf("%f ", &x)
			if(rc != nil){
				return false
			}
			(*m)[row][column] = x
		}
	}
	fmt.Print("\n")
	return true
}

// Вывод типов

func (arr int_array_t)print(){
	fmt.Print("Массив [", len(arr), "]: ")
	for i:=0;i<len(arr);i++{
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
func (arr string_array_t)print(){
	/*if len(arr)<1{
		fmt.Println("[WARNING] Empty Array!")
		return
	} */
	fmt.Print("Массив [", len(arr), "]: ")
	for i:=0;i<len(arr);i++{
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func (m float_matrix_t)print(){
	if len(m)<1{
		fmt.Println("[WARNING] Empty Matrix!")
		return
	} 
	fmt.Print("Матрица [", len(m), " x ", len(m[0]), "]:\n")
	for i:=0;i<len(m);i++{
		for j:=0;j<len(m[i]);j++{
			fmt.Print(m[i][j], " ")
		}
		fmt.Print("\n")
	}
}