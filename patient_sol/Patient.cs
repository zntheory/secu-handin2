using System.Text.Json;
namespace patient;

public class Patient
{
    public int Id { get; set; }
    public int Data { get; set; }
    public int Port { get; set; }
    private int hospitalPort = 8080;
    private int maxdata = 15000;
    private int[] myShares = new int[3];
    private List<int> shares = new();
    private List<string> ports = ["8081", "8082", "8083"];
    private Random rand = new();

    public Patient(int id, int data)
    {
        Id = id;
        Data = maxdata - data;
        Port = hospitalPort + Id;
        Split();
    }

    private void Split()
    {
        int temp = Data;

        for (int i = 0; i < ports.Count - 1; i++)
        {
            myShares[i] = rand.Next(maxdata);
            temp -= myShares[i];
        }

        myShares[2] = temp % maxdata;
    }

    public void AppendShare(int share) { shares.Add(share); }

    public int SumShares() { return myShares[0] + shares.Sum(); }

    public async Task SendShares()
    {
        int count = 1;

        foreach (string port in ports)
        {
            // Ignore if this is my own port tihi
            if (port.Equals(Port.ToString())) continue;

            var handler = new HttpClientHandler();
            handler.ServerCertificateCustomValidationCallback = (message, cert, chain, err) => true; // Ignore cert since I don't have a good one
            using HttpClient client = new(handler);
            var content = new StringContent(JsonSerializer.Serialize(myShares[count]), System.Text.Encoding.UTF8, "application/json");
            var response = await client.PostAsync($"https://localhost:{port}/Patient/Message", content);
            var result = await response.Content.ReadAsStringAsync();
            Console.WriteLine("Response: " + result);
            count++;
        }
    }

    public async Task SendToHospital()
    {
        int sumShares = SumShares();
        using HttpClient client = new();
        var handler = new HttpClientHandler();
        handler.ServerCertificateCustomValidationCallback = (message, cert, chain, errors) => true; // Actual good certs cost money, so... let's pretend.
        using var secureClient = new HttpClient(handler);
        
        var content  = new StringContent(sumShares.ToString(), System.Text.Encoding.UTF8, "application/json");
        var response = await secureClient.PostAsync($"https://localhost:{hospitalPort}/Hospital/Aggregate", content);

        if (!response.IsSuccessStatusCode)
        {
            Console.WriteLine($"Error sending sum {sumShares} to SECU Hospital");
            var result = await response.Content.ReadAsStringAsync();
            Console.WriteLine("Response: " + result);
        }
        else { Console.WriteLine($"Sent sum {sumShares} to SECU Hospital"); }
    }
    
}