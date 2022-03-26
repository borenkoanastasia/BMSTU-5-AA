package main 

/*
#include <time.h>
*/
import "C"
import "math/rand"
import "fmt"

func line(length int) string{
	var str string = ""
	for i:=0;i<length;i++{
		str = str + "_"
	}
	return str
}
func title(length int) string{
	var str string = ""
	for i:=0;i<(length - 4)/2;i++{
		str = str + " "
	}
	str = str + "MENU"
	for i:=0;i<(length - 4)/2;i++{
		str = str + " "
	}
	return str
}
type vector_t struct{
	elem[] float32
	size int
}
func print_vector(v vector_t){
	fmt.Print("Вектор [", v.size, "]: ")
	for i:=0;i<v.size;i++{
		fmt.Print(v.elem[i], " ")
	}
	fmt.Println()
}

func make_empty_vector(size int)  vector_t{
    var v vector_t
    v.elem = make([]float32, size)
    v.size = size
    return v
}


func sort_bubble_test(array vector_t){
	var change int = 0
	var buf float32
	var crack = array.size
	var i int
	var j int
		for i=1;i< crack;i++{
			for j=0;j < crack - i;j++{
				if (array.elem[j + 1] < array.elem[j]){
					buf = array.elem[j + 1]
					array.elem[j+1] = array.elem[j] 
					array.elem[j] = buf
				    change++;
				}
				//fmt.Println(C.clock())
			}
		}
    //print_vector(array)
	//fmt.Println("Перестановки", change)
	//fmt.Println("Время", t2 - t1, "\n")
}

func sort_bubble(array vector_t){
	var change int = 0
	var buf float32
	var crack = array.size
	var i int
	var j int
	t1:=C.clock()
		for i=1;i< crack;i++{
			for j=0;j < crack - i;j++{
				if (array.elem[j + 1] < array.elem[j]){
					buf = array.elem[j + 1]
					array.elem[j+1] = array.elem[j] 
					array.elem[j] = buf
				    change++;
				}
				//fmt.Println(C.clock())
			}
		}
	t2:=C.clock()
    //print_vector(array)
	fmt.Println("Перестановки", change)
	fmt.Println("Время", t2 - t1, "\n")
}
	
func RandVector(size int) vector_t {
    var b vector_t = make_empty_vector(size)
    for i := 0; i < b.size;i++ {
        b.elem[i] = float32(rand.Intn(100))
    }
    return b
}
func SortVector(size int) vector_t {
    var b vector_t = make_empty_vector(size)
    for i := 0; i < b.size;i++ {
        b.elem[i] = float32(i+1)
    }
    return b
}
func ReverseVector(size int) vector_t {
    var b vector_t = make_empty_vector(size)
    for i := 0; i < b.size;i++ {
        b.elem[i] = float32(b.size - i)
    }
    return b
}

func manual_testing(){
    var N int = 10000
	var v1 vector_t = make_empty_vector(N)
	var v2 vector_t = make_empty_vector(N)
	var v3 vector_t = make_empty_vector(N)

	v1 = SortVector(N)
	v2 = ReverseVector(N)
	v3 = RandVector(N)

	fmt.Println()
	fmt.Printf("Векторы длины %d:\n", N)

	sort_bubble_test(v1)
	sort_bubble_test(v1)
	sort_bubble_test(v1)
	sort_bubble_test(v1)
	sort_bubble_test(v1)
	sort_bubble_test(v1)
	//print_vector(v1)
	//print_vector(v2)
	//print_vector(v3)

	fmt.Println("Сортируем:\n\n")

	fmt.Println("Реверс тест:")
	sort_bubble(v2)
	
	fmt.Println("Случайный тест:")
	sort_bubble(v3)


	fmt.Println("Сорт тест:")
	sort_bubble(v1)

	fmt.Println()
	fmt.Println("Результат сортировки пузырьком")
	//print_vector(v1)
	//print_vector(v2)
	//print_vector(v3)
	fmt.Println()

}
// Ввод строки
func input_string()string{
	var str1 string
    fmt.Scanf("%s\n", &str1)
	return str1
}
func main(){var length int = 111
	var choise string
    for ;true;{
		fmt.Println(line(length))
		fmt.Println()
		fmt.Println(title(length))
		fmt.Println(line(length))

		fmt.Println("\t1.Testing")
		fmt.Println("\tAnother choice is exit")

		choise = input_string()

		fmt.Println(line(length))

		if	(choise == "1"){
			manual_testing()
		} 
    }
}
