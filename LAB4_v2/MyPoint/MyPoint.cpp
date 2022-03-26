#include "MyPoint.hpp"

MyPoint::MyPoint()
{}
MyPoint::MyPoint(double c_x, double c_y): x(c_x), y(c_y)
{}
MyPoint::~MyPoint()
{}

double MyPoint::getX()
{
    return x;
}
double MyPoint::getY()
{
    return y;
}

void MyPoint::setX(double c_x)
{
    x = c_x;
}
void MyPoint::setY(double c_y)
{
    y = c_y;
}