using GenshinAcademy.ArtiRater.Web.Data;
using Microsoft.EntityFrameworkCore;

namespace GenshinAcademy.ArtiRater.Web
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var builder = WebApplication.CreateBuilder(args);

            // Add services to the container.

            string dbConnection = "";

            builder.Services.AddControllers();
            builder.Services.AddDbContext<ArtiRaterContext>(x =>
            {
                //x.UseSqlServer("Server=(localdb)\\mssqllocaldb;Database=BookmarksOnline;Trusted_Connection=True;");
                x.UseNpgsql(dbConnection);
            });
            var app = builder.Build();

           // app.UseHttpsRedirection();

            app.MapControllers();

            app.Run();
        }
    }
}