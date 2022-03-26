package main

//import "fmt"
/*
type node_t struct {
	elem int
	next *node_t
}

type list_t struct{
	head *node_t
}

func create_node(elem int)*node_t{
	var node *node_t = new(node_t)
	node.next = nil
	node.elem = elem
	return node
}

func (l list_t)print(){
	var cur = l.head
	fmt.Print("List: ")
	for ;cur.next != nil;cur = cur.next{
    	fmt.Print(cur.elem, " ")
	}
	fmt.Println()
}
func (l list_t)include(num int)bool{
	if (l.head==nil)	{
		return false
	}
	for cur:=l.head; cur.next != nil;cur=cur.next{
		if cur.elem == num{
			return true
		}
	}
	/*
	for i:= 0;i<arr.size;i++{
		if (num == arr.elems[i]){
			return true
		}
	}
	return false
}

func (l *list_t)push(elem int){
	var new_node = create_node(elem)
	if l.head == nil{
		l.head = new_node
		//fmt.Println(l, l.head, l.head == nil, elem, new_node)
		return
	}
	var cur = l.head
	for ;cur.next != nil;cur = cur.next{}
	cur.next = new_node
	//fmt.Println(l, cur)
}

func (l *list_t)pop()int{
	if (l.head == nil){
		return -1
	}
	var elem = l.head.elem 
	//var del_node = l.head
	//defer delete(del_node) ???
	l.head = l.head.next
	return elem
}


type int_array_t struct {
	elems[] int
	//size int
}

func make_empty_int_array(size int)  int_array_t{
    var v int_array_t
    v.elems = make([]int, size)
    //v.size = size
    return v
}

func (arr int_array_t)print_array(){
	for i:=0;i<len(arr.elems);i++{}
}

type ant_t struct{
	way int_array_t   //не пройденные вершины
	route int_array_t //пройденные вершины
}

type ant_colony_t struct{
	graph matrix_t
	pheromonths matrix_t
	alpha float32
	beta float32
	Q float32
	p float32
}

type matrix_t struct{
	elem[][] float32
	rows int 
	columns int
}

func make_empty_matrix(rows int, columns int)  matrix_t{
    var m matrix_t
    m.elem = make([][]float32, rows)
    m.rows = rows
    m.columns = columns
    for row:= 0; row < m.rows; row++{
        m.elem[row] = make([]float32, columns)
    }
    return m
}

func input_matrix(m *matrix_t)bool{
	var rows, columns int
	var x float32
	var rc error
	fmt.Print("Введите кол-во строк: ")
	_, rc = fmt.Scanf("%d\n", &rows)
	if(rc != nil){
		return false
	}
	fmt.Print("Введите кол-во столбцов: ")
	_, rc = fmt.Scanf("%d\n", &columns)
	if(rc != nil){
		return false
	}
	*m = make_empty_matrix(rows, columns)

	for row:=0;row<rows;row++{
		for column:=0;column<columns;column++{
			_, rc = fmt.Scanf("%f", &x)
			//fmt.Scanf("%f ", &x)
			if(rc != nil){
				return false
			}
			(*m).elem[row][column] = x
		}
	}
	fmt.Print("\n")
	return true
}
func print_matrix(m matrix_t){
	fmt.Print("Матрица [", m.rows, " x ", m.columns, "]:\n")
	for i:=0;i<m.rows;i++{
		for j:=0;j<m.columns;j++{
			fmt.Print(m.elem[i][j], " ")
		}
		fmt.Print("\n")
	}
}*/