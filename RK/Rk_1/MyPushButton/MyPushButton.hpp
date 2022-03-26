#ifndef MYPUSHBUTTON_HPP
#define MYPUSHBUTTON_HPP

#include <QPushButton>
#include "../ObserverInterface.hpp"

class MyPushButton : public QPushButton, public IObject
{
Q_OBJECT
    QColor curColor;
public:
    void update(ISubject *subject);
    MyPushButton(QWidget *parent = nullptr);
    ~MyPushButton();
//public signals:
//    void triggered();
};

#endif