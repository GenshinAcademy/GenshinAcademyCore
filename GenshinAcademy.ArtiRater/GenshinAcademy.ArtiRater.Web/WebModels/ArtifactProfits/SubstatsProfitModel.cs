using GenshinAcademy.ArtiRater.Web.Models;
using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels.ArtifactProfits
{
    public class SubstatsProfitModel : ArtifactProfitModel
    {

        [JsonPropertyName("ATK")]
        public int AttackFlat { get; set; }

        [JsonPropertyName("ATK_P")]
        public int AttackPercentage { get; set; }

        [JsonPropertyName("HP")]
        public int HealthFlat { get; set; }

        [JsonPropertyName("HP_P")]
        public int HealthPercentage { get; set; }

        [JsonPropertyName("CD")]
        public int CritDamage { get; set; }

        [JsonPropertyName("CR")]
        public int CritRate { get; set; }

        [JsonPropertyName("EM")]
        public int ElementalMastery { get; set; }

        [JsonPropertyName("DEF")]
        public int DefenseFlat { get; set; }

        [JsonPropertyName("DEF_P")]
        public int DefensePercentage { get; set; }

        [JsonPropertyName("ER")]
        public int EnergyRecharge { get; set; }

        public SubstatsProfitModel(ArtifactStatsProfit dataModel) : base("substats")
        {
            AttackFlat = dataModel.Attack;
            AttackPercentage = dataModel.AttackPercentage;
            HealthFlat = dataModel.Health;
            HealthPercentage = dataModel.HealthPercentage;
            DefenseFlat = dataModel.Defense;
            DefensePercentage = dataModel.DefensePercentage;
            EnergyRecharge = dataModel.EnergyRecharge;
            CritDamage = dataModel.CritDamage;
            CritRate = dataModel.CritRate;
            ElementalMastery = dataModel.ElementalMastery;
        }
    }
}
