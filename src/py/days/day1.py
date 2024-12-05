import json

from days.day import Day


class Day1(Day):

    def process_input(self, input_data: str) -> str:

        lines: list[str] = input_data.splitlines()
        left_list: list[int] = []
        right_list: list[int] = []

        for line in lines:
            numbers: list[str] = line.split()
            if len(numbers) > 0:
                left_list.append(int(numbers[0]))
                right_list.append(int(numbers[1]))

        right_list.sort()
        left_list.sort()

        total_distance: int = 0
        total_similarity: float = 0

        for i in range(len(left_list)):
            total_distance += abs(left_list[i] - right_list[i])

        for left in left_list:
            found: int = 0
            for right in right_list:
                if left == right:
                    found += 1
            total_similarity += left * found

        data = {
            "part_one": total_distance,
            "part_two": total_similarity
        }

        return json.dumps(data)