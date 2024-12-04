#include <iostream>
#include <string>
#include <sstream>
#include "DayFactory.h"

int main(int argc, char* argv[] ) {

    if (argc < 2) {
        std::cout << "No day number provided" << std::endl;
        return 1;
    }

    std::ostringstream buffer;
    buffer << std::cin.rdbuf();
    std::string input = buffer.str();

    const std::string dayNumber = argv[1];

    if (const std::unique_ptr<Day> day = DayFactory::createDay(dayNumber); day != nullptr) {
        std::cout << day->execute(input) << std::endl;
    } else {
        std::cout << "Invalid day number" << std::endl;
    }

    return 0;
}
