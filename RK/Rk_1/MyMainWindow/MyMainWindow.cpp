#include "MyMainWindow.hpp"

MyMainWindow::MyMainWindow(QMainWindow *parent): QMainWindow(parent), 
    ui(new Ui::MainWindow()), locale(new MyLocale())
{
    ui->setupUi(this);
    setColorButton();
    setValidators();
}

MyMainWindow::~MyMainWindow()
{
    //delete FoneColorButton;
    delete DrawColorButton;
    delete ui;
}

void MyMainWindow::setValidators()
{
    ui->rotateParamX->setValidator(locale->getDoubleValidator());
    ui->rotateParamY->setValidator(locale->getDoubleValidator());
    ui->rotateParamZ->setValidator(locale->getDoubleValidator());


    ui->x_start->setValidator(locale->getDoubleValidator());
    ui->x_end->setValidator(locale->getDoubleValidator());
    ui->x_step->setValidator(locale->getDoubleValidator());
    ui->y_start->setValidator(locale->getDoubleValidator());
    ui->y_end->setValidator(locale->getDoubleValidator());
    ui->y_step->setValidator(locale->getDoubleValidator());
}
double MyMainWindow::getRotateX()
{
    return locale->toDouble(ui->rotateParamX->text());
}
double MyMainWindow::getRotateY()
{
    return locale->toDouble(ui->rotateParamY->text());
}
double MyMainWindow::getRotateZ()
{
    return locale->toDouble(ui->rotateParamZ->text());
}
double MyMainWindow::getXStart()
{
    return locale->toDouble(ui->x_start->text());
}
double MyMainWindow::getXEnd()
{
    return locale->toDouble(ui->x_end->text());
}
double MyMainWindow::getXStep()
{
    return locale->toDouble(ui->x_step->text());
}
double MyMainWindow::getZStart()
{
    return locale->toDouble(ui->y_start->text());
}
double MyMainWindow::getZEnd()
{
    return locale->toDouble(ui->y_end->text());
}
double MyMainWindow::getZStep()
{
    return locale->toDouble(ui->y_step->text());
}

QPushButton *MyMainWindow::getRotateButton()
{
    return ui->rotateButton;
}
QPushButton *MyMainWindow::getDrawButton()
{
    return ui->drawButton;
}
QPushButton *MyMainWindow::getClearButton()
{
    return ui->clearButton;
}
MyPushButton *MyMainWindow::getDrawColorButton()
{
    return DrawColorButton;
}

void MyMainWindow::setPicture(MyLabel *canvas)
{
    ui->canvasLayout->addWidget(canvas);
}

void MyMainWindow::setColorButton()
{
    DrawColorButton = new MyPushButton();
    ui->ColorLayout->addWidget(DrawColorButton);
}
void MyMainWindow::setAlgButton(QVector<QString> names)
{
    AlgorithmButton = new MyToolButton(names);
    ui->functionLayuot->addWidget(AlgorithmButton);
}
QString MyMainWindow::getAlg()
{
    return AlgorithmButton->defaultAction()->text();
}
