package vrcapi

func (s *Session) InstanceLaunch(worldId string, instanceId InstanceId) error {

	reqUrl := EndpointInstanceLaunch(worldId, instanceId)

	_, err := s.request("GET", reqUrl, "", nil)
	if err != nil {
		return err
	}

	return nil
}
