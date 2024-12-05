namespace AdventOfCode2024;

public class Day5 : IDay
{
    public int Day => 5;

    public string Execute(string input)
    {
        var lines = input.Split('\n', StringSplitOptions.RemoveEmptyEntries).Select(x => x.Trim()).ToList();
        var pageOrderRules = lines.Where(x => x.Contains('|')).Select(GetPageOrders).ToList();
        var updates = lines.Where(l => l.Contains(',')).Select(GetPageUpdates).ToList();

        HashSet<int> correctUpdates = new();

        for (int i = 0; i < updates.Count; i++)
        {
            int[] pages = updates[i];
            bool isUpdateCorrect = true;

            for(int c = 0; c < pages.Length; c++)
            {
                int currentPage = pages[c];

                for(int s = 0; s < pages.Length; s++)
                {
                    if (s == c) continue;
                    int otherPage = pages[s];

                    int before = c < s ? currentPage : otherPage;
                    int after = c < s ? otherPage : currentPage;

                    bool isCorrect = pageOrderRules.Exists(x => x.Before == before && x.After == after);

                    if (isCorrect)
                    {
                        // Console.WriteLine(c < s ?
                        //     $"{currentPage} is correctly before {otherPage}" :
                        //     $"{currentPage} is correctly after {otherPage}");
                    }
                    else
                    {
                        // Console.WriteLine(c < s ?
                        //     $"{otherPage} is incorrectly after {currentPage}" :
                        //     $"{otherPage} is incorrectly before {currentPage}"
                        // );

                        isUpdateCorrect = false;
                    }
                }
            }

            if (isUpdateCorrect)
            {
                correctUpdates.Add(i);
            }
        }

        List<int> middleNumbers = new();

        foreach(var idx in correctUpdates)
        {
            var pageNumbers = updates[idx];
            var middleNumber = pageNumbers[pageNumbers.Length / 2];
            middleNumbers.Add(middleNumber);
        }

        var partOne = middleNumbers.Sum();

        HashSet<int> fixedUpdates = new();

        for(int u = 0; u < updates.Count; u++)
        {
            if (correctUpdates.Contains(u)) continue;

            List<int> incorrectPages = updates[u].ToList();

            incorrectPages.Sort((a, b) =>
            {
                bool isCorrect = pageOrderRules.Exists(x => x.Before == a && x.After == b);
                return isCorrect ? -1 : 1;
            });

            updates[u] = incorrectPages.ToArray();
            fixedUpdates.Add(u);
        }

        middleNumbers.Clear();

        foreach(var idx in fixedUpdates)
        {
            var pageNumbers = updates[idx];
            var middleNumber = pageNumbers[pageNumbers.Length / 2];
            middleNumbers.Add(middleNumber);
        }

        var partTwo = middleNumbers.Sum();

        return System.Text.Json.JsonSerializer.Serialize(new
        {
            PartOne = partOne,
            PartTwo = partTwo
        });
    }

    int[] GetPageUpdates(string x)
    {
        return x.Split(',').Select(int.Parse).ToArray();
    }

    PageOrder GetPageOrders(string x)
    {
        var seg = x.Split('|');
        return new PageOrder(int.Parse(seg[0]), int.Parse(seg[1]));
    }

    public record PageOrder(int Before, int After);
}