package schema

type Joke struct {
    ID   int `json:"id"`
    Type   string `json:"type"`
    Setup    string    `json:"setup"`
	Punchline   string    `json:"punchline"`
}
