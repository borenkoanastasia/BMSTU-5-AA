package main

import "fmt"

import (
    //"bytes"
    "os"
    "github.com/wcharczuk/go-chart" //exposes "chart"
)


func auto_print_title(){

	fmt.Printf("|%6s|%24s|%24s|%24s|%24s|\n", "Длина", "Левенштейн мат.", "Левенштейн рек.", 
											  "Левенштейн рек. с кешем", "Дамерау-Левенштейн рек.")
}

func auto_testing(){
	var exp_count int = 10
	var res [10]test_result_t
	var LRecursionKesh = getLRecursionKesh()



	auto_print_title()


	for i := 1; i <= exp_count; i++{
		res[i-1].len = i*10
		res[i-1].times[0]= get_time(StringType1Bytes(i), StringType2Bytes(i), LMatrix)
		res[i-1].times[1] = get_time(StringType1Bytes(i), StringType2Bytes(i), LRecursion)
		res[i-1].times[2] = get_time(StringType1Bytes(i), StringType2Bytes(i), LRecursionKesh)
		res[i-1].times[3] = get_time(StringType1Bytes(i), StringType2Bytes(i), DLRecursion)

		print_result_auto(res[i-1])

	}
}

func graph_testing(){
	var exp_count int = 10
	var LRecursionKesh = getLRecursionKesh()


	var resfloatx[] float64 = make([]float64, 10)
	var resfloaty1[] float64 = make([]float64, 10)
	var resfloaty2[] float64 = make([]float64, 10)
	var resfloaty3[] float64 = make([]float64, 10)
	var resfloaty4[] float64 = make([]float64, 10)


	for i := 1; i <= exp_count; i++{
		resfloatx[i-1] = float64(i*10)
		resfloaty1[i-1] = float64(get_time(StringType1Bytes(i), StringType2Bytes(i), LMatrix))
		resfloaty2[i-1] = float64(get_time(StringType1Bytes(i), StringType2Bytes(i), LRecursion))
		resfloaty3[i-1] = float64(get_time(StringType1Bytes(i), StringType2Bytes(i), LRecursionKesh))
		resfloaty4[i-1] = float64(get_time(StringType1Bytes(i), StringType2Bytes(i), DLRecursion))


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
                    Name:    "Левенштейн мат.",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty1,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Левенштейн рек.",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty2,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Левенштейн рек. с кешем",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty3,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Дамерау-Левенштейн рек.",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty4,//[]float64{1.0, 2.0, 3.0, 4.0},
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
                    Name:    "Левенштейн мат.",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty1,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Левенштейн рек. с кешем",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty3,//[]float64{1.0, 2.0, 3.0, 4.0},
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
	
}
