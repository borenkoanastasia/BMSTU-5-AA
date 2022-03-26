package main

import "fmt"
import "bufio"
import "os"

func print_error_message(){
	fmt.Println("Введенный вектор не корректен. Выход в меню.")
}

func manual_testing(){
	var v_src vector_t
	var rc = input_vector(&v_src)
	if (!rc){
		print_error_message()
		stdin := bufio.NewReader(os.Stdin)
		stdin.ReadString('\n')
		return
	}	
	fmt.Println()
	fmt.Println("Введенный вектор:")
	print_vector(v_src)

	var v1 vector_t = vector_copy(v_src)
	var v2 vector_t = vector_copy(v_src)
	var v3 vector_t = vector_copy(v_src)

	sort_bubble(v1)
	sort_insert(v2)
	sort_choice(v3)

	fmt.Println()
	fmt.Println("Результат сортировки пузырьком")
	print_vector(v1)
	fmt.Println()
	fmt.Println("Результат сортировки вставками")
	print_vector(v2)
	fmt.Println()
	fmt.Println("Результат сортировки выбором")
	print_vector(v3)
	fmt.Println()

}