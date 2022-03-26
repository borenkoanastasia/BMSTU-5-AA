#include "MyLabel.hpp"

MyLabel::MyLabel()
{
    setFixedSize(QSize(1350, 1000));
}
MyLabel::~MyLabel()
{}

void MyLabel::setImage(QImage *image)
{
    setPixmap(QPixmap::fromImage(*(image)));
}

void MyLabel::mousePressEvent(QMouseEvent *event)
{
    emit mouseAddPointEvent(MyPoint(event->x(), event->y()));
}