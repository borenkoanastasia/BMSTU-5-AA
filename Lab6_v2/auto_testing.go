package main

import "fmt"
import "math/rand"
//import "time"

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


/*func print_tests_func(d dict_t, test_s test_t, f(func(dict_t, int)person_t), func_name string){
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
}*/

func get_test_matrix(size int)float_matrix_t{
	var m = make_float_matrix(size, size)
	for i:=0;i<size;i++{
		for j:=i+1;j < size;j++{
			m[j][i] = float64(rand.Intn(100))
			m[i][j] = m[j][i]
		}
	}
	return m
}
/*

    0 & 0 & 1790 & 200 & 1900 & 63 & 1659 & 1820 & 1395 & 2382 & 649 \\
    1 & 1790 & 0 & 1573 & 2435 & 1515 & 714 & 892 & 2193 & 1590 & 1003 \\
    2 & 200 & 1573 & 0 & 833 & 392 & 2404 & 962 & 902 & 141 & 1123 \\
    3 & 1900 & 2435 & 833 & 0 & 2283 & 1652 & 2362 & 2262 & 1512 & 2166 \\
    4 & 63 & 1515 & 392 & 2283 & 0 & 1322 & 290 & 1305 & 2100 & 969 \\
    5 & 1659 & 714 & 2404 & 1652 & 1322 & 0 & 256 & 78 & 2236 & 2041 \\ 
    6 & 1820 & 892 & 962 & 2362 & 290 & 256 & 0 & 1180 & 1547 & 1279 \\ 
    7 & 1395 & 2193 & 902 & 2262 & 1305 & 78 & 1180 & 0 & 1640 & 1161 \\
    8 & 2382 & 1590 & 141 & 1512 & 2100 & 2236 & 1547 & 1640 & 0 & 2212 \\
    9 & 649 & 1003 & 1123 & 2166 & 969 & 2041 & 1279 & 1161 & 2212 & 0 \\
	*/
func get_test_params_matrix()float_matrix_t{
	var size = 10
	var m = make_float_matrix(size, size)
	m[0] = []float64{0, 1790, 200, 1900, 63, 1659, 1820, 1395, 2382, 649}
	m[1] = []float64{1790, 0, 1573, 2435, 1515, 714, 892, 2193, 1590, 1003}
	m[2] = []float64{200, 1573, 0, 833, 392, 2404, 962, 902, 141, 1123}
	m[3] = []float64{1900, 2435, 833, 0, 2283, 1652, 2362, 2262, 1512, 2166}
	m[4] = []float64{63, 1515, 392, 2283, 0, 1322, 290, 1305, 2100, 969}
	m[5] = []float64{1659, 714, 2404, 1652, 1322, 0, 256, 78, 2236, 2041}
	m[6] = []float64{1820, 892, 962, 2362, 290, 256, 0, 1180, 1547, 1279}
	m[7] = []float64{1395, 2193, 902, 2262, 1305, 78, 1180, 0, 1640, 1161}
	m[8] = []float64{2382, 1590, 141, 1512, 2100, 2236, 1547, 1640, 0, 2212}
	m[9] = []float64{649, 1003, 1123, 2166, 969, 2041, 1279, 1161, 2212, 0}
	return m
}
func autotesting_parameters(){

//	s1 := rand.NewSource(time.Now().UnixNano())
//	r1 := rand.New(s1)

	//var index int = r1.Intn(len(dictionary))
	//var text_snilses =[TEST_COUNT]test_t{{test:0, name:"err"}, 
										    //{test:dictionary[0].snils, name:"first"}, 
					  				  	    //{test:dictionary[index].snils, name:"rand"},
										    //{test:dictionary[len(dictionary)-1].snils, name:"last"}} 

	//var names = [FUNC_COUNT]string{"all_search","bin_search"}//,"seg_search"}
	//var funcs = [FUNC_COUNT](func(dict_t, int)person_t){all_search, bin_search}//, seg_search}

	fmt.Printf("%13s&%13s&%13s&%13s&%13s\\\\\n", "alpha", "beta", "p", "Результат", "Разница")
	line(60)

	m:=get_test_params_matrix()

	//var opt_way, way int_array_t
	var lenght, cur_lenght float64
	search_all(m, &lenght)

	for i:=0;i<11;i++{
		alpha := float64(i)/10
		beta := 1 - alpha
		for k:=0;k<11;k++{
			p := float64(k)/10
			q:=20.0
			var colony colony_t;
			var tmax = 100

			colony.generate_colony(m, alpha, beta, q, p, tmax)
			ant_alg(colony, &cur_lenght)	
			fmt.Printf("%13.1f&%13.1f&%13.1f&%13.1f&%13.1f\\\\\n", alpha, beta, p, cur_lenght, cur_lenght-lenght)
		}
	}
	fmt.Println()
}


func autotesting(){

//	s1 := rand.NewSource(time.Now().UnixNano())
//	r1 := rand.New(s1)

	//var index int = r1.Intn(len(dictionary))
	//var text_snilses =[TEST_COUNT]test_t{{test:0, name:"err"}, 
										    //{test:dictionary[0].snils, name:"first"}, 
					  				  	    //{test:dictionary[index].snils, name:"rand"},
										    //{test:dictionary[len(dictionary)-1].snils, name:"last"}} 

	//var names = [FUNC_COUNT]string{"all_search","bin_search"}//,"seg_search"}
	//var funcs = [FUNC_COUNT](func(dict_t, int)person_t){all_search, bin_search}//, seg_search}

	fmt.Printf("%13s|%13s|%13s|\n", "size", "func_name", "time")
	line(60)

	for i:=2;i<11;i+=1{
		m := get_test_matrix(i)
		fmt.Printf("%13d|%13s|%13f|\n", i, "all_search", get_time(m, search_all))
		fmt.Printf("%13d|%13s|%13f|\n", i, "ant_search", get_time(m, AllAntsAlg))
	}
	fmt.Println()
}


func graph_testing(){
	//var exp_count int = 10
	var resfloatx[] float64 = make([]float64, 9)
	var resfloaty1[] float64 = make([]float64, 9)
	var resfloaty2[] float64 = make([]float64, 9)
	
	for i := 1; i < 10; i++{
		test_matrix := get_test_matrix(i)
		resfloatx[i-1] = float64(i)
		resfloaty1[i-1] = get_time(test_matrix, search_all)
		resfloaty2[i-1] = get_time(test_matrix, AllAntsAlg)
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
                    Name:    "Полный перебор",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty1,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                chart.ContinuousSeries{
                    Name:    "Муравьиный алгоритм",
                    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                    YValues: resfloaty2,//[]float64{1.0, 2.0, 3.0, 4.0},
                },
                //chart.ContinuousSeries{
                //    Name:    "Дамерау-Левенштейн рек.",
                //    XValues: resfloatx,//[]float64{1.0, 2.0, 3.0, 4.0},
                //    YValues: resfloaty4,//[]float64{1.0, 2.0, 3.0, 4.0},
                //},
            },
        }

	/*graph2 := chart.Chart{
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
        }*/
    //buffer := bytes.NewBuffer([]byte{})
    //err := graph.Render(chart.PNG, buffer)
    //fmt.Print( err)
    
	graph1.Elements = []chart.Renderable{
		chart.Legend(&graph1),
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph1.Render(chart.PNG, f)
	    
	/*graph2.Elements = []chart.Renderable{
		chart.Legend(&graph2),
	}

	f2, _ := os.Create("output2.png")
	defer f2.Close()
	graph2.Render(chart.PNG, f2)*/
	
}