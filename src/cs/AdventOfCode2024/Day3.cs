using System.Text.RegularExpressions;

namespace AdventOfCode2024;

public partial class Day3 : IDay
{
    public string Execute(string input)
    {
        var regex = InstructionsRegex();
        var matches = regex.Matches(input);
        var total = 0;
        var isEnabled = true;

        foreach (Match match in matches)
        {
            if (match.Value == "do()")
            {
                isEnabled = true;
            }
            else if (match.Value == "don't()")
            {
                isEnabled = false;
            }
            else
            {
                if (isEnabled)
                {
                    string[] values = match.Value[4..^1].Split(',');
                    total += int.Parse(values[0]) * int.Parse(values[1]);
                }

            }
        }

        return total.ToString();
    }

    [GeneratedRegex(@"mul\(\d+,\d+\)|do\(\)|don't\(\)")]
    private static partial Regex InstructionsRegex();
}