package postHandlers

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/UI/Authorized"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
)

func HandleAddEvent(w http.ResponseWriter, r *http.Request, account *Authorized.Account, postService UserServices.IPostService) {
	if r.Method == http.MethodGet {
		htmlTemplate, err := os.ReadFile("Presentation/UI/resources/addForm/addEvent.html")
		if err != nil {
			http.Error(w, "Ошибка при обращении к шаблону", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.New("events").Parse(string(htmlTemplate))
		if err != nil {
			http.Error(w, "Ошибка при парсинге шаблона", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		date := r.FormValue("date")
		description := r.FormValue("description")

		_, err := postService.CreateEvent(account.Id, date, description, account.Key)
		if err != nil {
			http.Error(w, "Ошибка при добавлении события", http.StatusInternalServerError)
			return
		}

		redirectURL := url.URL{
			Path:     "/authorized",
			RawQuery: fmt.Sprintf("userId=%d&userKey=%s", account.Id, url.QueryEscape(account.Key)),
		}
		http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
		return
	}
	http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
}
