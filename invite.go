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

	err = s.request("POST", EndpointInviteUser(userId), "application/json", body, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) InviteMySelf(instanceId string) (response SentNotification, err error) {
	err = s.request("POST", EndpointInviteMyself(instanceId), "", nil, &response)
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

	err = s.request("POST", EndpointRequestInvite(userId), "application/json", body, &response)
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

	err = s.request("POST", EndpointRespondInvite(notifyId), "application/json", body, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) ListInviteMessages(userId string, messageType MessageType) (response []Message, err error) {
	err = s.request("GET", EndpointListInviteMessages(userId, messageType), "", nil, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) GetInviteMessages(userId string, messageType MessageType, slot int) (response Message, err error) {
	err = s.request("GET", EndpointGetInviteMessage(userId, messageType, slot), "", nil, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) UpdateInviteMessage(userId string, messageType MessageType, slot int, message string) (response []Message, err error) {
	body, err := json.Marshal(UpdateInviteMessageRequest{
		Message: message,
	})

	if err != nil {
		return response, err
	}

	err = s.request("PUT", EndpointUpdateInviteMessage(userId, messageType, slot), "application/json", body, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *Session) ResetInviteMessage(userId string, messageType MessageType, slot int) (response []Message, err error) {
	err = s.request("DELETE", EndpointResetInviteMessage(userId, messageType, slot), "", nil, &response)
	if err != nil {
		return response, err
	}

	return response, err
}
