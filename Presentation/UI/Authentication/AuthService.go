package Authentication

import (
	"WB2/Application/Contracts/UserServices"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type AuthService struct {
	getService UserServices.IGetService
}

func CreateAuthService(getService UserServices.IGetService) AuthService {
	return AuthService{getService}
}

func (service AuthService) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		showForm(w)
		return
	}

	if r.Method == http.MethodPost {
		formData, err := processForm(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Получены данные: ID = %d, Key = %s\n", formData.UserID, formData.UserKey)
		// TODO: доделываем авторизированное окно

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Данные успешно получены!"))

		if err != nil {
			http.Error(w, "Ошибка при отправке ответа: "+err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func showForm(w http.ResponseWriter) {
	htmlTemplate, err := os.ReadFile("Presentation/UI/resources/authentication.html")
	if err != nil {
		http.Error(w, "Ошибка при обращении к шаблону", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("form").Parse(string(htmlTemplate))
	if err != nil {
		http.Error(w, "Ошибка при парсинге шаблона", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
		return
	}
}

type FormData struct {
	UserID  int
	UserKey string
}

func processForm(r *http.Request) (FormData, error) {
	err := r.ParseForm()
	if err != nil {
		return FormData{}, fmt.Errorf("ошибка при парсинге формы: %w", err)
	}

	userIdStr := r.Form.Get("userId")
	userKey := r.Form.Get("userKey")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return FormData{}, fmt.Errorf("некорректный формат ID: %w", err)
	}

	return FormData{UserID: userId, UserKey: userKey}, nil
}
