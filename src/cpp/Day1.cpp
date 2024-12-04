#include "Day1.h"
#include <string>
#include <sstream>
#include <vector>
#include <iostream>
#include <nlohmann/json.hpp>

using nlohmann::json;

std::string Day1::execute(std::string input) const {

    std::vector<int> leftList;
    std::vector<int> rightList;

    std::istringstream iss(input);
    std::string line;

    while (std::getline(iss, line)) {
        std::istringstream lineStream(line);
        int left, right;
        if (lineStream >> left >> right) {
            leftList.push_back(left);
            rightList.push_back(right);
        } else {
            std::cerr << "Invalid line: " << line << std::endl;
        }
    }

    std::ranges::sort(leftList);
    std::ranges::sort(rightList);

    auto totalDistance = 0;

    for (size_t i = 0; i < leftList.size(); i++) {
        totalDistance += std::abs(leftList[i] - rightList[i]);
    }

    auto similarityScore = 0;

    for (int l : leftList) {
        auto leftInRight = 0;
        for (int r : rightList) {
            if (l == r) {
                leftInRight++;
            }
        }

        similarityScore += l * leftInRight;
    }

    json resp;
    resp["similarity_score"] = similarityScore;
    resp["total_distance"] = totalDistance;

    return resp.dump();
}