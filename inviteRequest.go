package vrcapi

import "encoding/json"

func (s *Session) InviteRequest(instanceId, userId string) error {
	body, err := json.Marshal(InviteRequest{
		InstanceId:  instanceId,
		MessageSlot: 0,
	})

	if err != nil {
		return err
	}

	_, err = s.request("POST", EndpointInviteRequest(userId), "application/json", body)
	if err != nil {
		return err
	}

	return nil
}
