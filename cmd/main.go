package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amandaanjolin/desafio-go-web/internal/loader"
	"github.com/amandaanjolin/desafio-go-web/internal/tickets/handler"
	"github.com/amandaanjolin/desafio-go-web/internal/tickets/repository"
	"github.com/amandaanjolin/desafio-go-web/internal/tickets/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	// env
	// ...

	// application
	// - config
	cfg := &ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "./docs/db/tickets.csv",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	db, err := loader.LoadTickets(a.dbFile)
	if err != nil {
		return
	}

	// Cria repository, service e handler
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// Configura o roteador chi
	r := chi.NewRouter()
	r.Get("/tickets/getByCountry/{country}", h.CountByCountry)
	r.Get("/tickets/getByTimePeriod", h.CountByTimePeriod)
	r.Get("/tickets/getAverageByCountry/{country}", h.PercentageByCountry)

	log.Println("Servidor iniciado em :8080")
	http.ListenAndServe(":8080", r)

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
