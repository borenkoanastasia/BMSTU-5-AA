#ifndef POINT_HPP 
#define POINT_HPP 

#include <QVector>

#define MY_X 0
#define MY_Y 1
#define MY_Z 2
#define M3D 4
#define N3D 3

class Point 
{
    QVector<double> el;
    int size;
public:
    Point();
    Point(int n_size);
    Point(QVector<double> data);
    ~Point();
    int getSize();
    double getCurEl(int i);
    void setCurEl(int i, double data);
    void output();
    QVector<double> getEls();
};

#endif