package models

type Character struct {
	Name          string `json:"name"`
	FullName      string `json:"fullname"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Rarity        string `json:"rarity"`
	Element       string `json:"element"`
	Weapontype    string `json:"weapontype"`
	Substat       string `json:"substat"`
	Gender        string `json:"gender"`
	Body          string `json:"body"`
	Association   string `json:"association"`
	Region        string `json:"region"`
	Affiliation   string `json:"affiliation"`
	Birthdaymmdd  string `json:"birthdaymmdd"`
	Birthday      string `json:"birthday"`
	Constellation string `json:"constellation"`
	Cv            cv     `json:"cv"`
	Costs         costs  `json:"costs"`
	Images        images `json:"images"`
	URL           url    `json:"url"`
	Version       string `json:"version"`
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
