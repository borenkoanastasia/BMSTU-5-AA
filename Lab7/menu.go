package main

import "fmt"

func line(length int) string{
	var str string = ""
	for i:=0;i<length;i++{
		str = str + "_"
	}
	return str
}
func title(length int) string{
	var str string = ""
	for i:=0;i<(length - 4)/2;i++{
		str = str + " "
	}
	str = str + "MENU"
	for i:=0;i<(length - 4)/2;i++{
		str = str + " "
	}
	return str
}

func menu(){
	var length int = 111
	var choise string
	var dictionary = read_file("dict.txt");
    for ;true;{
		fmt.Println(line(length))
		fmt.Println()
		fmt.Println(title(length))
		fmt.Println(line(length))

		fmt.Println("\t1.Протестировать")
		//fmt.Println("\t2.График")
		fmt.Println("\tИначе - выход")

		choise = input_string()

		fmt.Println(line(length))

		if	(choise == "1"){
		//} else if	(choise == "2"){
			autotesting(dictionary);
		} else if	(choise == "2"){
			graph_testing(dictionary)
			break
		} else {
			break
		}
    }
}
