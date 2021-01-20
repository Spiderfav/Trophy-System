package user

import "encoding/json"

type user struct {
	Oauth        string
	RefreshToken string
}

type message struct {
	MessageStart string
	Body         string
}

func me(u user, m message) {
	var usersUrl string = "https://us-prof.np.community.playstation.net/userProfile/v1/users/"
	var endpoint string = "me/profile2?fields=npId,onlineId,avatarUrls,plus,aboutMe,languagesUsed,trophySummary(@default,progress,earnedTrophies),isOfficiallyVerified,personalDetail(@default,profilePictureUrls),personalDetailSharing,personalDetailSharingRequestMessageFlag,primaryOnlineStatus,presences(@titleInfo,hasBroadcastData),friendRelation,requestMessageFlag,blocking,mutualFriendsCount,following,followerCount,friendsCount,followingUsersCount&avatarSizes=m,xl&profilePictureSizes=m,xl&languagesUsedLanguageSet=set3&psVitaTitleIcon=circled&titleIconSize=s"
	header := message{"Authorisation", "Bearer " + u.Oauth}
	b, err := json.Marshal(header)
	if err != nil {
		panic(err)
	}

}
