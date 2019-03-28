package score

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (score *ScoreServer) Osz2BmsubmitGetid(c echo.Context) (err error) {
	query := c.QueryParams()
	fmt.Println("query", query)

	username := c.QueryParam("u")
	user, err := score.db.GetUserByName(username)
	if err != nil {
		return
	}

	fmt.Println(user)

	// 1. new beatmap set id
	// 2. comma-separated array of new beatmap ids
	// 3. whether or not it's a full upload
	// 4. submission quota remaining
	// 5. whether or not this will result in a bubble pop (optional)
	// 6. whether or not this map is approved (optional)
	result := fmt.Sprintf("0\n%d\n13,14\n1\n\n", 12)
	return c.String(http.StatusOK, result)
}

func (score *ScoreServer) Osz2BmsubmitUpload(c echo.Context) (err error) {
	return
}
