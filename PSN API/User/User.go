package user

import (
	"io/ioutil"
	"net/http"
)

type User struct {
	Oauth        string
	RefreshToken string
}

type Message struct {
	MessageStart string
	Body         string
}

func me(oauth string, token string) (data []byte) {

	userMade := User{oauth, token}
	var usersURL string = "https://us-prof.np.community.playstation.net/userProfile/v1/users/"
	var endpoint string = "me/profile2?fields=npId,onlineId,avatarUrls,plus,aboutMe,languagesUsed,trophySummary(@default,progress,earnedTrophies),isOfficiallyVerified,personalDetail(@default,profilePictureUrls),personalDetailSharing,personalDetailSharingRequestMessageFlag,primaryOnlineStatus,presences(@titleInfo,hasBroadcastData),friendRelation,requestMessageFlag,blocking,mutualFriendsCount,following,followerCount,friendsCount,followingUsersCount&avatarSizes=m,xl&profilePictureSizes=m,xl&languagesUsedLanguageSet=set3&psVitaTitleIcon=circled&titleIconSize=s"

	req, err := http.NewRequest("GET", usersURL+endpoint, nil)
	req.Header.Add("Authorisation", "Bearer "+userMade.Oauth)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	data, _ = ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return data

}
