package main

import "fmt"
import "math/rand"
import "time"

//const MAX_QUEUE_LEN = 100000
const CUTE_LINE_LEN = 40



const (
	LATTINICA = "abcdefghijklmnopqrstuvwxyz"
	KIRILLICA = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
)
type bricks_matrix_t [][]brick_t 

func (brick_matrix bricks_matrix_t)print(){
	print_cute_line(CUTE_LINE_LEN)
	fmt.Print("bricks_matrix")
	print_cute_line(CUTE_LINE_LEN)
	fmt.Println()
	for i:=0;i<len(brick_matrix[0]);i++{
		for j:=0;j<len(brick_matrix);j++{
			fmt.Print(brick_matrix[j][i].fixed, "  ")
		}
		fmt.Println()
	}
	print_cute_line(CUTE_LINE_LEN + len("bricks_matrix") + CUTE_LINE_LEN)
	fmt.Println()

}

type len_limits_t struct{
	min int 
	max int
}

type system_parameters_t struct{
	width int 
	height int 
	builders_count int
	len_limits len_limits_t

	flag_end_of_works bool

	bricks_matrix bricks_matrix_t
	builders_borders [][]int  
	builders_works_status []int

	request_queue chan string

	dispetcher_time int 
	builder_time int
}

type int_coordinates_t struct{
	x int 
	y int
}

type float_coordinates_t struct{
	x float64
	y float64
}
type brick_t struct{
	fixed bool 
	builder_name int
}

func make_random_string(lim len_limits_t)string{
	rand.Seed(time.Now().UnixNano())
	var string_len = lim.min  + rand.Intn(lim.max)
	var cur_str string = ""
	for i:=0;i<string_len;i++{
		var rand_symbol_index = rand.Intn(len(LATTINICA))
		cur_str = cur_str + string(LATTINICA[rand_symbol_index])
	}
	return cur_str
}

func make_start_queue(queue_len int, lim len_limits_t)chan string{
	var start_queue chan string = make(chan string, queue_len)
	for i:=0;i<queue_len;i++{
		start_queue <- make_random_string(lim)
	}
	close(start_queue)
	return start_queue
}

func make_builders_weight(builders_count int, weight int)[]int{
	var builders_weight = make([]int, builders_count + 1)
    var b_w int = weight/builders_count
	var flag bool = false
    if (b_w*builders_count != weight){
		b_w +=1
		flag = true
	}
	var cur_weight = 0
	builders_weight[0] = cur_weight
	for i:=1;i<builders_count+1;i++{
		cur_weight += b_w 
		//print()
		if (flag && (b_w-1)*(builders_count-i) == weight - cur_weight){
			b_w -= 1
			flag = false
		}
		builders_weight[i] = cur_weight
	}
	return builders_weight
}

func make_builders_borders(builders_count int, weight int)[][]int{
    var builders_borders [][]int = make([][]int, builders_count)
	//builders_borders
	var builders_weight = make_builders_weight(builders_count, weight)
	
	for i:=0;i<len(builders_weight)-1;i++{
		builders_borders[i] = make([]int, 2)
		builders_borders[i][0] = builders_weight[i]
		builders_borders[i][1] = builders_weight[i+1]
	}
	return builders_borders
}

func print_cute_line(l int){
	for i:=0;i < l;i++{
		fmt.Print("-")
	}
}

func make_bricks_matrix(width int, height int)bricks_matrix_t{
	var bricks_matrix bricks_matrix_t = make(bricks_matrix_t, height)
	for i:=0;i<height;i++{
		bricks_matrix[i] = make([]brick_t, width)
		for j:=0;j<width;j++{
			bricks_matrix[i][j] = brick_t{false,0}
		}
	}
	return bricks_matrix
}

func init_system_parameters(w int, h int, bc int, len_lim len_limits_t, dispetcher_time int, builder_time int)system_parameters_t{
	var system_parameters system_parameters_t

	system_parameters.width = w 
	system_parameters.height = h 
	system_parameters.builders_count = bc
	system_parameters.len_limits = len_lim

	system_parameters.flag_end_of_works = false

	system_parameters.bricks_matrix = make_bricks_matrix(w, h) 
	system_parameters.builders_borders = make_builders_borders(bc, w) 
	system_parameters.builders_works_status = make([]int, bc)
	system_parameters.request_queue = make_start_queue(w*h, len_lim)

	system_parameters.dispetcher_time = dispetcher_time
	system_parameters.builder_time = builder_time


	return system_parameters
}

func (sp system_parameters_t)print(){
	fmt.Println("W x H \t\t\t", sp.width, "x", sp.height)
	fmt.Println("FLAG end works \t\t", sp.flag_end_of_works)
	fmt.Println("BUILDERS COUNT \t\t", sp.builders_count)
	fmt.Println("String Len Limits:\t", sp.len_limits)
	fmt.Println("Quene_start \t\t", sp.request_queue)

	sp.bricks_matrix.print()
} 
