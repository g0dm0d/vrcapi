package vrcapi

var (
	EndpointVRC = "https://vrchat.com/"
	EndpointAPI = "https://api.vrchat.cloud/api/1/"

	EndpointLogin = EndpointAPI + "auth/user"
	Endpoint2FA   = EndpointAPI + "auth/twofactorauth/totp/verify"

	EndpointInstance = func(worldID string, instanceID InstanceId) string {
		return EndpointAPI + "instances/" + worldID + ":" + instanceID.String()
	}

	// EndpointInstanceLaunch https://vrchatapi.github.io/tutorials/instances/
	EndpointInstanceLaunch = func(worldID string, instanceID InstanceId) string {
		return EndpointVRC + "home/launch?" + "worldId=" + worldID + "&" +
			"instanceId=" + instanceID.String()
	}

	EndpointInviteRequest = func(userID string) string {
		return EndpointAPI + "invite/" + userID
	}

	EndpointFriendRequest = func(userID string) string {
		return EndpointAPI + "user/" + userID + "/friendRequest"
	}
)
