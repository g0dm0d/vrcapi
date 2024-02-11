package vrcapi

import "strconv"

var (
	EndpointVRC = "https://vrchat.com/"
	EndpointAPI = "https://api.vrchat.cloud/api/1/"

	EndpointLogin = EndpointAPI + "auth/user"
	Endpoint2FA   = EndpointAPI + "auth/twofactorauth/totp/verify"

	// INSTANCES

	EndpointGetInstance = func(worldId string, instanceId InstanceId) string {
		return EndpointAPI + "instances/" + worldId + ":" + instanceId.String()
	}

	EndpointInstance = func(worldId string, instanceId InstanceId) string {
		return EndpointAPI + "instances/" + worldId + ":" + instanceId.String()
	}

	// EndpointInstanceLaunch https://vrchatapi.github.io/tutorials/instances/
	EndpointInstanceLaunch = func(worldId string, instanceId InstanceId) string {
		return EndpointVRC + "home/launch?" + "worldId=" + worldId + "&" +
			"instanceId=" + instanceId.String()
	}

	// INVITE

	EndpointInviteUser = func(userId string) string {
		return EndpointAPI + "invite/" + userId
	}

	EndpointInviteMyself = func(instance string) string {
		return EndpointAPI + "invite/myself/to/" + instance
	}

	EndpointRequestInvite = func(userId string) string {
		return EndpointAPI + "requestInvite/" + userId
	}

	EndpointRespondInvite = func(notifyId string) string {
		return EndpointAPI + "invite/" + notifyId + "/response"
	}

	EndpointListInviteMessages = func(userId string, messageType MessageType) string {
		return EndpointAPI + "message/" + userId + "/" + string(messageType)
	}

	EndpointGetInviteMessage = func(userId string, messageType MessageType, slot int) string {
		return EndpointAPI + "message/" + userId + "/" + string(messageType) + "/" + strconv.Itoa(slot)
	}

	EndpointUpdateInviteMessage = func(userId string, messageType MessageType, slot int) string {
		return EndpointAPI + "message/" + userId + "/" + string(messageType) + "/" + strconv.Itoa(slot)
	}

	EndpointResetInviteMessage = func(userId string, messageType MessageType, slot int) string {
		return EndpointAPI + "message/" + userId + "/" + string(messageType) + "/" + strconv.Itoa(slot)
	}

	// AVATARS

	EndpointGetOwnAvatar = func(userID string) string {
		return EndpointAPI + "user/" + userID + "/avatar"
	}

	// FRIENDS

	EndpointListFriends = EndpointAPI + "auth/user/friends?offline=true"

	EndpointFriendRequest = func(userID string) string {
		return EndpointAPI + "user/" + userID + "/friendRequest"
	}

	EndpointCheckFriendStatus = func(userID string) string {
		return EndpointAPI + "user/" + userID + "/friendStatus"
	}

	EndpointUnfriend = func(userID string) string {
		return EndpointAPI + "auth/user/friends/" + userID
	}
)
