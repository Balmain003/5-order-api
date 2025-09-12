package auth

import (
	"authorizate/config"
	"authorizate/middleware"
	"authorizate/pkg/jwt"
	"authorizate/pkg/req"
	"authorizate/pkg/res"
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

type LoginHandler struct {
	*config.Config
	sessions map[string]string
	mtx      sync.Mutex
}

func NewHandler(cfg *config.Config) *http.ServeMux {
	handler := &LoginHandler{
		Config:   cfg,
		sessions: make(map[string]string),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /verify", handler.VerifyNumber)
	mux.HandleFunc("POST /login", handler.VerifyCode)

	authMiddleware := middleware.AuthMiddleware(cfg)
	mux.Handle("GET /profile", authMiddleware(http.HandlerFunc(handler.ProfileHandler)))

	return mux
}

func (h *LoginHandler) VerifyNumber(w http.ResponseWriter, r *http.Request) {
	body, err := req.HandleBody[NumberRequest](w, r)
	if err != nil {
		return
	}
	matched, _ := regexp.MatchString(`^\+?[1-9]\d{1,14}$`, body.Phone)
	if !matched {
		res.Json(w, "Invalid phone number format", http.StatusBadRequest)
		return
	}
	session := res.RandStringRunes(10)

	h.mtx.Lock()
	h.sessions[session] = body.Phone
	h.mtx.Unlock()

	data := NumberResponse{
		Session: session,
		Code:    "5432",
	}
	res.Json(w, data, 200)
}

func (h *LoginHandler) VerifyCode(w http.ResponseWriter, r *http.Request) {
	body, err := req.HandleBody[VerifyCodeRequest](w, r)
	if err != nil {
		return
	}
	if body.Code != "5432" {
		res.Json(w, "Invalid code", http.StatusBadRequest)
		return
	}
	h.mtx.Lock()
	phone, exists := h.sessions[body.Session]
	h.mtx.Unlock()
	if !exists {
		res.Json(w, "Invalid session", http.StatusBadRequest)
		return
	}
	token, err := jwt.NewJwt(h.Config.Numb.Secret).CreateJwt(phone)
	if err != nil {
		res.Json(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	res.Json(w, VerifyCodeResponse{
		Token: token,
	}, 200)
}

func (h *LoginHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	phone, _ := r.Context().Value(middleware.PhoneContextKey).(string)

	fmt.Fprintln(w, "Phone from context:", phone)
}
