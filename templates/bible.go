package templates

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/tursodatabase/go-libsql"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type BibleId struct {
	Id   string
	Name string
}

type Lang struct {
	id              string
	name            string
	nameLocal       string
	script          string
	scriptDirection string
}

type BookId = BibleId

type ChapId = BibleId

var (
	language_code string
	language      string
	bibleapi      bool
)

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Languages() []Lang {
	bibleapi = false

	langIds := []Lang{}

	if bibleapi {
		url := "https://api.scripture.api.bible/v1/bibles"

		bibleClient := http.Client{
			Timeout: time.Second * 20,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Println(err)
			//return []Lang{}
		}

		req.Header.Set("api-key", os.Getenv("API_KEY"))

		res, getErr := bibleClient.Do(req)
		if getErr != nil {
			log.Println(getErr)
			//return []Lang{}
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Println(readErr)
			//return []Lang{}
		}

		var body1 map[string]interface{}
		jsonErr := json.Unmarshal(body, &body1)
		if jsonErr != nil {
			log.Println(jsonErr)
			//return []Lang{}
		}
		bibleData := body1["data"].([]interface{}) //[0].(map[string]interface{})

		for i := 0; i < len(bibleData); i++ {
			langIds = append(langIds, Lang{
				id:        bibleData[i].(map[string]interface{})["language"].(map[string]interface{})["id"].(string),
				name:      bibleData[i].(map[string]interface{})["language"].(map[string]interface{})["name"].(string),
				nameLocal: bibleData[i].(map[string]interface{})["language"].(map[string]interface{})["nameLocal"].(string),
			})
		}
	}
	dbName := "file:./test.db"

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
		//os.Exit(1)
	}
	rows, err := db.Query("select language_code, language from translations")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&language_code, &language)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(language_code, language)
		langIds = append(langIds, Lang{
			id:        language_code,
			name:      language,
			nameLocal: language,
		})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//fmt.Println(bibleData["name"].(string))

	dupremlangs := removeDuplicate(langIds)

	sort.Slice(dupremlangs, func(i, j int) bool {
		return dupremlangs[i].name < dupremlangs[j].name
	})
	return dupremlangs
}

func Bibleid(langid string) []BibleId {
	bibleIds := []BibleId{}

	if bibleapi {
		url := "https://api.scripture.api.bible/v1/bibles?language=" + langid

		bibleClient := http.Client{
			Timeout: time.Second * 20,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
			return []BibleId{}
		}

		req.Header.Set("api-key", os.Getenv("API_KEY"))

		res, getErr := bibleClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
			return []BibleId{}
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
			return []BibleId{}
		}

		var body1 map[string]interface{}
		jsonErr := json.Unmarshal(body, &body1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
			return []BibleId{}
		}
		bibleData := body1["data"].([]interface{}) //[0].(map[string]interface{})

		for i := 0; i < len(bibleData); i++ {
			//if bibleData[i].(map[string]interface{})["language"].(map[string]interface{})["name"].(string) == "English" {
			//fmt.Println(bibleData[i].(map[string]interface{})["name"].(string))
			//}
			bibleIds = append(bibleIds, BibleId{
				Id:   bibleData[i].(map[string]interface{})["id"].(string),
				Name: bibleData[i].(map[string]interface{})["name"].(string),
			})
		}
	}
	//fmt.Println(bibleData["name"].(string))
	dbName := "file:./test.db"

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
		//os.Exit(1)
	}
	rows, err := db.Query("select id, name from translations where language_code = ?", langid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&language_code, &language)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(language_code, language)
		bibleIds = append(bibleIds, BibleId{
			Id:   language_code,
			Name: language,
		})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return bibleIds
}

func Bookid(biblId string) []BookId {
	bookIds := []BookId{}
	if bibleapi {
		url := fmt.Sprintf("https://api.scripture.api.bible/v1/bibles/%s/books", biblId)

		bibleClient := http.Client{
			Timeout: time.Second * 20,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
			return []BookId{}
		}

		req.Header.Set("api-key", os.Getenv("API_KEY"))

		res, getErr := bibleClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
			return []BookId{}
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
			return []BookId{}
		}

		var body1 map[string]interface{}
		jsonErr := json.Unmarshal(body, &body1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
			return []BookId{}
		}
		bibleBookData := body1["data"].([]interface{}) //[0].(map[string]interface{})

		for i := 0; i < len(bibleBookData); i++ {
			//if bibleData[i].(map[string]interface{})["language"].(map[string]interface{})["name"].(string) == "English" {
			//fmt.Println(bibleData[i].(map[string]interface{})["name"].(string))
			//}
			bookIds = append(bookIds, BookId{
				Id:   bibleBookData[i].(map[string]interface{})["id"].(string),
				Name: bibleBookData[i].(map[string]interface{})["name"].(string),
			})
		}
		//fmt.Println(bibleData["name"].(string))
	}

	dbName := "file:./test.db"

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
		//os.Exit(1)
	}
	rows, err := db.Query("select distinct book_id, book from verses s where s.translation_id = ?", biblId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&language_code, &language)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(language_code, language)
		bookIds = append(bookIds, BookId{
			Id:   language_code,
			Name: language,
		})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return bookIds
}

