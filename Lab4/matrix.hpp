#ifndef MATRIX_HPP
#define MATRIX_HPP

#include <iostream>
#include <vector>

typedef std::vector<double> vector_t;
typedef std::vector<std::vector<double>> matrix_t;

vector_t multiplicateVectorAndMatrix(matrix_t matrix, vector_t vector);
vector_t make_vector(double x, double y);

//matrix_t makeMatrix(double x, double y, int operation);
matrix_t multiplicateMatrixes(matrix_t a, matrix_t b);

matrix_t makeTestMatrix1();
void output_matrix(matrix_t m);
void makeLeftTriangleMatrix(matrix_t &m);
void makeRightTriangleMatrix(matrix_t &m);

#endif