#include "MyLocale.hpp"

MyLocale::MyLocale(QLocale *parent)
{
    setValidator();
}

MyLocale::~MyLocale()
{
    delete intValidator;
    delete doubleValidator;
}

void MyLocale::setValidator()
{
	intValidator = new QIntValidator();
    doubleValidator = new QDoubleValidator();
	doubleValidator->setNotation(QDoubleValidator::StandardNotation);
	doubleValidator->setLocale(*this);
}

QIntValidator *MyLocale::getIntValidator()
{
    return intValidator;
}

QDoubleValidator *MyLocale::getDoubleValidator()
{
    return doubleValidator;
}