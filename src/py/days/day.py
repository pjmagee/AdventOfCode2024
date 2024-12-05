from abc import ABC, abstractmethod


class Day(ABC):

    @abstractmethod
    def process_input(self, input_data: str) -> str:
        pass
