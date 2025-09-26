package verify

import (
	"fmt"
	"net/http"
	"net/smtp"
	"validation-api/configs"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct {
	*configs.EmailVerifyConfig
}

type VerifyHandlerDeps struct {
	*configs.EmailVerifyConfig
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps)  {
	handler := &VerifyHandler{
		EmailVerifyConfig: deps.EmailVerifyConfig,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request)  {
		fmt.Println("Send")
		mail := &email.Email{
			To: []string{"test@example.com"},
			Subject: "Тема письма",
			Text: []byte("Тело письма."),
		}
		mail.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Email, handler.Password, handler.Address))
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request)  {
		fmt.Println("Verify")
	}
}