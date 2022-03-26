package main 

import "sync"
import "time"
//import "fmt"
//import "math/rand"

const SLEEP_COEF = 1000000

//
//var RAND = rand.New(rand.NewSource(10))

var bricks_matrix_mutex sync.Mutex
var builders_borders_mutex sync.Mutex
var builders_works_status_mutex sync.Mutex 
var flag_end_of_works_mutex sync.Mutex

func director(system_parameters system_parameters_t){

	var app_ch = app_start(&system_parameters)	
	time.Sleep(10)

	//(*BRICKS_MATRIX)[0][0].fixed = true
	//(*BRICKS_MATRIX)[0][1].fixed = true
	//(*BRICKS_MATRIX)[1][0].fixed = true
	//(*BRICKS_MATRIX)[5][5].fixed = true

	//(*BRICKS_MATRIX).print()

	//draw_brick()

	for i:=0;i<system_parameters.builders_count;i++{
		go builder(&system_parameters, i)
	}
	go dispetcher(&system_parameters)


	for ;!system_parameters.flag_end_of_works;{}
/*	for i := range system_parameters.request_queue{
		fmt.Println(i)
	}*/

	var f bool = <- app_ch
	if (f){}


	print("I am director! I AM POWER!!!\n")
}


