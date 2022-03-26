package main

import "bufio"
import "os"
import "fmt"
import "strconv"

type dict_t []person_t

type person_t struct{
	name string
    snils int
}

const DICTSIZE = 2000;

func read_person(file *os.File, fileScanner*bufio.Scanner)person_t{
    var ch person_t
	var snils string

	fileScanner.Scan() 
	snils = fileScanner.Text()
	ch.snils, _ = strconv.Atoi(snils) 
	fileScanner.Scan()
	ch.name = fileScanner.Text()
	fileScanner.Scan()
	ch.name = ch.name + " " + fileScanner.Text()

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error while reading file: ", err)
	}
	//fmt.Println(ch)

	return ch
}

func read_dict(file *os.File)dict_t{
	fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanWords) 
	var dict dict_t = make(dict_t, DICTSIZE)
	var ch person_t = read_person(file, fileScanner)
	for i:=0;ch.snils != 0 && DICTSIZE != i;i++{
		    dict[i] = ch 
			ch = read_person(file, fileScanner)
	//		fmt.Println(ch)
	}
	return dict
}

func read_file(filename string)dict_t{

	// open the file
	file, err := os.Open(filename)

	//handle errors while opening
	if err != nil {
		fmt.Println("Error when opening file: ", err)
    }
	pers := read_dict(file)

	file.Close()

	return pers
}