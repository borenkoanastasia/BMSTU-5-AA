#include "MainWindow/MyMainWindow.hpp"

int main(int argc, char **argv)
{
	QApplication app(argc, argv); // Обработчик событий - создание
	MyMainWindow *window = new MyMainWindow; // Создаем элемент нашего класса (т.е. наше главное окно)
	window->show(); // Показываем наше главное окно
	return app.exec(); // Обработчик событий - запуск*/
}
