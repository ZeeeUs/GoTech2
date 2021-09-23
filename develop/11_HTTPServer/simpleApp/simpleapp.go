package simpleapp

import (
	"github.com/ZeeeUs/GoTech2/develop/11_HTTPServer/config"
	"github.com/ZeeeUs/GoTech2/develop/11_HTTPServer/model"
	"github.com/ZeeeUs/GoTech2/develop/11_HTTPServer/store"
	"log"
	"net/http"
)

type simpleApp struct {
	config *config.Config
	router *http.ServeMux
	store  *store.Store
	events []model.Event
}

func New(config *config.Config) *simpleApp {
	return &simpleApp{
		config: config,
		router: http.NewServeMux(),
	}
}

func (s *simpleApp) Start() error {
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	log.Println("Starting simpleApp server")
	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *simpleApp) configureStore() error {
	st := store.New(s.config.Store)
	events, err := st.Open()
	if err != nil {
		return err
	}

	s.store = st
	s.events = events

	return nil
}

func (s *simpleApp) configureRouter() {
	s.router.Handle("/events", logMiddleware(http.HandlerFunc(s.getEvent)))
	s.router.Handle("/create_event", logMiddleware(http.HandlerFunc(s.createEvent)))
}

