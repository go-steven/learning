// run it with:
//   g++.exe -c main.cpp ./log4cpp -o main.o -I"./include" -I"E:/Dev-Cpp/MinGW64/include" -I"E:/Dev-Cpp/MinGW64/x86_64-w64-mingw32/include" -I"E:/Dev-Cpp/MinGW64/lib/gcc/x86_64-w64-mingw32/4.8.1/include" -I"E:/Dev-Cpp/MinGW64/lib/gcc/x86_64-w64-mingw32/4.8.1/include/c++" 
//   g++.exe main.o -o main.exe -L"E:/Dev-Cpp/MinGW64/lib" -L"E:/Dev-Cpp/MinGW64/x86_64-w64-mingw32/lib" -static-libgcc
//   main.exe
#include <iostream>
#include "log4cpp/Category.hh"
#include "log4cpp/OstreamAppender.hh"
#include "log4cpp/BasicLayout.hh"
#include "log4cpp/Priority.hh"

using namespace std;

int main(int argc, char** argv) {
	log4cpp::OstreamAppender* osAppender = new log4cpp::OstreamAppender("osAppender", &cout);
	osAppender->setLayout(new log4cpp::BasicLayout());
	log4cpp::Category& root =log4cpp::Category::getRoot();
	root.addAppender(osAppender);
	root.setPriority(log4cpp::Priority::DEBUG);
	root.error("Hello log4cpp in aError Message!");
	log4cpp::Category::shutdown();   

	return 0;
}
