package main

import "fmt"
import "bufio"
import "os"

type vector_t struct{
	elem[] float32
	size int
}

func make_empty_vector(size int)  vector_t{
    var v vector_t
    v.elem = make([]float32, size)
    v.size = size
    return v
}

func input_vector_elems(v *vector_t)bool{
	var x float32
	var rc error
	for i:=0;i<v.size;i++{
			_, rc = fmt.Scanf("%f", &x)
			//fmt.Scanf("%f ", &x)
			if(rc != nil){
				return false
			}
			(*v).elem[i] = x
	}
	return true
}

func vector_copy(v vector_t)vector_t{
	var v_res vector_t = make_empty_vector(v.size)
	for i := 0; i < v.size; i++{
		v_res.elem[i] = v.elem[i]
	}
	return v_res
}

func input_vector(v *vector_t)bool{
	var size, con int
	var rc error
	var r bool

	fmt.Print("Введите кол-во элементов: ")
	_, rc = fmt.Scanf("%d\n", &size)
	if(rc != nil){
		return false
	}
	
	fmt.Print("Сгенерировать автоматически (1 - да, иначе - нет): ")
	_, rc = fmt.Scanf("%d\n", &con)
	if(rc != nil || con != 1){
		if (rc != nil){
			stdin := bufio.NewReader(os.Stdin)
			stdin.ReadString('\n')
		}
		*v = make_empty_vector(size)
		r = input_vector_elems(v)
		return r
	}
	*v = RandVector(size)

	fmt.Print("\n")
	return true
}
func print_vector(v vector_t){
	fmt.Print("Вектор [", v.size, "]: ")
	for i:=0;i<v.size;i++{
		fmt.Print(v.elem[i], " ")
	}
	fmt.Println()
}

// Ввод строки
func input_string()string{
	var str1 string
    fmt.Scanf("%s\n", &str1)
	return str1
}
