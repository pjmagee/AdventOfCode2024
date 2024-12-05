#ifndef DAYFACTORY_H
#define DAYFACTORY_H

#include "Day.h"
#include <memory>
#include <string>

class DayFactory {
public:
    static std::unique_ptr<Day> createDay(int arg);
};

#endif // DAYFACTORY_H
