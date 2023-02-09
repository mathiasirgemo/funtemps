package funfacts

type FunFact struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {
	funFacts := FunFact{
		Sun: []string{
			"Temperatur i Solens kjerne	15 000 000C°",
			"Temperatur på ytre lag av Solen	5778K",
		},
		Luna: []string{
			"Temperatur på Månens overflate om natten	-183°C",
			"Temperatur på Månens overflate om dagen",
		},
		Terra: []string{
			"Høyeste temperatur målt på Jordens overflate	56.7°C",
			"Laveste temperatur målt på Jordens overflate	-89,4°C",
			"Temperatur i Jordens indre kjerne	9392K",
		},
	}

	switch about {
	case "Sun":
		return funFacts.Sun
	case "luna":
		return funFacts.Luna
	case "terra":
		return funFacts.Terra
	default:
		return []string{"Ingen funfacts funnet.."}
	}
}