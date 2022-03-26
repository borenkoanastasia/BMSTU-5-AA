#include "MyAction.hpp"

MyAction::MyAction(QObject *parent): QAction(parent)
{
    connect(this, SIGNAL(clicked),this,SLOT(slotClicked()));
}
MyAction::MyAction(QString name): QAction(name)
{
    connect(this, SIGNAL(triggered()),this,SLOT(slotClicked()));
}

MyAction::~MyAction()
{}

void MyAction::slotClicked()
{
    emit selected(this);
}
