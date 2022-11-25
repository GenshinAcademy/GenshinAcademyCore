using System.Text.Json.Serialization;

namespace GenshinAcademy.ArtiRater.Web.WebModels
{
    public abstract class ArtifactProfitModel
    {
        [JsonIgnore]
        public string Key { get; }

        protected ArtifactProfitModel(string key)
        {
            Key = key;
        }
    }
}
