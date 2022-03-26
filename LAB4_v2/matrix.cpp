#include "matrix.hpp"

matrix_t make_matrix(int rows, int columns){
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
    std::cout << "строки: ";
    std::cin >> rows;
    std::cout << "столбцы: ";
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
    std::cout << "Матрица: " << m.rows << "x" << m.columns << std::endl;
    for (int i = 0; i < m.rows; i++)
    {
        for (int j = 0; j < m.columns; j++){
            std::cout << m.elems[i][j] << " ";
        }
        std::cout << std::endl;
    }
}

void standart_average(matrix_t &m1, double &res)
{
    res = 0;
    for (int i = 0; i < m1.rows; i++){
        for (int j = 0; j < m1.columns; j++){
            res += m1.elems[i][j];
        }
    }
    res = res / m1.rows/m1.columns;
}
void parall_row_average(double &res, double **elems, int rows, int columns, int index, int step)
{
    double r = 0;
    for (int i = index; i < rows; i+=step){
        for (int j = 0; j < columns; j++){
            r += elems[i][j];
        }
    }
    MuTeX.lock();
    res += r;
    MuTeX.unlock();
}
void parall_column_average(double &res, double **elems, int rows, int columns, int index, int step)
{
    double r = 0;
    for (int i = 0; i < rows; i++){
        for (int j = index; j < columns; j+=step){
            r += elems[i][j];
        }
    }
    MuTeX.lock();
    res += r;
    MuTeX.unlock();
}

double control_thread(matrix_t &m1, 
    void func(double &, double **, int, int, int, int), int thread_count)
{
    std::thread th[thread_count];
    double res = 0;
    for (int i = 0; i < thread_count; i++)
    {
        th[i] = std::thread(func, std::ref(res), m1.elems, m1.rows, m1.columns, i, (thread_count));
    }
    double del = m1.columns*m1.rows;
    for (int i = 0; i < thread_count; i++)
    {
        th[i].join();
    }
    
    return res / del;
}

void generate_rand_matrix(matrix_t &m1){
    for (int i = 0; i < m1.rows; i++){
        for (int j = 0; j < m1.columns; j++){
            m1.elems[i][j] = std::rand();
        }
    }
}