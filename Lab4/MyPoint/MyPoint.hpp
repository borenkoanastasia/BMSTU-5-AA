#ifndef MYPOINT_H
#define MYPOINT_H



class MyPoint
{
private:
    double x;
    double y;
public:
    MyPoint();
    MyPoint(double c_x, double c_y);
    ~MyPoint();
    double getX();
    double getY();
    void setX(double c_x);
    void setY(double c_y);
};

#endif