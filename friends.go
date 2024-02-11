package vrcapi

func (s *Session) ListFriends(userId string) (response LimitedUser, err error) {
	err = s.request("GET", EndpointListFriends, "", nil, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) SendFriendRequest(userId string) (response Notification, err error) {
	err = s.request("POST", EndpointFriendRequest(userId), "", nil, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) DeleteFriendRequest(userId string) (response Success, err error) {
	err = s.request("DELETE", EndpointFriendRequest(userId), "", nil, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) CheckFriendStatus(userId string) (response FriendStatus, err error) {
	err = s.request("GET", EndpointCheckFriendStatus(userId), "", nil, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) Unfriend(userId string) (response Success, err error) {
	err = s.request("DELETE", EndpointUnfriend(userId), "", nil, &response)
	if err != nil {
		return response, err
	}

	return response, err
}
