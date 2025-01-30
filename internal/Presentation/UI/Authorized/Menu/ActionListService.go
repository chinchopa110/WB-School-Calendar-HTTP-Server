package Menu

import (
	UserServices2 "WB2/internal/Application/Contracts/UserServices"
	"WB2/internal/Presentation/UI/Authorized"
	getHandlers2 "WB2/internal/Presentation/UI/Authorized/getHandlers"
	postHandlers2 "WB2/internal/Presentation/UI/Authorized/postHandlers"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ActionListService struct {
	getService  UserServices2.IGetService
	postService UserServices2.IPostService
	account     *Authorized.Account
}

func CreateActionListService(getService UserServices2.IGetService, postService UserServices2.IPostService) ActionListService {
	return ActionListService{getService, postService, nil}
}

func (service *ActionListService) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/authorized":
		service.showMenu(w, r)
	case "/events/day":
		getHandlers2.HandleShowDayEvents(w, service.account, service.getService)
	case "/events/week":
		getHandlers2.HandleShowWeekEvents(w, service.account, service.getService)
	case "/events/month":
		getHandlers2.HandleShowMonthEvents(w, service.account, service.getService)
	case "/add-event":
		postHandlers2.HandleAddEvent(w, r, service.account, service.postService)
	case "/update-date":
		postHandlers2.HandleUpdateEventDate(w, r, service.account, service.postService)
	case "/update-description":
		postHandlers2.HandleUpdateEventDescription(w, r, service.account, service.postService)
	case "/delete-event":
		postHandlers2.HandleDeleteEvent(w, r, service.account, service.postService)
	default:
		http.NotFound(w, r)
	}
}

func (service *ActionListService) showMenu(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("userId")
	userKey := r.URL.Query().Get("userKey")
	log.Printf("Получены параметры: ID = %s, Key = %s\n", userIdStr, userKey)

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "некорректный формат ID", http.StatusInternalServerError)
	}
	service.account = &Authorized.Account{Id: userId, Key: userKey}

	htmlTemplate, err := os.ReadFile("internal/Presentation/UI/resources/menu.html")
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
		http.Error(w, "Ошибка при выполнении шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
