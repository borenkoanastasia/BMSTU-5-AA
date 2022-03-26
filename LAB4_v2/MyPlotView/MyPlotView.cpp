#include "MyPlotView.hpp"

MyPlotView::MyPlotView(QWidget *parent) : QChartView(parent), chart(new QChart(nullptr,  Qt::WindowFlags()))
{
    setDefaultSettings();
}

void MyPlotView::setDefaultSettings()
{
    //legend()->hide(); 
    setRenderHint(QPainter::Antialiasing);
}

MyPlotView::~MyPlotView()
{
    clear();
}

void MyPlotView::clear()
{
    chart->removeAllSeries();
    series.clear();
}

void MyPlotView::makePlot(QVector<MyPoint>points, QString name)
{
    QLineSeries *seria = new QLineSeries();
    for (int i = 0; i < points.size(); i++)
    {
        seria->append(points[i].getX(), points[i].getY());
    }

    seria->setName(name);
    chart->addSeries(seria);
    seria->setVisible(true);
    series.append(seria);
}

void MyPlotView::updateView()
{
    setChart(chart);
}

void MyPlotView::drawPlot(QVector<QVector<MyPoint>> points, QVector<QString> names, QString plotName)
{
    clear();
    for (int i = 0; i < points.size(); i++)
    {
        makePlot(points[i], names[i]);
    }
    chart->createDefaultAxes();
    updateView();
}

void MyPlotView::addWindow(QString string)
{
    QDialog *d = new QDialog();
    
    QGridLayout *mainLayout = new QGridLayout();
    mainLayout->addWidget(this);
    d->setLayout(mainLayout);

    d->setWindowTitle(string);
    d->show();
}

void MyPlotView::setTitle(QString string)
{
    chart->setTitle(string);   
}
