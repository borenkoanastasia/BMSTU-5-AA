#ifndef MYPLOT_H
#define MYPLOT_H

#include <cmath>

#include <QChart>
#include <QChartView>
#include <QLineSeries>
#include <QValueAxis>
#include <QGridLayout>
#include <QDialog>

#include "../MyPoint/MyPoint.hpp"

using namespace QtCharts;

class MyPlotView : public QChartView
{
    Q_OBJECT
    QVector<QLineSeries *>series;// = new QLineSeries();
    QChart *chart;
    void setDefaultSettings();

    void makePlot(QVector<MyPoint>points, QString name);
    //void makeSeries(QVector<MyPoint>points, QString name);
    void updateView();
public:
    explicit MyPlotView(QWidget *parent = nullptr);
    ~MyPlotView();
    void clear();
    void drawPlot(QVector<QVector<MyPoint>>points, QVector<QString>names, QString plotName);
    void addWindow(QString string);
    void setTitle(QString string = QString(tr("График зависимости времени отрисовки от размеров объекта")));
};

#endif