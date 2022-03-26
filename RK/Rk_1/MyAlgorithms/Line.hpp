#ifndef LINE_HPP
#define LINE_HPP 

#include "Point.hpp"
#include "Matrix.hpp"
#include "../MyImage/MyImage.hpp"

class Line
{
    QVector<Point> points;
    Point z_intersection;
public:
    Line();
    ~Line();
    void addPoint(Point p);
    QVector<Point> getPoints(Point p);
    Point getPoint(int i);
    int getSize();
    void setPointZIntersection(Point p);
    Point getPointZIntersection();
    void transformPoints(Matrix transform);
    void getXYminXYmax(double &x_min, double &x_max, double &y_min, double &y_max);

    void draw(double *y_min, double *y_max, bool invert, MyImage *drawer);
    bool add_smal_edge(int x_start, int y_start, int x_end, int y_end, double *y_min, double *y_max);
    void addEdge(double *y_min, double *y_max);

    void output();
};

#endif