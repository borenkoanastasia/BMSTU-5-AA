#include "matrix.hpp"
#include "MyPlotView/MyPlotView.hpp"
#include <QVector>
#include <QString>
#include <QApplication>

std::mutex MuTeX;


void print_message(){
    //std::cout << '_'*111 << std::endl;
    std::cout << "МЕНЮ" << std::endl;
    //std::cout << '_'*111 << std::endl;
    std::cout << "\t1. Ручное тестирование\n";
    std::cout << "\t2. Автотестирование\n";
    std::cout << "\t3. График\n";
    std::cout << "\t4. Выход\n";
    std::cout << "Ваш выбор: ";
}

void manual_testing()
{
    matrix_t m1;
    input_matrix(m1);
    if (m1.rows < 1 || m1.columns < 1){
        std::cout << "В матрице нет элементов. Нельзя расчитать среднее арифметическое.\n";
        return;
    }
    double res;
    standart_average(m1, res);

    control_thread(m1, parall_row_average, 4);

    std::cout << "Ответ: \nСтандартный алгоритм: " << res << std::endl;
    std::cout << "Распараллеленый по строкам (4 потока): " << control_thread(m1,parall_row_average, 4) << std::endl;
    std::cout << "Распараллеленый по столбцам (4 потока): " << control_thread(m1,parall_column_average, 4) << std::endl;
    output_matrix(m1);
}

std::chrono::duration<double> get_standart_time(int &rows, int &columns){
    std::chrono::duration<double> dif;
    matrix_t m = make_matrix(rows, columns);
    //clock_t start, end, dif = 0;
    double res = 0;
    for (int i = 0; i < REPEATS; i++){
        auto start = std::chrono::system_clock::now();
        standart_average(m, res);
        auto end = std::chrono::system_clock::now();
        dif += end - start;
    }
    dif /= REPEATS;
    return dif;
}

std::chrono::duration<double>   get_par_rows_time(int &rows, int &columns, int thread_count){
    std::chrono::duration<double> dif;


    matrix_t m = make_matrix(rows, columns);
    //clock_t start, end, dif = 0;
    double res = 0;
    for (int i = 0; i < REPEATS; i++){
        auto start = std::chrono::system_clock::now();
        control_thread(m,parall_row_average, thread_count);
        auto end = std::chrono::system_clock::now();
        dif += end - start;
    }
    dif /= REPEATS;
    return dif;
}
std::chrono::duration<double>  get_par_columns_time(int &rows, int &columns, int thread_count){

    
    std::chrono::duration<double> dif;


    matrix_t m = make_matrix(rows, columns);
    //clock_t start, end, dif = 0;
    double res = 0;
    for (int i = 0; i < REPEATS; i++){
        auto start = std::chrono::system_clock::now();
        control_thread(m,parall_column_average, thread_count);
        auto end = std::chrono::system_clock::now();
        dif += end - start;
    }
    dif /= REPEATS;
    return dif;
}

void auto_testing()
{

    for (int i = 0; i < 20; i++){
        get_standart_time(i, i);
    }
    int one = 1;
    printf("Таблица (в сек) (Выбор %0.0f)\n", get_par_rows_time(one, one, one).count()+2);

    printf("%13s|%10s", "Длина", "stand");
    for (int j = 1; j < MAX_THREADS; j*=2){
        printf("|%-5d%8s", j, "стр.");
        printf("|%-5d%9s", j, "стол.");
    }
    printf("\n");
    //std::cout << std::endl;
    for (int i = 100; i < MAX_ROWS; i+=100){

        printf("%8d|%10f", i, get_standart_time(i, i).count());
        for (int j = 1; j < MAX_THREADS; j*=2){
            auto t1 = get_par_rows_time(i, i, j).count();
            printf("|%10f", get_par_rows_time(i, i, j).count());
            printf("|%10f", get_par_columns_time(i, i, j).count());
        }
        printf("\n");
    };
}

