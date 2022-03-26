#ifndef MATRIX_HPP
#define MATRIX_HPP

#include <iostream>
#include <vector>
#include <thread>
#include <mutex>

#include <chrono>
#include <ctime>    

#include <stdio.h>

#define REPEATS 200
#define MAX_ROWS 1001
#define MAX_COLUMNS 100
#define MAX_THREADS 33

extern std::mutex MuTeX;

struct matrix_t{
    double ** elems;
    int rows;
    int columns;
};

matrix_t make_matrix(int rows, int columns);
void input_matrix(matrix_t &m1);
void output_matrix(matrix_t &m1);

void standart_average(matrix_t &m1, double &res);
void parall_row_average(double &res, double **elems, int rows, int columns, int index, int step);
void parall_column_average(double &res, double **elems, int rows, int columns, int index, int step);
double control_thread(matrix_t &m1, 
    void func(double &, double **, int, int, int, int), int thread_count);

void generate_rand_matrix(matrix_t &m1);

#endif