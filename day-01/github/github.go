package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const githubApiUrl = "https://api.github.com/users/%s"

func main() {
	name, repos, err := githubInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Printf("Name: %s, Public repos: %d\n", name, repos)
}

func demo() {
	res, err := http.Get("https://api.github.com/users/tebeka")

	if err != nil {
		log.Fatalf("Get: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Status not OK: %s", res.Status)
	}

	fmt.Printf("Content-type: %s\n", res.Header.Get("Content-type"))
	io.Copy(os.Stdout, res.Body)
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatalf("error: can't copy %s", err)
	}

	/* JSON <-> Go
	true/false ,-> true/false
	string <-> string
	null <-> nil
	number <-> float64, int64, int, uint64, uint, ...
	array <-> [[]interface{}
	object <-> map[string]interface{}

	JSON -> io.Reader -> json.NewDecoder -> Decode -> Go
	JSON -> []byte -> json.Unmarshal -> Go
	Go -> json.Marshal -> []byte -> JSON
	*/
	var r reply
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode %s", err)
	}
	fmt.Printf("Name: %s, Public repos: %d\n", r.Name, r.PublicRepos)
	fmt.Printf("r: %#v\n", r)

}

type reply struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func githubInfo(login string) (string, int, error) {
	// var r struct {
	// 	Name        string `json:"name"`
	// 	PublicRepos int    `json:"public_repos"`
	// }

	var r reply
	res, err := http.Get(fmt.Sprintf(githubApiUrl, url.PathEscape(login)))
	if err != nil {
		return "", 0, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("status not OK: %s", res.Status)
	}

	dec := json.NewDecoder(res.Body)

	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.PublicRepos, nil
}
