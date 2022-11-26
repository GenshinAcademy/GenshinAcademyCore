using GenshinAcademy.ArtiRater.Web.Data;
using GenshinAcademy.ArtiRater.Web.Models;
using GenshinAcademy.ArtiRater.Web.WebModels;
using Microsoft.AspNetCore.Mvc;

namespace GenshinAcademy.ArtiRater.Web.Controllers
{
    [Route("characters")]
    public class CharactersController : Controller
    {
        private ArtiRaterContext _context;
        private ArtifactRaterRepository _repository;

        public CharactersController(ArtiRaterContext context)
        {
            _context = context;
            _repository = new ArtifactRaterRepository(_context);
        }

        [HttpGet("list")]
        public async Task<IActionResult> GetCharactersList([FromQuery] string element)
        {
            if (!string.IsNullOrEmpty(element))
            {
                IEnumerable<Character> foundCharacters = await _repository.GetCharactersByElementAsync(element, false);
                return Ok(foundCharacters.Select(x => new CharacterModel(x)).ToArray());
            }

            IEnumerable<Character> chars = await _repository.GetAllCharactersListAsync();
            return Ok(chars.Select(x => new CharacterModel(x)).ToArray());
        }

        [HttpGet("artifacts")]
        public async Task<IActionResult> GetCharactersProfits([FromQuery]string character, [FromQuery] string element)
        {
            if (!string.IsNullOrEmpty(character))
            {
                Character foundCharacter = await _repository.GetCharacterByIdAsync(character);
                return Ok(new CharacterWithProfitsModel(foundCharacter, foundCharacter.StatsProfit));
            }
            if (!string.IsNullOrEmpty(element))
            {
                IEnumerable<Character> foundCharacters = await _repository.GetCharactersByElementAsync(element, true);
                return Ok(foundCharacters.Select(x => new CharacterWithProfitsModel(x, x.StatsProfit)).ToArray());
            }

            IEnumerable<Character> chars = await _repository.GetAllCharactersWithProfitsAsync();
            return Ok(chars.Select(x => new CharacterWithProfitsModel(x, x.StatsProfit)).ToArray());
        }
    }
}
