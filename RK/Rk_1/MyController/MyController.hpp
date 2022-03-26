#ifndef MYCONTROLLER_HPP
#define MYCONTROLLER_HPP

#include <cmath>
#include <time.h>
#include <QPicture>
#include <QPixmap>
#include <QApplication>
#include <QtGlobal>
#include "../MyAlgorithms/MyAlgorithmFloatingHorizont.hpp"
#include "../MyMainWindow/MyMainWindow.hpp"
#include "../MyColorInput/MyColorInput.hpp"
#include "../MyImage/MyImage.hpp"

class MyController : QObject
{
Q_OBJECT
private:
    MyMainWindow *mainWindow;
    MyAlgorithmShortedEdges *curAlg;
    MyToolButton *algorithmButton;

    MyColorInput *inputDrawColor;

    MyImage *image;
    Matrix current_matrix;
    QVector<Line> line;

    void setConnections();
    void setLabel();
public:
    MyController(MyMainWindow *main);
    ~MyController();

public slots:

    void draw();
    void rotate();
    void clear();

    void getDrawColor();
};

#endif