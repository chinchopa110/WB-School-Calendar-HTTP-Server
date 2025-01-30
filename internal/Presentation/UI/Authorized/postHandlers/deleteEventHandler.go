package postHandlers

import (
	"WB2/internal/Application/Contracts/UserServices"
	"WB2/internal/Presentation/UI/Authorized"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HandleDeleteEvent(w http.ResponseWriter, r *http.Request, account *Authorized.Account, postService UserServices.IPostService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusBadRequest)
		return
	}
	eventIdStr := r.Form.Get("eventId")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		http.Error(w, "некорректный формат ID", http.StatusInternalServerError)
		return
	}

	log.Printf("Удаляем событие %d\n", eventId)
	_, err = postService.DeleteEvent(account.Id, eventId, account.Key)
	if err != nil {
		http.Error(w, "Ошибка при удалении события", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/authorized?userId=%d&userKey=%s", account.Id, account.Key), http.StatusSeeOther)
}
