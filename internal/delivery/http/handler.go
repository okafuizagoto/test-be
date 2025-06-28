package http

import (
	"errors"
	"log"
	"net/http"

	"gold-gym-be/pkg/response"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	// Jika tidak ditemukan, jangan diubah.
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	// Health Check
	r.HandleFunc("", defaultHandler).Methods("GET")
	r.HandleFunc("/", defaultHandler).Methods("GET")

	// Tambahan Prefix di depan API endpoint
	router := r.PathPrefix("/gold-gym").Subrouter()

	router.HandleFunc("", defaultHandler).Methods("GET")
	router.HandleFunc("/", defaultHandler).Methods("GET")

	sub := router.PathPrefix("/v2").Subrouter()

	// Routes
	goldgym := sub.PathPrefix("/userdata").Subrouter()

	goldgym.HandleFunc("", s.Goldgym.GetGoldGym).Methods("GET")
	goldgym.HandleFunc("", s.Goldgym.InsertGoldGym).Methods("POST")
	goldgym.HandleFunc("", s.Goldgym.UpdateGoldGym).Methods("PUT")
	goldgym.HandleFunc("", s.Goldgym.DeleteGoldGym).Methods("DELETE")

	goldgym.HandleFunc("/login", s.Auth.LoginUser).Methods("POST")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Example Service API"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp   *response.Response
		err    error
		errRes response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	err = errors.New("404 Not Found")

	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   404,
			Msg:    "404 Not Found",
			Status: true,
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = 404
		resp.Error = errRes
		return
	}
}
