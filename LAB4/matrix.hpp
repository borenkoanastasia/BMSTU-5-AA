#ifndef MATRIX_HPP
#define MATRIX_HPP

#include <iostream>
#include <vector>
#include <thread>
#include <mutex>

std::mutex MuTeX;

struct matrix_t{
    double ** elems;
    int rows;
    int columns;
};

void input_matrix();
void output_matrix();

void standart_multiplicate(matrix_t &m1, matrix_t &m2, matrix_t &m3);
void parall_row_multiplicate(matrix_t &m1, matrix_t &m2, matrix_t &m3, int step);
void parall_column_multiplicate(matrix_t &m1, matrix_t &m2, matrix_t &m3, int step);
matrix_t control_thread(matrix_t &m1, matrix_t &m2, 
    matrix_t func(matrix_t &, matrix_t &, matrix_t &, int), int thread_count);

#endif