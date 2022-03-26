#include "matrix.hpp"

matrix_t &make_matrix(int rows, int columns){
    matrix_t m;
    m.rows = rows;
    m.columns = columns;
    m.elems = (double **)calloc(m.rows, sizeof(double *));
    for (int i = 0; i < m.rows; i++)
    {
        m.elems[i] = (double *)calloc(m.columns, sizeof(double));
    }
    return m;
}

void input_matrix(matrix_t &m)
{
    double buf;
    int rows, columns;
    std::cout << "Введите размер матрицы[rowsxcolumns]:\n";
    std::cout << "rows:";
    std::cin >> rows;
    std::cout << "columns:";
    std::cin >> columns;
    std::vector<double> v1;
    m = make_matrix(rows, columns);
    for (int i = 0; i < m.rows; i++)
    {
        m.elems[i] = (double *)calloc(m.columns, sizeof(double));
        for (int j = 0; j < m.columns; j++){
            std::cin >> m.elems[i][j];
        }
    }
}

void output_matrix(matrix_t &m){
    for (int i = 0; i < m.rows; i++)
    {
        for (int j = 0; j < m.columns; j++){
            std::cout << m.elems[i][j];
        }
        std::cout << std::endl;
    }
}

void standart_multiplicate(matrix_t &m1, matrix_t &m2, matrix_t &m3)
{
    matrix_t m3 = make_matrix(m1.rows, m2.columns);
    for (int i = 0; i < m1.rows; i++){
        for (int j = 0; j < m2.rows; j++){
            for (int k = 0; k < m2.columns; k++){
                m3.elems[i][j] += m1.elems[i][j]*m2.elems[i][j];
            }
        }
    }
}
void parall_row_multiplicate(matrix_t &m1, matrix_t &m2, matrix_t &m3, int step)
{
    double res;
    matrix_t m3 = make_matrix(m1.rows, m2.columns);
    for (int i = 0; i < m1.rows; i+=step){
        for (int j = 0; j < m2.rows; j++){
            res = 0;
            for (int k = 0; k < m2.columns; k++){
                m3.elems[i][j] += m1.elems[i][j]*m2.elems[i][j];
            }
            MuTeX.lock();
            m3.elems[i][j] = res;
            MuTeX.unlock();
        }
    }
}
void parall_column_multiplicate(matrix_t &m1, matrix_t &m2, matrix_t &m3, int step)
{
    double res;
    matrix_t m3 = make_matrix(m1.rows, m2.columns);
    for (int i = 0; i < m1.rows; i++){
        for (int j = 0; j < m2.rows; j+=step){
            res = 0;
            for (int k = 0; k < m2.columns; k++){
                res += m1.elems[i][j]*m2.elems[i][j];
            }
            MuTeX.lock();
            m3.elems[i][j] = res;
            MuTeX.unlock();
        }
    }
}

matrix_t control_thread(matrix_t &m1, matrix_t &m2, 
    matrix_t func(matrix_t &, matrix_t &, matrix_t &, int), int thread_count)
{
    std::thread th[thread_count];
    matrix_t m3;
    for (int i = 0; i < thread_count; i++)
    {
        th[i] = std::thread(func(m1, m2, m3, thread_count));
    }
    for (int i = 0; i < thread_count; i++)
    {
        th[i].join();
    }
    return m3;
}

void generate_matrix(matrix_t m1){
    for (int i = 0; i < m1.rows; i++){
        for (int j = 0; j < m1.columns; j++){
        std::rand();
        }
    }
}