package main 

import "time"
//import "fmt"

func get_request(system_parameters system_parameters_t)string{
	var req = <- system_parameters.request_queue
	return req
}

func get_place_half_way(start int, end int, system_parameters *system_parameters_t, i int, plus bool, x_res*int, y_res*int)bool{
	var condition bool= start < end
	if (!plus){
		buf:= start - 1
		start = end - 1
		end = buf
		condition = start>end
		//print(condition)
	}
	if (!condition)	{
		return false
	}
	for j:= start;condition;{
		var con = true 
		if (i %2 == 0 && i != 0 && j > 0){ //system_parameters.width){
			con = (system_parameters.bricks_matrix[i-1][j].fixed == true && 
				system_parameters.bricks_matrix[i-1][j-1].fixed == true)
		}
		if (i %2 == 0 && i != 0 && j==0){
			con = (system_parameters.bricks_matrix[i-1][j].fixed == true)
		}
		if (i %2 == 1){
			/*if (j == 0){
				con = system_parameters.bricks_matrix[i-1][j].fixed == true && system_parameters.bricks_matrix[i-1][j+1].fixed == true
			}else */
			if (j==system_parameters.width-1){
				con = system_parameters.bricks_matrix[i-1][j].fixed == true
			}else{
				con = (system_parameters.bricks_matrix[i-1][j+1].fixed == true && 
					system_parameters.bricks_matrix[i-1][j].fixed == true)
				}
		}
		if (system_parameters.bricks_matrix[i][j]).fixed == false && con{
			*x_res = i
			*y_res =j
			return true
		}
		if (plus){
			j++
			condition = j < end
		} else {
			j--
			condition = j > end
		}
	}
	return false
}

func get_request_place(system_parameters *system_parameters_t, builder_name int, x_res*int,  y_res*int)bool{
	var borders = system_parameters.builders_borders[builder_name]
	var middle = borders[0]+borders[1]
	middle /=2

	for i:=0;i<system_parameters.height;i++{
		var ans = get_place_half_way(middle, borders[1], system_parameters, i, true, x_res, y_res)
		if (ans){
			return ans
		}
		ans = get_place_half_way(borders[0], middle, system_parameters, i, false, x_res, y_res)
		if (ans){
			return ans
		}
	}
	return false
}


func lock_builder_system_semaphore(builder_name int){
	//SYSTEM_MUTEX.Lock()
	SYSTEM_SEMAFORE[builder_name].Lock()    // к системе обратились
	//SYSTEM_MUTEX.Unlock()
}
func unlock_builder_system_semaphore(builder_name int){
	//SYSTEM_MUTEX.Lock()
	SYSTEM_SEMAFORE[builder_name].Unlock()    // систему отпустили
	//SYSTEM_MUTEX.Unlock()
}

func builder(system_parameters *system_parameters_t, builder_name int){
	print("BUILDER START\n")
	var x_res, y_res int
	var req string
	var con bool
	var time_to_sleep time.Duration

	ec := !system_parameters.flag_end_of_works
	for ;ec;{
		req = get_request(*system_parameters)
 		con = false
		for ;!con;{
			BIG_MUTEX.Lock()
			con = get_request_place(system_parameters, builder_name, &x_res, &y_res)
			if (con){
				system_parameters.bricks_matrix[x_res][y_res].fixed = true
				system_parameters.bricks_matrix[x_res][y_res].builder_name = builder_name
				system_parameters.builders_works_status[builder_name] +=1
				system_parameters.builders_wait_status[builder_name] = false
			} else {
				system_parameters.builders_wait_status[builder_name] = true
			}
			BIG_MUTEX.Unlock()

			if (con){
				time_to_sleep = time.Duration(len(req)*SLEEP_COEF*system_parameters.builder_time)/1000
				if (time_to_sleep > 1000*1000*1000*2) {
					time_to_sleep = 1000*1000*1000*2
				}
				time.Sleep(time_to_sleep)
				if (builder_name == 1){
					time.Sleep(time_to_sleep)
				}
			}
			BIG_MUTEX.Lock()
			draw_brick()
			BIG_MUTEX.Unlock()
			end_cond := system_parameters.flag_end_of_works
			if end_cond{
				break
			}
		}
		ec = !system_parameters.flag_end_of_works
	}
}