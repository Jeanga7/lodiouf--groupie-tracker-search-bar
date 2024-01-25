package gt

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var BaseURL = "https://groupietrackers.herokuapp.com/api/"
var FileNames = []string{"vues/index.html", "vues/_header.html", "vues/_footer.html", "vues/_search_part.html"}
var ErrorFileNames = []string{"vues/error.html", "vues/_header.html", "vues/_footer.html", "vues/_search_part.html"}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(errors.New("PAGE NOT FOUND"), w, 404)
		return
	}

	switch r.Method {
	case "POST":
		ErrorPage(errors.New("METHODE NOT ALLOWED"), w, 405)
		return
		// r.ParseForm()
		// searchQuery := r.FormValue("search")

		// fmt.Fprintf(w, "Vous avez recherche: %s", searchQuery)
		// return
	default:
		tStart := time.Now()
		var data AllData
		parseDone := make(chan struct{})

		go func() {
			Parse(w)
			close(parseDone)
		}()

		<-parseDone

		data.Artists = *endpoints["artists"].(*[]Artist)
		data.Locations = *endpoints["locations"].(*Location)
		data.Dates = *endpoints["dates"].(*Date)
		// data.Relations = RedableRelations(*endpoints["relation"].(*Relation))
		data.Relations = *endpoints["relation"].(*Relation)
		// fmt.Println(Geocoding(strings.Join(UniqueCities(data.Locations.Index), ",")))

		
		for _, lot := range ChunkCities(UniqueCities(data.Locations.Index)) {
			for _, v := range lot {
				Geocoding(v)
			}
		}

		for _, val := range data.Artists {
			// data.FinalArtists = append(data.FinalArtists, FinalArtist{val.Id, val.Image, val.Name, []string{strings.Join(val.Members, ", ")}, val.CreationDate, FormatDates([]string{val.FirstAlbum})[0], data.Locations.Index[val.Id-1], data.Dates.Index[val.Id-1], data.Relations.Index[val.Id-1]})
			data.FinalArtists = append(data.FinalArtists, FinalArtist{val.Id, val.Image, val.Name, []string{strings.Join(val.Members, ", ")}, val.CreationDate, val.FirstAlbum, data.Locations.Index[val.Id-1], data.Dates.Index[val.Id-1], data.Relations.Index[val.Id-1]})
		}

		t, err := (template.ParseFiles(FileNames...))
		if err != nil {
			ErrorPage(err, w, http.StatusInternalServerError)
			return
		}
		tmpl := template.Must(t, err)

		err = tmpl.Execute(w, data)
		if err != nil {
			ErrorPage(err, w, http.StatusInternalServerError)
			return
		}

		elapsed := time.Since(tStart)
		log.Printf("Parsing took %.4fs\n", elapsed.Seconds())
	}
}

func ChunkCities(a []string) [][]string {
	var ch = 5
	var CitiesSlice []string
	result := make([][]string, 0, len(a)/ch)
	for len(a) >= ch {
		CitiesSlice, a = a[:ch], a[ch:]
		result = append(result, CitiesSlice)
	}
	if len(a) > 0 {
		result = append(result, a[:])
	}
	return result
}

func UniqueCities(i LocationIndex) []string {
	var citiesBool = make(map[string]bool)
	for _, loc := range i {
		for _, city := range loc.Locations {
			citiesBool[city] = true
		}
	}
	var citiesList []string
	for city := range citiesBool {
		citiesList = append(citiesList, city)
	}
	return citiesList
}

func Geocoding(cities string) string {
	url := "https://trueway-geocoding.p.rapidapi.com/Geocode?address=" + cities + "&language=en"
	fmt.Println("url", url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "0fca6dcab2mshc493be6db351be8p13df6bjsne0859659b3c0")
	req.Header.Add("X-RapidAPI-Host", "trueway-geocoding.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
	return string(body)
}

func ErrorPage(err error, w http.ResponseWriter, code int) {
	if err != nil {
		errorName := map[int]string{
			http.StatusNotFound:            http.StatusText(http.StatusNotFound),
			http.StatusMethodNotAllowed:    http.StatusText(http.StatusMethodNotAllowed),
			http.StatusInternalServerError: http.StatusText(http.StatusInternalServerError),
			http.StatusBadRequest:          http.StatusText(http.StatusBadRequest),
		}

		w.WriteHeader(code)

		tmpl, err := template.ParseFiles(ErrorFileNames...)
		if err != nil {
			panic(err)
		}
		tmpl = template.Must(tmpl, err)
		data := map[string]string{
			"error_code": strconv.Itoa(code),
			"error_name": errorName[code],
		}
		tmpl.Execute(w, data)
	}
}
