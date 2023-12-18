package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()
	fmt.Println(githubInfo(ctx, "tebeka"))
}

// githubInfo returns name and number of public repos for login
func githubInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	// resp,err := http.Get(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	defer resp.Body.Close()

	// fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// var r Reply

	var r struct { // anonymous struct
		Name string
		// Public_Repos int
		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.NumRepos, nil
}

/*
type Reply struct {
	Name string
	// Public_Repos int
	NumRepos int `json:"public_repos"`
}
*/

/* JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int32, int64, int, uint8, ...
array <-> []any ([]interface{})
object <-> map[string]any, struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// )

// const githubApiUrl = "https://api.github.com/users/%s"

// func main() {
// 	name, repos, err := githubInfo("tebeka")
// 	if err != nil {
// 		log.Fatalf("error: %s", err)
// 	}
// 	fmt.Printf("Name: %s, Public repos: %d\n", name, repos)
// }

// func demo() {
// 	res, err := http.Get("https://api.github.com/users/tebeka")

// 	if err != nil {
// 		log.Fatalf("Get: %s", err)
// 	}

// 	if res.StatusCode != http.StatusOK {
// 		log.Fatalf("Status not OK: %s", res.Status)
// 	}

// 	fmt.Printf("Content-type: %s\n", res.Header.Get("Content-type"))
// 	io.Copy(os.Stdout, res.Body)
// 	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
// 		log.Fatalf("error: can't copy %s", err)
// 	}

// 	/* JSON <-> Go
// 	true/false ,-> true/false
// 	string <-> string
// 	null <-> nil
// 	number <-> float64, int64, int, uint64, uint, ...
// 	array <-> [[]interface{}
// 	object <-> map[string]interface{}

// 	JSON -> io.Reader -> json.NewDecoder -> Decode -> Go
// 	JSON -> []byte -> json.Unmarshal -> Go
// 	Go -> json.Marshal -> []byte -> JSON
// 	*/
// 	var r reply
// 	dec := json.NewDecoder(res.Body)
// 	if err := dec.Decode(&r); err != nil {
// 		log.Fatalf("error: can't decode %s", err)
// 	}
// 	fmt.Printf("Name: %s, Public repos: %d\n", r.Name, r.PublicRepos)
// 	fmt.Printf("r: %#v\n", r)

// }

// type reply struct {
// 	Name        string `json:"name"`
// 	PublicRepos int    `json:"public_repos"`
// }

// func githubInfo(login string) (string, int, error) {
// 	// var r struct {
// 	// 	Name        string `json:"name"`
// 	// 	PublicRepos int    `json:"public_repos"`
// 	// }

// 	var r reply
// 	res, err := http.Get(fmt.Sprintf(githubApiUrl, url.PathEscape(login)))
// 	if err != nil {
// 		return "", 0, err
// 	}

// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		return "", 0, fmt.Errorf("status not OK: %s", res.Status)
// 	}

// 	dec := json.NewDecoder(res.Body)

// 	if err := dec.Decode(&r); err != nil {
// 		return "", 0, err
// 	}
// 	return r.Name, r.PublicRepos, nil
// }
