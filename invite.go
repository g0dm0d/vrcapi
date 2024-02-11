package vrcapi

import "encoding/json"

func (s *Session) InviteUser(instanceId, userId string) (response InviteResponse, err error) {
	body, err := json.Marshal(InviteRequest{
		InstanceId:  instanceId,
		MessageSlot: 0,
	})

	if err != nil {
		return response, err
	}

	resp, err := s.request("POST", EndpointInviteUser(userId), "application/json", body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) InviteMySelf(instanceId string) (response SentNotification, err error) {
	resp, err := s.request("POST", EndpointInviteMyself(instanceId), "", nil)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) RequestInvite(userId string) (response Notification, err error) {
	body, err := json.Marshal(RequestInviteRequest{
		MessageSlot: 0,
	})

	if err != nil {
		return response, err
	}

	resp, err := s.request("POST", EndpointRequestInvite(userId), "application/json", body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) RespondInvite(notifyId string) (response Notification, err error) {
	body, err := json.Marshal(InviteResponse{
		MessageSlot: 0,
	})

	if err != nil {
		return response, err
	}

	resp, err := s.request("POST", EndpointRespondInvite(notifyId), "application/json", body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) ListInviteMessages(userId string, messageType MessageType) (response []Message, err error) {
	body, err := json.Marshal(InviteResponse{
		MessageSlot: 0,
	})

	if err != nil {
		return response, err
	}

	resp, err := s.request("GET", EndpointListInviteMessages(userId, messageType), "application/json", body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
