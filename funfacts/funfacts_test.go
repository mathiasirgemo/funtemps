package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	expectedSunFacts := []string{
		"Temperatur i Solens kjerne	15 000 000C°",
    	"Temperatur på ytre lag av Solen	5778K",
	}

	sunFacts := GetFunFacts("Sun")
	if !reflect.DeepEqual(sunFacts, expectedSunFacts) {
		t.Errorf("Expected %v, but got %v", expectedSunFacts, sunFacts)
	}

	expectedLunaFacts := []string{
		"Temperatur på Månens overflate om natten	-183°C",
    	"Temperatur på Månens overflate om dagen",
	}

	lunaFacts := GetFunFacts("Luna")
	if !reflect.DeepEqual(lunaFacts, expectedLunaFacts) {
		t.Errorf("Expected %v, but got %v", expectedLunaFacts, lunaFacts)
	}

	expectedTerraFacts := []string{
		"Høyeste temperatur målt på Jordens overflate	56.7°C",
    	"Laveste temperatur målt på Jordens overflate	-89,4°C",
    	"Temperatur i Jordens indre kjerne	9392K",
	}

	terraFacts := GetFunFacts("Terra")
	if !reflect.DeepEqual(terraFacts, expectedTerraFacts) {
		t.Errorf("Expected %v, but got %v", expectedTerraFacts, terraFacts)
	}

	expectedDefaultFacts := []string{"Ingen funfacts funnet.."}

	defaultFacts := GetFunFacts("unknown")
	if !reflect.DeepEqual(defaultFacts, expectedDefaultFacts) {
		t.Errorf("Expected %v, but got %v", expectedDefaultFacts, defaultFacts)
	}
}