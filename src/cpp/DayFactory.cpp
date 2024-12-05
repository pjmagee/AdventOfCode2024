#include "DayFactory.h"
#include "Day.h"
#include "Day1.h"
#include "Day2.h"
#include "Day3.h"

std::unique_ptr<Day> DayFactory::createDay(const int arg)
{
    switch (arg)
    {
    case 1:
        return std::make_unique<Day1>();
    case 2:
        return std::make_unique<Day2>();
    case 3:
        return std::make_unique<Day3>();
    default:
        return nullptr;
    }
}
