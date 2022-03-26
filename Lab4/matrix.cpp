#include "matrix.hpp"

vector_t make_vector(double x, double y)
{
	vector_t result = {x, y, 1};
	return result;
}

vector_t multiplicateVectorAndMatrix(matrix_t matrix, vector_t vector)
{
	vector_t result = {0, 0, 0};
	for (int i = 0; i < (int) matrix.size(); i++)
	{
		for (int j = 0; j < (int) matrix.size(); j++)
		{
			result[i] += matrix[j][i] * vector[j];
			//std::cout << result[i] << " " << matrix[j][i] << " " << vector[j] << std::endl;
			//std::cout << i << " " << j << std::endl;
		}
	}
	//std::cout << std::endl;
	return result;
}

matrix_t makeTestMatrix1()
{
    matrix_t m = {{ 2, 1, 0,  0},
                  {-1, 1, 4, 13},
                  { 1, 2, 3, 14}};
    return m;
}

void makeLeftTriangleMatrix(matrix_t &m)
{
    double cur_del;
    for (int i = 0; i < m.size(); i++){
        cur_del = m[i][0];
        for (int k = 0; k < m[0].size(); k++){
            m[i][k] = m[i][k]/cur_del;
        }
        for (int j = i + 1; j < m.size(); j++){
            cur_del = m[j][0];
            for (int k = 0; k < j; k++){
                m[j][k] = m[j][k]/cur_del - m[0][k];   
            }
        }
    }
}
void makeRightTriangleMatrix(matrix_t &m)
{
    double cur_del_1, cur_del_2;
    for (int i = 0; i < m.size(); i++){
        cur_del_1 = m[i][0];
        for (int k = 0; k < m[0].size(); k++){
            m[i][k] = m[i][k]/cur_del_1;
        }
        for (int j = i + 1; j < m.size(); j++){
            cur_del_2 = m[i][i];
            for (int k = 0; k < m.size(); k++){
                std::cout << cur_del_2 << " " << i << " " << j << " " << k << std::endl << std::endl;
                if (m[i][k] * m[0][k] > 0){
                    m[i][k] = m[i][k]/cur_del_2 - m[0][k];
                }
                else{
                    m[i][k] = m[i][k]/cur_del_2 + m[0][k];
                }   
            }
            output_matrix(m);
        }
    }
}
// складываются строки r1 + r2 и записываются в r1 
void complicate_rows(matrix_t m, int r1, int r2){}

void output_matrix(matrix_t m)
{
    //std::cout << std::endl;
    for (int i = 0; i < m.size(); i++){
        for (int k = 0; k < m[0].size(); k++){
            std::cout << m[i][k] << " ";
        }
        std::cout << std::endl;
    }
    std::cout << std::endl;
}

matrix_t multiplicateMatrixes(matrix_t a, matrix_t b)
{
	matrix_t matrix = {{0, 0, 0}, {0, 0, 0}, {0, 0, 0}};
	for (int i = 0; i < (int) matrix.size(); i++)
	{	
		for (int j = 0; j < (int) matrix.size(); j++)
		{
			for (int k = 0; k < (int) matrix.size(); k++)
			{
				matrix[j][k] += a[j][i] * b[i][k];
			}
		}
	}
	return matrix;
}