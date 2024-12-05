
// check if FILE was piped to STDIN

using AdventOfCode2024;

if (Console.IsInputRedirected)
{
    int day = int.Parse(args[0]);
    string input = Console.In.ReadToEnd();

    Dictionary<int, IDay> days = new()
    {
        { 1, new Day1() },
        { 2, new Day2() },
        { 3, new Day3() },
    };

    Console.WriteLine(days.ContainsKey(day) ? days[day].Execute(input) : "Day not found");
}
else
{
    Console.WriteLine("No input was piped to STDIN");
}