package main

import "fmt"
import "time"




func main(){
	fmt.Println( "abc", string_encryption_Cesardz("abc", 2, 2))



	
	fmt.Println("Hey! Now we will work with conveer.")

	var request_amount int

        fmt.Println("How many request we will test?")
	_, err := fmt.Scanf("%d", &request_amount)
	if (err != nil){
		fmt.Println("[ERROR] ", err)
		return
	}

	fmt.Println("So, we will test conveer on", request_amount, "requests")

	fmt.Println("START")

	/*// buf time
	g := 1
	for i:=0;i<1000000;i++{
		g = g*i
	}*/

	var channel = setup(request_amount)
	var res request_array_t = conveyorRun(request_amount, channel)

	fmt.Println("END")
	fmt.Println("Data:")

	res.print_request_log()
	res.print_request_graph()

    fmt.Println("Show each request log?(1 -yes, other -no)")
	var r int
	fmt.Scanf("%d", &r)
	if (r==1){
    	res.print_each_request_table()
	}

	request_amount *=10
	var t1, t2 time.Duration
	var conv_time1, lean_time2 time.Duration
	testing(request_amount, &t2, &t1)
	var repeats = 10

	for j:=0;j<repeats;j++{
		testing(request_amount, &t1, &t2)
		conv_time1 += (t1)
		lean_time2 += (t2)
	}
	conv_time1 /= time.Duration(repeats)
	lean_time2 /= time.Duration(repeats)

	fmt.Println("On ", request_amount, " request we have results:")
	fmt.Println("Conveer time = ", lean_time2)
	fmt.Println("Lenear  time = ", conv_time1)

	
    fmt.Println("Draw compare graph?(1 -yes, other -no)")
	fmt.Scanf("%d", &r)
	if (r==1){
    	print_comp_graph()
	}
}
