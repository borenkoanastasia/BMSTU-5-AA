#ifndef MYMAINWINDOW_H
#define MYMAINWINDOW_H

#include <QMainWindow>
#include <QRadioButton>
#include <QVBoxLayout>
#include "ui_main_window_ui.h"
#include "../MyPushButton/MyPushButton.hpp"
#include "../MyToolButton/MyToolButton.hpp"
#include "../MyLabel/MyLabel.hpp"
#include "../MyLocale/MyLocale.hpp"

class MyMainWindow : public QMainWindow
{
private:
    Ui::MainWindow *ui;
    MyLocale *locale;
    MyPushButton *DrawColorButton;
    MyToolButton *AlgorithmButton;

    void setValidators();
    //void setSplitters();
    //void addWidgets();
    //void constractWidgets();
    void setConnects();
public:
    explicit MyMainWindow(QMainWindow *parent = 0);
    ~MyMainWindow();

    QPushButton *getRotateButton();
    QPushButton *getDrawButton();
    QPushButton *getClearButton();
    
    MyPushButton *getDrawColorButton();

    double getRotateX();
    double getRotateY();
    double getRotateZ();
    double getXStart();
    double getXEnd();
    double getXStep();
    double getZStart();
    double getZEnd();
    double getZStep();

    void setPicture(MyLabel *canvas);
    void setAlgButton(QVector<QString> names);
    void setColorButton();

    QString getAlg();
};

#endif