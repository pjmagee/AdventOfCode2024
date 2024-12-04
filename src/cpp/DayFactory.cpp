#include "DayFactory.h"
#include "Day.h"
#include "Day1.h"
#include "Day2.h"
#include "Day3.h"

std::unique_ptr<Day> DayFactory::createDay(const std::string& arg)
{
    if (arg == "1")
    {
        return std::make_unique<Day1>();
    }

    if (arg == "2")
    {
        return std::make_unique<Day2>();
    }

    if (arg == "3")
    {
        return std::make_unique<Day3>();
    }

    return nullptr;
}
