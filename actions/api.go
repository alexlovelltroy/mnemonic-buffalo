package actions

import (
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/gobuffalo/buffalo"
)

var wordlistLength = strconv.Itoa(len(wordlist))
var wordlistCapacity = strconv.Itoa(cap(wordlist))

func init() {
	rand.Seed(52)
}

func getColor() string {
	return os.Getenv("COLOR")
}

func oneword() string {
	return wordlist[rand.Intn(len(wordlist))]
}

// APIHandler is a default handler to serve up
// a the list of words.
func APIHandler(c buffalo.Context) error {
	intWordCount, err := strconv.Atoi(c.Param("count"))

	if err != nil {
		if intWordCount >= 200 {
			return c.Render(500, r.String("The API doesn't support wordlists higher than 200 items long."))
		}
		return c.Render(200, r.JSON(map[string]string{"Hello": "Friend", "Version": getColor(), "Wordlist Length": wordlistLength, "Wordlist Capacity": wordlistCapacity}))
	}
	var words = make([]string, intWordCount)
	for i := range words {
		words[i] = oneword()
	}
	return c.Render(200, r.JSON(map[string]string{"count": strconv.Itoa(intWordCount), "words": strings.Join(words[:], " ")}))

}
