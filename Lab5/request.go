package main

import ("github.com/wcharczuk/go-chart" 
//exposes "chart"
)
import "time"
import "fmt"
import "math/rand"
import "os"


/*
func (pt processing_time)set_start(start time.Time){
	pt.start = start
}
func (pt processing_time)set_end(end time.Time){
	pt.end = end
}
func (pt processing_time)get_start()time.Time{
	return pt.start
}
func (pt processing_time)get_end()time.Time{
	return pt.end
}
func (pt processing_time)get_time()time.Duration{
	return pt.end.Sub(pt.start)
}*/


type processing_time struct {
	start time.Time
	end time.Time
}
type request_t struct {
	id int
	str string
	start_wait time.Time
	steps[STEPS_COUNT] processing_time
}

func make_request(id int) request_t {
	var req request_t
	var len int = rand.Intn(50) + 50
	req.id = id
	req.str = generate_encryption_Xor_key(len)
	return req
}

type request_array_t struct {
	elems[] request_t
	size int
}

func make_empty_request_array(size int)  request_array_t{
    var v request_array_t
    v.elems = make([]request_t, size)
    v.size = size
    return v
}
/*
func input_vector_elems(v *request_array_t)bool{
	var x float32
	var rc error
	for i:=0;i<v.size;i++{
			_, rc = fmt.Scanf("%f", &x)
			//fmt.Scanf("%f ", &x)
			if(rc != nil){
				return false
			}
			(*v).elems[i]. = x
	}
	return true
}*/

func request_array_copy(v request_array_t)request_array_t{
	var v_res request_array_t = make_empty_request_array(v.size)
	for i := 0; i < v.size; i++{
		v_res.elems[i] = v.elems[i]
	}
	return v_res
}
/*
func input_vector(v *request_array_t)bool{
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
}*/

func (v request_array_t)print_steps_time(){
	
	var min[STEPS_COUNT] int64
	var max[STEPS_COUNT] int64
	var cur[STEPS_COUNT] int64
	var sum[STEPS_COUNT] int64
	var av[STEPS_COUNT] int64


	for i:=0; i < STEPS_COUNT;i++{
		min[i] = int64(v.elems[i].steps[i].start.Sub(v.elems[i].start_wait))
	}

	for i:=0; i < v.size; i++{
		for j:=0;j < STEPS_COUNT;j++{
	    	cur[j] = int64(v.elems[i].steps[j].end.Sub(v.elems[i].steps[j].start))
			
			if (cur[j] < min[j]){
				min[j] = cur[j]
			}
			if (cur[j] > max[j]){
				max[j] = cur[j]
			}
			sum[j] += cur[j]
		}
	}

    fmt.Printf("%8s|%8s|%8s|%8s\n", "Этап", "min", "max", "av")
	for i:= 0; i < STEPS_COUNT; i++{
		av[i] = sum[i]/int64(v.size)
		fmt.Printf("%8d|%8d|%8d|%8d\n", i+1, min[i], max[i], av[i])
	}
}

func (v request_array_t)print_quenue_time(){
	
	var min[STEPS_COUNT-1] int64
	var max[STEPS_COUNT-1] int64
	var cur[STEPS_COUNT-1] int64
	var sum[STEPS_COUNT-1] int64
	var av[STEPS_COUNT-1] int64


	for i:=0; i < STEPS_COUNT-1;i++{
		min[i] = int64(v.elems[i].steps[i].start.Sub(v.elems[i].start_wait))
	}

	for i:=0; i < v.size; i++{
		for j:=0;j < STEPS_COUNT-1;j++{
			/*if (j == 0){
				cur[j] = int64(v.elems[i].steps[0].start.Sub(v.elems[i].start_wait))
			} else{*/
				cur[j] = int64(v.elems[i].steps[j+1].start.Sub(v.elems[i].steps[j].end))
			//}
			
			if (cur[j] < min[j]){
				min[j] = cur[j]
			}
			if (cur[j] > max[j]){
				max[j] = cur[j]
			}
			sum[j] += cur[j]
		}
	}

    fmt.Printf("%8s|%8s|%8s|%8s\n", "Очередь", "min", "max", "av")
	for i:= 0; i < STEPS_COUNT-1; i++{
		av[i] = sum[i]/int64(v.size)
		fmt.Printf("%8d|%8d|%8d|%8d\n", i+1, min[i], max[i], av[i])
	}
}

