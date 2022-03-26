#include "MyException.hpp"

MyException::~MyException()
{}

void MyException::makeException(QString error)
{
    label = new QLabel(error);
    Ok = new QPushButton(tr("Ок"));
    Exception = new QWidget();
    connect(Ok, SIGNAL(clicked()), this, SLOT(delete_this()));

    QGridLayout *mainLayout = new QGridLayout;
    mainLayout->addWidget(label);
    mainLayout->addWidget(Ok);
    setLayout(mainLayout);

    setWindowTitle(tr("Сообщение об ошибке"));
    show();
}

void MyException::delete_this()
{
    setVisible(false);
}