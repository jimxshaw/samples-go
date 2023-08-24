package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	name, repos, err := githubInfo(ctx, "jimxshaw")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(name, repos)
}

func githubInfo(ctx context.Context, login string) (string, int, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	//resp, err := http.Get(url)
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

	//fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: can't copy - %s", err)
	// }

	// var r Reply

	// Anonymous struct, typically used when working with APIs like this.
	var r struct {
		Name string
		// Public_Repos int // The field needs _ underscore to match the JSON field
		// or use the json tag for the field.
		PublicRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	//fmt.Println(r)
	// fmt.Printf("%#v\n", r)

	return r.Name, r.PublicRepos, nil
}

// type Reply struct {
// 	Name string
// 	// Public_Repos int // The field needs _ underscore to match the JSON field
// 	// or use the json tag for the field.
// 	PublicRepos int `json:"public_repos"`
// }

/* JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int32, int64, int, uint8, ...
array <-> []any or []interface{} before Go 1.18
object <-> map[string]any, struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte (in-memory sequence of bytes read from a DB or a file) -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte (in-memory sequence of bytes read from a DB or a file) -> JSON: json.Marshal
*/
