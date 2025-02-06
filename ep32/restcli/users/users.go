package users

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type SupportInfo struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}
type RetUser struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type RetUserOne struct {
	Data    RetUser     `json:"data"`
	Support SupportInfo `json:"support"`
}

type RetUserList struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
	Data       []RetUser   `json:"data"`
	Support    SupportInfo `json:"support"`
}

func Get(id int) (RetUserOne, error) {

	url := config.BaseURL + "/" + strconv.Itoa(id)

	resp, err := http.Get(url)
	if err != nil {
		return RetUserOne{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return RetUserOne{}, errors.New("Error: " + strconv.Itoa(resp.StatusCode))
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return RetUserOne{}, err
	}

	var retUser RetUserOne
	err = json.Unmarshal(content, &retUser)
	if err != nil {
		return RetUserOne{}, err
	}
	return retUser, nil

}

func List(page int) (RetUserList, error) {
	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return RetUserList{}, err
	}
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	baseURL.RawQuery = params.Encode()
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return RetUserList{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return RetUserList{}, errors.New("Error: " + strconv.Itoa(resp.StatusCode))
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return RetUserList{}, err
	}
	var retUsers RetUserList
	err = json.Unmarshal(content, &retUsers)
	if err != nil {
		return RetUserList{}, err
	}
	return retUsers, nil
}
