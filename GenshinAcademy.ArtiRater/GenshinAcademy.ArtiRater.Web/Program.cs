using GenshinAcademy.ArtiRater.Web.Data;
using Microsoft.AspNetCore;
using Microsoft.EntityFrameworkCore;

namespace GenshinAcademy.ArtiRater.Web
{
    public class Program
    {
        public static void Main(string[] args)
        {
            IWebHost host = BuildWebHost(args);
            host.Run();
        }

        public static IWebHost BuildWebHost(string[] args) =>
            WebHost.CreateDefaultBuilder(args)
            .UseStartup<Startup>()
            .UseWebRoot("static")
            .Build();
    }
}