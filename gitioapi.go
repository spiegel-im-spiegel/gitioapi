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
type ApiError struct {
	Status string
	Err    error
}

// Error reports.
func (e *ApiError) Error() string {
	if len(e.Status) > 0 {
		return e.Err.Error() + " (" + e.Status + ")"
	}
	return e.Err.Error()
}

func NewApiError(status string, err error) error {
	if err == nil {
		return nil
	}
	return &ApiError{Status: status, Err: err}
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
// Web API: curl -i http://git.io -F "url=https://github.com/technoweenie" -F "code=t"
func Encode(prm *Param) (string, error) {
	//shortening url
	resp, err := http.PostForm("http://git.io", prm.GetUrlValuse())
	if err != nil {
		return "", NewApiError("", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", NewApiError("", err)
	}

	result := resp.Header.Get("Location")
	status := string(body)
	if status != prm.Url {
		return result, NewApiError(status, os.ErrInvalid)
	} else {
		return result, nil
	}
}
