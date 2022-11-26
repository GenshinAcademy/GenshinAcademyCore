using GenshinAcademy.ArtiRater.Web.Data;
using Microsoft.EntityFrameworkCore;

namespace GenshinAcademy.ArtiRater.Web
{
    public class Startup
    {
        public IConfiguration Configuration { get; set; }

        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public void ConfigureServices(IServiceCollection services)
        {
            services.AddRouting(x => x.LowercaseUrls = true);
            services.AddControllers(x =>
            {
                x.EnableEndpointRouting = false;
            });
            services.AddDbContext<ArtiRaterContext>(x =>
            {
                x.UseNpgsql(Configuration["ConnectionStrings:Postgres"]);
            });
            services.AddCors();
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
                app.UseStatusCodePages();
                //app.UseHttpsRedirection();
            }
            else
            {
                app.UseHsts();
            }

            app.UseCors(cpb =>
            {
                cpb.AllowAnyOrigin();
                cpb.WithMethods(new[] { "GET", "POST", "PUT", "PATCH", "DELETE" });
                cpb.AllowAnyHeader();
            });

            app.UseStaticFiles();
            app.UseRouting();

            app.UseAuthentication();
            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}