func (v request_array_t)print_system_time(){
	var system_time_min int64 = -1
	var system_time_max int64 = 0
	var system_time_sum int64 = 0

	for i:=0; i < v.size; i++{
		var cur_time =  int64(v.elems[i].steps[2].end.Sub(v.elems[i].steps[0].start))
		if (system_time_min > cur_time || system_time_min == -1){
			system_time_min = cur_time
		}
		if (system_time_max < cur_time){
			system_time_max = cur_time
		}
		system_time_sum += cur_time
	}
	fmt.Println()
	fmt.Printf("%8s|%8s|%8s|%8s\n", "Система", "min", "max", "av")
	fmt.Printf("%8s|%8d|%8d|%8d\n", " ", system_time_min, system_time_max, system_time_sum/int64(v.size))
}

func (v request_array_t)print_request_log(){    
    fmt.Println("Время в nsec")
	v.print_quenue_time()
	v.print_steps_time()
	v.print_system_time()
}

func print_title(){
	fmt.Printf("%3s | %9s |", "id", "go conv")
	for i:=0;i<STEPS_COUNT;i++{
		fmt.Printf("%9s | %9s |", "1 step st", "2 step en")
	}
	fmt.Printf("\n")
}
func print_line(id int, start time.Time, steps[STEPS_COUNT] processing_time, abs_start time.Time){
	fmt.Printf("%3d | %9d |", id, int64(start.Sub(abs_start)))
	for i:=0;i<STEPS_COUNT;i++{
		fmt.Printf("%9d | %9d |", int64(steps[i].start.Sub(abs_start)), int64(steps[i].end.Sub(abs_start)))
	}
	fmt.Printf("\n")
}

func (v request_array_t)print_each_request_table(){   
	var abs_start = v.elems[0].start_wait 
	fmt.Println("Table of requests log")
	print_title()
	for i:=0; i < v.size; i++{
		print_line(v.elems[i].id, v.elems[i].start_wait, v.elems[i].steps, abs_start)
	}
}
/*
func print_request_array(v request_array_t){
	fmt.Print("Вектор [", v.size, "]: ")
	for i:=0;i<int(v.size);i++{
		fmt.Print(v.elems[i].str, " ")
	}
	fmt.Println()
}
*/



