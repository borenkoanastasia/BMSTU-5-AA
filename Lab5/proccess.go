package main

import "math/rand"
//import "fmt"

//string_encryption_Cesar
func get_string_encryption_Cesar(sh int) func(string) string {
	var code byte = byte(sh)
	return func(str1 string) string {
		var str2 string = ""
		var r byte
		for i := 0; i < len(str1); i++ {
			r = str1[i] + code
			if (r > 'z'){
				r = r-'z' + 'a'
			}
			str2 = str2 + string([]byte{r})
			//fmt.Println(r," ", str2, string([]byte{r}))
		}
		return str2
	}
}

//string_encryption_Xor
func get_string_encryption_Xor(sh string) func(string) string {
	var code string = sh
	return func(str1 string) string {
		var str2 string = ""
		var r byte
		for i := 0; i < len(str1); i++ {
			r = r - 'a'
			r = (str1[i]-'a') ^ (code[i]-'a')
			r = r + 'a'
			str2 = str2 + string([]byte{r})
			//fmt.Println(r, str2)
		}
		return str2
	}
}

func generate_encryption_Xor_key(max_str_len int) string {
	var r int = 0;
	var max_int int = 'z'-'a'
	var res string = ""
	for i:=0;i<max_str_len;i++ {
		r = rand.Intn(max_int) + 'a'
		res = res + string(string([]byte{byte(r)}))
	}
	return res
}

//string_encryption_Atbash
func string_encryption_Atbash(str1 string) string {
	var str2 string = ""
	var r byte
	for i := 0; i < len(str1); i++ {
		r = (str1[i])
		r = r - 'a'
		r = 'z' - r
		str2 = str2 + string([]byte{r})

	}
	return str2
}
