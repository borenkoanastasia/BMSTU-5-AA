#include "MyController.hpp"


MyController::MyController(MyMainWindow *main): mainWindow(main), current_matrix(Matrix(M3D, M3D)),
    inputDrawColor(new MyColorInput()),// inputFoneColor(new MyColorInput), 
    image(new MyImage)//, table(new MyTableModel()), Catter(new MyShape())
{
    curAlg = new MyAlgorithmShortedEdges();
    setConnections();
    setLabel();
    current_matrix.make_single_matrix();
    current_matrix.output();
    //mainWindow->setTable(table);
    mainWindow->setAlgButton(curAlg->getNames());
}

MyController::~MyController()
{
    delete inputDrawColor;
    //delete inputFoneColor;
    delete mainWindow;
}

void MyController::draw()
{
    printf("start\n");
    image->clean();
    printf("clear\n");
    QString funcName = mainWindow->getAlg();
    printf("get alg\n");
    parameters param = make_param(mainWindow->getXStart(), mainWindow->getXEnd(), mainWindow->getXStep(), 
                                mainWindow->getZStart(), mainWindow->getZEnd(), mainWindow->getZStep());
    printf("make param\n");
    curAlg->draw(image, funcName, current_matrix, param);
    printf("draw end\n");
    image->updateDisplay();
    printf("update image\n");
}
void MyController::rotate()
{
    QString funcName = mainWindow->getAlg();
    Point p = get_rotate_point(mainWindow->getRotateX(), mainWindow->getRotateY(), mainWindow->getRotateZ());
    current_matrix = curAlg->rotate(current_matrix, p);
    current_matrix.output();
    draw();
}

void MyController::clear()
{
    current_matrix.make_single_matrix();
    image->clean();
}

void MyController::setLabel()
{
    mainWindow->setPicture(image->getDisplay());
}

void MyController::setConnections()
{
    connect(mainWindow->getDrawColorButton(), SIGNAL(clicked()),this, SLOT(getDrawColor()));
    inputDrawColor->addObject(mainWindow->getDrawColorButton());
    inputDrawColor->addObject(image);
    inputDrawColor->setNewColor(Qt::darkRed);

    connect(mainWindow->getDrawButton(), SIGNAL(clicked()), this, SLOT(draw()));
    connect(mainWindow->getRotateButton(), SIGNAL(clicked()), this, SLOT(rotate()));
    connect(mainWindow->getClearButton(), SIGNAL(clicked()), this, SLOT(clear()));
}
/*
void MyController::getFoneColor()
{
    inputFoneColor->getNewColor();
}*/

void MyController::getDrawColor()
{
    inputDrawColor->getNewColor();
}
