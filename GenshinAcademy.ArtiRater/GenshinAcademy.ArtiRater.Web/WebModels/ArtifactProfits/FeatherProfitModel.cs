using GenshinAcademy.ArtiRater.Web.Models;
using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels.ArtifactProfits
{
    public class FeatherProfitModel : ArtifactProfitModel
    {
        [JsonPropertyName("ATK")]
        public int AttackFlat { get; set; }

        public FeatherProfitModel(ArtifactStatsProfit dataModel) : base("feather")
        {
            AttackFlat = dataModel.Attack;
        }
    }
}
