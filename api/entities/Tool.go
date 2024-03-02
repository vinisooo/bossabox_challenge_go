package entities

type Tool struct {
	Id    		int      `json:"id"`
	Title 		string   `json:"title"`
	Link  		string   `json:"link"`
	Description string   `json:"description"`
	Tags 		[]string `json:"tags"`
}