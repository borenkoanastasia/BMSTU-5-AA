package main

import "time"
import "fmt"
/*
def generate_new_borders(work_res, weight_borders, last_borders):

//...
    cur = weight_borders[0] 
    borders.append(0)
    for i in range(len(work_res)):
        if (work_res[i][1] == STATUS_BUILDER_ACTIVE_WAIT):
            cur += (weight_borders[1]-weight_borders[0]) * max_v/summ
        else:
            cur += (weight_borders[1]-weight_borders[0]) * work_res[i][0]/summ
        borders.append(cur)
    res = []
    print(borders)
    for i in range(len(borders)-1):
        res.append([int(borders[i]),int(borders[i+1])])
    return res

def get_end_cond(bricks_matrix):
    last_line_id = len(bricks_matrix) -1
    con = True
    for i in range(len(bricks_matrix[last_line_id])):
        if((bricks_matrix[last_line_id][i]) == 0):
            con = False 
            return con
    return con
        
def dispetcher():
    global FLAG_END_OF_WORK,BUILDERS_BORDERS
    con = True
    width_borders = [0, W]
    while con:
        sleep(0.15)
        MUTeX_BUILDERS_BORDERS.acquire()
        BUILDERS_BORDERS = generate_new_borders(BUILDERS_WORK_STATUS,width_borders, BUILDERS_BORDERS)
        for i in range(BUILDERS_COUNT):
            BUILDERS_WORK_STATUS[i][0] = 0
        MUTeX_BUILDERS_BORDERS.release()

        if (get_end_cond(BRICKS_MATRIX)):
            FLAG_END_OF_WORK = END_OF_WORK
            con = False
    print("FLAG_END_OF_WORK",FLAG_END_OF_WORK)
        */

func check_end_bricks_matrix(bricks_matrix bricks_matrix_t)bool{
	var condition bool = false
	for i:=0;i<len(bricks_matrix);i++{
		for j:=0;j<len(bricks_matrix[i]);j++{
			if (bricks_matrix[i][j].fixed == false){
				condition = true
				return condition
			}
		}
	}
	return condition
}

func get_max_work_res(work_res []int)int{
	var max int = work_res[0]
	for i:=0;i<len(work_res);i++{
		if (work_res[i] > max){
			max = work_res[i]
		}
	}
	return max
}

func get_sum_work_res(work_res []int, max int) int{
	var summa int
	for i:=0;i<len(work_res);i++{
		summa += work_res[i]
		if (work_res[i] == 0){
			summa +=max
		}
	}
	return summa
}

func generate_new_borders(work_res []int, width int)[][]int{
	var res [][]int = make([][]int, len(work_res))
	var borders []int = make([]int, len(work_res) + 1)

	var max = get_max_work_res(work_res)
	if (max == 0){
		max = 1
	}
	var summa = get_sum_work_res(work_res, max)

	var cur float64 = 0
	borders[0] = 0

	for i:=1;i<len(work_res)+1;i++{
		
		cur += float64(work_res[i-1]) / float64(summa)
		if (work_res[i-1] == 0){
			cur += float64(max)/float64(summa)
		}
	//	fmt.Println(cur, summa, width, max, work_res[i-1])
		borders[i] = int(cur * float64(width))
	}
	for i:=0;i<len(work_res);i++{
		res[i]=make([]int, 2)
		res[i] = []int{borders[i], borders[i+1]}
	}
	//fmt.Println(borders)
	return res
}

func dispetcher(system_parameters *system_parameters_t){
	var con bool = true
	time.Sleep(100*SLEEP_COEF)
	for ;con;{
		time.Sleep(time.Duration(SLEEP_COEF * system_parameters.dispetcher_time))

		builders_borders_mutex.Lock()
		builders_works_status_mutex.Lock()
		var new_borders = generate_new_borders(system_parameters.builders_works_status, system_parameters.width)
		fmt.Println(system_parameters.builders_works_status,"\t", system_parameters.builders_borders, "\t", new_borders)
		system_parameters.builders_borders = new_borders
		for i:=0;i<len(system_parameters.builders_works_status);i++{
			system_parameters.builders_works_status[i] = 0
		}
		builders_works_status_mutex.Unlock()
		builders_borders_mutex.Unlock()
	//	system_parameters.bricks_matrix.print()

		con = check_end_bricks_matrix(system_parameters.bricks_matrix)
		
	}
	//flag_end_of_works_mutex.Lock()
	system_parameters.flag_end_of_works = true
	//flag_end_of_works_mutex.Unlock()
	print("I am dispetcher!\n")
}

