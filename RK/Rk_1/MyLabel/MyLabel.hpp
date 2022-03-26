#ifndef MYLABEL_HPP
#define MYLABEL_HPP

#include <QLabel>
#include <QMouseEvent>
#include "../MyPoint/MyPoint.hpp"

class MyLabel : public QLabel
{
    Q_OBJECT
public:
    MyLabel();
    ~MyLabel();
    void setImage(QImage *image);
    void mousePressEvent(QMouseEvent *event);
signals:
    void mouseAddPointEvent(MyPoint event);
};

#endif