func (v request_array_t)print_request_graph(){    
	var time_steps[3][] float64
	var x[] float64 = make([]float64, v.size)
	for i:=0;i<3;i++{
		time_steps[i] = make([]float64, v.size)
	}

	for i:=0; i < v.size; i++{
		for j:=0;j < 3;j++{
			if (j == 0){
				time_steps[j][i] = float64(v.elems[i].steps[0].start.Sub(v.elems[i].start_wait))
			} else{
				time_steps[j][i] = float64(v.elems[i].steps[j].start.Sub(v.elems[i].steps[j-1].end))
		    }
	    }
		x[i] = float64(i)
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
				Name:    "Очередь 2",
				XValues: x,//[]float64{1.0, 2.0, 3.0, 4.0},
				YValues: time_steps[1],//[]float64{1.0, 2.0, 3.0, 4.0},
			},
			chart.ContinuousSeries{
				Name:    "Очередь 3",
				XValues: x,//[]float64{1.0, 2.0, 3.0, 4.0},
				YValues: time_steps[2],//[]float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}

	graph1.Elements = []chart.Renderable{
		chart.Legend(&graph1),
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph1.Render(chart.PNG, f)
	    
}

/*
	var resfloatx[] float64 = make([]float64, 10)
	var resfloaty1[] float64 = make([]float64, 10)
	var resfloaty2[] float64 = make([]float64, 10)
	var resfloaty3[] float64 = make([]float64, 10)
	var resfloaty4[] float64 = make([]float64, 10)
	var resfloaty5[] float64 = make([]float64, 10)
	var resfloaty6[] float64 = make([]float64, 10)
	
	for i := 1; i < 11; i++{
		resfloatx[i-1] = float64(i*10) 
		m1 = RandMatrix(i*10, i*10)
		m2 = RandMatrix(i*10, i*10)
		resfloaty1[i-1] = get_time(m1, m2, multiplicate_matrix_norm)
		resfloaty2[i-1] = get_time(m1, m2, multiplicate_matrix_vinograd)
		resfloaty3[i-1] = get_time(m1, m2, multiplicate_matrix_fast_vinograd)
		//print_row_table(i*10, t1, t2, t3)
	}
	graph1 := chart.Chart{
	    Background: chart.Style{
		Padding: chart.Box{
		    Top:  20,
	            Left: 20,
		},
	    },
            Series: []chart.Series{
                chart.ContinuousSeries{
                    Name:    "Стандартный алг.",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty1,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Алг. Винограда",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty2,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Оптимизированный алг. Винограда",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty3,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
            },
        }
	for i := 1; i < 11; i++{
		resfloatx[i-1] = float64(i*10) 
		m1 = RandMatrix(i*10, i*10)
		m2 = RandMatrix(i*10, i*10)
		resfloaty4[i-1] = get_time(m1, m2, multiplicate_matrix_norm)
		resfloaty5[i-1] = get_time(m1, m2, multiplicate_matrix_vinograd)
		resfloaty6[i-1] = get_time(m1, m2, multiplicate_matrix_fast_vinograd)
	}
	
	graph2 := chart.Chart{
	    Background: chart.Style{
		Padding: chart.Box{
		    Top:  20,
	            Left: 20,
		},
	    },
            Series: []chart.Series{
                chart.ContinuousSeries{
                    Name:    "Стандартный алг.",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty1,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Алг. Винограда",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty2,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Оптимизированный алг. Винограда",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty3,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
            },
        }


	graph3 := chart.Chart{
	    Background: chart.Style{
		Padding: chart.Box{
		    Top:  20,
	            Left: 20,
		},
	    },
            Series: []chart.Series{
                chart.ContinuousSeries{
                    Name:    "Алг. Винограда (чет.)",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty2,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Алг. Винограда (нечет.)",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty5,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
            },
        }

	graph4 := chart.Chart{
	    Background: chart.Style{
		Padding: chart.Box{
		    Top:  20,
	            Left: 20,
		},
	    },
            Series: []chart.Series{
                chart.ContinuousSeries{
                    Name:    "Оптимизированный алг. Винограда",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty3,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Оптимизированный алг. Винограда",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty6,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
            },
        }
    
	graph1.Elements = []chart.Renderable{
		chart.Legend(&graph1),
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph1.Render(chart.PNG, f)
	    
	graph2.Elements = []chart.Renderable{
		chart.Legend(&graph2),
	}

	f2, _ := os.Create("output2.png")
	defer f2.Close()
	graph2.Render(chart.PNG, f2)
	
	
	graph3.Elements = []chart.Renderable{
		chart.Legend(&graph3),
	}

	f3, _ := os.Create("output3.png")
	defer f3.Close()
	graph3.Render(chart.PNG, f3)
	
	
	graph4.Elements = []chart.Renderable{
		chart.Legend(&graph4),
	}

	f4, _ := os.Create("output4.png")
	defer f4.Close()
	graph4.Render(chart.PNG, f4)
*/
