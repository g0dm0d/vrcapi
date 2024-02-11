package vrcapi

import (
	"fmt"
	"time"
)

type Privacy string

const (
	Public      Privacy = ""
	FriendsPlus Privacy = "~hidden(%s)"
	Friends     Privacy = "~friends(%s)"
	InvitePlus  Privacy = "~private(%s)~canRequestInvite"
	Invite      Privacy = "~private(%s)"
)

type Region string

const (
	EU  Region = "eu"
	USW Region = "usw"
	USE Region = "us"
	JP  Region = "jp"
)

type MessageType string

const (
	DefaultMessage  MessageType = "message"
	Response        MessageType = "response"
	Request         MessageType = "request"
	RequestResponse MessageType = "requestResponse"
)

type InstanceId struct {
	Name    string
	Privacy Privacy
	Region  Region
	OwnerID string
	UUID    string
}

func (i *InstanceId) String() string {
	return i.Name + fmt.Sprintf(string(i.Privacy), i.OwnerID) + fmt.Sprintf("~region(%s)", i.Region) + fmt.Sprintf("~nonce(%s)", i.UUID)
}

type InviteRequest struct {
	InstanceId  string `json:"instanceId"`
	MessageSlot int    `json:"messageSlot"`
}

type RequestInviteRequest struct {
	MessageSlot int `json:"messageSlot"`
}

type InviteResponse struct {
	MessageSlot int `json:"messageSlot"`
}

type TwoFactorAuthCode struct {
	Code string `json:"code"`
}

type Verify2FAResult struct {
	Verified bool `json:"verified"`
}

type SentNotification struct {
	Id             string `json:"id"`
	RecieverUserId string `json:"recieverUserId"`
	SenderUserId   string `json:"senderUserId"`
	Type           string `json:"type"`
	Message        string `json:"message"`
	Details        string `json:"details"`
	CreatedAt      string `json:"created_at"`
}

type Notification struct {
	Id           string    `json:"id"`
	SenderUserId string    `json:"senderUserId"`
	Type         string    `json:"type"`
	Message      string    `json:"message"`
	Details      string    `json:"details"`
	Seen         bool      `json:"seen"`
	CreatedAt    time.Time `json:"created_at"`
}

