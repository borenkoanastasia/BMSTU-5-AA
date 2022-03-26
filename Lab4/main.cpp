#include <iostream>
#include <chrono>
#include <thread>
#include <mutex>
#include <vector>
 
#include "matrix.hpp"

std::mutex g_lock;

struct matrix{
     std::vector<float> elems;
     int rows;
     int columns;
};

void threadFunction()
{
     g_lock.lock();
 
     std::cout << "entered thread " << std::this_thread::get_id() << std::endl;
     std::this_thread::sleep_for(std::chrono::seconds(rand()%10));
     std::cout << "leaving thread " << std::this_thread::get_id() << std::endl;
 
     g_lock.unlock();
}

int main()
{
     matrix_t m = makeTestMatrix1();
     output_matrix(m);
     makeRightTriangleMatrix(m);
     std::cout << "\n";
     output_matrix(m);
     return 0;
}