if(args.Length != 1)
{
    if (OperatingSystem.IsWindows())
    {
        Console.Error.WriteLine("PS> <input> | dotnet run --project AdventOfCode2024 -- <day>");
    }
    else if (OperatingSystem.IsLinux())
    {
        Console.Error.WriteLine("$ dotnet run --project AdventOfCode2024 -- <day> < <input>");
    }

    return;
}

int day = int.Parse(args[0]);

string input = Console.IsInputRedirected ?  Console.In.ReadToEnd() : args[0];

if (string.IsNullOrWhiteSpace(input))
{
    Console.Error.WriteLine("No input was piped to STDIN");
}

Dictionary<int, IDay> days = new();

foreach(var type in typeof(IDay).Assembly.GetTypes().Where(x => x.GetInterfaces().Contains(typeof(IDay))))
{
    IDay? implmentation = (IDay)Activator.CreateInstance(type)!;
    days.Add(implmentation.Day, implmentation);
}

Console.Out.WriteLine(days.ContainsKey(day) ? days[day].Execute(input) : "Day not found");