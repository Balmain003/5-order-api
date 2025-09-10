package auth

type NumberRequest struct {
	Number string `json:"number"`
}
type NumberResponse struct {
	Session string `json:"sessionId"`
	Code    string `json:"code"`
}
