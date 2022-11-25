using GenshinAcademy.ArtiRater.Web.Models;
using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels.ArtifactProfits
{
    public class GobletProfitModel : ArtifactProfitModel
    {
        
        [JsonPropertyName("ATK_P")]
        public int AttackPercentage { get; set; }

        [JsonPropertyName("HP_P")]
        public int HealthPercentage { get; set; }

        [JsonPropertyName("DEF_P")]
        public int DefensePercentage { get; set; }

        [JsonPropertyName("EM")]
        public int ElementalMastery { get; set; }

        [JsonPropertyName("PHYS")]
        public int PhysicalDamage { get; set; }

        [JsonPropertyName("ELEM")]
        public int ElementalDamage { get; set; }

       public GobletProfitModel(ArtifactStatsProfit dataModel) : base("goblet")
        {
            AttackPercentage = dataModel.AttackPercentage;
            HealthPercentage = dataModel.HealthPercentage;
            DefensePercentage = dataModel.DefensePercentage;
            ElementalMastery = dataModel.ElementalMastery;
            PhysicalDamage = dataModel.PhysicalDamage;
            ElementalDamage = dataModel.ElementalDamage;
        }
    }
}
