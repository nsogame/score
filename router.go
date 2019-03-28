package score

import (
	"github.com/labstack/echo"
)

func (score *ScoreServer) router(web *echo.Echo) {
	web.POST("/web/osu-submit-modular-selector.php", score.SubmitModularHandler)
	web.GET("/web/osu-osz2-getscores.php", score.GetScores)

	web.GET("/web/osu-osz2-bmsubmit-getid.php", score.Osz2BmsubmitGetid)
	web.POST("/web/osu-osz2-bmsubmit-upload.php", score.Osz2BmsubmitUpload)
}
