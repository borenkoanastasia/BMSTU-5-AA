#ifndef MYIMAGE_HPP
#define MYIMAGE_HPP

#include <QImage>
#include <QPainter>
#include <QtGlobal>
#include <QCoreApplication>
#include <QColor>
#include "../MyLabel/MyLabel.hpp"
#include "../ObserverInterface.hpp"
#include "../MyPoint/MyPoint.hpp"
#include "../MyException/MyException.hpp"

class MyImage : public IObject
{
    //QImage *base = nullptr;
    QImage *image = nullptr;
    QColor foneColor = Qt::white;
    QColor drawColor;
    MyLabel *display;
    MyPoint *lastEvent = nullptr;
    
public:
    MyImage();
    MyImage(int width, int height);
    ~MyImage();
    void setSizes(int width, int height);
    QImage *getImage();
    QColor getPixel(int x, int y);
    void updateDisplay();
    void setPixel(int x, int y, QColor color);
    void setPoint(MyPoint center);
    void setLine(MyPoint start, MyPoint end);
    void setRect(MyPoint start, MyPoint end);
    QColor getFone();

    void update(ISubject *subject);
    void changeColors(QColor newDraw);

    void newMouseEvent(MyPoint event);

    MyLabel *getDisplay();

    void makeStartFone();
    void clean();

    int getWidth();
    int getHeight();
};

#endif

