package gitioapi_test

import (
	"fmt"
	"github.com/spiegel-im-spiegel/gitioapi"
)

func ExampleDayNumber() {
	shortUrl, status := gitioapi.Encode(&gitioapi.Param{Url: "https://github.com/spgl", Code: "spgl"})
	fmt.Print(shortUrl, status.Status)
	// http://git.io/spgl:
	// 201 Created
}
