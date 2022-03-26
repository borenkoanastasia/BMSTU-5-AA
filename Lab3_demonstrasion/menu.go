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
    for ;true;{
		fmt.Println(line(length))
		fmt.Println()
		fmt.Println(title(length))
		fmt.Println(line(length))

		fmt.Println("\t1.Manual testing")
		fmt.Println("\t2.Auto   testing")
		fmt.Println("\tAnother choice is exit")

		choise = input_string()

		fmt.Println(line(length))

		if	(choise == "1"){
			manual_testing()
		} else if	(choise == "2"){
			auto_testing()
		} else {
			break
		}
    }
}
