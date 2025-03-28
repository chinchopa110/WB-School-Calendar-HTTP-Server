package postHandlers

import (
	"WB2/internal/Application/Contracts/UserServices"
	"WB2/internal/Presentation/UI/Authorized"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

func HandleUpdateEventDate(w http.ResponseWriter, r *http.Request, account *Authorized.Account, postService UserServices.IPostService) {
	htmlTemplate, err := os.ReadFile("web/resources/updateForms/date.html")
	if err != nil {
		http.Error(w, "Ошибка при обращении к шаблону", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("update_date").Parse(string(htmlTemplate))
	if err != nil {
		http.Error(w, "Ошибка при обращении к шаблону", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		eventIdStr := r.URL.Query().Get("eventId")
		eventId, err := strconv.Atoi(eventIdStr)
		if err != nil {
			http.Error(w, "некорректный формат ID", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.Execute(w, struct {
			Account *Authorized.Account
			EventId int
		}{Account: account, EventId: eventId})
		if err != nil {
			http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusBadRequest)
		return
	}

	date := r.Form.Get("date")
	eventIdStr := r.Form.Get("eventId")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		http.Error(w, "некорректный формат ID", http.StatusInternalServerError)
		return
	}

	log.Printf("Обновляем дату события %d, на дату = %s\n", eventId, date)
	_, err = postService.UpdateEventDate(account.Id, eventId, date, account.Key)
	if err != nil {
		http.Error(w, "Ошибка при обновлении события", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/authorized?userId=%d&userKey=%s", account.Id, account.Key), http.StatusSeeOther)

}
