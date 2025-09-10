package auth

import (
	"authorizate/pkg/req"
	"authorizate/pkg/res"
	"net/http"
	"regexp"
)

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /verify", VerifyNumber())
	return mux
}

func VerifyNumber() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[NumberRequest](&w, r)
		if err != nil {
			return
		}
		matched, _ := regexp.MatchString(`^\+?[1-9]\d{1,14}$`, body.Number)
		if !matched {
			res.Json(w, "Invalid phone number format", http.StatusBadRequest)
			return
		}
		session := res.RandStringRunes(10)
		data := NumberResponse{
			Session: session,
			Code:    "5432",
		}
		res.Json(w, data, 200)
	}
}
