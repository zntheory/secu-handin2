namespace hospital;

public class Hospital
{
    public int ShareCount { get; set; }
    private int _shares;

    public Hospital()
    {
        _shares = 0;
        ShareCount = 0;
    }

    public void Aggregate(int amount) // Additive
    {
        _shares += amount;
        ShareCount++;
        Console.WriteLine($"Received share {amount}. Got {ShareCount} out of 3 shares.");
        Console.WriteLine(ShareCount < 3 ? $"Aggregation so far: {_shares}" : $"Aggregation result: {_shares}");
    }
}