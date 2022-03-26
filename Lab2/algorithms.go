package main

/*
	Список функций - алгоритмов
	multiplicate_matrix_norm
	multiplicate_matrix_vinograd
	multiplicate_matrix_fast_vinograd
*/

func multiplicate_matrix_norm(m1 matrix_t, m2 matrix_t, rc *bool)matrix_t{
	var m3 matrix_t = make_empty_matrix(m1.rows, m2.columns)
	*rc = true
	if (m2.rows != m1.columns){
		*rc = false
		return m3;
	}
	for i := 0;i < m1.rows;i++{
		for j := 0; j < m2.columns;j++{
			for k := 0; k < m1.columns;k++{
				m3.elem[i][j] += m1.elem[i][k] * m2.elem[k][j]
			}
		}
	}
	return m3
}


func multiplicate_matrix_vinograd(m1 matrix_t, m2 matrix_t, rc *bool)matrix_t{
	var m3 matrix_t = make_empty_matrix(m1.rows, m2.columns)
	*rc = true
	if (m2.rows != m1.columns){
		*rc = false
		return m3;
	}

	var rows_factor vector_t = make_empty_vector(m1.rows)
	var columns_factor vector_t = make_empty_vector(m2.columns)

	for i:=0;i<m1.rows;i++{
		for j:=1;j<m1.columns;j+=2{
			rows_factor.elem[i] += m1.elem[i][j - 1] * m1.elem[i][j]
		}
	}
	for i:=0;i<m2.columns;i++{
		for j:=1;j<m2.rows;j+=2{
			columns_factor.elem[i] += m2.elem[j - 1][i] * m2.elem[j][i]
		}
	}

	for i := 0;i < m1.rows;i++{
		for j := 0; j < m2.columns;j++{
			m3.elem[i][j] = - rows_factor.elem[i] - columns_factor.elem[j]
			for k := 0; k < m1.columns/2;k++{
				m3.elem[i][j] += (m1.elem[i][k*2] + m2.elem[2*k+1][j])*(m1.elem[i][k*2+1] + m2.elem[2*k][j])
			}
		}
	}

	if (m1.columns % 2 == 1){
		for i:= 0; i<m1.rows;i++{
			for j:=0;j<m2.columns;j++{
				m3.elem[i][j] += m1.elem[i][m1.columns-1]*m2.elem[m1.columns - 1][j]
			}
		}
	}

	return m3
}


func multiplicate_matrix_fast_vinograd(m1 matrix_t, m2 matrix_t, rc *bool)matrix_t{
	var m3 matrix_t = make_empty_matrix(m1.rows, m2.columns)
	*rc = true
	if (m2.rows != m1.columns){
		*rc = false
		return m3;
	}

	var rows_factor vector_t = make_empty_vector(m1.rows)
	var columns_factor vector_t = make_empty_vector(m2.columns)

	for i:=0;i<m1.rows;i++{
		for j:=1;j<m1.columns;j+=2{
			rows_factor.elem[i] += m1.elem[i][j - 1] * m1.elem[i][j]
		}
	}
	for i:=0;i<m2.columns;i++{
		for j:=1;j<m2.rows;j+=2{
			columns_factor.elem[i] += m2.elem[j - 1][i] * m2.elem[j][i]
		}
	}

	var flag bool = false
	if (m1.columns % 2 == 1){
		flag = true
	}

	for i := 0;i < m1.rows;i++{
		for j := 0; j < m2.columns;j++{
			m3.elem[i][j] = - rows_factor.elem[i] - columns_factor.elem[j]
			for k := 0; k < m1.columns/2;k++{
				m3.elem[i][j] += (m1.elem[i][k*2] + m2.elem[2*k+1][j])*(m1.elem[i][k*2+1] + m2.elem[2*k][j])
			}
			if (flag){
				m3.elem[i][j] += m1.elem[i][m1.columns-1]*m2.elem[m1.columns - 1][j]
			}
		}
	}
	return m3
}
