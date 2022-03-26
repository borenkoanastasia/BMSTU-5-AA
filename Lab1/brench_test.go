package main

import "testing"

func BenchmarkLM_10S(b *testing.B){
    //var data_len int
    //var data [][2] string
    //read_file("test", data, &data_len);
	var s1 string = "aaaaaaaaaa"
	var s2 string = "aaaaaaaaaa"
    b.ResetTimer()
    for i:=0;i<b.N;i++{
        LMatrix(&s1, &s2, 3, 3)//data[0][0], data[0][1])
    }
}
func BenchmarkLR_10S(b *testing.B){
    //var data_len int
    //var data [][2] string
    //read_file("test", data, &data_len);
	var s1 string = "aaaaaaaaaa"
	var s2 string = "aaaaaaaaaa"
    b.ResetTimer()
    for i:=0;i<b.N;i++{
        LRecursion(&s1, &s2, 3, 3)//data[0][0], data[0][1])
    }
}

func BenchmarkDLR_10S(b *testing.B){
    //var data_len int
    //var data [][2] string
    //read_file("test", data, &data_len);
	var s1 string = "aaaaaaaaaa"
	var s2 string = "aaaaaaaaaa"
    b.ResetTimer()
    for i:=0;i<b.N;i++{
        DLRecursion(&s1, &s2, 3, 3)//data[0][0], data[0][1])
    }
}

func BenchmarkLRK_10S(b *testing.B){
    var _ map[string] int = make(map[string]int)
    //var data_len int
    //var data [][2] string
    //read_file("test", data, &data_len);

	var s1 string = "aaaaaaaaaa"
	var s2 string = "aaaaaaaaaa"
    //b.ResetTimer()
    for i:=0;i<b.N;i++{
		var lrk = getLRecursionKesh()
		lrk(&s1, &s2, 3, 3)//data[0][0], data[0][1])
    }
}

func BenchmarkLM_10D(b *testing.B){
    //var data_len int
    //var data [][2] string
    //read_file("test", data, &data_len);
	var s1 string = "aaaaaaaaaa"
	var s2 string = "bbbbbbbbbb"
    b.ResetTimer()
    for i:=0;i<b.N;i++{
        LMatrix(&s1, &s2, 3, 3)//data[0][0], data[0][1])
    }
}
func BenchmarkLR_10D(b *testing.B){
    //var data_len int
    //var data [][2] string
    //read_file("test", data, &data_len);
	var s1 string = "aaaaaaaaaa"
	var s2 string = "bbbbbbbbbb"
    b.ResetTimer()
    for i:=0;i<b.N;i++{
        LRecursion(&s1, &s2, 3, 3)//data[0][0], data[0][1])
    }
}

func BenchmarkDLR_10D(b *testing.B){
    //var data_len int
    //var data [][2] string
    //read_file("test", data, &data_len);
	var s1 string = "aaaaaaaaaa"
	var s2 string = "bbbbbbbbbb"
    b.ResetTimer()
    for i:=0;i<b.N;i++{
        DLRecursion(&s1, &s2, 3, 3)//data[0][0], data[0][1])
    }
}

func BenchmarkLRK_10D(b *testing.B){
    var _ map[string] int = make(map[string]int)
    //var data_len int
    //var data [][2] string
    //read_file("test", data, &data_len);

	var s1 string = "aaaaaaaaaa"
	var s2 string = "bbbbbbbbbb"
    //b.ResetTimer()
    for i:=0;i<b.N;i++{
		var lrk = getLRecursionKesh()
		lrk(&s1, &s2, 3, 3)//data[0][0], data[0][1])
    }
}


//go test -bench=. -benchmem > test_res.txt
