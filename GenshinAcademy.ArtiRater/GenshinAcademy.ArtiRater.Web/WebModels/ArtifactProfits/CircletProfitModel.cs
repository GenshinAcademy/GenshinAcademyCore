using GenshinAcademy.ArtiRater.Web.Models;
using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels.ArtifactProfits
{
    public class CircletProfitModel : ArtifactProfitModel
    {
        [JsonPropertyName("ATK_P")]
        public int AttackPercentage { get; set; }

        [JsonPropertyName("DEF_P")]
        public int DefensePercentage { get; set; }

        [JsonPropertyName("HP_P")]
        public int HealthPercentage { get; set; }

        [JsonPropertyName("EM")]
        public int ElementalMastery { get; set; }

        [JsonPropertyName("CR")]
        public int CritRate { get; set; }

        [JsonPropertyName("CD")]
        public int CritDamage { get; set; }

        [JsonPropertyName("HEAL")]
        public int Heal { get; set; }

        public CircletProfitModel(ArtifactStatsProfit dataModel) : base("circlet")
        {
            AttackPercentage = dataModel.AttackPercentage;
            DefensePercentage = dataModel.DefensePercentage;
            HealthPercentage = dataModel.HealthPercentage;
            ElementalMastery = dataModel.ElementalMastery;
            CritRate = dataModel.CritRate;
            CritDamage = dataModel.CritDamage;
            Heal = dataModel.Heal;
        }
    }
}
