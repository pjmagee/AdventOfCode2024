import importlib
import inspect
import os
import sys

from days.day import Day


def load_days() -> dict[str, type[Day]]:

    base_path = os.path.dirname(os.path.abspath(__file__))
    folder_path = os.path.join(base_path, "days")

    modules = [
        f[:-3]  # remove .py
        for f in os.listdir(folder_path)
        if f.endswith(".py") and f != "__init__.py"
    ]

    day_classes = {}

    for module_name in modules:
        full_module_name = f"days.{module_name}"
        module = importlib.import_module(full_module_name)

        for name, obj in inspect.getmembers(module, inspect.isclass):
            if issubclass(obj, Day) and obj is not Day:
                day_classes[name] = obj

    return day_classes


def main(day: int) -> None:
    if sys.stdin.isatty():
        print("No input data.")
        sys.exit(1)
    else:
        input_data: str = sys.stdin.read()

        days_modules = load_days()

        for module in days_modules.items():
            day_class: type[Day] = module[1]
            if day_class.__name__.endswith(str(day)):
                day_instance = day_class()
                if day_instance.process_input(input_data):
                    print(day_instance.process_input(input_data))
                    sys.exit(0)

        print(f"Day {day} not found.")
        sys.exit(1)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python main.py <day>")
        sys.exit(1)

    try:
        day = int(sys.argv[1])
    except ValueError:
        print("The day argument must be an integer.")
        sys.exit(1)

    main(day)
