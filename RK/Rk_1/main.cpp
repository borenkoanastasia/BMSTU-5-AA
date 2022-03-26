#include "MyMainWindow/MyMainWindow.hpp"
#include "MyController/MyController.hpp"
#include <QApplication>

int main(int argc, char ** argv)
{
	QApplication app(argc, argv);
	MyMainWindow *window = new MyMainWindow();
	MyController *controller = new MyController(window);
	//controller->drawCircleHistogram();
	//controller->drawEllipseHistogram();
	window->show();
	return app.exec();
}