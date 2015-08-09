# Git.io Web API Package

## Git.io Web API

Refer to “[Git.io: GitHub URL Shortener](https://github.com/blog/985-git-io-github-url-shortener)”.

## Code Generation

- [cURL as DSL — cURL as DSL 1.0 documentation](https://shibukawa.github.io/curl_as_dsl/)
    - [shibukawa/curl_as_dsl](https://github.com/shibukawa/curl_as_dsl) (coded by golang)

## Example

```go:example_test.go
package gitioapi_test

import (
	"fmt"
	"github.com/spiegel-im-spiegel/gitioapi"
)

func ExampleEncode() {
	shortUrl, _ := gitioapi.Encode(&gitioapi.Param{Url: "https://github.com/technoweenie", Code: "t"})
	fmt.Print(shortUrl)
	// http://git.io/t:
}
```

## License

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed.ja)
