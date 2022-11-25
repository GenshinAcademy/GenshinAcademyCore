using GenshinAcademy.ArtiRater.Web.Models;
using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels.ArtifactProfits
{
    public class SandsProfitModel : ArtifactProfitModel
    {

        [JsonPropertyName("ATK_P")]
        public int AttackPercentage { get; set; }

        [JsonPropertyName("HP_P")]
        public int HealthPercentage { get; set; }

        [JsonPropertyName("DEF_P")]
        public int DefensePercentage { get; set; }

        [JsonPropertyName("EM")]
        public int ElementalMastery { get; set; }

        [JsonPropertyName("ER")]
        public int EnergyRecharge { get; set; }

        public SandsProfitModel(ArtifactStatsProfit dataModel) : base("sands")

        {
            AttackPercentage = dataModel.AttackPercentage;
            HealthPercentage = dataModel.HealthPercentage;
            DefensePercentage = dataModel.DefensePercentage;
            ElementalMastery = dataModel.ElementalMastery;
            EnergyRecharge = dataModel.EnergyRecharge;
        }
    }
}
