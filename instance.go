package vrcapi

func (s *Session) InstanceLaunch(worldId string, instanceId InstanceId) error {
	err := s.request("GET", EndpointInstanceLaunch(worldId, instanceId), "", nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) GetInstance(worldId string, instanceId InstanceId) error {
	err := s.request("GET", EndpointGetInstance(worldId, instanceId), "", nil, nil)
	if err != nil {
		return err
	}

	return nil
}
