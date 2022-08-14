package characters

import common "genshinacademycore"

const (
	//Minimal possible level of character
	MinCharacterLevel CharaterLevel = 1
	//Maximum possible level of character
	MaxCharacterLevel CharaterLevel = 90
)

// Character names as constant string values

const (
	Aether    CharacterName = "Aether"
	Albedo    CharacterName = "Albedo"
	Aloy      CharacterName = "Aloy"
	Amber     CharacterName = "Amber"
	Ayaka     CharacterName = "Kamisato Ayaka"
	Ayato     CharacterName = "Kamisato Ayato"
	Barbara   CharacterName = "Barbara"
	Beidou    CharacterName = "Beidou"
	Bennett   CharacterName = "Bennett"
	Childe    CharacterName = "Childe"
	Chongyun  CharacterName = "Chongyun"
	Diluc     CharacterName = "Diluc"
	Diona     CharacterName = "Diona"
	Eula      CharacterName = "Eula"
	Fischl    CharacterName = "Fischl"
	Ganyu     CharacterName = "Ganyu"
	Gorou     CharacterName = "Gorou"
	HuTao     CharacterName = "Hu Tao"
	Heizou    CharacterName = "Shikanoin Heizo"
	Itto      CharacterName = "Arataki Itto"
	Jean      CharacterName = "Jean"
	Kazuha    CharacterName = "Kaedehara Kazuha"
	Keqing    CharacterName = "Keqing"
	Kokomi    CharacterName = "Sangonomiya Kokomi"
	Kaeya     CharacterName = "Kaeya"
	Klee      CharacterName = "Klee"
	Lisa      CharacterName = "Lisa"
	Lumine    CharacterName = "Lumine"
	Miko      CharacterName = "Yae Miko"
	Mona      CharacterName = "Mona"
	Ningguang CharacterName = "Ningguang"
	Noelle    CharacterName = "Noelle"
	Qiqi      CharacterName = "Qiqi"
	Raiden    CharacterName = "Raiden Shogun"
	Razor     CharacterName = "Razor"
	Rosaria   CharacterName = "Rosaria"
	Sara      CharacterName = "Kujou Sara"
	Sayu      CharacterName = "Sayu"
	Shenhe    CharacterName = "Shenhe"
	Shinobu   CharacterName = "Kuki Shinobu"
	Sucrose   CharacterName = "Sucrose"
	Thoma     CharacterName = "Thoma"
	Traveler  CharacterName = "Traveler"
	Venti     CharacterName = "Venti"
	Xiao      CharacterName = "Xiao"
	Xiangling CharacterName = "Xiangling"
	Xingqiu   CharacterName = "Xingqiu"
	Xinyan    CharacterName = "Xinyan"
	Yanfei    CharacterName = "Yanfei"
	Yelan     CharacterName = "Yelan"
	Yoimiya   CharacterName = "Yoimiya"
	YunJin    CharacterName = "Yun Jin"
	Zhongli   CharacterName = "Zhogli"
)

type CharacterName string
type CharaterLevel uint8

type Character struct {
	common.Entity
	Name CharacterName `json:"name"`

	stats map[uint8]CharacterStats
}
