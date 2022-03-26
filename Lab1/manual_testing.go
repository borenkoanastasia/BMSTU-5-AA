package main

import "fmt"

func input_start_data(str1 *string, str2 *string){
	fmt.Println("Введите первую строку:");
    *str1 = input_string()
	fmt.Println("Введите вторую строку:");
    *str2 = input_string()
    *str1 = " " + *str1
    *str2 = " " + *str2
}

func manual_testing_test(str1 string, str2 string){
}


func manual_testing(){
	var str1, str2 string
	input_start_data(&str1, &str2)
    var len1 int = len(str1)
    var len2 int = len(str2)
	var LRecursionKesh = getLRecursionKesh()

	fmt.Println("Матрица:");
    var ans1 int = LMatrix_print(&str1, &str2, len1, len2)
    var ans2 int = LRecursion(&str1, &str2, len1 - 1, len2 - 1)
    var ans3 int = LRecursionKesh(&str1, &str2, len1 - 1, len2 - 1)
    var ans4 int = DLRecursion(&str1, &str2, len1 - 1, len2 - 1)

	var res [2]test_result_t
	res[0].times[0] = get_time(str1, str2,LMatrix)
	res[1].times[0] = (float64)(ans1)

	res[0].times[1] = get_time(str1, str2,LRecursion)
	res[1].times[1] = (float64)(ans2)

	res[0].times[2] = get_time(str1, str2,LRecursionKesh)
	res[1].times[2] = (float64)(ans3)

	res[0].times[3] = get_time(str1, str2,DLRecursion)
	res[1].times[3] = (float64)(ans4)

	print_result_manual(res)
}

