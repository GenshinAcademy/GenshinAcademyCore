package db

import (
	"context"
	"fmt"
	"ga/internal/db/entity"
	"ga/internal/db/mapper"
	"ga/internal/db/repository"
	"ga/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

type PostgresConfig struct {
	Host         string
	Port         uint16
	Username     string
	Password     string
	DatabaseName string
}

func (cfg *PostgresConfig) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DatabaseName,
		cfg.Port,
	)
}

type Provider struct {
	db *gorm.DB
}

func NewPostgresProvider(cfg *PostgresConfig) (*Provider, error) {
	db, err := gorm.Open(postgres.Open(cfg.String()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Provider{db}, nil
}

func (p *Provider) GetCharacterRepository() *repository.CharacterRepository {
	return repository.NewCharacterRepository(p.db, mapper.NewCharacterMapper())
}

func (p *Provider) GetArtifactProfitsRepository() *repository.ArtifactProfitsRepository {
	return repository.NewArtifactProfitsRepository(p.db)
}

func (p *Provider) GetNewsRepository() *repository.NewsRepository {
	return repository.NewNewsRepository(p.db, mapper.NewNewsMapper())
}

func (p *Provider) GetTableRepository() *repository.TableRepository {
	return repository.NewTableRepository(p.db, mapper.NewTableMapper())
}

func (p *Provider) Migrate(ctx context.Context) error {
	if err := CreateEnumType(p.db, "element", []types.Element{
		types.UndefinedElement,
		types.Pyro,
		types.Hydro,
		types.Geo,
		types.Anemo,
		types.Electro,
		types.Cryo,
		types.Dendro,
	}); err != nil {
		return err
	}

	if err := CreateEnumType(p.db, "rarity", []types.Rarity{
		types.Rare,
		types.Epic,
		types.Legendary,
	}); err != nil {
		return err
	}

	if err := CreateEnumType(p.db, "weapon_type", []types.WeaponType{
		types.UndefinedWeapon,
		types.Sword,
		types.Claymore,
		types.Polearm,
		types.Bow,
		types.Catalyst,
	}); err != nil {
		return err
	}

	if err := CreateEnumType(p.db, "artifact_slot", []types.ArtifactSlot{
		types.SubStats,
		types.Flower,
		types.Plume,
		types.Sands,
		types.Goblet,
		types.Circlet,
	}); err != nil {
		return err
	}

	if err := p.db.WithContext(ctx).AutoMigrate(&entity.Character{}); err != nil {
		return err
	}
	if err := p.db.WithContext(ctx).AutoMigrate(&entity.News{}); err != nil {
		return err
	}
	if err := p.db.WithContext(ctx).AutoMigrate(&entity.Table{}); err != nil {
		return err
	}
	if err := p.db.WithContext(ctx).AutoMigrate(&entity.ArtifactProfits{}); err != nil {
		return err
	}
	if err := p.db.WithContext(ctx).AutoMigrate(&entity.CharacterIcons{}); err != nil {
		return err
	}

	return nil
}

func CreateEnumType[T any](db *gorm.DB, title string, enums []T) error {
	var enumsStr []string
	for _, element := range enums {
		enumsStr = append(enumsStr, fmt.Sprintf("'%v'", element))
	}
	if err := db.Exec(fmt.Sprintf(`DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = '%s') THEN
        CREATE TYPE %s AS ENUM (%s);
    END IF;
END$$;`,
		title,
		title,
		strings.Join(enumsStr, ","),
	)).Error; err != nil {
		return err
	}
	return nil
}
