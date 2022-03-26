package main

import (
    "fmt"
    //"math/rand"
    "time"
    )


type matrix_t struct{
    elem[][] float32
    rows int 
    columns int
}
type quene_time_t struct{
    start time.Time 
    end time.Time
}
type request_t struct{
    id int
    res string
    sorce string
    qt1 quene_time_t
    qt2 quene_time_t
    qt3 quene_time_t
}
type queue_t struct{
    waiting [](*request_t)
    last int
}

// Работа с матрицей
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
}
// Работа с очередью
func make_queue(amount int) *queue_t{
    var queue *queue_t = new(queue_t)
    queue.waiting = make([](*request_t), amount, amount)
    queue.last = -1
    return queue
}
func (q *queue_t) push(elem *request_t){
    if q.last !=  len(q.waiting)-1{
        q.waiting[q.last+1] = elem
        q.last+=1
    }
}
func (q *queue_t) pop() *request_t{
    if (q.last < 0){
        return nil
    }
    elem := q.waiting[0]
    q.waiting = q.waiting[1:]
    q.last -= 1
    return elem
}
func make_start_queue(amount int, max_str_len int) *queue_t{
    var queue *queue_t = make_queue(amount)
    for i := 0;i < amount;i++{
        var req *request_t = new(request_t)
        req.id = i
        req.sorce = generate_encryption_Xor_key(max_str_len)
        queue.push(req)
    }
    return queue
}

// Ввод строки
func input_string()string{
    var str1 string
    fmt.Scanf("%s\n", &str1)
    return str1
}
