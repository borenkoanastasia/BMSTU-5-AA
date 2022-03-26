package main

//import "fmt"

func sort_bubble(array vector_t){
	for i:=1;i< array.size;i++{
		for j:=0;j < array.size - i;j++{
			if (array.elem[j + 1] < array.elem[j]){
				array.elem[j + 1], array.elem[j] = array.elem[j], array.elem[j + 1]
			}
		}
	}
}


func sort_insert(array vector_t){
	for i:=0;i<array.size;i++{
		for j:= i;j > 0 && array.elem[j - 1] > array.elem[j];j--{
			array.elem[j], array.elem[j - 1] = array.elem[j - 1], array.elem[j]
		}
	}
}

func sort_choice(array vector_t){
	for i:=0;i<array.size;i++{
		var min_i int = i

		for j:=i+1;j< array.size;j++{
			if(array.elem[j] <= array.elem[min_i]){
				min_i = j
			}
		}

		if(min_i !=i){
			array.elem[i], array.elem[min_i] = array.elem[min_i], array.elem[i]
		}
	}
}
