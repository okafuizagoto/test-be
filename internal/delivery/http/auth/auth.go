package goldgym

import (
	"errors"
	"gold-gym-be/pkg/response"
	"log"
	"net/http"
)

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("testDelivery")
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	user, password, ok := r.BasicAuth()
	if !ok {
		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, errors.New("403 Forbidden"))
		return
	}
	log.Println("testDelivery2")

	result, metadata, err := h.goldgymSvc.LoginUser(ctx, user, password, r.RemoteAddr)
	log.Println("testDelivery3")
	if err != nil {
		// Return error message with HTTP 200 OK
		resp.SetError(err, http.StatusOK)

		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = metadata

	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
