#ifndef MYALGORITHMSHOTEDGES_HPP
#define MYALGORITHMSHOTEDGES_HPP

#include <cmath>
#include <iostream>
#include <unistd.h>
#include "MyAlgorithm.hpp"
#include "horizontal.hpp"
//#using namespace std;

class MyAlgorithmShortedEdges: public MyAlgorithm//, QObject
{
    //Q_OBJECT

    QVector<QString> names;
public:
    MyAlgorithmShortedEdges();
    ~MyAlgorithmShortedEdges();
    
    QVector<QString> getNames();
    void makeNames();
    void draw(MyImage *drawer, QString funcName, Matrix current_rotate, parameters param);
    Matrix rotate(Matrix current_rotate, Point rotate_point);
};

#endif