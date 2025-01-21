package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

type EdtUser struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type CrtRespUser struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Job       string `json:"job"`
	CreatedAt string `json:"createdAt"`
}

type UpdRespUser struct {
	Name      string `json:"name"`
	Job       string `json:"job"`
	UpdatedAt string `json:"updatedAt"`
}

func Create(user EdtUser) (CrtRespUser, error) {
	client := &http.Client{Timeout: 15 * time.Second}

	data, err := json.Marshal(user)
	if err != nil {
		return CrtRespUser{}, err
	}
	req, err := http.NewRequest(http.MethodPost,
		"https://reqres.in/api/users", bytes.NewBuffer(data))
	if err != nil {
		return CrtRespUser{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return CrtRespUser{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return CrtRespUser{}, errors.New("Error: " + strconv.Itoa(resp.StatusCode))
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return CrtRespUser{}, err
	}
	var crtRespUser CrtRespUser
	err = json.Unmarshal(content, &crtRespUser)
	if err != nil {
		return CrtRespUser{}, err
	}
	return crtRespUser, nil

}

func Update(id int, user EdtUser) (UpdRespUser, error) {
	client := &http.Client{Timeout: 15 * time.Second}

	data, err := json.Marshal(user)
	if err != nil {
		return UpdRespUser{}, err
	}
	req, err := http.NewRequest(http.MethodPut,
		"https://reqres.in/api/users/"+strconv.Itoa(id), bytes.NewBuffer(data))
	if err != nil {
		return UpdRespUser{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return UpdRespUser{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return UpdRespUser{}, errors.New("Error: " + strconv.Itoa(resp.StatusCode))
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return UpdRespUser{}, err
	}
	var updRespUser UpdRespUser
	err = json.Unmarshal(content, &updRespUser)
	if err != nil {
		return UpdRespUser{}, err
	}
	return updRespUser, nil
}

func Delete(id int) error {
	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest(http.MethodDelete,
		"https://reqres.in/api/users/"+strconv.Itoa(id), nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return errors.New("Error: " + strconv.Itoa(resp.StatusCode))
	}
	return nil
}
