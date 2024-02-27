package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func bibleid() []bibleId {
	url := "https://api.scripture.api.bible/v1/bibles"

	bibleClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("api-key", os.Getenv("API_KEY"))

	res, getErr := bibleClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var body1 map[string]interface{}
	jsonErr := json.Unmarshal(body, &body1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	bibleData := body1["data"].([]interface{}) //[0].(map[string]interface{})

	bibleIds := []bibleId{}
	for i := 0; i < len(bibleData); i++ {
		//if bibleData[i].(map[string]interface{})["language"].(map[string]interface{})["name"].(string) == "English" {
		//fmt.Println(bibleData[i].(map[string]interface{})["name"].(string))
		//}
		bibleIds = append(bibleIds, bibleId{
			Id:   bibleData[i].(map[string]interface{})["id"].(string),
			Name: bibleData[i].(map[string]interface{})["name"].(string),
		})
	}
	//fmt.Println(bibleData["name"].(string))

	return bibleIds
}
