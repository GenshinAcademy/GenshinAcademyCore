using GenshinAcademy.ArtiRater.Web.Models;
using GenshinAcademy.ArtiRater.Web.Models.Enums;
using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels
{
    public class CharacterWithProfitsModel : CharacterModel
    {

        [JsonPropertyName("stats_profit")]
        public IReadOnlyDictionary<string, object> StatsProft { get; set; }

        public CharacterWithProfitsModel(Character dataModel, IEnumerable<ArtifactStatsProfit> profits) : base(dataModel)
        {
            Name = dataModel.Name;
            Element = dataModel.Element;
            IconUrl = dataModel.IconUrl;
            var dictionary = new Dictionary<string, object>(6);
            foreach (ArtifactStatsProfit profit in profits)
            {
                ArtifactProfitModel model = profit.Type switch
                {
                    ArtifactStatProfitType.Circlet => new ArtifactProfits.CircletProfitModel(profit),
                    ArtifactStatProfitType.Feather => new ArtifactProfits.FeatherProfitModel(profit),
                    ArtifactStatProfitType.Flower => new ArtifactProfits.FlowerProfitModel(profit),
                    ArtifactStatProfitType.Goblet => new ArtifactProfits.GobletProfitModel(profit),
                    ArtifactStatProfitType.Sands => new ArtifactProfits.SandsProfitModel(profit),
                    ArtifactStatProfitType.SubStats => new ArtifactProfits.SubstatsProfitModel(profit),
                    _ => throw new Exception($"Unknown type of profit type {((int)profit.Type)}")
                };
                dictionary.Add(model.Key, model);
            }
            StatsProft = dictionary;
        }
    }
}
