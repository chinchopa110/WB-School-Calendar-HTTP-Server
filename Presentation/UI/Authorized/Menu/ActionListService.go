package Menu

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/UI/Authorized"
	"WB2/Presentation/UI/Authorized/getHandlers"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ActionListService struct {
	getService  UserServices.IGetService
	postService UserServices.IPostService
	account     *Authorized.Account
}

func CreateActionListService(getService UserServices.IGetService, postService UserServices.IPostService) ActionListService {
	return ActionListService{getService, postService, nil}
}

func (service *ActionListService) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/authorized":
		service.showMenu(w, r)
	case "/add-event":
		service.handleAddEvent(w)
	case "/events/day":
		getHandlers.HandleShowDayEvents(w, service.account, service.getService)
	case "/events/week":
		getHandlers.HandleShowWeekEvents(w, service.account, service.getService)
	case "/events/month":
		getHandlers.HandleShowMonthEvents(w, service.account, service.getService)
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

	htmlTemplate, err := os.ReadFile("Presentation/UI/resources/menu.html")
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

func (service ActionListService) handleAddEvent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Выбрано добавление события"))
	if err != nil {
		http.Error(w, "Ошибка при отправке ответа: "+err.Error(), http.StatusInternalServerError)
	}
}
