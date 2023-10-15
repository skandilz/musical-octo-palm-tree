package api

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Twitter struct {
	Token string
	Name  string
	Id    string
}

func NewTwitter(token string) *Twitter {
	twitter := &Twitter{Token: token}
	twitter.Init()
	return twitter
}

func (t *Twitter) Init() {
	t.Name = t.GetNameFromToken()
	t.Id = t.GetIdFromName(t.Name)
}

func (t *Twitter) GetIdFromName(name string) (id string) {
	csrf := generateCSRF()
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://twitter.com/i/api/graphql/G3KGOASz96M-Qu0nwmGXNg/UserByScreenName?variables=%7B%22screen_name%22%3A%22"+name+"%22%2C%22withSafetyModeUserFields%22%3Atrue%7D&features=%7B%22hidden_profile_likes_enabled%22%3Atrue%2C%22hidden_profile_subscriptions_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22subscriptions_verification_info_is_identity_verified_enabled%22%3Atrue%2C%22subscriptions_verification_info_verified_since_enabled%22%3Atrue%2C%22highlights_tweets_tab_ui_enabled%22%3Atrue%2C%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%7D&fieldToggles=%7B%22withAuxiliaryUserLabels%22%3Afalse%7D", nil)
	req.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	req.Header.Add("cookie", fmt.Sprintf("ct0=%s; auth_token=%s;", csrf, t.Token))
	req.Header.Add("x-csrf-token", csrf)
	req.Header.Add("x-twitter-auth-type", "OAuth2Session")

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var data UserByScreenName
	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}

	return data.Data.User.Result.RestID
}

func (t *Twitter) GetNameFromToken() (name string) {
	csrf := generateCSRF()
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.twitter.com/1.1/account/settings.json?include_mention_filter=true&include_nsfw_user_flag=true&include_nsfw_admin_flag=true&include_ranked_timeline=true&include_alt_text_compose=true&ext=ssoConnections&include_country_code=true&include_ext_dm_nsfw_media_filter=true&include_ext_sharing_audiospaces_listening_data_with_followers=true", nil)
	req.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	req.Header.Add("cookie", fmt.Sprintf("ct0=%s; auth_token=%s;", csrf, t.Token))
	req.Header.Add("x-csrf-token", csrf)
	req.Header.Add("x-twitter-auth-type", "OAuth2Session")

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var data Settings
	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}

	return data.ScreenName
}

func (t *Twitter) GetFollowersFromId(id string) (_followers []Follower) {
	csrf := generateCSRF()
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://twitter.com/i/api/graphql/ihMPm0x-pC35X86L_nUp_Q/Followers?variables=%7B%22userId%22%3A%22"+id+"%22%2C%22count%22%3A20%2C%22includePromotedContent%22%3Afalse%7D&features=%7B%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22responsive_web_home_pinned_timelines_enabled%22%3Atrue%2C%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22tweetypie_unmention_optimization_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22graphql_is_translatable_rweb_tweet_is_translatable_enabled%22%3Atrue%2C%22view_counts_everywhere_api_enabled%22%3Atrue%2C%22longform_notetweets_consumption_enabled%22%3Atrue%2C%22responsive_web_twitter_article_tweet_consumption_enabled%22%3Afalse%2C%22tweet_awards_web_tipping_enabled%22%3Afalse%2C%22freedom_of_speech_not_reach_fetch_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Atrue%2C%22longform_notetweets_rich_text_read_enabled%22%3Atrue%2C%22longform_notetweets_inline_media_enabled%22%3Atrue%2C%22responsive_web_media_download_video_enabled%22%3Afalse%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D", nil)
	req.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	req.Header.Add("cookie", fmt.Sprintf("ct0=%s; auth_token=%s;", csrf, t.Token))
	req.Header.Add("x-csrf-token", csrf)
	req.Header.Add("x-twitter-auth-type", "OAuth2Session")

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var data Followers
	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}

	var followers []Follower

	for i := 0; i < len(data.Data.User.Result.Timeline.Timeline.Instructions[3].Entries); i++ {
		element := data.Data.User.Result.Timeline.Timeline.Instructions[3].Entries[i]
		if element.Content.Typename != "TimelineTimelineItem" {
			continue
		}

		result := element.Content.ItemContent.UserResults.Result

		if result.Legacy.CanDm != true {
			continue
		}

		followers = append(followers, Follower{Id: result.RestID, Name: result.Legacy.ScreenName})
	}

	return followers
}

func (t *Twitter) SendMessage(id string, message string) (status bool) {
	csrf := generateCSRF()
	client := &http.Client{}
	body := &NewMessage{
		ConversationID:    fmt.Sprintf("%s-%s", t.Id, id),
		RecipientIds:      false,
		RequestID:         uuid.NewString(),
		Text:              message,
		CardsPlatform:     "Web-12",
		IncludeCards:      1,
		IncludeQuoteCount: true,
		DmUsers:           false,
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "https://twitter.com/i/api/1.1/dm/new2.json?ext=mediaColor%2CaltText%2CmediaStats%2ChighlightedLabel%2ChasNftAvatar%2CvoiceInfo%2CbirdwatchPivot%2CsuperFollowMetadata%2CunmentionInfo%2CeditControl&include_ext_alt_text=true&include_ext_limited_action_results=true&include_reply_count=1&tweet_mode=extended&include_ext_views=true&include_groups=true&include_inbox_timelines=true&include_ext_media_color=true&supports_reactions=true", strings.NewReader(string(jsonBody)))
	req.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	req.Header.Add("cookie", fmt.Sprintf("ct0=%s; auth_token=%s;", csrf, t.Token))
	req.Header.Add("x-csrf-token", csrf)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-twitter-auth-type", "OAuth2Session")

	res, _ := client.Do(req)

	return res.StatusCode == 200
}

func generateCSRF() (csrf string) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	const length = 32
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		randomDigit := random.Intn(16)
		result[i] = fmt.Sprintf("%X", randomDigit)[0]
	}

	return strings.ToLower(string(result))
}
