#ifndef DAY1_H
#define DAY1_H

#include "Day.h"
#include <string>

class Day1 final : public Day {
public:
    [[nodiscard]] std::string execute(std::string input) const override;
};

#endif
