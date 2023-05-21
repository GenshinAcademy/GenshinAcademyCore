package models

import "ga/pkg/genshindb_wrapper/enums"

type Character struct {
	Id           uint              `json:"id"`
	Name         string            `json:"name"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	WeaponText   enums.WeaponText  `json:"weaponText"`
	WeaponType   enums.WeaponType  `json:"weaponType"`
	BodyType     enums.BodyType    `json:"bodyType"`
	Gender       enums.Gender      `json:"gender"`
	QualityType  enums.QualityType `json:"qualityType"`
	Rarity       enums.Rarity      `json:"rarity"`
	Birthdaymmdd string            `json:"birthdaymmdd"`
	Birthday     string            `json:"birthday"`
	ElementText  enums.ElementText `json:"elementText"`
	ElementType  enums.ElementType `json:"elementType"`
	Affiliation  string            `json:"affiliation"`
	Association  string            `json:"association"`
	Region       string            `json:"region"`
	// TODO: Enums
	SubstatText string `json:"substatType"`
	// TODO: Enums
	SubstatType string `json:"substatText"`
	// TODO: Enums
	Constellation string `json:"constellation"`
	Cv            cv     `json:"cv"`
	Costs         costs  `json:"costs"`
}

type CharacterWeb struct {
	Character
	Images  images `json:"images"`
	URL     url    `json:"url"`
	Version string `json:"version"`
}

type images struct {
	Card            string `json:"card"`
	Portrait        string `json:"portrait"`
	Icon            string `json:"icon"`
	Sideicon        string `json:"sideicon"`
	Cover1          string `json:"cover1"`
	Cover2          string `json:"cover2"`
	HoyolabAvatar   string `json:"hoyolab-avatar"`
	Nameicon        string `json:"nameicon"`
	Nameiconcard    string `json:"nameiconcard"`
	Namegachasplash string `json:"namegachasplash"`
	Namegachaslice  string `json:"namegachaslice"`
	Namesideicon    string `json:"namesideicon"`
}

type costs struct {
	Ascend1 []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"ascend1"`
	Ascend2 []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"ascend2"`
	Ascend3 []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"ascend3"`
	Ascend4 []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"ascend4"`
	Ascend5 []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"ascend5"`
	Ascend6 []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"ascend6"`
}

type cv struct {
	English  string `json:"english"`
	Chinese  string `json:"chinese"`
	Japanese string `json:"japanese"`
	Korean   string `json:"korean"`
}

type url struct {
	Fandom string `json:"fandom"`
}
