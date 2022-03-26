#ifndef M_HPP
#define M_HPP

#include <QVector>
#include <cmath>
#include "Point.hpp"

class Matrix
{
    int rows;
    int columns;
    QVector<QVector<double>> el;
    void make_x_rotate(double x);
    void make_y_rotate(double y);
    void make_z_rotate(double z);
public:
    Matrix(int n_rows, int n_columns);
    Matrix(Point p);
    ~Matrix();

    int getRows();
    int getColumns();
    Point getPoint();
    QVector<QVector<double>> getMatrix();
    double getCurEl(int row, int column);
    void setCurEl(int row, int column, double data);

    void make_single_matrix();
    void make_scale_matrix(Point p);
    void make_rotate_matrix(Point p);
    void make_transfer_matrix(Point p);
    void add_center(Point center);

    void output();

    Matrix operator*(Matrix m1);
};
#endif