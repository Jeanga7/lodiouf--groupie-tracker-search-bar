package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type LocationDataList []LocationData

type LocationData struct {
	DisplayName string `json:"display_name"`
	Lon         string `json:"lon"`
	Lat         string `json:"lat"`
}

func Geocoding(city string) (LocationDataList, error) {
	token := "pk.2f226e3203c56af77056c10e939e3a0a"
	url := fmt.Sprintf("https://us1.locationiq.com/v1/search?key=%s&q=%s&format=json", token, city)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error making HTTP request: %v", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading HTTP response: %v", err)
		return nil, err
	}

	var locationDataList LocationDataList
	err = json.Unmarshal(body, &locationDataList)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return nil, err
	}

	for _, locationData := range locationDataList {
		filename := locationData.DisplayName + ".json"
		file, err := os.Create(filename)
		if err != nil {
			log.Printf("Error creating file: %v", err)
			continue
		}
		defer file.Close()

		err = json.NewEncoder(file).Encode(locationData)
		if err != nil {
			log.Printf("Error encoding JSON: %v", err)
		}
	}

	return locationDataList, nil
}

func main() {
	addresses := []string{
		"vienna-austria", "prague-czechia", "werchter-belgium", "warsaw-poland", "michigan-usa", "omaha-usa", "rosemont-usa", "san_francisco-usa", "los_angeles-usa", "lausanne-switzerland", "arras-france", "melbourne-australia", "lyon-france", "birmingham-uk", "sanya-china", "houston-usa", "new_orleans-usa", "imola-italy", "riyadh-saudi_arabia", "mexico_city-mexico", "madrid-spain", "groningen-netherlands", "dallas-usa", "victoria-australia", "new_south_wales-australia", "freyming_merlebach-france", "burswood-australia", "taipei-taiwan", "buenos_aires-argentina", "nagoya-japan", "bratislava-slovakia", "monchengladbach-germany", "amsterdam-netherlands", "quebec-canada", "eindhoven-netherlands", "bilbao-spain", "illinois-usa", "monterrey-mexico", "paris-france", "kansas_city-usa", "brooklyn-usa", "krakow-poland", "lima-peru", "pagney_derriere_barine-france", "salem-germany", "munich-germany", "st_gallen-switzerland", "manila-philippines", "athens-greece", "san_isidro-argentina", "bogota-colombia", "merkers-germany", "florida-usa", "boulogne_billancourt-france", "oslo-norway", "queensland-australia", "yogyakarta-indonesia", "toronto-usa", "pico_rivera-usa", "lisbon-portugal", "aarhus-denmark", "west_melbourne-australia", "sydney-australia", "columbia-usa", "turku-finland", "budapest-hungary", "arizona-usa", "oregon-usa", "montreal-canada", "scheessel-germany", "mannheim-germany", "san_jose-costa_rica", "zaragoza-spain", "atlanta-usa", "mumbai-india", "berwyn-usa", "georgia-usa", "california-usa", "las_vegas-usa", "seattle-usa", "amityville-usa", "abu_dhabi-united_arab_emirates", "nimes-france", "windsor-canada", "brisbane-australia", "aalborg-denmark", "santiago-chile", "detroit-usa", "seville-spain", "gdynia-poland", "london-uk", "cardiff-uk", "barcelona-spain", "boston-usa", "glasgow-uk", "grand_rapids-usa", "newark-usa", "philadelphia-usa", "penrose-new_zealand", "stockholm-sweden", "brixton-uk", "oulu-finland", "canton-usa", "washington-usa", "oklahoma-usa", "berlin-germany", "cuxhaven-germany", "papeete-french_polynesia", "sochaux-france", "minneapolis-usa", "dubai-united_arab_emirates", "ostrava-czechia", "indianapolis-usa", "belo_horizonte-brazil", "vancouver-canada", "osaka-japan", "new_york-usa", "st_louis-usa", "hershey-usa", "inglewood-usa", "brasilia-brazil", "noumea-new_caledonia", "del_mar-usa", "minnesota-usa", "leipzig-germany", "skanderborg-denmark", "colorado-usa", "oakland-usa", "rotselaar-belgium", "texas-usa", "copenhagen-denmark", "aberdeen-uk", "jakarta-indonesia", "changzhou-china", "cincinnati-usa", "recife-brazil", "chicago-usa", "frankfurt-germany", "maine-usa", "gothenburg-sweden", "burriana-spain", "pittsburgh-usa", "antwerp-belgium", "north_carolina-usa", "dunedin-new_zealand", "milan-italy", "uniondale-usa", "alabama-usa", "playa_del_carmen-mexico", "nevada-usa", "napoca-romania", "madison-usa", "toronto-canada", "utah-usa", "dublin-ireland", "huizhou-china", "saitama-japan", "auckland-new_zealand", "landgraaf-netherlands", "wellington-new_zealand", "seoul-south_korea", "hong_kong-china", "charlotte-usa", "frauenfeld-switzerland", "manchester-uk", "klagenfurt-austria", "hamburg-germany", "missouri-usa", "new_hampshire-usa", "florence-italy", "bangkok-thailand", "willemstad-netherlands_antilles", "sao_paulo-brazil", "westcliff_on_sea-uk", "south_carolina-usa", "rio_de_janeiro-brazil", "montreal-usa", "anaheim-usa", "massachusetts-usa", "porto_alegre-brazil", "minsk-belarus", "doha-qatar", "cleveland-usa", "nickelsdorf-austria", "pennsylvania-usa", "sion-switzerland", "la_plata-argentina", "dusseldorf-germany", "zurich-switzerland",
	}

	for _, address := range addresses {
		locationDataList, err := Geocoding(address)
		jsonData, err := json.MarshalIndent(locationDataList, "", " ")
		if err != nil {
			log.Printf("Error JSON: %v", err)
			continue
		}
		fmt.Println(string(jsonData))
		time.Sleep(2 * time.Second)
	}
}
