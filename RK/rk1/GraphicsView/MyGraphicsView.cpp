#include "MyGraphicsView.hpp"


MyGraphicsView::MyGraphicsView(QWidget *parent): QGraphicsView(parent)
{
	scene = new QGraphicsScene();
	this->setScene(scene);

	this->setHorizontalScrollBarPolicy(Qt::ScrollBarAlwaysOff);
	this->setVerticalScrollBarPolicy(Qt::ScrollBarAlwaysOff);  
	this->setAlignment(Qt::AlignCenter);
	this->setSizePolicy(QSizePolicy::Expanding, QSizePolicy::Expanding);
	this->setMinimumHeight(100);
	this->setMinimumWidth(100);

	builder = new QGraphicsItemGroup();
	scene->addItem(builder);


	timer = new QTimer();			   // Инициализируем Таймер
	timer->setSingleShot(true);


	connect(timer, SIGNAL(timeout()), this, SLOT(slotAlarmTimer()));
	timer->start(50);	
}

void MyGraphicsView::slotAlarmTimer()
{
	deleteItemsFromGroup(builder);

	QSize cur_size = this->size();
	int width = cur_size.width();
	int height = cur_size.height();



	scene->setSceneRect(0,0,width,height);
}

void MyGraphicsView::resizeEvent(QResizeEvent *event)
{
	timer->start(50);   // Как только событие произошло стартуем таймер для отрисовки
	QGraphicsView::resizeEvent(event);  // Запускаем событие родителького класса
}


MyGraphicsView::~MyGraphicsView()
{}
/*
void MyGraphicsView::
{}
*/
void MyGraphicsView::deleteItemsFromGroup(QGraphicsItemGroup *group)
{

    scene->removeItem(group);
	/*foreach( QGraphicsItem *item, group) {
	   if(item->group() == group ) {
		  delete item;
	   }
	}*/
	//scene->destroyItemGroup(group);
	/*foreach( QGraphicsItem *item, scene->items(group->boundingRect())) {
	   if(item->group() == group ) {
		  delete item;
	   }
	}*/
	/*)
	delete group;
	group = new QGraphicsItemGroup();
	scene->addItem(group);*/
}
/*
void MyGraphicsView::mousePressEvent(QMouseEvent *event)
{
	emit pressMouse(event);
}*/
/*
void MyGraphicsView::restartFish()
{
	//DrawFish::MyPoint p = DrawFish::MyPoint(0, 0);
	timer->start(0);
}*/

