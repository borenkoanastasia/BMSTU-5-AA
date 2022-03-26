#include "MyImage.hpp"

MyImage::MyImage()
{
    display = new MyLabel();
    setSizes(display->size().width(), display->size().height());
    makeStartFone();
    display->setImage(image);
}

void MyImage::updateDisplay()
{
    display->setImage(image);
}

MyImage::~MyImage()
{
    if (image)
        delete image;
}

QColor MyImage::getFone()
{
    return foneColor;
}

void MyImage::makeStartFone()
{
    for (int i = 0; i < image->size().width(); i++)
    {
        for (int j = 0; j < image->size().height(); j++)   
        {
            image->setPixelColor(i, j, QColor(foneColor));
        }
    }
}

void MyImage::changeColors(QColor newDraw)
{
    for (int i = 0; i < image->size().width(); i++)
    {
        for (int j = 0; j < image->size().height(); j++)   
        {
            if (image->pixelColor(i, j).value() == drawColor.value())
            {
                image->setPixelColor(i, j, newDraw);
            }
        }
    }
                
    drawColor = newDraw;
}

void MyImage::setSizes(int width,int height)
{
    if (image)
        delete image;
    image = new QImage(QSize(width, height), QImage::Format_RGB32);
}

void MyImage::setPixel(int x, int y, QColor color)
{
    image->setPixelColor(x, y, color);
}

QColor MyImage::getPixel(int x, int y)
{
    return image->pixelColor(x, y);
}
QImage *MyImage::getImage()
{
    return image;
}
void MyImage::setLine(MyPoint start, MyPoint end)
{
    if (!image)
        return;
    QPainter painter(image);
    QPen pen(drawColor);
    painter.setPen(pen);
    painter.drawLine(QPoint(start.getX(), start.getY()), QPoint(end.getX(), end.getY()));
}
void MyImage::setRect(MyPoint start, MyPoint end)
{
    if (!image)
        return;
    QPainter painter(image);
    QPen pen(drawColor);
    painter.setPen(pen);
    painter.drawLine(QPoint(start.getX(), start.getY()), QPoint(start.getX(), end.getY()));
    painter.drawLine(QPoint(start.getX(), start.getY()), QPoint(end.getX(), start.getY()));
    painter.drawLine(QPoint(end.getX(), end.getY()), QPoint(end.getX(), start.getY()));
    painter.drawLine(QPoint(end.getX(), end.getY()), QPoint(start.getX(), end.getY()));
}
void MyImage::setPoint(MyPoint center)
{
    if (!image)
        return;
    image->setPixelColor(center.getX(), center.getY(), drawColor);
    //updateDisplay();
/*    QPainter painter(image);
    QPen pen(drawColor);
    painter.setPen(pen);
    painter.drawEllipse(center.getX(), center.getY(), 1, 1);
*/}

void MyImage::update(ISubject *subject)
{
    MyException *exp = new MyException();
    QColor newDraw;
    printf("update\n");
//QColor newFone, QColor newCatter, QColor newCatLine, QColor newLine
    newDraw = subject->getColor();
    /*if (newDraw.value() == foneColor.value())
    {
        subject->setNewColor(drawColor);
        exp->makeException(QString("Цвета должны различаться.\n"));
        return;
    }*/
    changeColors(newDraw);
    display->setImage(image);
}

MyLabel *MyImage::getDisplay()
{
    return display;
}

void MyImage::newMouseEvent(MyPoint event)
{
    setPoint(MyPoint(event.getX(), event.getY()));
    if(lastEvent)
    {
        setLine(MyPoint(lastEvent->getX(), lastEvent->getY()), MyPoint(event.getX(), event.getY()));
    }
    lastEvent = new MyPoint(event.getX(), event.getY());
}

void MyImage::clean()
{
    for (int i = 0; i < image->size().width(); i++)
    {
        for (int j = 0; j < image->size().height(); j++)   
        {
            image->setPixelColor(i, j, foneColor);
        }
    }
    display->setImage(image);
}

int MyImage::getWidth()
{
    return image->width();
}
int MyImage::getHeight()
{
    return image->height();
}