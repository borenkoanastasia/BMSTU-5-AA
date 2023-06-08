package main

import "time"
import "os"
//import "fmt"
import ("github.com/wcharczuk/go-chart" 
//exposes "chart"
)
const STEPS_COUNT = 3

var CESAR_CODE = get_string_encryption_Cesar(13)
var XOR_CODE = get_string_encryption_Xor(generate_encryption_Xor_key(100))
var ATBASH_CODE = string_encryption_Atbash

type procFunc func(string) string

func proc(input <-chan request_t, procID int, f procFunc) <-chan request_t {
	output := make(chan request_t)

	go func() {
		defer close(output)
		for arg := range input {
			arg.steps[procID].start = time.Now()

			//time.Sleep(time.Duration(100) * time.Microsecond)
			//fmt.Println(arg.id, procID, time.Now())
			arg.str = f(arg.str)
			arg.steps[procID].end = time.Now()
			output <- arg //value

	        time.Sleep(10*time.Microsecond)
		}
	}()
	
	return output
}

func setup(n int) <-chan request_t {
	channel := make(chan request_t)

	var req []request_t = make([]request_t, n)

	for i:=0;i<n;i++{
		req[i] = make_request(i)
	}

	go func(req []request_t) {
		defer close(channel)
		for i := 0; i < n; i++ {
			//req := make_request(i)
			req[i].start_wait = time.Now()
			channel <- req[i]
		}
	}(req)
	return channel
}


func conveyorRun(count int, channel <- chan request_t) request_array_t {
    //var mutex sync.Mutex

	var steps[STEPS_COUNT]func(string)string 
	steps[0] = CESAR_CODE
	steps[1] = XOR_CODE
	steps[2] = ATBASH_CODE


	for i:=0;i<STEPS_COUNT;i++{
		channel = proc(channel, i, steps[i])
	}


	var res_arr_request request_array_t = make_empty_request_array(count)

	var i int = 0
	for res := range channel {
		res_arr_request.elems[i] = res
		i++
	}

	return res_arr_request
}


func lenear(input <-chan request_t){
	var time_start, time_end time.Time
	var sum, count int64
	sum = 0
	count = 0
	for req := range input{
		req.str = CESAR_CODE(req.str)
		req.str = ATBASH_CODE(req.str)
		req.str = XOR_CODE(req.str)
		sum = sum + int64(time_end.Sub(time_start))
		count++
		//fmt.Println(count, time.Now())
		time.Sleep(30*time.Microsecond)
	}
}

func testing(n int, leantime *time.Duration, convtime *time.Duration){
	channel1 :=setup(n)
	channel2 :=setup(n)

	time.Sleep(10*time.Millisecond)

	time1 := time.Now()
	lenear(channel1)
	time2 := time.Now()
	time3 := time.Now()
	conveyorRun(n, channel2)
	time4 := time.Now()

	*leantime = time2.Sub(time1)
	*convtime = time4.Sub(time3)
}

func print_comp_graph(){    
	const REPEATS = 1003
	var time1[] float64 =make([]float64, REPEATS/50+1)
	var time2[] float64 =make([]float64, REPEATS/50+1)
	var t1, t2 time.Duration
	var x[] float64 = make([]float64, REPEATS/50+1)

	for i:=1; i < REPEATS/2; i+=50{
		testing(i, &t1, &t2)
    }
	for i:=1; i < REPEATS; i+=50{
		x[i/50] = float64(i)
		for j:=0;j<10;j++{
		    testing(i, &t1, &t2)
		    time1[i/50] += float64(t1)
		    time2[i/50] += float64(t2)
		}
		time1[i/50] /= 10
		time2[i/50] /= 10
    }


	graph1 := chart.Chart{
	Background: chart.Style{
	Padding: chart.Box{
		Top:  20,
			Left: 20,
			},
		},
		Series: []chart.Series{/*
			chart.ContinuousSeries{
				Name:    "Очередь 1",
				XValues: x,//[]float64{1.0, 2.0, 3.0, 4.0},
				YValues: time_steps[0],//[]float64{1.0, 2.0, 3.0, 4.0},
			},*/
			chart.ContinuousSeries{
				Name:    "Linear",
				XValues: x,//[]float64{1.0, 2.0, 3.0, 4.0},
				YValues: time1,//[]float64{1.0, 2.0, 3.0, 4.0},
			},
			chart.ContinuousSeries{
				Name:    "Conveer",
				XValues: x,//[]float64{1.0, 2.0, 3.0, 4.0},
				YValues: time2,//[]float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}

	graph1.Elements = []chart.Renderable{
		chart.Legend(&graph1),
	}

	f, _ := os.Create("output2.png")
	defer f.Close()
	graph1.Render(chart.PNG, f)
	    
}