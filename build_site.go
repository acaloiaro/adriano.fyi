package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	build()
}

// In this example `messageId` is a value that
// needs to be stored in an environment variable to be
// used by the command that's going to be executed
func build() error {
	binary, err := exec.LookPath("hugo")
	if err != nil {
		return err
	}

	modePtr := flag.String("mode", "", "set application mode to either 'serve' or 'build'")

	flag.Parse()

	cmd := exec.Command(binary, *modePtr)
	env := os.Environ()

	lat, lon, currentLocationName := getCurrentLocation()
	env = append(env, fmt.Sprintf("HUGO_PARAMS.CurrentLat=%f", lat))
	env = append(env, fmt.Sprintf("HUGO_PARAMS.CurrentLon=%f", lon))
	env = append(env, fmt.Sprintf("HUGO_PARAMS.CurrentLocationName=%s", currentLocationName))

	cmd.Env = env

	cmdOut, _ := cmd.StdoutPipe()
	cmdErr, _ := cmd.StderrPipe()

	startErr := cmd.Start()
	if startErr != nil {
		return startErr
	}

	// read stdout and stderr
	stdOutput, _ := ioutil.ReadAll(cmdOut)
	errOutput, _ := ioutil.ReadAll(cmdErr)

	fmt.Printf("STDOUT: %s\n", stdOutput)
	fmt.Printf("ERROUT: %s\n", errOutput)

	err = cmd.Wait()
	return err
}

type point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func getCurrentLocation() (lat, lon float64, currentLocationName string) {
	//{{ $.Scratch.Set "description" (printf "I'm current in %s" $place.properties.label) }}
	url := "https://adriano.fyi/current_location"

	client := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	p := point{}
	jsonErr := json.Unmarshal(body, &p)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	lat, lon = p.Lat, p.Lon
	currentLocationName = getCurrentLocationName(lon, lat)

	return
}

type feature struct {
	PlaceName string `json:"place_name"`
}

type peliasResponse struct {
	Features []feature `json:"features"`
}

func getCurrentLocationName(lon, lat float64) (locationName string) {
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	url := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%f,%f.json?access_token=%s&types=locality,place", lon, lat, os.Getenv("MAPBOX_API_KEY"))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	r := peliasResponse{}
	jsonErr := json.Unmarshal(body, &r)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	locationName = r.Features[0].PlaceName

	return
}