void graph_testing(int argc, char ** argv)
{
	QApplication app(argc, argv);
    auto graph = new MyPlotView(); 

    MyPoint p, p1r, p1c, p2r, p2c, p4r, p4c, p8r, p8c, p16r, p16c, p32r, p32c;

    for (int i = 0; i < 20; i++){
        get_standart_time(i, i);
    }
    //std::cout << std::endl;
    QVector<MyPoint> v, v1r, v1c, v2r, v2c, v4r, v4c, v8r, v8c, v16r, v16c, v32r, v32c;
    QVector<QVector<MyPoint>> res;
    for (int i = 100; i < MAX_ROWS; i+=100){
        p.setX(i);
        p1r.setX(i);
        //p1c.setX(i);
        p2r.setX(i);
        //p2c.setX(i);
        p4r.setX(i);
        //p4c.setX(i);
        p8r.setX(i);
        //p8c.setX(i);
        p16r.setX(i);
        //p16c.setX(i);
        p32r.setX(i);
        //p32c.setX(i);

        p.setY(get_standart_time(i, i).count());
        p1r.setY(get_par_rows_time(i, i, 1).count());
        //p1c.setY(get_par_columns_time(i, i, 1).count());
        p2r.setY(get_par_rows_time(i, i, 2).count());
        //p2c.setY(get_par_columns_time(i, i, 2).count()i);
        p4r.setY(get_par_rows_time(i, i, 4).count());
        //p4c.setY(get_par_columns_time(i, i, 4).count());
        p8r.setY(get_par_rows_time(i, i, 8).count());
        //p8c.setY(get_par_columns_time(i, i, 8).count());
        p16r.setY(get_par_rows_time(i, i, 16).count());
        //p16c.setY(get_par_columns_time(i, i, 16).count());
        p32r.setY(get_par_rows_time(i, i, 32).count());
        //p32c.setY(get_par_columns_time(i, i, 32).count());

        v.append(p);
        v1r.append(p1r);
        //v1c.append(p1c);
        v2r.append(p2r);
        //v2c.append(p2c);
        v4r.append(p4r);
        //v4c.append(p4c);
        v8r.append(p8r);
        //v8c.append(p8c);
        v16r.append(p16r);
        //v16c.append(p16c);
        v32r.append(p32r);
       // v32c.append(p32c);
    };
    res.append(v);
    res.append(v1r);
    //res.append(v1c);
    res.append(v2r);
    //res.append(v2c);
    res.append(v4r);
    //res.append(v4c);
    res.append(v8r);
    //res.append(v8c);
    res.append(v16r);
    //res.append(v16c);
    res.append(v32r);
    //res.append(v32c);
    QVector<QString>names;

    names.append(QString(MyPlotView::tr("Станд. алгоритм")));
    names.append(QString(MyPlotView::tr("1 поток строки")));
    //names.append(QString(MyPlotView::tr("1 поток столбцы")));
    names.append(QString(MyPlotView::tr("2 потока строки")));
    //names.append(QString(MyPlotView::tr("2 потока столбцы")));
    names.append(QString(MyPlotView::tr("4 потока строки")));
    //names.append(QString(MyPlotView::tr("4 потока столбцы")));
    names.append(QString(MyPlotView::tr("8 потоков строки")));
    //names.append(QString(MyPlotView::tr("8 потоков столбцы")));
    names.append(QString(MyPlotView::tr("16 потоков строки")));
    //names.append(QString(MyPlotView::tr("16 потоков стол.")));
    names.append(QString(MyPlotView::tr("32 потока строки")));
    //names.append(QString(MyPlotView::tr("32п. стол.")));

    graph->drawPlot(res, names, QString(MyPlotView::tr("Время")));
    graph->addWindow("");
    app.exec();
}


// Очищение stdin

void clearInputBuf(void) 
{ 
      int garbageCollector; 
      while ((garbageCollector = getchar()) != '\n' && garbageCollector != EOF) 
      {}
}


void menu(int argc, char ** argv)
{
    int choise;
    for (;true;){
        print_message();
        std::cin >> choise;
        switch(choise){
            case 1:
                manual_testing();
                break;
            case 2:
                auto_testing();
                break;
            case 3:
                graph_testing(argc, argv);
                break;
            default:
                return;
        }
        clearInputBuf();
    };
}

int main(int argc, char ** argv){
    menu(argc, argv);
    return 0;
}