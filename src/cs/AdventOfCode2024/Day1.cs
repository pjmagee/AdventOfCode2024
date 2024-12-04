using System.Text.Json;
using System.Text.Json.Serialization;

namespace AdventOfCode2024;

public class Day1 : IDay
{
    class Response
    {
        [JsonPropertyName("total_distance")]
        public int TotalDistance { get; set; }

        [JsonPropertyName("total_similarity")]
        public int TotalSimilarity { get; set; }
    }

    public string Execute(string input)
    {
        var lines = input.Split("\n");

        List<int> left = new List<int>();
        List<int> right = new List<int>();

        foreach (var line in lines)
        {
            var segments = line.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            left.Add(int.Parse(segments[0]));
            right.Add(int.Parse(segments[1]));
        }

        left.Sort();
        right.Sort();

        Response response = new Response();
        response.TotalDistance = 0;

        for(int i = 0; i < left.Count; i++)
        {
            response.TotalDistance += Math.Abs(left[i] - right[i]);
        }

        response.TotalSimilarity = 0;

        foreach (var l in left)
        {
            var leftInRight = 0;

            foreach (var r in right)
            {
                if(l == r)
                {
                    leftInRight++;
                }
            }

            response.TotalSimilarity += l * leftInRight;
        }

        return JsonSerializer.Serialize(response);
    }
}