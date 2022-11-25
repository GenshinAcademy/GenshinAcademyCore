using GenshinAcademy.ArtiRater.Web.Models.Enums;

namespace GenshinAcademy.ArtiRater.Web.Models
{
    public class ArtifactStatsProfit
    {
        public int Id { get; set; }

        public Character? OwnerCharacter { get; set; }
        public int OwnerCharacterId { get; set; }

        public int Attack { get; set; }

        public int AttackPercentage { get; set; }

        public int Health { get; set; }

        public int HealthPercentage { get; set; }

        public int Defense { get; set; }

        public int DefensePercentage { get; set; }

        public int ElementalMastery { get; set; }

        public int EnergyRecharge { get; set; }

        public int ElementalDamage { get; set; }

        public int CritRate { get; set; }

        public int CritDamage { get; set; }

        public int PhysicalDamage { get; set; }

        public int Heal { get; set; }

        public ArtifactStatProfitType Type { get; set; }
    }
}
