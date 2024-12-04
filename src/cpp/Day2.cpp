#include "Day2.h"
#include <iostream>
#include <sstream>
#include <utility>
#include <vector>

std::string Day2::execute(std::string input) const
{
    std::vector<Report> reports;
    std::istringstream stream(input);
    std::string line;
    while (std::getline(stream, line))
    {
        reports.push_back(Report(line));
    }

    int safeReports = 0;

    for (const Report& report : reports)
    {
        if (report.isSafe())
        {
            safeReports++;
        }
        else
        {
            for (int level = 0; level < report.levels.size(); level++)
            {
                std::vector<int> newLevels;
                for (int i = 0; i < report.levels.size(); i++)
                {
                    if (i != level)
                        newLevels.push_back(report.levels[i]);
                }

                Report adjusted(newLevels);

                if (adjusted.isSafe())
                {
                    safeReports++;
                    break;
                }
            }
        }
    }

    return std::to_string(safeReports);
}

Report::Report(std::string input)
{
    std::vector<int> numbers;
    std::istringstream lineStream(input);
    int number;
    while (lineStream >> number)
    {
        numbers.push_back(number);
    }
    this->levels = numbers;
}

Report::Report(std::vector<int> levels)
{
    this->levels = std::move(levels);
}

bool Report::isSafe() const
{
    bool increasing = true;
    bool decreasing = true;

    for (int level = 1; level < levels.size(); level++)
    {
        auto diff = levels[level] - levels[level - 1];
        if (diff < 0) increasing = false;
        if (diff > 0) decreasing = false;
    }

    bool withinRange = true;

    for (int level = 1; level < levels.size(); level++)
    {
        auto diff = abs(levels[level] - levels[level - 1]);

        if (diff < 1 || diff > 3)
        {
            withinRange = false;
            break;
        }
    }

    return (increasing || decreasing) && withinRange;
}
