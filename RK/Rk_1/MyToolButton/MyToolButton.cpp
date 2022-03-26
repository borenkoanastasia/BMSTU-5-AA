#include "MyToolButton.hpp"

MyToolButton::MyToolButton(QVector<QString> names, QWidget *parent) : QToolButton(parent), menu(new QMenu())
{
    makeButton(names);
}

MyToolButton::~MyToolButton()
{}

void MyToolButton::setActions(QVector<QString> names)
{
    for (int i = 0; i < names.size();i++)
    {
        //printf("i = %d\n", i);
        actions.append(new MyAction(names[i]));
    }
    connectActions();
}

void MyToolButton::makeMenu()
{
    for (int i = 0; i < actions.size();i++)
    {
        //printf("i = %d\n", i);
        menu->addAction(actions[i]);
    }
    //printf("finish create menu ToolButton\n");
}
void MyToolButton::makeButton(QVector<QString> names)
{
    setActions(names);
    makeMenu();
    setPopupMode(QToolButton::MenuButtonPopup);
    setMenu(menu);
    //printf("actions size = %d\n", actions.size());
    setDefaultAction(actions[0]);
    //printf("finish create QToolButton\n");
}

void MyToolButton::connectActions()
{
    for (int i = 0; i < actions.size(); i++)
    {
        connect(actions[i], SIGNAL(selected(MyAction *)), this, 
            SLOT(slotDefaultActionChange(MyAction *)));
    }
}
void MyToolButton::slotDefaultActionChange(MyAction *action)
{
    QString str1 = action->text();
    QByteArray ba = str1.toLocal8Bit();
    const char *c_str2 = ba.data();
    //printf("%s\n", c_str2);
    for (int i = 0; i < actions.size(); i++)
    {
        if(actions[i]->text() == action->text())
        {
            setDefaultAction(action);
        }
    }
}
