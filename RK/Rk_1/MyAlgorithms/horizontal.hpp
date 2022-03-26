#ifndef HORIZONTAL_HPP
#define HORIZONTAL_HPP

#include <cmath>
#include <QVector>
//#include "matrix.hpp"
#include "Line.hpp"
#include "Matrix.hpp"
#include "../MyImage/MyImage.hpp"


struct one_parameter
{
    double start;
    double end;
    double step;
};

struct parameters
{
    one_parameter x;
    one_parameter z;
};

parameters make_param(double x_start, double x_end, double x_step, double z_start, double z_end, double z_step);
Point get_center_point(MyImage *drawer);
Point get_rotate_point(double x, double y, double z);
void show_alg(Matrix current_rotate, MyImage *drawer, parameters param, double (*func)(double, double));

#endif