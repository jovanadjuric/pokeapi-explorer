package pokeapi

type Pokemon struct {
	Name  *string `json:"name"`
	Stats []struct {
		Stat struct {
			Name *string `json:"name"`
		} `json:"stat"`
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name *string `json:"name"`
		}
		Slot int `json:"slot"`
	}
	BaseExperience int `json:"base_experience"`
	Height         int `json:"height"`
	Weight         int `json:"weight"`
	Id             int `json:"id"`
}
