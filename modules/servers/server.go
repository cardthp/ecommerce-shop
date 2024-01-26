package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/cardthp/ecommerce-shop/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type IServer interface {
	Start()
}

type server struct {
	app *fiber.App
	cfg config.IConfig
	db  *sqlx.DB
}

func NewServer(cfg config.IConfig, db *sqlx.DB) IServer {
	return &server{
		cfg: cfg,
		db:  db,
		app: fiber.New(fiber.Config{
			AppName:      cfg.App().Name(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeout(),
			WriteTimeout: cfg.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}
}

func (s *server) Start() {
	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// go rounties is not do by line from 1 to 10 but it will work when have some trigger
	go func() {
		_ = <-c // _ is if c have receiver something it will input to _ parameter ( parameter _ is assign for ignored / not used for return anymore)
		log.Printf("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	//Listen to host:port
	log.Printf("Server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}
