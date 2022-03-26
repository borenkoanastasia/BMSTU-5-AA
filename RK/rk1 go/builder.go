package main 

import "time"
//import "fmt"

func get_request(system_parameters system_parameters_t)string{
	var req = <- system_parameters.request_queue
	return req
}

/*
            if (j %2 == 0 and j != 0 and i+1 < W):
                con = (BRICKS_MATRIX[j-1][i] == 1 and BRICKS_MATRIX[j-1][i+1] == 1)
            if (j %2 == 0 and j != 0 and i+1 == W):
                con = (BRICKS_MATRIX[j-1][i] == 1)
            if (j %2 == 1):
                if (i == 0):
                    con = BRICKS_MATRIX[j-1][i] == 1
                elif (i==W):
                    con = BRICKS_MATRIX[j-1][i-1] == 1
                else:
                    con = (BRICKS_MATRIX[j-1][i-1] == 1 and BRICKS_MATRIX[j-1][i] == 1)
            if (BRICKS_MATRIX[j][i]) == 0 and con:
                #MUTeX_BRICKS_MATRIX.acquire()
                #BRICKS_MATRIX[j][i] = 2
                #MUTeX_BRICKS_MATRIX.release()
                return i, j, True
*/

func get_place_half_way(start int, end int, system_parameters *system_parameters_t, i int, plus bool, x_res*int, y_res*int)bool{
	var condition bool= start < end
	if (!plus){
		buf:= start - 1
		start = end - 1
		end = buf
		condition = start>end
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
			if (j == 0){
				con = system_parameters.bricks_matrix[i-1][j].fixed == true && system_parameters.bricks_matrix[i-1][j+1].fixed == true
			}else if (j==system_parameters.width-1){
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
			//print(j, "end", end, "start", start, "plus", plus, "con", con, "j<end", j<end, "\n")
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
		/*for j:=middle;j<borders[1];j++{
            var con = true 
            if (i %2 == 0 && i != 0 && j > 0){ //system_parameters.width){
                con = (system_parameters.bricks_matrix[i-1][j].fixed == true && 
					system_parameters.bricks_matrix[i-1][j-1].fixed == true)
			}
            if (i %2 == 0 && i != 0 && j==0){
                con = (system_parameters.bricks_matrix[i-1][j].fixed == true)
			}
            if (i %2 == 1){
                if (j == 0){
                    con = system_parameters.bricks_matrix[i-1][j].fixed == true && system_parameters.bricks_matrix[i-1][j+1].fixed == true
				}else if (j==system_parameters.width-1){
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
		}
		for j:=middle-1;j>borders[0]-1;j--{
            var con = true 
            if (i %2 == 0 && i != 0 && j > 0){ //system_parameters.width){
                con = (system_parameters.bricks_matrix[i-1][j].fixed == true && 
					system_parameters.bricks_matrix[i-1][j-1].fixed == true)
				//print("con", con, "i", i, "j", j, "i-1j", system_parameters.bricks_matrix[i-1][j].fixed, "i-1j+1", system_parameters.bricks_matrix[i-1][j+1].fixed, "\n")
			}
            if (i %2 == 0 && i != 0 && j==0){
                con = (system_parameters.bricks_matrix[i-1][j].fixed == true)
			}
            if (i %2 == 1){
                if (j == 0){
                    con = system_parameters.bricks_matrix[i-1][j].fixed == true && system_parameters.bricks_matrix[i-1][j+1].fixed == true
				}else if (j==system_parameters.width-1){
                    con = system_parameters.bricks_matrix[i-1][j].fixed == true
				}else{
					con = (system_parameters.bricks_matrix[i-1][j+1].fixed == true && 
						system_parameters.bricks_matrix[i-1][j].fixed == true)
					}
			}
            if (system_parameters.bricks_matrix[i][j]).fixed == false && con{
				*x_res = i
				*y_res = j
                return true
			}
		}
		*/
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


func builder(system_parameters *system_parameters_t, builder_name int){
	for ;!system_parameters.flag_end_of_works;{
		var req = get_request(*system_parameters)
 		var con = true
		for ;con;{
			var x_res, y_res int
			bricks_matrix_mutex.Lock()
			builders_borders_mutex.Lock()
			con = get_request_place(system_parameters, builder_name, &x_res, &y_res)
			if (con){
				system_parameters.bricks_matrix[x_res][y_res].fixed = true
				system_parameters.bricks_matrix[x_res][y_res].builder_name = builder_name
				builders_works_status_mutex.Lock()
				system_parameters.builders_works_status[builder_name] +=1
				builders_works_status_mutex.Unlock()
			}
			builders_borders_mutex.Unlock()
			bricks_matrix_mutex.Unlock()
			if (con){
				time.Sleep(time.Duration(len(req)*SLEEP_COEF*system_parameters.builder_time))
				if (builder_name == 1){
					time.Sleep(time.Duration(len(req)*SLEEP_COEF*system_parameters.builder_time))
				}
			}
			draw_brick()
		}
	}
	print("I am builder!\n")
}