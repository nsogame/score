package score

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type ScoreServer struct {
	config *Config
	db     *gorm.DB
	router http.Handler
}

func NewInstance(config *Config) (score *ScoreServer, err error) {
	// db
	db, err := gorm.Open(config.DbProvider, config.DbConnection)
	if err != nil {
		return
	}

	// router
	router := handler()

	score = &ScoreServer{
		config: config,
		db:     db,
		router: router,
	}
	return
}

func (score *ScoreServer) close() {
	score.db.Close()
}

func (score *ScoreServer) Run() {
	defer score.close()
	log.Println("starting...")
	server := &http.Server{
		Handler: score.router,
		Addr:    score.config.BindAddr,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
