package main

//import "fmt"

func all_search(d dict_t, s int)person_t{
	var p person_t
	p.snils = s

	for i:=0;i<len(d);i++{
		if (s == d[i].snils){
			p.name = d[i].name
			break
		}
	}
	return p
}

func bin_search(d dict_t, s int)person_t{
	if (len(d) < 1){
		return person_t{name: "", snils: s}
	}
	var middle = len(d)/2
	if (d[middle].snils < s){
		return bin_search(d[middle:], s)
	} else	if (d[middle].snils > s){
		return bin_search(d[:middle], s)
	}
	return d[middle]
}

func make_seg(d dict_t, seg_count int)[][]person_t{
    var seg [][]person_t = make([][]person_t, 10)

	for i:=0;i<10;i++{
		for j:=0;j<len(d);j++{
			if (d[j].snils / 10000 %10 == i){
				seg[i] = append(seg[i], d[j])
			}
	}}
	return seg
}

func get_seg_search(d dict_t, seg_count int) (func(dict_t, int)person_t){
	seg := make_seg(d, seg_count)
	return func(d dict_t, s int)person_t{
        var p person_t
		var index = s / 10000 %10
		if (index >= len(seg) || index <0){
			p = person_t{name:"", snils:s}
		} else {
			p = bin_search(seg[index], s)
		}
		return p
	}
}
/*
func seg_search1(d dict_t, s int, seg_count int)person_t{
	//var pers string = d[0].name

	seg := make_seg(d, seg_count)
	var start int = -1
	var end int = -1 

	for i:=0;i<len(seg)-1;i++{
		//fmt.Println(seg[i], d[seg[i]].snils, seg[i+1], d[seg[i+1]].snils, s)
		if (d[seg[i]].snils <= s && d[seg[i+1]].snils > s){
			start = seg[i]
			end = seg[i+1]
			break
		}
	}
	//fmt.Println(seg[len(seg)-1], d[seg[len(seg)-1]].snils, len(d)-1, d[len(d)-1].snils, s)
	if (d[seg[len(seg)-1]].snils < s && d[len(d)-1].snils >= s){
		start =  seg[len(seg)-1]
		end = len(d)
	}
	if (start == -1){
		return person_t{name:"", snils:s}
	}
	//fmt.Println(start, end)
	p := bin_search(d[start:end], s)


	return p
}*/
