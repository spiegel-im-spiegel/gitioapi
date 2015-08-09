/**
 * API for Git.io
 *
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/deed.ja
 */

//Git.io API Package.
package gitioapi

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Error reports an error and status.
type Status struct {
	Status string
	Err    error
}

// Is this Error?
func (s *Status) IsError() bool {
	if s.Err != nil {
		return true
	}
	return false
}

// Error reports.
func (s *Status) Error() string {
	if s.IsError() {
		if len(s.Status) > 0 {
			return s.Err.Error() + " (" + s.Status + ")"
		}
		return s.Err.Error()
	}
	return ""
}

//Parameter for Git.io API
type Param struct {
	Url  string
	Code string
}

//Get url.Values object from Param object.
func (prm *Param) GetUrlValuse() url.Values {
	v := url.Values{}
	if len(prm.Url) > 0 {
		v.Add("url", prm.Url)
	}
	if len(prm.Code) > 0 {
		v.Add("code", prm.Code)
	}
	return v
}

//Shorten GitHub Domain URL.
func Encode(prm *Param) (string, Status) {
	//shortening url
	resp, err := http.PostForm("http://git.io", prm.GetUrlValuse())
	if err != nil {
		return "", Status{Status: "", Err: err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", Status{Status: "", Err: err}
	}

	result := resp.Header.Get("Location")
	if string(body) != prm.Url {
		return result, Status{Status: string(body), Err: os.ErrInvalid}
	} else {
		return result, Status{Status: "", Err: nil}
	}
}
