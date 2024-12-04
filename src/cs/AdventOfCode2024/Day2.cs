public class Day2 : IDay
{
    public string Execute(string input)
    {
        var reports = input.Split('\n');

        int safeReports = 0;

        foreach (var report in reports)
        {
            var levels = report.Split(' ').Select(int.Parse).ToArray();

            if (IsSafe(levels))
            {
                safeReports++;
            }
            else
            {
                for(int level = 0; level < levels.Length; level++)
                {
                    var newLevels = levels.Take(level).Concat(levels.Skip(level + 1)).ToArray();

                    if (IsSafe(newLevels))
                    {
                        safeReports++;
                        break;
                    }
                }
            }

        }

        return safeReports.ToString();
    }

    private bool IsSafe(int[] levels)
    {
        bool increasing = true;
        bool decreasing = true;

        for(int level = 1; level < levels.Length; level++)
        {
            var diff = levels[level] - levels[level - 1];
            if (diff < 0) increasing = false;
            if (diff > 0) decreasing = false;
        }

        bool withinRange = true;

        for(int level = 1; level < levels.Length; level++)
        {
            var diff = Math.Abs(levels[level] - levels[level - 1]);

            if (diff < 1 || diff > 3)
            {
                withinRange = false;
                break;
            }
        }

        return (increasing || decreasing) && withinRange;
    }
}