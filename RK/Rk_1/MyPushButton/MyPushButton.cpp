#include "MyPushButton.hpp"

MyPushButton::MyPushButton(QWidget *parent) : QPushButton(parent)
{
    setAutoFillBackground(true);
    curColor = Qt::black;
    
    QPalette palette = this->palette();
    palette.setColor(QPalette::Button, curColor);
    this->setPalette(palette);
}

MyPushButton::~MyPushButton()
{}

void MyPushButton::update(ISubject *subject)
{
    curColor = subject->getColor();

    QPalette palette = this->palette();
    palette.setColor(QPalette::Button, curColor);
    this->setPalette(palette);
}