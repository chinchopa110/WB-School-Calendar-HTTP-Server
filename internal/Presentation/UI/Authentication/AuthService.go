package Authentication

import (
	"WB2/internal/Application/Contracts/UserServices"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
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
		showForm(w, "")
		return
	}

	if r.Method == http.MethodPost {
		formData, err := processForm(r)
		if err != nil {
			showForm(w, "Неверный формат данных")
			return
		}

		log.Printf("Получены данные: ID = %d, Key = %s\n", formData.UserID, formData.UserKey)
		err = service.getService.Authentication(formData.UserID, formData.UserKey)
		if err != nil {
			showForm(w, "Неверный ID или Key")
			return
		}

		redirectURL := url.URL{
			Path:     "/authorized",
			RawQuery: fmt.Sprintf("userId=%d&userKey=%s", formData.UserID, url.QueryEscape(formData.UserKey)),
		}

		http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
		return

	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func showForm(w http.ResponseWriter, errorMessage string) {
	htmlTemplate, err := os.ReadFile("internal/Presentation/UI/resources/authentication.html")
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
	err = tmpl.Execute(w, struct {
		ErrorMessage string
	}{ErrorMessage: errorMessage})

	if err != nil {
		http.Error(w, "Ошибка при выполнении шаблона: "+err.Error(), http.StatusInternalServerError)
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
		log.Printf("ошибка при парсинге формы: %v", err)
		return FormData{}, fmt.Errorf("ошибка при парсинге формы: %w", err)
	}

	userIdStr := r.Form.Get("userId")
	userKey := r.Form.Get("userKey")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("некорректный формат ID: %v", err)
		return FormData{}, fmt.Errorf("некорректный формат ID: %w", err)
	}

	return FormData{UserID: userId, UserKey: userKey}, nil
}
