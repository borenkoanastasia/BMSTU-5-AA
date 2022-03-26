package main 

import "fmt"
import "bufio"
import "os"

func print_error_message(){
	fmt.Println("Введенные матрицы не корректны. Выход в меню.")
}

func input_two_matrixes(m1 *matrix_t, m2 *matrix_t)bool {
    stdin := bufio.NewReader(os.Stdin)
	fmt.Println("Введите матрицу 1:")
	var rc bool = input_matrix(m1)
	if (!rc){

        stdin.ReadString('\n')
		return rc
	}
	fmt.Println("Введите матрицу 2:")
	rc = input_matrix(m2)

	stdin.ReadString('\n')
	return rc
}

func manual_testing(){
	var m1, m2 matrix_t
	var rc bool
	rc = input_two_matrixes(&m1, &m2)
	if !(rc){
		print_error_message()
		return
	}
	var m3_1 matrix_t = multiplicate_matrix_norm(m1, m2, &rc)
	if !(rc){
		print_error_message()
		return
	}
	var m3_2 matrix_t = multiplicate_matrix_vinograd(m1, m2, &rc)

	//fmt.Println("Результат работы обычного алгоритма умножения матриц:")
	//print_matrix(m3_1)
	//fmt.Println("Результат работы алгоритма Винограда умножения матриц:")
	//print_matrix(m3_2)
	//fmt.Println("Виноград удался")
	var m3_3 matrix_t = multiplicate_matrix_fast_vinograd(m1, m2, &rc)

	fmt.Println("Результат работы обычного алгоритма умножения матриц:")
	print_matrix(m3_1)
	fmt.Println("Результат работы алгоритма Винограда умножения матриц:")
	print_matrix(m3_2)
	fmt.Println("Результат работы оптимизированного алгоритма Винограда умножения матриц:")
	print_matrix(m3_3)
}