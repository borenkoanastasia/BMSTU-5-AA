#ifndef MYEXCEPTION_HPP
#define MYEXCEPTION_HPP

#include <QtGlobal>
#include <QLabel>
#include <QPushButton>
#include <QDialog>
#include <QGridLayout>
#include <QDialog>

class MyException : public QDialog
{
Q_OBJECT
public slots:
    void delete_this();
public:
    ~MyException();
    void makeException(QString error);
private:

    QLabel *label;
    QPushButton *Ok;
    QWidget *Exception;
};

#endif