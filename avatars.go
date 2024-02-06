package vrcapi

import "encoding/json"

func (s *Session) GetOwnAvatar(userId string) (response *Avatar, err error) {
	reqUrl := EndpointGetOwnAvatar(userId)

	resp, err := s.request("GET", reqUrl, "", nil)
	if err != nil {
		return &Avatar{}, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return &Avatar{}, err
	}

	return response, err
}
