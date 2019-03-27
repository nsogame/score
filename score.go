package score

import (
	"log"
	"net/http"
	"time"

	"git.iptq.io/nso/common"
)

type ScoreServer struct {
	config *Config
	db     *common.DB
	rds    *common.RedisAPI
	router http.Handler
}

func NewInstance(config *Config) (score *ScoreServer, err error) {
	// db
	db, err := common.ConnectDB(config.DbProvider, config.DbConnection)
	if err != nil {
		return
	}

	// router

	score = &ScoreServer{
		config: config,
		db:     db,
	}

	router := score.Handlers()
	score.router = router
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
