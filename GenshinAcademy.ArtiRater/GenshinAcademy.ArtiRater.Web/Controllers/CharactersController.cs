using GenshinAcademy.ArtiRater.Web.Data;
using GenshinAcademy.ArtiRater.Web.WebModels;
using Microsoft.AspNetCore.Mvc;

namespace GenshinAcademy.ArtiRater.Web.Controllers
{
    [Route("characters")]
    public class CharactersController : Controller
    {
        private ArtiRaterContext _context;
        private ArtiRaterRepository _repository;

        public CharactersController(ArtiRaterContext context)
        {
            _context = context;
            _repository = new ArtiRaterRepository(_context);
        }

        [HttpGet("list")]
        public async Task<IActionResult> GetCharactersList()
        {
            var chars = await _repository.GetCharacters();
            return Ok(chars.Select(x => new CharacterModel(x)).ToArray());
        }

        [HttpGet("artifacts")]
        public async Task<IActionResult> GetCharactersProfits()
        {
            var chars = await _repository.GetCharactersWithProfits();
            return Ok(chars.Select(x => new CharacterWithProfitsModel(x, x.StatsProfit)).ToArray());
        }

        [HttpGet("artifacts/{character}")]
        public async Task<IActionResult> GetCharactersList([FromRoute]string character)
        {
            await Task.CompletedTask;
            return Ok();
        }
    }
}
