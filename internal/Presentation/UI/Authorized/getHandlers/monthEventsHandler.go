package getHandlers

import (
	"WB2/internal/Application/Contracts/UserServices"
	"WB2/internal/Application/Domain"
	"WB2/internal/Presentation/UI/Authorized"
	"html/template"
	"net/http"
	"os"
)

func HandleShowMonthEvents(w http.ResponseWriter, account *Authorized.Account, getService UserServices.IGetService) {
	htmlTemplate, err := os.ReadFile("internal/Presentation/UI/resources/eventsResponses/events.html")
	if err != nil {
		http.Error(w, "Ошибка при обращении к шаблону", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("events").Parse(string(htmlTemplate))
	if err != nil {
		http.Error(w, "Ошибка при парсинге шаблона", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	events, err := getService.EventsForMonth(account.Id, account.Key)
	if err != nil {
		http.Error(w, "Ошибка при получении событий: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, struct {
		Account *Authorized.Account
		Events  []Domain.Event
		Period  string
	}{
		Events:  events,
		Period:  "месяц",
		Account: account})

	if err != nil {
		http.Error(w, "Ошибка при выполнении шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
