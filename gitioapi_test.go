package gitioapi

import (
	"errors"
	"net/url"
	"os"
	"testing"
)

type errorTestCase struct { //Test case for Error
	status  string
	err     error
	errMsg  string
}

var errorTests []errorTestCase //Test cases for Error

type parmsTestCase struct { //Test case for Param
	url    string
	code   string
	res    string
	errMsg string
}

var parmsTests []parmsTestCase //Test cases for Param

func TestMain(m *testing.M) {
	//Test cases for Error
	err1 := errors.New("error 1 !")
	err2 := errors.New("error 2 !")
	errorTests = []errorTestCase{
		{status: "status", err: err1, errMsg: "error 1 ! (status)"},
		{status: "", err: err2, errMsg: "error 2 !"},
		{status: "status", err: nil, errMsg: ""},
		{status: "", err: nil, errMsg: ""},
	}
	//Test cases for Param
	parmsTests = []parmsTestCase{
		{url: "https://github.com/spiegel-im-spiegel", code: "", res: "http://git.io/vOj52", errMsg: ""},
		{url: "https://github.com/spgl", code: "spgl", res: "http://git.io/spgl", errMsg: ""},
		{url: "http://www.baldanders.info/", code: "", res: "", errMsg: "invalid argument (Must be a GitHub.com URL.)"},
		{url: "", code: "spgl", res: "", errMsg: "invalid argument (Invalid url: )"},
		{url: "", code: "", res: "", errMsg: "invalid argument (Invalid url: )"},
	}

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}

func TestError(t *testing.T) {
	for _, testCase := range errorTests {
		err := NewApiError(testCase.status, testCase.err)
		if err == nil {
			if len(testCase.errMsg) > 0 {
				t.Error("Error Status  = 'false', want 'true'.")
			}
		} else {
			msg := err.Error()
			if msg != testCase.errMsg {
				t.Errorf("Status  = %v, want %v.", msg, testCase.errMsg)
			}
		}
	}
}

func TestParams(t *testing.T) {
	var prm Param
	var values url.Values
	for _, testCase := range parmsTests {
		//case x-1
		prm = Param{Url: testCase.url, Code: testCase.code}
		values = prm.GetUrlValuse()
		if values.Get("url") != testCase.url {
			t.Errorf("values[\"url\"]  = %v, want %v.", values.Get("url"), testCase.url)
		}
		if values.Get("code") != testCase.code {
			t.Errorf("values[\"code\"]  = %v, want %v.", values.Get("code"), testCase.code)
		}
		//case x-2
		prm = Param{Url: testCase.url}
		values = prm.GetUrlValuse()
		if values.Get("url") != testCase.url {
			t.Errorf("values[\"url\"]  = %v, want %v.", values["url"], testCase.url)
		}
		if len(values.Get("code")) > 0 {
			t.Errorf("values[\"code\"]  = %v, want empty.", values.Get("code"))
		}
		//case x-3
		prm = Param{Code: testCase.code}
		values = prm.GetUrlValuse()
		if len(values.Get("url")) > 0 {
			t.Errorf("values[\"url\"]  = %v, want empty.", values.Get("url"))
		}
		if values.Get("code") != testCase.code {
			t.Errorf("values[\"code\"]  = %v, want %v.", values.Get("code"), testCase.code)
		}
	}
}

func TestEncode(t *testing.T) {
	for _, testCase := range parmsTests {
		prm := Param{Url: testCase.url, Code: testCase.code}
		result, err := Encode(&prm)
		if result != testCase.res {
			t.Errorf("Encode()  = %v, want %v.", result, testCase.res)
		}
		if err == nil {
			if len(testCase.errMsg) > 0 {
				t.Error("Status of Encode() = false, want true.")
			}
		} else {
			if err.Error() != testCase.errMsg {
				t.Errorf("Status of Encode() = %v, want %v.", err.Error(), testCase.errMsg)
			}
		}
	}
}
