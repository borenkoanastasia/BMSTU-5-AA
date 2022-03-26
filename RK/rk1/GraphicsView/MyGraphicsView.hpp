#ifndef MYGRAPHICVIEW_H
#define MYGRAPHICVIEW_H

#include <QGraphicsItemGroup>
#include <QGraphicsSceneWheelEvent>
#include <QGraphicsView>
#include <QTimer>



class MyGraphicsView: public QGraphicsView
{
	Q_OBJECT
public: 
	explicit MyGraphicsView(QWidget *parent = 0);
	~MyGraphicsView();


	//void mousePressEvent(QMouseEvent *event);
	//void wheelEvent(QGraphicsSceneWheelEvent *wheelEvent);

signals:
	void pressMouse(QMouseEvent *event);

private:

	QGraphicsScene *scene;

	QGraphicsItemGroup *builder;

	void resizeEvent(QResizeEvent *event);

	QTimer *timer;

	void deleteItemsFromGroup(QGraphicsItemGroup *group);

private slots:
	void slotAlarmTimer();  // слот для обработчика переполнения таймера в нём будет производиться перерисовка виджета
};

#endif