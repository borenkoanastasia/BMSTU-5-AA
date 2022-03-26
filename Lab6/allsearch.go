package main 

//import "fmt"

type string_array_t struct {
	elems[] string
	//size int
}

func make_empty_string_array(size int)  string_array_t{
    var v string_array_t
    v.elems = make([]string, size)
    //v.size = size
    return v
}

func (arr string_array_t)print_array(){
	for i:=0;i<len(arr.elems);i++{}
}

func get_factorial(n int) int{
	var res int = 1
	for i:=1;i<n+1;i++{
		res = res*i
	}
	return res
}

func generate_routs_arr(n int)string_array_t{
	var len = get_factorial(n)
	var routs = make_empty_string_array(len)

	for i:=0;i<len;i++{
		routs[i] = ""
	}
	return routs
}




/*
func work_one_node(min_weight *float32, weight float32, deap int, flag_min_len bool, flag_fail_found *bool,  node int, first_node int,
	               contiguity *matrix_t, passed_nodes *list_t){
	//fmt.Println(*passed_nodes)
	if (deap == contiguity.rows-1){
		(passed_nodes).push(node)
		passed_nodes.print()
		fmt.Println("len = ", weight,  ", min_len = ", min_weight)
		if (contiguity.elem[node][first_node] == -1){
			return
		}
		weight += contiguity.elem[node][first_node]
		if (flag_min_len){
			if (weight > *min_weight){
				return
			}
		}
		flag_min_len = true
		*min_weight = weight
		return
	}

	for i:=0;i < contiguity.rows;i++{
		if ((*contiguity).elem[node][i] == -1){
			continue
		}
		if (passed_nodes.include(i)){
			continue
		}
		*flag_fail_found = false
		work_one_node(min_weight, weight + (*contiguity).elem[node][i], deap+1, flag_min_len, flag_fail_found, i, first_node, contiguity, passed_nodes)
		if (*flag_fail_found == true){
			passed_nodes.pop()
		}		else{
	        (passed_nodes).push(node)
		}
	}
}

/*
func get_len_for_one_node(node int, contiguity *matrix_t, passed_nodes *array_t, min_len float32) float32{
	(*passed_nodes).elems[node] = 1
	for i:=0; i<contiguity.columns; i++{
		fmt.Println(node, i, passed_nodes,(*contiguity).elem[node][i] )
		if ((*contiguity).elem[node][i] == -1){
			continue
		}
		if (passed_nodes.include(i)){
			continue
		}
		cur_len += get_len_for_one_node(i, contiguity, passed_nodes, min_len)
		fmt.Println(node, i, cur_len)
		if (cur_len == -1){
			continue
		}
		min_len += cur_len + (*contiguity).elem[node][i] 
		flag_change = true
		fmt.Println(node, i, cur_len)
	}
	if (flag_change){
		return min_len
	}
	return -1
}*/
/*
func all_search(contiguity matrix_t, pn *list_t) float32{
	//pn := make_empty_array(contiguity.rows)
	//var pn list_t
	pn.head = nil
	var len float32 = 0
	var flag_change bool = false
    work_one_node(&len, 0, 0, false, &flag_change, 0, 0, &contiguity, pn)//= get_len_for_one_node(0, &contiguity, &pn, cur_len)
	//fmt.Println(*pn)

	return len
	/*for i:= 1; i < contiguity.columns;i++{
		pn := make_empty_array(contiguity.rows)
		cur_len = 0
		cur_len = get_len_for_one_node(i, &contiguity, &pn, cur_len)
		if (cur_len<min_len){
			min_len = cur_len
		}
	}
}*/