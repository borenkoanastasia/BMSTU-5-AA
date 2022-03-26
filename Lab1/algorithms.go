package main

/*
Список функций - реализаций алгоритмов:
	LRecursion(s1 *string, s2 *string, i int, j int)
	DLRecursion(s1 *string, s2 *string, i int, j int)
	getLRecursionKesh()(func (s1 *string, s2 *string, i int, j int)
	LMatrix(str1 *string, str2 *string, i int, j int)
	LMatrix_print(str1 *string, str2 *string, i int, j int)
*/


// Левенштейн рекурсия
func LRecursion(s1 *string, s2 *string, i int, j int) int{
	if (i == 0 && j == 0){
		return 0
	} else if (i == 0 ){
		return j
	} else if (j == 0){
		return i
	} else if (i > 0 && j > 0){
		var delta int = 1 
		if ((*s1)[i] == (*s2)[j]){
			delta = 0
		}
		return min3(LRecursion(s1, s2, i, j - 1) + 1, LRecursion(s1, s2, i - 1, j) + 1, LRecursion(s1, s2, i - 1, j -1) + delta)
	}
	return -1
}

// Дамерау-Левенштейн рекурсия
func DLRecursion(s1 *string, s2 *string, i int, j int) int{
	if (i == 0 && j == 0){
		return 0
	} else if (i == 0 ){
		return j
	} else if (j == 0){
		return i
	} else if (i > 0 && j > 0 && (*s1)[i] == (*s2)[j - 1] && (*s1)[i - 1] == (*s2)[j]){
		var delta int = 1 
		if ((*s1)[i] == (*s2)[j]){
			delta = 0
		}
		return min4(DLRecursion(s1, s2, i, j - 1) + 1, DLRecursion(s1, s2, i - 1, j) + 1, DLRecursion(s1, s2, i - 1, j - 1) + delta,
				    DLRecursion(s1, s2, i - 2, j - 2) + delta)
	} else if (i > 0 && j > 0){
		var delta int = 1 
		if ((*s1)[i] == (*s2)[j]){
			delta = 0
		}
		return min3(DLRecursion(s1, s2, i, j - 1) + 1, DLRecursion(s1, s2, i - 1, j) + 1, DLRecursion(s1, s2, i - 1, j -1) + delta)
	}
	//fmt.Println("DLRECURSION: Impossible critical ERROR")
	return -1
}

// Левенштейн рекурсия с кешем
func getLRecursionKesh()(func (s1 *string, s2 *string, i int, j int) int){
	var delta int = 1 
	var ok bool 
	var Cur int
	var M map[string] int = make(map[string]int)
	var map_administrator = func (s string, ans int){
		var v int
		var ok bool
		v, ok = M[s]
		if (!ok){
			M[s] = ans
		} else if (v > ans){
			M[s] = ans
		}
	}
	return func (s1 *string, s2 *string, i int, j int) int{
		if (i == 0 && j == 0){
			map_administrator(((*s1)[:i] + (*s2)[:j]), 0)
			return 0
		} else if (i == 0 ){
			map_administrator(((*s1)[:i] + (*s2)[:j]), j)
			return j
		} else if (j == 0){
			map_administrator(((*s1)[:i] + (*s2)[:j]), i)
			return i
		} else if (i > 0 && j > 0){
			if ((*s1)[i] == (*s2)[j]){
				delta = 0
			}
			Cur, ok = M[(*s1)[:i]+ (*s2)[:j]]
			if (!ok){
				Cur =  min3(DLRecursion(s1, s2, i, j - 1) + 1, DLRecursion(s1, s2, i - 1, j) + 1, DLRecursion(s1, s2, i - 1, j -1) + delta)
				map_administrator((*s1)[:i]+ (*s2)[:j], Cur)
			}
			return Cur
		}
		return -1
	}
}

// Левенштейн матрица
// Получить матрицу
func get_levenshtein_matrix(str1 string, str2 string) Matrix{
    var m Matrix = make_empty_matrix(len(str1), len(str2))   //str1 - vertical word, str2 - horisontal word
    for row := 0; row < len(str1); row++{
        for column := 0; column < len(str2); column++{
            if (row == 0 && column == 0){
                m.matrix[row][column] = 0 
            } else if (row == 0) {
                m.matrix[row][column] = m.matrix[row][column - 1] + 1 
            } else if (column == 0) {
                m.matrix[row][column] = m.matrix[row - 1][column] + 1 
            } else {
                var delta int = 1 // for diag step
                if (str1[row] == str2[column]){
                    delta = 0
                }

                var con1 int = m.matrix[row - 1][column] + 1 // vertical step
                var con2 int = m.matrix[row][column - 1] + 1 // horisontal step
                var con3 int = m.matrix[row - 1][column - 1] + delta// diag step
                
                m.matrix[row][column] = min3(con1, con2, con3)
            }
        }
    }
    return m
}

// Решить без вывода матрицы
func LMatrix(str1 *string, str2 *string, i int, j int) int {
	var m Matrix = get_levenshtein_matrix(*str1, *str2)
	return m.matrix[len(*str1) - 1][len(*str2) - 1]
}

// Решить с выводом матрицы
func LMatrix_print(str1 *string, str2 *string, i int, j int) int {
	var m Matrix = get_levenshtein_matrix(*str1, *str2)
	
	print_levenshtein_matrix(*str1, *str2, m)
	return m.matrix[len(*str1) - 1][len(*str2) - 1]
}
