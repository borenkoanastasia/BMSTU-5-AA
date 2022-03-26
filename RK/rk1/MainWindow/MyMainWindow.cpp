#include "MyMainWindow.hpp"

MyMainWindow::MyMainWindow(QWidget *parent) : QMainWindow(parent), ui(new Ui::MainWindow), myPicture(new MyGraphicsView)
//                         В header'e мы задали parent = 0 по умолчанию
//                                            Вызываем конструктор родителя (он инициализирует свойства родителя)
//                                                                 инициализируем переменную ui как укаатель на объект Ui::MainWindow
{
	ui->setupUi(this); // ui содает виджеты на нашем окне

	//setMyValidator();
	setMyPicture();
	setConnect();
	setSplitters();

	//ui->tabWidget->setTabVisible(0, true);

	// Инициализируем все наши виджеты
}

void MyMainWindow::setSplitters()
{
	/*QList<int> s1;
	s1.push_back(20);
	s1.push_back(1000);
	ui->splitter_3->setSizes(s1);

	QList<int> s2;
	s2.push_back(20);
	s2.push_back(1000);
	ui->splitter->setSizes(s2);
	QList<int> s3;
	s3.push_back(1000);
	s3.push_back(10);
	ui->splitter_2->setSizes(s3);*/
}


void MyMainWindow::setConnect()
{
	//connect(myPicture, SIGNAL(pressMouse(QMouseEvent *)), this, SLOT(slotChangeCenter(QMouseEvent *)));
}


void MyMainWindow::setMyPicture()
{
	ui->graphic_layout->addWidget(myPicture);
}
/*
void MyMainWindow::setMyValidator()
{
	myValidator = new QDoubleValidator();
	myValidator->setNotation(QDoubleValidator::StandardNotation);
	myLocale = new QLocale();
	QLocale::setDefault(*myLocale);
	myValidator->setLocale(*myLocale);
	
}
*/
MyMainWindow::~MyMainWindow()
{
	delete ui;
	delete myPicture;
	delete myValidator;
	delete myLocale;
	// Удаляем все наши виджеты
}

void MyMainWindow::slotChangeCenter(QMouseEvent *event)
{
	QPoint position = event->pos();
}
void MyMainWindow::slotTransfer()
{

//	myPicture->Fish.transfer(matrix);

	//myPicture->drawFish();
	/*
	transferToCenter();
	qreal x = myLocale->toDouble(ui->transform_x->text());
	qreal y = myLocale->toDouble(ui->transform_y->text());

	matrix_t matrix = makeMatrix(x, y, TRANSFER);
	myPicture->Fish.transfer(matrix);
	transferOutCenter();
  
	myPicture->drawFish();*/
}
void MyMainWindow::slotRotate()
{

	/*
	transferToCenter();
	qreal x = myLocale->toDouble(ui->rotate_angle->text());
	qreal y = 0;

	matrix_t matrix = makeMatrix(x, y, ROTATE);
	myPicture->Fish.transfer(matrix);
	transferOutCenter();
	myPicture->drawFish();*/
}

void MyMainWindow::slotScale()
{
}


void MyMainWindow::slotArhiveLast()
{
}

void MyMainWindow::slotArhiveNext()
{
}

void MyMainWindow::slotAtStart()
{
	//myPicture->restartFish();
}