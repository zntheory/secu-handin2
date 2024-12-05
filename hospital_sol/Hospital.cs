namespace hospital;

public class Hospital
{
    public int ShareCount { get; private set; }
    private int _shares, _maxdata;

    public Hospital()
    {
        _shares = 0;
        ShareCount = 0;
        _maxdata = 15000;
    }

    public void Aggregate(int amount) // Additive
    {
        _shares += amount;
        ShareCount++;
        Console.WriteLine($"Received share {amount}. Got {ShareCount} out of 3 shares.");
        Console.WriteLine(ShareCount < 3 ? "" : $"Aggregation result: {_shares}");
    }
}