using GenshinAcademy.ArtiRater.Web.Models;
using GenshinAcademy.ArtiRater.Web.Models.Enums;
using GenshinAcademy.ArtiRater.Web.WebModels.ArtifactProfits;
using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels
{
    public class CharacterWithProfitsModel : CharacterModel
    {

        [JsonPropertyName("stats_profit")]
        public IReadOnlyDictionary<string, ArtifactProfitModel> StatsProft { get; set; }

        public CharacterWithProfitsModel(Character dataModel, IEnumerable<ArtifactStatsProfit> profits) : base(dataModel)
        {
            Name = dataModel.Name;
            Element = dataModel.Element;
            IconUrl = dataModel.IconUrl;
            var dictionary = new Dictionary<string, ArtifactProfitModel>(6);
            foreach (ArtifactStatsProfit profit in profits)
            {
                ArtifactProfitModel model = profit.Type switch
                {
                    ArtifactStatProfitType.Circlet  => new CircletProfitModel(profit),
                    ArtifactStatProfitType.Feather  => new FeatherProfitModel(profit),
                    ArtifactStatProfitType.Flower   => new FlowerProfitModel(profit),
                    ArtifactStatProfitType.Goblet   => new GobletProfitModel(profit),
                    ArtifactStatProfitType.Sands    => new SandsProfitModel(profit),
                    ArtifactStatProfitType.SubStats => new SubstatsProfitModel(profit),
                    _ => throw new Exception($"Unknown type of profit type {((int)profit.Type)}")
                };
                dictionary.Add(model.Key, model);
            }
            StatsProft = dictionary;
        }
    }
}
