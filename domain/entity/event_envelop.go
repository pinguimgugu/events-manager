package entity

type EventEnvelop struct {
	Name string `json:"name"`
	Metadata map[string]interface{} `json:"metadata"`
}