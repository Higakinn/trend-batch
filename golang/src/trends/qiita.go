package trends

import (
	"os"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
	"regexp"
)

type Response []struct {
	IsNewArrival bool `json:"isNewArrival"`
	Node         struct {
		EncryptedID string    `json:"encryptedId"`
		LikesCount  int       `json:"likesCount"`
		LinkURL     string    `json:"linkUrl"`
		PublishedAt time.Time `json:"publishedAt"`
		Title       string    `json:"title"`
		UUID        string    `json:"uuid"`
		Author      struct {
			ProfileImageURL string `json:"profileImageUrl"`
			URLName         string `json:"urlName"`
			Name            string `json:"name"`
		} `json:"author"`
		Organization    interface{}   `json:"organization"`
		FollowingLikers []interface{} `json:"followingLikers"`
		Tags            []struct {
			URLName string `json:"urlName"`
			Name    string `json:"name"`
		} `json:"tags"`
	} `json:"node"`
}

type Qiita struct {
	Name string
}

func test(url string) (string){
	resp, err := http.Get(url)
	if err != nil {
			fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
			fmt.Println("http read error")
			return "error"
	}

	src := []byte(string(body))
	re := regexp.MustCompile(`"(https://twitter.com/.*?)"`)
	group := re.FindSubmatch(src)
	// fmt.Println(re.FindAllString(src, -1))
	// fmt.Println(group[0])
	return string(group[0])
	// for k, v := range group {
	// 	fmt.Printf("%d. %s\n", k, v)
	// }
	// fmt.Println(match)
	// fmt.Println(src)
}

func (q *Qiita) FetchTrends() ([]TrendResponse, error){
	url := os.Getenv("TREND_QIITA_URL")
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
			log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
			log.Fatal(err)
	}

	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("json.Unmarshal err=%s", err.Error())
	}
	fmt.Println(response[0].Node.Title)
	fmt.Println(response[0].Node.LinkURL)

	for _,k := range response {
		test("https://qiita.com/" + k.Node.Author.URLName)
	}

	test("https://qiita.com/higakin")
	test("https://qiita.com/" + response[0].Node.Author.URLName)

	var results []TrendResponse
	for _,k := range response {
		res := TrendResponse{
			Title: k.Node.Title,
			Url: k.Node.LinkURL,
			TwitterUrl: test("https://qiita.com/" + k.Node.Author.URLName),
		}
		results = append(results, res)
	}
	return results,nil
}


