#ifndef DAY_H
#define DAY_H

#include <string>

class Day {
public:
    virtual ~Day() = default;
    [[nodiscard]] virtual std::string execute(std::string input) const = 0;
};

#endif // DAY_H
