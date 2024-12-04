#ifndef DAY2_H
#define DAY2_H

#include <vector>
#include "Day.h"

class Day2 final : public Day {

public:
    [[nodiscard]] std::string execute(std::string input) const override;

};

class Report final {
public:
    explicit Report(std::string input);
    explicit Report(std::vector<int> levels);
    [[nodiscard]] bool isSafe() const;
    std::vector<int> levels;
};

#endif
