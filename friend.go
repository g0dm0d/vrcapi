package vrcapi

func (s *Session) FriendRequest(userId string) error {
	_, err := s.request("POST", EndpointFriendRequest(userId), "", nil)
	if err != nil {
		return err
	}

	return nil
}
