package gitioapi_test

import (
	"fmt"
	"github.com/spiegel-im-spiegel/gitioapi"
)

func ExampleDayNumber() {
	shortUrl, _ := gitioapi.Encode(&gitioapi.Param{Url: "https://github.com/spgl", Code: "spgl"})
	fmt.Print(shortUrl)
	// http://git.io/spgl:
}
