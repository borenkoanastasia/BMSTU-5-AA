package main

import "fmt"

import (
    //"bytes"
    "os"
    "github.com/wcharczuk/go-chart" //exposes "chart"
)

func print_title_table(){
	fmt.Printf("|%5s|%32s|%32s|%32s|\n", "Длина", "Пузырек", "Вставки", "Выбор")
	fmt.Printf("|%5s|%10s|%10s|%10s|%10s|%10s|%10s|%10s|%10s|%10s|\n", "", "Сорт.мас.", "Рев.мас.", "Случ.мас.", "Сорт.мас.", "Рев.мас.", "Случ.мас.", "Сорт.мас.", "Рев.мас.", "Случ.мас.")
}
func print_row_table(i int, t1 float64, t2 float64, t3 float64, t4 float64, t5 float64, t6 float64, t7 float64, t8 float64, t9 float64){
	fmt.Printf("|%5d|%10.3f|%10.3f|%10.3f|%10.3f|%10.3f|%10.3f|%10.3f|%10.3f|%10.3f|\n", i, t1, t2, t3, t4, t5, t6, t7, t8, t9)
}

func auto_testing(){
	var size int
	var v1, v2, v3 vector_t
	var t1, t2, t3, t4, t5, t6, t7, t8, t9 float64

	print_title_table()
	for i := 1; i < 21;i+=1{
		size = i*10
		v1 = SortVector(size)
		get_time(v1, sort_bubble)
	}
	for i := 1; i < 21;i++{
		size = i*100
		v1 = SortVector(size)
		v2 = ReverseVector(size)
		v3 = RandVector(size)
		
		t1 = get_time(v1, sort_bubble)
		t4 = get_time(v2, sort_bubble)
		t7 = get_time(v3, sort_bubble)
		
		t2 = get_time(v1, sort_insert)
		t5 = get_time(v2, sort_insert)
		t8 = get_time(v3, sort_insert)
		
		t3 = get_time(v1, sort_choice)
		t6 = get_time(v2, sort_choice)
		t9 = get_time(v3, sort_choice)
		
		print_row_table(size, t1, t4, t7, t2, t5, t8, t3, t6, t9)
	}
}

func graph_testing(){


	var size int
	var v1, v2, v3 vector_t
	//var t1, t2, t3, t4, t5, t6, t7, t8, t9 float64

	var resfloatx[] float64 = make([]float64, 20)
	var resfloaty1[] float64 = make([]float64, 20)
	var resfloaty2[] float64 = make([]float64, 20)
	var resfloaty3[] float64 = make([]float64, 20)
	var resfloaty4[] float64 = make([]float64, 20)
	var resfloaty5[] float64 = make([]float64, 20)
	var resfloaty6[] float64 = make([]float64, 20)
	var resfloaty7[] float64 = make([]float64, 20)
	var resfloaty8[] float64 = make([]float64, 20)
	var resfloaty9[] float64 = make([]float64, 20)

	for i := 1; i < 21;i+=1{
		size = i*10
		v1 = SortVector(size)
		get_time(v1, sort_bubble)
	}
	//print_title_table()
	for i := 1; i < 21;i++{
		size = i*10
		resfloatx[i-1] = float64(i*100)
		v1 = SortVector(size)
		v2 = ReverseVector(size)
		v3 = RandVector(size)
		resfloaty1[i-1] = get_time(v1, sort_bubble)
		resfloaty2[i-1] = get_time(v1, sort_insert)
		resfloaty3[i-1] = get_time(v1, sort_choice)
		resfloaty4[i-1] = get_time(v2, sort_bubble)
		resfloaty5[i-1] = get_time(v2, sort_insert)
		resfloaty6[i-1] = get_time(v2, sort_choice)
		resfloaty7[i-1] = get_time(v3, sort_bubble)
		resfloaty8[i-1] = get_time(v3, sort_insert)
		resfloaty9[i-1] = get_time(v3, sort_choice)
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
                    Name:    "Сортировка пузырьком",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty1,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Сортировка вставками",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty2,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Сортировка выбором",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty3,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
            },
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
                    Name:    "Сортировка пузырьком",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty4,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Сортировка вставками",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty5,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Сортировка выбором",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty6,//[]float64{1.0, 2.0, 3.0, 4.0},
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
                    Name:    "Сортировка пузырьком",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty7,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Сортировка вставками",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty8,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Сортировка выбором",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty9,//[]float64{1.0, 2.0, 3.0, 4.0},
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
                    Name:    "Отсортированный м.",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty1,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Реверстнутый м.",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty4,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Случайный массив",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty7,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
            },
        }
    //buffer := bytes.NewBuffer([]byte{})
    //err := graph.Render(chart.PNG, buffer)
    //fmt.Print( err)
    
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