type CurrentUser struct {
	AcceptedTOSVersion     int    `json:"acceptedTOSVersion"`
	AcceptedPrivacyVersion int    `json:"acceptedPrivacyVersion"`
	AccountDeletionDate    string `json:"accountDeletionDate"`
	AccountDeletionLog     []struct {
		Message           string    `json:"message"`
		DeletionScheduled time.Time `json:"deletionScheduled"`
		DateTime          time.Time `json:"dateTime"`
	} `json:"accountDeletionLog"`
	ActiveFriends                  []string  `json:"activeFriends"`
	AllowAvatarCopying             bool      `json:"allowAvatarCopying"`
	Bio                            string    `json:"bio"`
	BioLinks                       []string  `json:"bioLinks"`
	CurrentAvatar                  string    `json:"currentAvatar"`
	CurrentAvatarAssetUrl          string    `json:"currentAvatarAssetUrl"`
	CurrentAvatarImageUrl          string    `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageUrl string    `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatarTags              []string  `json:"currentAvatarTags"`
	DateJoined                     string    `json:"date_joined"`
	DeveloperType                  string    `json:"developerType"`
	DisplayName                    string    `json:"displayName"`
	EmailVerified                  bool      `json:"emailVerified"`
	FallbackAvatar                 string    `json:"fallbackAvatar"`
	FriendKey                      string    `json:"friendKey"`
	Friends                        []string  `json:"friends"`
	HasBirthday                    bool      `json:"hasBirthday"`
	HideContentFilterSettings      bool      `json:"hideContentFilterSettings"`
	UserLanguage                   string    `json:"userLanguage"`
	UserLanguageCode               string    `json:"userLanguageCode"`
	HasEmail                       bool      `json:"hasEmail"`
	HasLoggedInFromClient          bool      `json:"hasLoggedInFromClient"`
	HasPendingEmail                bool      `json:"hasPendingEmail"`
	HomeLocation                   string    `json:"homeLocation"`
	Id                             string    `json:"id"`
	IsFriend                       bool      `json:"isFriend"`
	LastActivity                   time.Time `json:"last_activity"`
	LastLogin                      time.Time `json:"last_login"`
	LastPlatform                   string    `json:"last_platform"`
	ObfuscatedEmail                string    `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail         string    `json:"obfuscatedPendingEmail"`
	OculusId                       string    `json:"oculusId"`
	GoogleId                       string    `json:"googleId"`
	PicoId                         string    `json:"picoId"`
	ViveId                         string    `json:"viveId"`
	OfflineFriends                 []string  `json:"offlineFriends"`
	OnlineFriends                  []string  `json:"onlineFriends"`
	PastDisplayNames               []struct {
		DisplayName string    `json:"displayName"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"pastDisplayNames"`
	Presence struct {
		AvatarThumbnail     string   `json:"avatarThumbnail"`
		DisplayName         string   `json:"displayName"`
		Groups              []string `json:"groups"`
		Id                  string   `json:"id"`
		Instance            string   `json:"instance"`
		InstanceType        string   `json:"instanceType"`
		IsRejoining         string   `json:"isRejoining"`
		Platform            string   `json:"platform"`
		ProfilePicOverride  string   `json:"profilePicOverride"`
		Status              string   `json:"status"`
		TravelingToInstance string   `json:"travelingToInstance"`
		TravelingToWorld    string   `json:"travelingToWorld"`
		World               string   `json:"world"`
	} `json:"presence"`
	ProfilePicOverride string   `json:"profilePicOverride"`
	State              string   `json:"state"`
	Status             string   `json:"status"`
	StatusDescription  string   `json:"statusDescription"`
	StatusFirstTime    bool     `json:"statusFirstTime"`
	StatusHistory      []string `json:"statusHistory"`
	SteamDetails       struct {
	} `json:"steamDetails"`
	SteamId                  string    `json:"steamId"`
	Tags                     []string  `json:"tags"`
	TwoFactorAuthEnabled     bool      `json:"twoFactorAuthEnabled"`
	TwoFactorAuthEnabledDate time.Time `json:"twoFactorAuthEnabledDate"`
	Unsubscribe              bool      `json:"unsubscribe"`
	UpdatedAt                time.Time `json:"updated_at"`
	UserIcon                 string    `json:"userIcon"`
}

type Avatar struct {
	AssetUrl       string `json:"assetUrl"`
	AssetUrlObject struct {
	} `json:"assetUrlObject"`
	AuthorId          string    `json:"authorId"`
	AuthorName        string    `json:"authorName"`
	CreatedAt         time.Time `json:"created_at"`
	Description       string    `json:"description"`
	Featured          bool      `json:"featured"`
	Id                string    `json:"id"`
	ImageUrl          string    `json:"imageUrl"`
	Name              string    `json:"name"`
	ReleaseStatus     string    `json:"releaseStatus"`
	Tags              []string  `json:"tags"`
	ThumbnailImageUrl string    `json:"thumbnailImageUrl"`
	UnityPackageUrl   string    `json:"unityPackageUrl"`
	UnityPackages     []struct {
		AssetUrl       string `json:"assetUrl"`
		AssetUrlObject struct {
		} `json:"assetUrlObject"`
		AssetVersion    int       `json:"assetVersion"`
		CreatedAt       time.Time `json:"created_at"`
		Id              string    `json:"id"`
		Platform        string    `json:"platform"`
		PluginUrl       string    `json:"pluginUrl"`
		PluginUrlObject struct {
		} `json:"pluginUrlObject"`
		UnitySortNumber int64  `json:"unitySortNumber"`
		UnityVersion    string `json:"unityVersion"`
	} `json:"unityPackages"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
}

type Message struct {
	CanBeUpdated             bool      `json:"canBeUpdated"`
	Id                       string    `json:"id"`
	Message                  string    `json:"message"`
	MessageType              string    `json:"messageType"`
	RemainingCooldownMinutes int       `json:"remainingCooldownMinutes"`
	Slot                     int       `json:"slot"`
	UpdatedAt                time.Time `json:"updatedAt"`
}