func Chapid(biblId string, bookId string) []ChapId {
	chapIds := []ChapId{}

	if bibleapi {
		url := fmt.Sprintf("https://api.scripture.api.bible/v1/bibles/%s/books/%s/chapters", biblId, bookId)

		bibleClient := http.Client{
			Timeout: time.Second * 20,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
			return []ChapId{}
		}

		req.Header.Set("api-key", os.Getenv("API_KEY"))

		res, getErr := bibleClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
			return []ChapId{}
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
			return []ChapId{}
		}

		var body1 map[string]interface{}
		jsonErr := json.Unmarshal(body, &body1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
			return []ChapId{}
		}
		bibleChapData := body1["data"].([]interface{}) //[0].(map[string]interface{})

		for i := 0; i < len(bibleChapData); i++ {
			//if bibleData[i].(map[string]interface{})["language"].(map[string]interface{})["name"].(string) == "English" {
			//fmt.Println(bibleData[i].(map[string]interface{})["name"].(string))
			//}
			chapIds = append(chapIds, ChapId{
				Id:   bibleChapData[i].(map[string]interface{})["id"].(string),
				Name: bibleChapData[i].(map[string]interface{})["number"].(string),
			})
		}
	}

	dbName := "file:./test.db"

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
		//os.Exit(1)
	}
	rows, err := db.Query("select distinct book_id||'.'||chapter, chapter from verses s where s.translation_id = ? and s.book_id = ?", biblId, bookId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&language_code, &language)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(language_code, language)
		chapIds = append(chapIds, ChapId{
			Id:   language_code,
			Name: language,
		})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//fmt.Println(bibleData["name"].(string))
	//chapIds.push(chapIds.shift())
	//fmt.Println(shiftfirsttoend(chapIds))
	if chapIds[0].Name == "intro" {
		chapIds = shiftfirsttoend(chapIds)
	}
	return chapIds
}

func Biblecontent(bibleId string, chapId string) string {
	biblecontent := ""
	if bibleapi {
		url := fmt.Sprintf("https://api.scripture.api.bible/v1/bibles/%s/passages/%s?content-type=html&include-notes=false&include-titles=true&include-chapter-numbers=false&include-verse-numbers=true&include-verse-spans=false&use-org-id=false", bibleId, chapId)
		//fmt.Println(bibleId, chapId)
		//fmt.Println(url)
		bibleClient := http.Client{
			Timeout: time.Second * 20,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
			return ""
		}

		req.Header.Set("api-key", os.Getenv("API_KEY"))

		res, getErr := bibleClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
			return ""
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
			return ""
		}

		var body1 map[string]interface{}
		jsonErr := json.Unmarshal(body, &body1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
			return ""
		}
		bibleContentData := body1["data"].(interface{}) //[0].(map[string]interface{})

		biblecontent = bibleContentData.(map[string]interface{})["content"].(string)
		//for i := 0; i < len(bibleBookData); i++ {
		//if bibleData[i].(map[string]interface{})["language"].(map[string]interface{})["name"].(string) == "English" {
		//fmt.Println(bibleData[i].(map[string]interface{})["name"].(string))
		//}
		//bookIds = append(bookIds, BookId{
		//	Id:   bibleBookData[i].(map[string]interface{})["id"].(string),
		//		Name: bibleBookData[i].(map[string]interface{})["name"].(string),
		//	})
		//	biblecontent
		//}
		//fmt.Println(bibleData["name"].(string))
		//fmt.Println(biblecontent)
	}

	bookId := strings.Split(chapId, ".")[0]

	chapterId := strings.Split(chapId, ".")[1]

	dbName := "file:./test.db"

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
		//os.Exit(1)
	}
	rows, err := db.Query("select verse,text from verses s where s.translation_id = ? and s.book_id = ? and s.chapter = ?", bibleId, bookId, chapterId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&language_code, &language)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(language_code, language)
		//chapIds = append(chapIds, ChapId{
		//	Id:   language_code,
		//	Name: language,
		//})
		biblecontent = fmt.Sprintf("%s<p>%s %s<br>", biblecontent, language_code, language)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//select verse,text from verses s where s.translation_id = ? and s.book_id = ? and s.chapter = ?

	return biblecontent

}

func shiftfirsttoend[T any](s []T) []T {
	if len(s) == 0 {
		return s
	}
	z := append(s, s[0])
	return z[1:]
}
