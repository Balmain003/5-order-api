package auth

type NumberRequest struct {
	Number string `json:"phone"`
}
type NumberResponse struct {
	Session string `json:"sessionId"`
	Code    string `json:"code"`
}

type VerifyCodeRequest struct {
	Session string `json:"session"`
	Code    string `json:"code"`
}

type VerifyCodeResponse struct {
	Token string `json:"token"`
}
