package gitioapi_test

import (
	"fmt"
	"github.com/spiegel-im-spiegel/gitioapi"
)

func ExampleEncode() {
	shortUrl, _ := gitioapi.Encode(&gitioapi.Param{Url: "https://github.com/technoweenie", Code: "t"})
	fmt.Print(shortUrl)
	// http://git.io/t
}

func ExampleDecode() {
	shortUrl, _ := gitioapi.Decode(&gitioapi.Param{Url: "http://git.io/t"})
	fmt.Print(shortUrl)
	// https://github.com/technoweenie
}
