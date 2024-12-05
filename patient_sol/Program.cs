using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using patient;

if (args.Length != 2)
{
    Console.WriteLine("Usage: patient <int id> <int data>\n - Where id is in range 1-3\n - Where data is in range 1-10000");
    return;
}

Patient patient = new Patient(int.Parse(args[0]), int.Parse(args[1]));
Console.WriteLine($"Patient ID: {patient.Id} is connecting...");

var bob = WebApplication.CreateBuilder(args);
bob.WebHost.UseUrls($"https://localhost:{patient.Port}");
bob.Services.AddControllers();
bob.Services.AddHttpClient();
bob.Services.AddSingleton(patient);
bob.Services.AddHttpsRedirection(options => { options.HttpsPort = patient.Port; });

var app = bob.Build();
if (app.Environment.IsDevelopment()) { app.UseDeveloperExceptionPage(); }
app.UseHttpsRedirection();
app.UseRouting();
app.MapControllers();

_ = app.RunAsync();
Console.WriteLine($"Patient {patient.Id} is online.\nPlease wait until SECU Hospital and the remaining patients are online before continuing.\n[Press ENTER to continue]");
Console.ReadLine();

Console.WriteLine($"Patient {patient.Id} is sending secrets...");
await patient.SendShares();
Console.WriteLine($"Wait until all secrets have been sent before continuing.\n[Press ENTER to continue].");
Console.ReadLine();

await patient.SendToHospital();
Console.WriteLine($"Patient {patient.Id} sent their sum of shares to SECU Hospital.\n- Closing connection -");