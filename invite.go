package vrcapi

import "encoding/json"

func (s *Session) InviteUser(instanceId, userId string) error {
	body, err := json.Marshal(InviteRequest{
		InstanceId:  instanceId,
		MessageSlot: 0,
	})

	if err != nil {
		return err
	}

	_, err = s.request("POST", EndpointInviteUser(userId), "application/json", body)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) InviteMySelf(instanceId string) error {
	_, err := s.request("POST", EndpointInviteMyself(instanceId), "", nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) RequestInvite(userId string) error {
	body, err := json.Marshal(RequestInviteRequest{
		MessageSlot: 0,
	})

	if err != nil {
		return err
	}

	_, err = s.request("POST", EndpointRequestInvite(userId), "application/json", body)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) RespondInvite(notifyId string) error {
	body, err := json.Marshal(InviteResponse{
		MessageSlot: 0,
	})

	if err != nil {
		return err
	}

	_, err = s.request("POST", EndpointRespondInvite(notifyId), "application/json", body)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) ListInviteMessages(userId string, messageType MessageType) (response []Message, err error) {
	body, err := json.Marshal(InviteResponse{
		MessageSlot: 0,
	})

	if err != nil {
		return []Message{}, err
	}

	resp, err := s.request("GET", EndpointListInviteMessages(userId, messageType), "application/json", body)
	if err != nil {
		return []Message{}, err
	}

	err = json.Unmarshal(resp, &response)

	return response, nil
}
