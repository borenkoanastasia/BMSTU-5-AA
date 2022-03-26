#include "Matrix.hpp"

Matrix::Matrix(int n_rows, int n_columns): rows(n_rows), columns(n_columns)
{
    el = QVector<QVector<double>>(rows);
    for (int i = 0; i < rows; i++)
    {
        for (int j = 0; j < columns; j++)
        {
            el[i].append(0);
        }
    }
}
Matrix::Matrix(Point p): rows(1), columns(p.getSize() + 1)
{
    el.append(p.getEls());
    el[0].append(1);
}

Matrix::~Matrix()
{}

void Matrix::make_single_matrix()
{
    for (int i = 0; i < rows; i++)
    {
        for (int j = 0; j < columns; j++)
        { 
            if (i == j)
                el[i][j] = 1;
            else 
                el[i][j] = 0;
        }
    }
} 

void Matrix::make_scale_matrix(Point p)
{
    make_single_matrix();
    int k = 0;
    if (p.getSize() < rows - 1)
    {
        return;
    }
    for (int i = 0; i < rows - 1; i++)
    {
        for (int j = 0; j < columns - 1; j++)
        { 
            if (i == j)
            {
                el[i][j] = p.getCurEl(i);
                k++;
            }
        }
    }
}

void Matrix::make_z_rotate(double x)
{
    make_single_matrix();
    printf("make single success rows = %d, columns = %d\n", el.size(), columns);
    el[0][0] = cos(x);
    el[0][1] = -sin(x);
    el[1][0] = sin(x);
    el[1][1] = cos(x);
}
void Matrix::make_y_rotate(double y)
{
    make_single_matrix();
    el[0][0] = cos(y);
    el[0][2] = -sin(y);
    el[2][0] = sin(y);
    el[2][2] = cos(y);
}
void Matrix::make_x_rotate(double z)
{
    make_single_matrix();
    el[1][1] = cos(z);
    el[1][2] = -sin(z);
    el[2][1] = sin(z);
    el[2][2] = cos(z);
}
void Matrix::make_rotate_matrix(Point p)
{
    int k = 0;
    if (p.getSize() < rows - 1)
    {
        return;
    }
    printf("start make_rotate\n");
    Matrix m1(M3D, M3D), m2(M3D, M3D), m3(M3D, M3D);
    //printf("s1\n");
    m1.make_x_rotate(p.getCurEl(MY_X));
    //printf("x\n");
    m2.make_y_rotate(p.getCurEl(MY_Y));
    //printf("y\n");
    m3.make_z_rotate(p.getCurEl(MY_Z));
    //printf("z\n");
    Matrix res = (m1*m2)*m3;
    res.output();
    el = res.getMatrix();
    output();
    printf("end make_rotate\n");
}

QVector<QVector<double>> Matrix::getMatrix()
{
    return el;
}

void Matrix::make_transfer_matrix(Point p)
{
    make_single_matrix();
    int k = 0;
    if (p.getSize() < rows - 1)
    {
        return;
    }
    for (int i = 0; i < rows - 1; i++)
    {
        el[rows - 1][i] = p.getCurEl(i);
        k++;
    }
    el[rows - 1][columns-1] = 1;
}
void Matrix::add_center(Point center)
{
    int k = 0;
    if (center.getSize() < columns - 1)
    {
        return;
    }
    for (int i = 0; i < columns - 1; i++)
    {
        //printf("i = %d, rows = %d, columns = %d, point size = %d\n", i,rows, columns, center.getSize());
        el[i][columns - 1] = center.getCurEl(i);
        k++;
    }
}
int Matrix::getRows()
{
    return rows;
}
int Matrix::getColumns()
{
    return columns;
}
Point Matrix::getPoint()
{
    if(rows != 1)
    {
        //printf("el = %d\n", el.size());
        return Point(0);
    }
    //printf("el = %d, el[0] = %d\n", el.size(), el[0].size());
    Point ans(N3D);
    for (int i = 0; i < el[0].size() - 1; i++)
    {
        ans.setCurEl(i, el[0][i]);
    }
    return ans;
}
double Matrix::getCurEl(int row, int column)
{
    return el[row][column];
}
void Matrix::setCurEl(int row, int column, double data)
{
    el[row][column] = data;
}

// this * m1
Matrix Matrix::operator*(Matrix m1)
{
    if (this->getColumns() != m1.getRows())
    {
        printf("Not multiplicated! left.rows = %d, right.columns = %d\n", el.size(), m1.getColumns());
        return Matrix(0, 0);
    }
    Matrix ans(this->getRows(), m1.getColumns());
    for (int i = 0; i < this->getRows(); i++)
    {
        for (int j = 0; j < m1.getColumns(); j++)
        {
            for (int k = 0; k < this->getColumns(); k++)
            {
                ans.setCurEl(i, j, ans.getCurEl(i, j) + this->getCurEl(i, k) * m1.getCurEl(k, j));
            }
        }
    }
    return ans;
}

void Matrix::output()
{
    printf("matrix size = %d\n", el.size());
    for (int i = 0; i < el.size();i++)
    {
        for (int j = 0; j < el[0].size();j++)
        {
            printf("%lf, ", el[i][j]);
        }
        printf("\n");
    }
    printf("end matrix\n");
}