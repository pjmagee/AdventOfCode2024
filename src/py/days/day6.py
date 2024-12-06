import asyncio
import json
import sys
from copy import deepcopy

from days.day import Day

sample_data = """
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."""


class Day6(Day):
    guard_position: tuple[int, int] = (-1, -1)
    grid: list[list[str]] = []
    valid_positions: list[str] = ['^', 'v', '<', '>', 'X']
    next_direction: dict[str, str] = {'^': '>', '>': 'v', 'v': '<', '<': '^'}
    forward_x: dict[str, int] = {'^': 0, '>': 1, 'v': 0, '<': -1}
    forward_y: dict[str, int] = {'^': -1, '>': 0, 'v': 1, '<': 0}
    positions: list[str] = ['^', 'v', '<', '>']
    obstruction: str = '⚙️'
    obstructions: str = [obstruction, '#']
    entries: dict[tuple[int, int], dict[str, int]] = {}  # y,x -> {direction: count}

    def process_input(self, input_data: str) -> str:
        resp = {
            "part_one": self.part_one(input_data),
            "part_two": self.part_two(input_data),
        }
        return json.dumps(resp)

    def predict_part_one(self):
        while True:
            if self._is_path_blocked():
                self._turn_right()
            else:
                out_of_bounds = self._move_forward()
                if out_of_bounds:
                    return

    def part_two(self, input_data: str) -> int:
        grid = self.create_grid(input_data)
        self.entries = {}
        obstacles: list[tuple[int, int]] = []

        for y in range(len(grid)):
            for x in range(len(grid[0])):
                if (y, x) in obstacles:
                    continue
                else:
                    cell = grid[y][x]
                    cell_blocked = cell in self.positions or cell == '#'
                    if cell_blocked:
                        continue
                    else:
                        self.grid = self.create_grid(input_data)
                        self.grid[y][x] = self.obstructions[0]
                        self._set_starting_position()
                        self.init_entries()

                        while True:
                            if self._is_path_blocked():
                                self._turn_right()
                            else:
                                out_of_bounds = self._move_forward()
                                if out_of_bounds:
                                    break
                                if self._stuck_in_loop():
                                    obstacles.append((y, x))
                                    # print(f"Obstacle at {y},{x} caused the guard to get stuck in a loop")
                                    break
        return len(obstacles)

    def part_one(self, input_data: str) -> int:
        self.grid = self.create_grid(input_data)
        self._set_starting_position()
        self.predict_part_one()
        distinct_positions = 0
        for row in self.grid:
            for cell in row:
                if cell in self.valid_positions:
                    distinct_positions += 1
        # self.print_grid()
        return distinct_positions

    @staticmethod
    def create_grid(input_data) -> list[list[str]]:
        grid: list[list[str]] = []
        lines: list[str] = input_data.splitlines()
        for line in lines:
            line = line.strip()
            if line == '':
                continue
            grid.append(list(line))
        return grid

    def print_grid(self):

        for row in self.grid:
            for cell in row:
                if cell in self.positions:
                    # make GREEN
                    print(f"\033[92m{cell}\033[0m", end='')
                elif cell == 'X':
                    # make ORANGE
                    print(f"\033[93m{cell}\033[0m", end='')
                elif cell == '#':
                    # make RED
                    print(f"\033[91m{cell}\033[0m", end='')
                else:
                    print(cell, end='')
            print()
        print()

    def _turn_right(self) -> None:
        (y, x) = self.guard_position
        self.grid[y][x] = self.next_direction[self.grid[y][x]]

    def _is_path_blocked(self) -> bool:
        y, x = self.guard_position
        direction = self.grid[y][x]
        next_x = x + self.forward_x[direction]
        next_y = y + self.forward_y[direction]

        if next_y < 0 or next_y >= len(self.grid):
            return False
        if next_x < 0 or next_x >= len(self.grid[0]):
            return False

        next_cell = self.grid[next_y][next_x]
        return next_cell in self.obstructions

    def _set_starting_position(self):
        for y in range(len(self.grid)):
            for x in range(len(self.grid[y])):
                if self.grid[y][x] in self.positions:
                    self.guard_position = (y, x)
                    return
        self.guard_position = (-1, -1)  # Guard is out of bounds

    def _move_forward(self) -> bool:

        (cy, cx) = self.guard_position
        direction = self.grid[cy][cx]
        next_x = cx + self.forward_x[direction]
        next_y = cy + self.forward_y[direction]

        self.grid[cy][cx] = 'X'
        self.guard_position = (next_y, next_x)

        if next_y < 0 or next_y >= len(self.grid):
            return True
        if next_x < 0 or next_x >= len(self.grid[0]):
            return True

        self.grid[next_y][next_x] = direction
        return False

    def _stuck_in_loop(self):
        y, x = self.guard_position
        direction = self.grid[y][x]
        self.entries[self.guard_position][direction] += 1
        return self.entries[self.guard_position][direction] > 2

    def init_entries(self):
        for y in range(len(self.grid)):
            for x in range(len(self.grid[0])):
                self.entries[(y, x)] = {
                    '^': 0,
                    'v': 0,
                    '<': 0,
                    '>': 0
                }


if __name__ == "__main__":
    day = Day6()
    print(day.process_input(sample_data))
