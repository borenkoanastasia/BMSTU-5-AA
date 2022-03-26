#include "MyColorInput.hpp"

void MyColorInput::addObject(IObject *obj)
{
    subscribers.append(obj);
}

void MyColorInput::removeObject(IObject *obj)
{
    subscribers.remove(subscribers.indexOf(obj));
}

void MyColorInput::update()
{
    for (int i = 0; i < subscribers.size(); i++)
    {
        subscribers[i]->update(this);
    }
}
void MyColorInput::setNewColor(QColor color)
{
    if(color.isValid())
    {
        newColor = color;
        update();
    }
}
QColor MyColorInput::getColor()
{
    return newColor;
}

void MyColorInput::getNewColor()
{
    QWidget *newWin = new QWidget();
    QColor color = QColorDialog::getColor(newColor, newWin);
    if( color.isValid() )
    {
        newColor = color;
        update();
    }
    delete newWin;
}

/*
Change ButtonColor
pButton->setAutoFillBackground(true);
QPalette palette = pButton->palette();
palette.setColor(QPalette::Window, QColor(Qt::blue));
pButton->setPalette(palette);
*/