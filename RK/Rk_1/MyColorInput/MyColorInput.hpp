#ifndef MYCOLORINPUT_HPP
#define MYCOLORINPUT_HPP

#include "../ObserverInterface.hpp"
#include <QColor>
#include <QColorDialog>
#include <stdio.h>

class MyColorInput : public ISubject
{
    QColor newColor;
    QVector<IObject *> subscribers;
    void update();

public:
    MyColorInput(){};
    ~MyColorInput(){};
    void addObject(IObject *obj);
    void removeObject(IObject *obj);
    void getNewColor();

    void setNewColor(QColor color);
    QColor getColor();
};
#endif