package score

import (
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gorilla/schema"
	"github.com/labstack/echo"
	"github.com/nsogame/common"
)

var (
	MaxMemory int64 = 1234567
	Decoder         = schema.NewDecoder()
)

type ScoreSubmission struct {
	ScoreEnc string `schema:"score"`
	IV       string `schema:"iv"`
	Password string `schema:"pass"`
	OsuVer   string `schema:"osuver"`

	Bmk string `schema:"bmk"`
	C1  string `schema:"c1"`
	Fs  string `schema:"fs"`
	Ft  string `schema:"ft"`
	I   string `schema:"i"`
	S   string `schema:"s"`
	X   string `schema:"x"`
}

func (score *ScoreServer) SubmitModularHandler(c echo.Context) (err error) {
	// files := c.MultipartForm()
	form, err := c.FormParams()
	if err != nil {
		return
	}

	var data ScoreSubmission
	err = Decoder.Decode(&data, form)
	if err != nil {
		return
	}

	fmt.Println("data", data)

	// decrypt the score with AES-CBC
	var key []byte
	key = []byte(fmt.Sprintf("osu!-scoreburgr---------%s", data.OsuVer))

	iv, err := base64.StdEncoding.DecodeString(data.IV)
	if err != nil {
		return
	}

	scoreEnc, err := base64.StdEncoding.DecodeString(data.ScoreEnc)
	if err != nil {
		return
	}

	block, err := common.NewCipher(key)
	if err != nil {
		return
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	scoreBytes := make([]byte, len(scoreEnc))
	mode.CryptBlocks(scoreBytes, scoreEnc)

	scoreData := strings.Split(string(scoreBytes), ":")
	// fileChecksum := scoreData[0]
	username := strings.Trim(scoreData[1], " ")

	// pull user out of the db
	user, err := score.db.GetUserByName(username)
	if err != nil {
		return
	}

	fmt.Println("user", user)

	return
}
