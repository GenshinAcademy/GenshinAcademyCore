namespace GenshinAcademy.ArtiRater.Web.Models
{
    public class Character
    {
        public int Id { get; set; }

        public string Name { get; set; }

        public string Element { get; set; }
        
        public Uri IconUrl { get; set; }

        public IReadOnlyCollection<ArtifactStatsProfit>? StatsProfit { get; set; }
    }
}
