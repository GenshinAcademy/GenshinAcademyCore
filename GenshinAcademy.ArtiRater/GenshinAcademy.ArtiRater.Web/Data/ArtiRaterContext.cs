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
                .HasMany(character => character.StatsProfit)
                .WithOne(profit => profit.OwnerCharacter)
                .HasForeignKey(profit => profit.OwnerCharacterId);
            modelBuilder.Entity<Character>().Property(character => character.IconUrl)
                .HasConversion<string>();
            modelBuilder.Entity<Character>()
                .Property(character => character.CharacterNameId)
                .IsRequired();
            modelBuilder.Entity<Character>()
                .HasIndex(character => character.CharacterNameId)
                .IsUnique();

            modelBuilder.Entity<ArtifactStatsProfit>()
                .Property(profit => profit.Type)
                .HasConversion<int>();
        }
    }
}
