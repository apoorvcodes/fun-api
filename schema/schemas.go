package schema

type Joke struct {
    ID   int `json:"id"`
    Type   string `json:"type"`
    Setup    string    `json:"setup"`
	Punchline   string    `json:"punchline"`
}


type Fact struct {
    Text string `json:"text"`
    Source string `json:"source"`
}