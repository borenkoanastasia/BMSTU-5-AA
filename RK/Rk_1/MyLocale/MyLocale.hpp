#ifndef MYVALIDATOR_H
#define MYVALIDATOR_H_H

#include <QLocale>
#include <QMainWindow>
#include <QValidator>

class MyLocale : public QLocale
{
private:
    QIntValidator *intValidator;
    QDoubleValidator *doubleValidator;
public:
    explicit MyLocale(QLocale *parent = 0);
    ~MyLocale();
    void setValidator();
    
    QIntValidator *getIntValidator();
    QDoubleValidator *getDoubleValidator();
};

#endif