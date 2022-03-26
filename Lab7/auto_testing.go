package main

import "fmt"
import "math/rand"
import "time"

import (
	"os"
    "github.com/wcharczuk/go-chart" //exposes "chart"

	//"github.com/wcharczuk/go-chart/v2"
	//"github.com/wcharczuk/go-chart/v2/drawing"
)

const TEST_COUNT = 4
const FUNC_COUNT = 2
const SEG_COUNT_MAX = 52//102
const SEG_TEXT_STEP = 50
const SEG_COUNT_MIN = 1

func print_test_func(d dict_t, test_s test_t, f(func(dict_t, int)person_t), func_name string){
	time1 := get_time(d, test_s.test, f)
	res1 := f(d, test_s.test)

	fmt.Printf("%13s|%13d|%13s|%13f|", func_name, test_s.test, test_s.name, time1)
	if (res1.name == ""){
		fmt.Printf("%13s|", "NONE" )
	}else{
		fmt.Printf("%13s|", res1.name )
	}
}

func print_line(len int){
	for i:=0;i<len;i++{
		fmt.Printf("-")
	}
	fmt.Println()
}

func test_func(dictionary dict_t, text_snilses [TEST_COUNT]test_t, f(func(dict_t, int)person_t), func_name string){
	for i:=0;i<len(text_snilses);i++{
		print_test_func(dictionary, text_snilses[i], f, func_name)
		//fmt.Println(seg_search1(dictionary, text_snilses[i], 10))
		fmt.Println()
	}
	print_line(60)
}

type test_t struct{
	test int
	name string
}

func autotesting(dictionary dict_t){

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var index int = r1.Intn(len(dictionary))
	var text_snilses =[TEST_COUNT]test_t{{test:0, name:"err"}, 
										    {test:dictionary[0].snils, name:"first"}, 
					  				  	    {test:dictionary[index].snils, name:"rand"},
										    {test:dictionary[len(dictionary)-1].snils, name:"last"}} 

	var names = [FUNC_COUNT]string{"all_search","bin_search"}//,"seg_search"}
	var funcs = [FUNC_COUNT](func(dict_t, int)person_t){all_search, bin_search}//, seg_search}

	fmt.Printf("%13s|%13s|%13s|%13s|%13s|\n", "func name", "src", "name", "time", "res")
	print_line(60)

	for i:=0;i<FUNC_COUNT;i++{
		test_func(dictionary, text_snilses, funcs[i], names[i])
	}
	for i:=SEG_COUNT_MIN;i<SEG_COUNT_MAX;i+=SEG_TEXT_STEP{
		fmt.Println("SEGMENTS COUNT = ", i)
		seg_search := get_seg_search(dictionary, i)
		//fmt.Println("SEGMENTS COUNT = ", i)
		test_func(dictionary, text_snilses, seg_search, "seg_search")
	}
	fmt.Println()
}


func graph_testing(dict dict_t){
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	//var exp_count int = 10
	var snils = dict[r1.Intn(len(dict))].snils


	var resfloatx[] float64 = make([]float64, 10)
	var resfloaty1[] float64 = make([]float64, 10)
	var resfloaty2[] float64 = make([]float64, 10)
	var resfloaty3[] float64 = make([]float64, 10)


	var res1, res2 float64
	for i := 1; i < 11; i++{
		seg_search := get_seg_search(dict, i*10)
		get_time(dict, snils, seg_search)
	}

	res1 = float64(get_time(dict, snils, all_search))
	res2 = float64(get_time(dict, snils, bin_search))

	for i := 1; i < 11; i++{
		seg_search := get_seg_search(dict, i*10)
		resfloatx[i-1] = float64(i*10)
		resfloaty1[i-1] = res1
		resfloaty2[i-1] = res2
		resfloaty3[i-1] = float64(get_time(dict, snils, seg_search))


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
                    Name:    "all_search",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty1,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "bin_search",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty2,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "seg_search",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty3,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                //chart.ContinuousSeries{
                //    Name:    "Дамерау-Левенштейн рек.",
                //    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                //    YValues: resfloaty4,//[]float64{1.0, 2.0, 3.0, 4.0},
                //},
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
                    Name:    "seg_search",
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
/*
func graph_testing() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var index int = r1.Intn(len(dictionary))
	var text_snilses =[TEST_COUNT]test_t{{test:0, name:"err"}, 
										    {test:dictionary[0].snils, name:"first"}, 
					  				  	    {test:dictionary[index].snils, name:"rand"},
										    {test:dictionary[len(dictionary)-1].snils, name:"last"}} 

	var names = [FUNC_COUNT]string{"all_search","bin_search"}//,"seg_search"}
	var funcs = [FUNC_COUNT](func(dict_t, int)person_t){all_search, bin_search}//, seg_search}

    //var x = make([]float64, 10)
	var value1, value2, value3 int

	chart.DefaultBackgroundColor = chart.ColorTransparent
	chart.DefaultCanvasColor = chart.ColorTransparent

	barWidth := 120
	for i:=0;i<len(text_snilses);i++{
		time1 := get_time(d, test_s.test, f)
		//print_test_func(dictionary, text_snilses[i], f, func_name)
		//fmt.Println(seg_search1(dictionary, text_snilses[i], 10))
		fmt.Println()
	}

	var (
		colorWhite          = drawing.Color{R: 241, G: 241, B: 241, A: 255}
		colorMariner        = drawing.Color{R: 60, G: 100, B: 148, A: 255}
		colorLightSteelBlue = drawing.Color{R: 182, G: 195, B: 220, A: 255}
		colorPoloBlue       = drawing.Color{R: 126, G: 155, B: 200, A: 255}
		colorSteelBlue      = drawing.Color{R: 73, G: 120, B: 177, A: 255}
	)

	stackedBarChart := chart.StackedBarChart{
		Title:      "Quarterly Sales",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 100,
			},
		},
		Width:      810,
		Height:     500,
		XAxis:      chart.StyleShow(),
		YAxis:      chart.StyleShow(),
		BarSpacing: 50,
		Bars: []chart.StackedBar{
			{
				Name:  "Q1",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "10",
						Value: 32,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "46K",
						Value: 46,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "48K",
						Value: 48,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "42K",
						Value: 42,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
				},
			},
			{
				Name:  "Q2",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "45K",
						Value: 45,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "60K",
						Value: 60,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "62K",
						Value: 62,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "53K",
						Value: 53,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
				},
			},
			{
				Name:  "Q3",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "54K",
						Value: 54,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "58K",
						Value: 58,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "55K",
						Value: 55,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "47K",
						Value: 47,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
				},
			},
			{
				Name:  "Q4",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "46K",
						Value: 46,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "70K",
						Value: 70,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "74K",
						Value: 74,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "60K",
						Value: 60,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
				},
			},
		},
	}

	pngFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}

	if err := stackedBarChart.Render(chart.PNG, pngFile); err != nil {
		panic(err)
	}

	if err := pngFile.Close(); err != nil {
		panic(err)
	}*/