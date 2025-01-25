package Authentication

import (
	"WB2/Application/Contracts/UserServices"
	"net/http"
	"os"
)

type AuthService struct {
	getService UserServices.IGetService
}

func CreateAuthService(getService UserServices.IGetService) AuthService {
	return AuthService{getService}
}

func (service AuthService) Handle(w http.ResponseWriter, r *http.Request) {

	html, err := os.ReadFile("Presentation/UI/resources/authentication.html")
	if err != nil {
		http.Error(w, "Ошибка при чтении файла"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = w.Write(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
