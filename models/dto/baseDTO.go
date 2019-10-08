package dto

type BaseDTO struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Active   bool   `json:"active"`
	ObjectId string `json:"accountid,omitempty"`
	Token    string `json:"token,omitempty"`
}
