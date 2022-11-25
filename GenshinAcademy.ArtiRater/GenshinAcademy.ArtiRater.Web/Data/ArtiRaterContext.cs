using GenshinAcademy.ArtiRater.Web.Models;
using Microsoft.EntityFrameworkCore;

namespace GenshinAcademy.ArtiRater.Web.Data
{
    public class ArtiRaterContext : DbContext
    {
        public DbSet<Character> Characters { get; set; }
        public DbSet<ArtifactStatsProfit> Profits { get; set; }

        public ArtiRaterContext(DbContextOptions<ArtiRaterContext> config) : base(config)
        {
            Database.EnsureCreated();
        }

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Character>()
                .HasMany(x => x.StatsProfit)
                .WithOne(y => y.OwnerCharacter)
                .HasForeignKey(y => y.OwnerCharacterId);

            modelBuilder.Entity<ArtifactStatsProfit>()
                .Property(y => y.Type)
                .HasConversion<int>();
        }
    }
}
