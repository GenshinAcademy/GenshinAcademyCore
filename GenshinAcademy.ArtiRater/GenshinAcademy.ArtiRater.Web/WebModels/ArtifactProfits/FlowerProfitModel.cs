using GenshinAcademy.ArtiRater.Web.Models;
using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels.ArtifactProfits
{
    public class FlowerProfitModel : ArtifactProfitModel
    {

        [JsonPropertyName("HP")]
        public int HealthFlat { get; set; }

        public FlowerProfitModel(ArtifactStatsProfit dataModel) : base("flower")
        {
            HealthFlat = dataModel.Health; 
        }
    }
}
