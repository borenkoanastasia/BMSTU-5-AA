#ifndef MYTOOLBUTTON_HPP
#define MYTOOLBUTTON_HPP

#include <QToolButton>
#include <QAction>
#include <QMenu>
#include <stdio.h>
#include "../MyAction/MyAction.hpp"

class MyToolButton : public QToolButton
{
Q_OBJECT
    QVector <MyAction *> actions;
    QMenu *menu;
    void setActions(QVector<QString> names);
    void makeMenu();
    void makeButton(QVector<QString> names);
    void clearActions();
public:
    explicit MyToolButton(QVector<QString> names, QWidget *parent = nullptr);
    ~MyToolButton();
    void connectActions();
public slots:
    void slotDefaultActionChange(MyAction *action);
};

#endif