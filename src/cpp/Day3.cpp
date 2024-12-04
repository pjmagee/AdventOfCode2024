#include "Day3.h"
#include <regex>

std::string Day3::execute(std::string input) const {

   std::regex pattern(R"(mul\(\d+,\d+\)|do\(\)|don't\(\))");
   std::smatch matches;

   std::string::const_iterator searchStart( input.cbegin() );

   auto result = 0;
   bool isEnabled = true;

   while ( regex_search ( searchStart, input.cend(), matches, pattern) ) {
      for (auto match : matches) {
         if (match.str() == "do()") {
            isEnabled = true;
         } else if (match.str() == "don't()") {
            isEnabled = false;
         } else {
            if (isEnabled) {
               std::string mul = match.str();

               auto left = std::stoi(mul.substr(4, mul.find(',')));
               auto right = std::stoi(mul.substr(mul.find(',') + 1, mul.size() - 1));

               result += left * right;
            }
         }
      }
      searchStart = matches.suffix().first;
   }

   return std::to_string(result);
}