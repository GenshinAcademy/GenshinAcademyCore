using GenshinAcademy.ArtiRater.Web.Data;
using Microsoft.AspNetCore;
using Microsoft.EntityFrameworkCore;

namespace GenshinAcademy.ArtiRater.Web
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            IWebHost host = BuildWebHost(args);
            await host.RunAsync();
        }

        public static IWebHost BuildWebHost(string[] args) =>
            WebHost.CreateDefaultBuilder(args)
            .UseStartup<Startup>()
            .Build();
    }
}