using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.AspNetCore.Builder;
using hospital;

var builderBob = WebApplication.CreateBuilder(args);
builderBob.Services.AddControllers();
builderBob.Services.AddSingleton<Hospital>();
builderBob.WebHost.UseUrls("https://localhost:5001");
builderBob.Services.AddEndpointsApiExplorer();
builderBob.Services.AddHttpsRedirection(options => { options.HttpsPort = 5001; });

var app = builderBob.Build();
app.UseHttpsRedirection();
app.UseRouting();
app.UseCors(builder => builder.AllowAnyOrigin().AllowAnyMethod().AllowAnyHeader());

app.MapControllers();
Console.WriteLine("- Hello World! Hospital is listening on port 5001");
app.Run();