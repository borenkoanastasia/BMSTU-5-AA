package main

import "time"
//import "fmt"


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

func get_sum_work_res(work_res []int, max int, builders_wait_status []bool) int{
	var summa int
	for i:=0;i<len(work_res);i++{
		if (builders_wait_status[i]){
			summa +=max
		} else {
			summa += work_res[i]
		}
	}
	return summa
}

func generate_new_borders(work_res []int, builders_wait_status []bool, width int)[][]int{
	//print("start", "\n")
	var res [][]int = DISPETCHER_RES_BUF//make([][]int, len(work_res))
	//print("good allock 1", "\n")
	var borders []int = make([]int, len(work_res) + 1) //DISPETCHER_BORDERS_BUF
	//unlock_dispetcher_system_semaphore()

	//print("good allock 2", "\n")
	var max = get_max_work_res(work_res)
	if (max == 0){
		max = 1
	}
	//print("get max work res SUCCESS", "\n")
	var summa = get_sum_work_res(work_res, max, builders_wait_status)
	if (summa == 0){
		//lock_dispetcher_system_semaphore()
		var resss = make_builders_borders(len(work_res),width)
		//unlock_dispetcher_system_semaphore()
		return resss
	}
	//print("get sum work res SUCCESS", summa, "\n")

	var cur float64 = 0
	borders[0] = 0

	for i:=1;i<len(work_res)+1;i++{
		
		if (builders_wait_status[i-1]){
			cur += float64(max)/float64(summa)
		} else {
			cur += float64(work_res[i-1]) / float64(summa)
		}
		if (int(cur) < 0){
			//var s int
			//fmt.Println(cur, int(cur), summa, width)
			*END_OF_WORK = true
			//fmt.Scan(s)
		}
		//fmt.Println(cur, summa, width, max, work_res[i-1], builders_wait_status)
		borders[i] = int(cur * float64(width))
	}
	borders[len(borders)-1] = width
	//print("get borders SUCCESS", "\n")
	//fmt.Print(borders, summa, max, work_res, "\n")
	for i:=0;i<len(work_res);i++{
		//print(len(work_res), " ", i, "\n")
		//lock_dispetcher_system_semaphore()
		//res[i] = make([]int, 2)
		//unlock_dispetcher_system_semaphore()
		//print("good memory allock\n")
		res[i] = []int{borders[i], borders[i+1]}
		//print(len(work_res), i, res[i], "\n")
	}
	//print("get res SUCCESS", "\n")
	//print(borders, summa, max, work_res, "\n")
	return res
}
func check_nessosory(work_res []int)bool{
	for i:=0;i<len(work_res);i++{
		if (work_res[i] != 0){
			return true
		}
	}
	return false
}



func lock_dispetcher_system_semaphore(){
	for i:=0;i < len(SYSTEM_SEMAFORE);i++{
		SYSTEM_SEMAFORE[i].Lock()    // к системе обратились
	}
}
func unlock_dispetcher_system_semaphore(){
	for i:=0;i < len(SYSTEM_SEMAFORE);i++{
		SYSTEM_SEMAFORE[i].Unlock()    // систему отпустили
	}
}

func dispetcher(system_parameters *system_parameters_t){
	print("DISPETCHER START\n")

	var con bool = true
	var new_borders[][]int
	for ;con;{

		var time_to_sleep = time.Duration(SLEEP_COEF * system_parameters.dispetcher_time)///1000

		time.Sleep(time_to_sleep)

		BIG_MUTEX.Lock()
		new_borders = generate_new_borders(system_parameters.builders_works_status,system_parameters.builders_wait_status,
			system_parameters.width)
		system_parameters.builders_borders = new_borders
		for i:=0;i<len(system_parameters.builders_works_status);i++{
			system_parameters.builders_works_status[i] = 0
		}
		BIG_MUTEX.Unlock()

		con = check_end_bricks_matrix(system_parameters.bricks_matrix)
		con = con && !system_parameters.flag_end_of_works
	}
	system_parameters.flag_end_of_works = true
}

