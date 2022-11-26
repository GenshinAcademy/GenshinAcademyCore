using GenshinAcademy.ArtiRater.Web.Models;
using Microsoft.EntityFrameworkCore;

namespace GenshinAcademy.ArtiRater.Web.Data
{
    public class ArtifactRaterRepository
    {
        private ArtiRaterContext _context;

        public ArtifactRaterRepository(ArtiRaterContext context)
        {
            _context = context;
        }

        public async Task<IEnumerable<Character>> GetAllCharactersListAsync()
        {
            Character[] result = await _context.Characters
                .AsNoTracking()
                .ToArrayAsync();

            return result;
        }

        public async Task<IEnumerable<Character>> GetAllCharactersWithProfitsAsync()
        {
            Character[] result = await _context.Characters
                .Include(x => x.StatsProfit)
                .AsNoTracking()
                .ToArrayAsync();

            return result;
        }

        public async Task<IEnumerable<Character>> GetCharactersByElementAsync(string element, bool includeProfits)
        {
            if (string.IsNullOrWhiteSpace(element))
            {
                throw new ArgumentException("Null or whitespace value", nameof(element));
            }

            string lowerElement = element.ToLower();
            IQueryable<Character> query = _context.Characters
                .Where(x => x.Element.ToLower() == lowerElement);

            if (includeProfits)
            {
                query = query.Include(character => character.StatsProfit);
            }

            Character[] result = await query
                .AsNoTracking()
                .ToArrayAsync();

            if (result.Length == 0)
            {
                throw new ArgumentException($"No characters found with provided element \"{element}\"", nameof(element));
            }

            return result;
        }

        public async Task<Character> GetCharacterByIdAsync(string characterId)
        {
            if (string.IsNullOrWhiteSpace(characterId))
            {
                throw new ArgumentException("Null or whitespace value", nameof(characterId));
            }

            Character? result = await _context.Characters
                .Include(character => character.StatsProfit)
                .AsNoTracking()
                .FirstOrDefaultAsync(character => character.CharacterNameId == characterId);

            if (result == null)
            {
                throw new ArgumentException($"Cannot find character by id \"{characterId}\"", nameof(characterId));
            }

            return result;
        }
    }
}
