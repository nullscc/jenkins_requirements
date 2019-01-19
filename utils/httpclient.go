package utils

import (
	"encoding/json"
	"fmt"
	"github.com/nullscc/jenkins_requirements/models"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Request(url string, token string, method string, form io.Reader) (interface{}, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, form)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("token", token)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, &StringError{
			Msg: fmt.Sprintf("请求: %s 错误，错误码: %d", url, response.StatusCode),
		}
	}

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data models.Response
	jerr := json.Unmarshal([]byte(result), &data)
	if jerr != nil {
		return nil, jerr
	}

	if data.Code != 200 {
		return nil, &StringError{
			Msg: data.Msg,
		}
	}

	response.Body.Close()
	return data.Info, nil
}

func PostSha1(org string, repo string, file_name string, sha1 string) (string, error) {
	http_url := "http://" + os.Getenv("URL") + "/api/requirement"
	form := url.Values{
		"org":       {org},
		"repo":      {repo},
		"file_name": {file_name},
		"sha1":      {sha1},
	}

	form_reader := strings.NewReader(form.Encode())
	_, err := Request(http_url, os.Getenv("REQUIREMENTS_TOKEN"), "POST", form_reader)
	if err != nil {
		return "", err
	}
	return "", nil
}

func GetSha1(org string, repo string, file_name string) (string, error) {
	http_url := "http://" + os.Getenv("URL") + "/api/requirement?org=" + org + "&repo=" + repo + "&file_name=" + file_name
	data, err := Request(http_url, os.Getenv("REQUIREMENTS_TOKEN"), "GET", nil)
	if err != nil {
		return "", err
	}
	ret_sha1 := data.(map[string]interface{})["sha1"].(string)
	return ret_sha1, nil
}
