#include "Point.hpp"

Point::Point()
{}

Point::Point(int n_size): size(n_size)
{
    for (int i = 0; i < size; i++)
    {
        el.append(0);
    }
}

Point::Point(QVector<double> data): el(data), size(data.size())
{}

Point::~Point()
{}

int Point::getSize()
{
    return el.size();
}
double Point::getCurEl(int i)
{
    return el[i];
}
void Point::setCurEl(int i, double data)
{
    el[i] = data;
}

QVector<double> Point::getEls()
{
    return el;
}

void Point::output()
{
    printf("point: ");
    for(int i = 0; i < el.size();i++)
    {
        printf("%lf, ", el[i]);
    }
    printf("\n");
}