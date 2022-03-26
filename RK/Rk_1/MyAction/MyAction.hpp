#ifndef MYACTION_HPP
#define MYACTION_HPP

#include <QAction>

class MyAction : public QAction
{
Q_OBJECT
public:
    MyAction(QObject *parent = nullptr);
    MyAction(QString name);
    ~MyAction();
public slots:
    void slotClicked();
signals:
    void selected(MyAction *it);
};

#endif