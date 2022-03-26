package main

import "math/rand"
import "time"

func main(){

	rand.Seed(time.Now().UnixNano())
	//var m = make_start_queue(10, len_limits_t{10, 10})
	//var m = make_random_string(len_limits_t{10, 10})
	/*var m = init_system_parameters(50, 80, 4, len_limits_t{8, 20}, 100, 1)
	
	//fmt.Println(m)

	m.print()

	director(m)*/
	app_start()

}