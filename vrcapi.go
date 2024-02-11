package vrcapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

type Session struct {
	Client *http.Client
	User   *CurrentUser
}

func New() (*Session, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error creating cookie jar:", err)
		return &Session{}, err
	}

	return &Session{
		Client: &http.Client{
			Jar: jar,
		},
	}, nil
}

func (s *Session) Auth(login, password, twoFactorSecret string) error {
	req, err := http.NewRequest("GET", EndpointLogin, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(login, password)
	req.Header.Set("User-Agent", "vrchat events bot - discord: @godmod, github: g0dm0d")

	_, err = s.Client.Do(req)
	if err != nil {
		return err
	}

	err = s.TwoFactorAuth(twoFactorSecret)
	if err != nil {
		return err
	}

	err = s.InitUser()
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) TwoFactorAuth(twoFactorSecret string) error {
	twoFactorCode, err := generateTOTP(twoFactorSecret)

	if err != nil {
		return err
	}

	body, err := json.Marshal(TwoFactorAuthCode{Code: twoFactorCode})
	if err != nil {
		return err
	}

	var response Verify2FAResult
	err = s.request("POST", Endpoint2FA, "application/json", body, &response)
	if err != nil {
		return err
	}

	if !response.Verified {
		return fmt.Errorf("2fa is unverified")
	}

	return nil
}

func (s *Session) InitUser() error {
	var response CurrentUser
	err := s.request("GET", EndpointLogin, "", nil, &response)
	if err != nil {
		return err
	}

	s.User = &response

	return nil
}

func (s *Session) request(method, urlStr, contentType string, body []byte, response interface{}) (err error) {
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "go vrcapi  - discord: @godmod, github: g0dm0d")

	if body != nil {
		req.Header.Add("Content-Type", contentType)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return err
		}
	case http.StatusBadGateway:
		return fmt.Errorf("bad gateway")
	case http.StatusTooManyRequests:
		// ratelimit
	default:
		var errorInfo Error
		err = json.Unmarshal(respBody, &errorInfo)
		if err != nil {
			return err
		}
		return fmt.Errorf(errorInfo.Error.Message)
	}

	return
}
