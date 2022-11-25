using GenshinAcademy.ArtiRater.Web.Models;
using Microsoft.EntityFrameworkCore;

namespace GenshinAcademy.ArtiRater.Web.Data
{
    public class ArtiRaterRepository
    {
        private ArtiRaterContext _context;

        public ArtiRaterRepository(ArtiRaterContext context)
        {
            _context = context;
        }

        public async Task<IEnumerable<Character>> GetCharacters()
        {
            var result = await _context.Characters
                .ToArrayAsync();
            return result;
        }

        public async Task<IEnumerable<Character>> GetCharactersWithProfits()
        {
            var result = await _context.Characters
                .Include(x => x.StatsProfit)
                .ToArrayAsync();

            return result;
        }
    }
}
