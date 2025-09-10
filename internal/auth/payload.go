package auth

type NumberRequest struct {
	Number string `json:"number"`
}
type NumberResponse struct {
	Session string `json:"session"`
	Code    string `json:"code"`
}

type VerifyCodeRequest struct {
	Session string `json:"session"`
	Code    string `json:"code"`
}

type VerifyCodeResponse struct {
	Token string `json:"token"`
}
