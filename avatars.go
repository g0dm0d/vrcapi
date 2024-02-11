package vrcapi

func (s *Session) GetOwnAvatar(userId string) (response *Avatar, err error) {
	err = s.request("GET", EndpointGetOwnAvatar(userId), "", nil, response)
	if err != nil {
		return &Avatar{}, err
	}

	return response, err
}
