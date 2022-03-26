#include "matrix.hpp"
#include <vector>

struct matrix_t{
    std::vector<std::vector<double>> elem;
};

void print_message(){
    std::cout << '_'*111;
    std::cout << "МЕНЮ";
    std::cout << '_'*111;
    std::cout << "\t1. Ручное тестирование";
    std::cout << "\t2. Автотестирование";
    std::cout << "\t3. Выход";
    std::cout << "Ваш выбор:";
}

void manual_testing()
{
}

void auto_testing()
{}

int menu()
{
    int choise;
    for (;true;){
        print_message();
        std::cin >> choise;
        switch(choise){
            case 1:
                manual_testing();
            case 2:
                auto_testing();
            default:
                break;
        }
    };
}

int main(){
    return 0;
}