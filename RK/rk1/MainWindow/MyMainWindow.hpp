#ifndef MYMAINWINDOW_H
#define MYMAINWINDOW_H

#include "ui_main_window_ui.h" //включить файл сгенеренный QtDesiner

#include "../GraphicsView/MyGraphicsView.hpp"
#include <QMainWindow>
#include <QMouseEvent>
#include <QDoubleValidator>

class MyMainWindow : public QMainWindow
{
	Q_OBJECT
public:
	explicit MyMainWindow(QWidget *arent = 0);
	~MyMainWindow();
private:
	Ui::MainWindow *ui; // Сгенеренный QtDesiner класс, описывающий внешность нашего окна (кастомайзер)
	MyGraphicsView *myPicture;
	QDoubleValidator *myValidator;
	QLocale *myLocale;
	void setConnect();
	void setMyPicture();
	void setMyValidator();
	void setSplitters();
public slots:
	void slotChangeCenter(QMouseEvent *event);
	void slotTransfer();
	void slotRotate();
	void slotScale();
	void slotArhiveLast();
	void slotArhiveNext();
	void slotAtStart();
};

#endif
