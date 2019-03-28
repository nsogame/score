package score

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nsogame/common"
)

type ScoreServer struct {
	config *Config
	db     *common.DB
	rds    *common.RedisAPI
	web    *echo.Echo
}

func NewInstance(config *Config) (score *ScoreServer, err error) {
	// db
	db, err := common.ConnectDB(config.DbProvider, config.DbConnection)
	if err != nil {
		return
	}

	// router
	web := echo.New()
	web.Debug = config.Debug

	web.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	web.Use(middleware.Recover())

	score = &ScoreServer{
		config: config,
		db:     db,
		web:    web,
	}
	score.router(web)
	return
}

func (score *ScoreServer) close() {
	score.db.Close()
}

func (score *ScoreServer) Run() {
	defer score.close()
	log.Fatal(score.web.Start(score.config.BindAddr))
}
