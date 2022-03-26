package main

import "time"
import "fmt"

func conveer(end chan *queue_t){
	var fsch chan *request_t = make(chan *request_t, 0)
	var ssch chan *request_t = make(chan *request_t, 0)
	var tsch chan *request_t = make(chan *request_t, 0)

	var xor_key string = generate_encryption_Xor_key(100)

    cesar := get_string_encryption_Cesar(13)  // сдвиг шифра (Цезарь цикличный)
    xor := get_string_encryption_Xor(xor_key) // 
	atbash := string_encryption_Atbash        //

	source := make_start_queue(101, 100)

    

	first_step := func() {
		var current_request *request_t
		for{
			current_request = <- fsch
    	    current_request.qt1.start = time.Now()
			current_request.res = cesar(current_request.sorce)
        	current_request.qt1.end = time.Now()
			//fmt.Println(11)
			ssch <- current_request
		}
	}
	second_step := func() {
		var current_request *request_t
		for{
			current_request = <- ssch
        	current_request.qt2.start = time.Now()
			current_request.res = atbash(current_request.res)
        	current_request.qt2.end = time.Now()
			//fmt.Println(22)
			tsch <- current_request
		}
	}
	third_step := func() {
		var current_request *request_t

	    res := make_queue(101)
		for{
			current_request = <- tsch
        	current_request.qt3.start = time.Now()
			current_request.res = xor(current_request.res)
        	current_request.qt3.end = time.Now()
			fmt.Println((*current_request).id )
            //mutex
			res.push(current_request)
			if ((*current_request).id == 100){
				fmt.Println((*current_request).id, 100)
				end <- res 
			}
		}
	}

	go first_step();
	go second_step();
	go third_step();
	fmt.Println("Все запустили")
	for {
		var r *request_t
		//fmt.Println(i)
		r = source.pop()
		fmt.Println(r)
		if (r == nil){
			break
		}
		select{
			case fsch <- r:
		}		
		//fmt.Println(i)
	}

	fmt.Println("Все выполнили")
}

func analysis(queue *queue_t){
	var first_waited time.Duration; var second_waited time.Duration; var third_waited time.Duration
	//first_waited = 0; second_waited = 0; third_waited = 0
	line := queue.waiting
	start := line[0].qt1.start
	fmt.Printf("Время начала\n")
	//fmt.Printf(line[0])
	for i:=0;i<len(line);i++{
		if line[i] != nil{
			fmt.Println(i, line[i].qt1.start.Sub(start),line[i].qt2.start.Sub(start), line[i].qt3.start.Sub(start))
		}}
	fmt.Printf("Время завершения\n")
	for i:=0;i<len(line);i++{
		if line[i] != nil{
			fmt.Println(i, line[i].qt1.end.Sub(start),line[i].qt2.end.Sub(start), line[i].qt3.end.Sub(start))
		}}
	fmt.Printf("Линии простаивали\n")
	for i:=0; i<len(line)-1;i++{
		first_waited += line[i+1].qt1.start.Sub(start)-line[i].qt1.end.Sub(start)
		second_waited+= line[i+1].qt2.start.Sub(start)-line[i].qt2.end.Sub(start)
		third_waited += line[i+1].qt3.start.Sub(start)-line[i].qt3.end.Sub(start)
	}
	fmt.Println(first_waited, second_waited, third_waited)
}

/*
func conv(amount int, wait chan int) *queue_t{
	uno := make(chan *cake, 5)
	dos := make(chan *cake, 5)
	tres := make(chan *cake, 5)
	line := new_queue(amount) 
	first := func(){
		for{
			select{
				case a := <- uno:
				//fmt.Printf("Cake num %d started dough\n", a.num)
				a.dough = true
				
				a.started_dough = time.Now()
				took_dough := 200
				time.Sleep(time.Duration(took_dough) * time.Millisecond)
				
				a.finished_dough = time.Now()
				dos <- a
			}
		}
	}
	
	second:= func(){
	for{
		select{
			case a := <- dos:
				//fmt.Printf("Cake num %d started topping\n", a.num)
				a.topping = true
				
				a.started_topping = time.Now()
				took_topping := 200
				time.Sleep(time.Duration(took_topping) * time.Millisecond)
				
				a.finished_topping = time.Now()
				tres <- a
		}
	}
}
	
	third := func(){
	for{
		select{
			case a := <- tres:
			//fmt.Printf("Cake num %d started decor\n", a.num)
			a.decor = true
			
			a.started_decor = time.Now()
			took_decor := 200
			time.Sleep(time.Duration(took_decor) * time.Millisecond)
			
			a.finished_decor = time.Now()
			line.push(a)
			if (a.num == amount){
			 wait <- 0 }
			
		}
	}
}
	
	go first()
	go second()
	go third()
	for i:=0; i<=amount; i++{
		a := new(cake)
		a.num = i
		uno <- a
	}
	return line
}
*/


/*
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
  "sync"
)

var start = time.Now()

func getTime() string {
	return time.Now().Format("15:04:05.99999999")
}

type procFunc func(int) int

func proc(input <-chan int, procName string, f procFunc, mutex *sync.Mutex) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for arg := range input {
      mutex.Lock()
			fmt.Print("time: ", getTime(), " in  ", procName, ": ", arg, "\n")
      mutex.Unlock()
			value := f(arg)
      mutex.Lock()
			fmt.Print("time: ", getTime(), " out ", procName, ": ", value, "\n")
      mutex.Unlock()
			output <- value
		}
	}()
	return output
}

func linearProc(arg int, procName string, f procFunc) int {
	fmt.Print("time: ", getTime(), " in  ", procName, ": ", arg, "\n")
	value := f(arg)
	fmt.Print("time: ", getTime(), " out ", procName, ": ", value, "\n")

	return value
}

func setup(n int, mutex *sync.Mutex) <-chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)

		for i := 0; i < n; i++ {
			value := rand.Intn(100)
			fmt.Print("time: ", getTime(), " send: ", value, "\n")
			channel <- value
		}
	}()
	return channel
}

func linearGenerate() int {
	value := rand.Intn(100)
	fmt.Print("time: ", getTime(), " send: ", value, "\n")
	return value
}

func stepA(arg int) int {
	time.Sleep(100)
	return arg + arg
}

func stepB(arg int) int {
	time.Sleep(100)
	return arg + arg
}

func stepC(arg int) int {
	time.Sleep(100)
	return arg + arg
}

func conveyorRun(count int) {
  var mutex sync.Mutex
	var channel = setup(count, &mutex)
	channel = proc(channel, "A", stepA, &mutex)
	channel = proc(channel, "B", stepB, &mutex)
	channel = proc(channel, "C", stepC, &mutex)

	for res := range channel {
    mutex.Lock()
		fmt.Print("time: ", getTime(), " res: ", res, "\n")
    mutex.Unlock()
	}
}

func main() {
  rand.Seed(time.Now().UnixNano())
	var count int
	fmt.Print("Enter setup count: ")
	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conveyorRun(count)
}
*/
