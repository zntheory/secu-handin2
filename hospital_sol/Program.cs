using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.AspNetCore.Builder;
using hospital;

// https://youtu.be/HdVg-2jn2OU?si=psQmPl3GkleABtrn&t=3
var builderBob = WebApplication.CreateBuilder(args);
builderBob.Services.AddControllers();
builderBob.Services.AddSingleton<Hospital>();
builderBob.Services.AddEndpointsApiExplorer();
builderBob.WebHost.UseUrls("https://localhost:8080");
builderBob.Services.AddHttpsRedirection(options => { options.HttpsPort = 8080; });

var app = builderBob.Build();
app.UseHttpsRedirection();
app.UseRouting();
app.UseCors(builder => builder.AllowAnyOrigin().AllowAnyMethod().AllowAnyHeader());
app.MapControllers();

Console.WriteLine("- Hello World! Hospital is listening on port 8080");

app.Run();