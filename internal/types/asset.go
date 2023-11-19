package types

type AssetType string

const (
	CharactersAsset      AssetType = "characters"
	CharactersIconsAsset AssetType = CharactersAsset + "/icons"
	TablesAsset          AssetType = "tables"
	NewsAsset            AssetType = "news"
	OpenGraphAsset       AssetType = "opengraph"
)
