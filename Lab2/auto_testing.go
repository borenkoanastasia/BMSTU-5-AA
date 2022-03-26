package main 
import "fmt"
import "os"
import "github.com/wcharczuk/go-chart" //exposes "chart"

func print_title_table(){
	fmt.Printf("|%3s|%15s|%15s|%20s|\n", "N", "Обычный алг.", "Виноград", "Оптимиз. Виноград")
}
func print_row_table(i int, t1 float64, t2 float64, t3 float64){
	fmt.Printf("|%3d|%15f|%15f|%20f|\n", i, t1, t2, t3)
}

func auto_testing(){
	var t1, t2, t3 float64
	var m1, m2 matrix_t

	fmt.Println()
	fmt.Println("Четный размер матриц")
	
	fmt.Println(line(6+24*4+6))

	print_title_table()
	for i := 1; i < 11; i++{
		m1 = RandMatrix(i*10, i*10)
		m2 = RandMatrix(i*10, i*10)
		t1 = get_time(m1, m2, multiplicate_matrix_norm)
		t2 = get_time(m1, m2, multiplicate_matrix_vinograd)
		t3 = get_time(m1, m2, multiplicate_matrix_fast_vinograd)
		print_row_table(i*10, t1, t2, t3)
	}
	fmt.Println(line(6+24*4+6))

	fmt.Println()

	fmt.Println("Нечетный размер матриц")
	fmt.Println(line(6+24*4+6))

	print_title_table()
	for i := 1; i < 11; i++{
		m1 = RandMatrix(i*10+1, i*10+1)
		m2 = RandMatrix(i*10+1, i*10+1)
		t1 = get_time(m1, m2, multiplicate_matrix_norm)
		t2 = get_time(m1, m2, multiplicate_matrix_vinograd)
		t3 = get_time(m1, m2, multiplicate_matrix_fast_vinograd)
		print_row_table(i*10 + 1, t1, t2, t3)
	}
}

func graph_testing(){
	var m1, m2 matrix_t


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
	
}
