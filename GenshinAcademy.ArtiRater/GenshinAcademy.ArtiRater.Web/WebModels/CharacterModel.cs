using GenshinAcademy.ArtiRater.Web.Models;
using GenshinAcademy.ArtiRater.Web.Models.Enums;
using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels
{
    public class CharacterModel
    {
        [JsonPropertyName("name")]
        public string Name { get; set; }

        [JsonPropertyName("element")]
        public string Element { get; set; }

        [JsonPropertyName("icon_url")]
        public Uri IconUrl { get; set; }

        public CharacterModel(Character dataModel)
        {
            Name = dataModel.Name;
            Element = dataModel.Element;
            IconUrl = dataModel.IconUrl;
        }
    }
}
