#include "MyAlgorithmFloatingHorizont.hpp"

double getYf1(double x, double z)
{
    return sin(x + z);
}
double getYf2(double x, double z)
{
    return x*x + z*z;
}
double getYf3(double x, double z)
{
    return x + z;
}

QVector<QString> MyAlgorithmShortedEdges::getNames()
{
    return names;
}
void MyAlgorithmShortedEdges::makeNames()
{
    names.append("y = sin(x + z)");
    names.append("y = x*x + z*z");
    names.append("y = x + z");
}
Matrix MyAlgorithmShortedEdges::rotate(Matrix current_rotate, Point rotate_point)
{
    Matrix new_rotate(M3D, M3D);
    new_rotate.make_rotate_matrix(rotate_point);
    current_rotate.output();
    current_rotate = current_rotate * new_rotate;
    new_rotate.output();
    current_rotate.output();
    return current_rotate;
}

void MyAlgorithmShortedEdges::draw(MyImage *drawer, QString funcName, Matrix current_rotate, parameters param)
{
    int index = names.indexOf(funcName);
    current_rotate.output();
    switch (index)
    {
        case 0:
            show_alg(current_rotate, drawer, param, getYf1);
            break;
        case 1:
            show_alg(current_rotate, drawer, param, getYf2);
            break;
        case 2:
            show_alg(current_rotate, drawer, param, getYf3);
            break;
    }
}

MyAlgorithmShortedEdges::MyAlgorithmShortedEdges()
{
    makeNames();
}

MyAlgorithmShortedEdges::~MyAlgorithmShortedEdges()
{}
