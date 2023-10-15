package api

type Settings struct {
	Protected                bool   `json:"protected"`
	ScreenName               string `json:"screen_name"`
	AlwaysUseHTTPS           bool   `json:"always_use_https"`
	UseCookiePersonalization bool   `json:"use_cookie_personalization"`
	SleepTime                struct {
		Enabled   bool        `json:"enabled"`
		EndTime   interface{} `json:"end_time"`
		StartTime interface{} `json:"start_time"`
	} `json:"sleep_time"`
	GeoEnabled                                   bool        `json:"geo_enabled"`
	Language                                     string      `json:"language"`
	DiscoverableByEmail                          bool        `json:"discoverable_by_email"`
	DiscoverableByMobilePhone                    bool        `json:"discoverable_by_mobile_phone"`
	DisplaySensitiveMedia                        bool        `json:"display_sensitive_media"`
	PersonalizedTrends                           bool        `json:"personalized_trends"`
	AllowMediaTagging                            string      `json:"allow_media_tagging"`
	AllowContributorRequest                      string      `json:"allow_contributor_request"`
	AllowAdsPersonalization                      bool        `json:"allow_ads_personalization"`
	AllowLoggedOutDevicePersonalization          bool        `json:"allow_logged_out_device_personalization"`
	AllowLocationHistoryPersonalization          bool        `json:"allow_location_history_personalization"`
	AllowSharingDataForThirdPartyPersonalization bool        `json:"allow_sharing_data_for_third_party_personalization"`
	AllowDmsFrom                                 string      `json:"allow_dms_from"`
	AlwaysAllowDmsFromSubscribers                interface{} `json:"always_allow_dms_from_subscribers"`
	AllowDmGroupsFrom                            string      `json:"allow_dm_groups_from"`
	TranslatorType                               string      `json:"translator_type"`
	TrendLocation                                []struct {
		Name        string `json:"name"`
		CountryCode string `json:"countryCode"`
		URL         string `json:"url"`
		Woeid       int    `json:"woeid"`
		PlaceType   struct {
			Name string `json:"name"`
			Code int    `json:"code"`
		} `json:"placeType"`
		Parentid int    `json:"parentid"`
		Country  string `json:"country"`
	} `json:"trend_location"`
	CountryCode                                     string      `json:"country_code"`
	NsfwUser                                        bool        `json:"nsfw_user"`
	NsfwAdmin                                       bool        `json:"nsfw_admin"`
	RankedTimelineSetting                           interface{} `json:"ranked_timeline_setting"`
	RankedTimelineEligible                          interface{} `json:"ranked_timeline_eligible"`
	AddressBookLiveSyncEnabled                      bool        `json:"address_book_live_sync_enabled"`
	UniversalQualityFilteringEnabled                string      `json:"universal_quality_filtering_enabled"`
	DmReceiptSetting                                string      `json:"dm_receipt_setting"`
	AltTextComposeEnabled                           interface{} `json:"alt_text_compose_enabled"`
	MentionFilter                                   string      `json:"mention_filter"`
	AllowAuthenticatedPeriscopeRequests             bool        `json:"allow_authenticated_periscope_requests"`
	ProtectPasswordReset                            bool        `json:"protect_password_reset"`
	RequirePasswordLogin                            bool        `json:"require_password_login"`
	RequiresLoginVerification                       bool        `json:"requires_login_verification"`
	ExtSharingAudiospacesListeningDataWithFollowers bool        `json:"ext_sharing_audiospaces_listening_data_with_followers"`
	Ext                                             struct {
		SsoConnections struct {
			R struct {
				Ok []interface{} `json:"ok"`
			} `json:"r"`
			TTL int `json:"ttl"`
		} `json:"ssoConnections"`
	} `json:"ext"`
	DmQualityFilter  string `json:"dm_quality_filter"`
	AutoplayDisabled bool   `json:"autoplay_disabled"`
	SettingsMetadata struct {
		IsEu string `json:"is_eu"`
	} `json:"settings_metadata"`
}

type UserByScreenName struct {
	Data struct {
		User struct {
			Result struct {
				Typename                   string `json:"__typename"`
				ID                         string `json:"id"`
				RestID                     string `json:"rest_id"`
				AffiliatesHighlightedLabel struct {
					Label struct {
						URL struct {
							URL     string `json:"url"`
							URLType string `json:"urlType"`
						} `json:"url"`
						Badge struct {
							URL string `json:"url"`
						} `json:"badge"`
						Description          string `json:"description"`
						UserLabelType        string `json:"userLabelType"`
						UserLabelDisplayType string `json:"userLabelDisplayType"`
					} `json:"label"`
				} `json:"affiliates_highlighted_label"`
				HasGraduatedAccess bool   `json:"has_graduated_access"`
				IsBlueVerified     bool   `json:"is_blue_verified"`
				ProfileImageShape  string `json:"profile_image_shape"`
				Legacy             struct {
					CanDm               bool   `json:"can_dm"`
					CanMediaTag         bool   `json:"can_media_tag"`
					CreatedAt           string `json:"created_at"`
					DefaultProfile      bool   `json:"default_profile"`
					DefaultProfileImage bool   `json:"default_profile_image"`
					Description         string `json:"description"`
					Entities            struct {
						Description struct {
							Urls []interface{} `json:"urls"`
						} `json:"description"`
					} `json:"entities"`
					FastFollowersCount      int           `json:"fast_followers_count"`
					FavouritesCount         int           `json:"favourites_count"`
					FollowersCount          int           `json:"followers_count"`
					FriendsCount            int           `json:"friends_count"`
					HasCustomTimelines      bool          `json:"has_custom_timelines"`
					IsTranslator            bool          `json:"is_translator"`
					ListedCount             int           `json:"listed_count"`
					Location                string        `json:"location"`
					MediaCount              int           `json:"media_count"`
					Name                    string        `json:"name"`
					NormalFollowersCount    int           `json:"normal_followers_count"`
					PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
					PossiblySensitive       bool          `json:"possibly_sensitive"`
					ProfileBannerURL        string        `json:"profile_banner_url"`
					ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
					ProfileInterstitialType string        `json:"profile_interstitial_type"`
					ScreenName              string        `json:"screen_name"`
					StatusesCount           int           `json:"statuses_count"`
					TranslatorType          string        `json:"translator_type"`
					Verified                bool          `json:"verified"`
					WantRetweets            bool          `json:"want_retweets"`
					WithheldInCountries     []interface{} `json:"withheld_in_countries"`
				} `json:"legacy"`
				Professional struct {
					RestID           string        `json:"rest_id"`
					ProfessionalType string        `json:"professional_type"`
					Category         []interface{} `json:"category"`
				} `json:"professional"`
				SuperFollowEligible   bool `json:"super_follow_eligible"`
				SmartBlockedBy        bool `json:"smart_blocked_by"`
				SmartBlocking         bool `json:"smart_blocking"`
				LegacyExtendedProfile struct {
				} `json:"legacy_extended_profile"`
				IsProfileTranslatable           bool `json:"is_profile_translatable"`
				HasHiddenLikesOnProfile         bool `json:"has_hidden_likes_on_profile"`
				HasHiddenSubscriptionsOnProfile bool `json:"has_hidden_subscriptions_on_profile"`
				VerificationInfo                struct {
					IsIdentityVerified bool `json:"is_identity_verified"`
					Reason             struct {
						Description struct {
							Text     string `json:"text"`
							Entities []struct {
								FromIndex int `json:"from_index"`
								ToIndex   int `json:"to_index"`
								Ref       struct {
									URL     string `json:"url"`
									URLType string `json:"url_type"`
								} `json:"ref"`
							} `json:"entities"`
						} `json:"description"`
						VerifiedSinceMsec    string `json:"verified_since_msec"`
						OverrideVerifiedYear int    `json:"override_verified_year"`
					} `json:"reason"`
				} `json:"verification_info"`
				HighlightsInfo struct {
					CanHighlightTweets bool   `json:"can_highlight_tweets"`
					HighlightedTweets  string `json:"highlighted_tweets"`
				} `json:"highlights_info"`
				BusinessAccount struct {
				} `json:"business_account"`
				CreatorSubscriptionsCount int `json:"creator_subscriptions_count"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

type Followers struct {
	Data struct {
		User struct {
			Result struct {
				Typename string `json:"__typename"`
				Timeline struct {
					Timeline struct {
						Instructions []struct {
							Type      string `json:"type"`
							Direction string `json:"direction,omitempty"`
							Entries   []struct {
								EntryID   string `json:"entryId"`
								SortIndex string `json:"sortIndex"`
								Content   struct {
									EntryType   string `json:"entryType"`
									Typename    string `json:"__typename"`
									ItemContent struct {
										ItemType    string `json:"itemType"`
										Typename    string `json:"__typename"`
										UserResults struct {
											Result struct {
												Typename                   string `json:"__typename"`
												ID                         string `json:"id"`
												RestID                     string `json:"rest_id"`
												AffiliatesHighlightedLabel struct {
												} `json:"affiliates_highlighted_label"`
												HasGraduatedAccess bool   `json:"has_graduated_access"`
												IsBlueVerified     bool   `json:"is_blue_verified"`
												ProfileImageShape  string `json:"profile_image_shape"`
												Legacy             struct {
													CanDm               bool   `json:"can_dm"`
													CanMediaTag         bool   `json:"can_media_tag"`
													CreatedAt           string `json:"created_at"`
													DefaultProfile      bool   `json:"default_profile"`
													DefaultProfileImage bool   `json:"default_profile_image"`
													Description         string `json:"description"`
													Entities            struct {
														Description struct {
															Urls []interface{} `json:"urls"`
														} `json:"description"`
													} `json:"entities"`
													FastFollowersCount      int           `json:"fast_followers_count"`
													FavouritesCount         int           `json:"favourites_count"`
													FollowersCount          int           `json:"followers_count"`
													FriendsCount            int           `json:"friends_count"`
													HasCustomTimelines      bool          `json:"has_custom_timelines"`
													IsTranslator            bool          `json:"is_translator"`
													ListedCount             int           `json:"listed_count"`
													Location                string        `json:"location"`
													MediaCount              int           `json:"media_count"`
													Name                    string        `json:"name"`
													NormalFollowersCount    int           `json:"normal_followers_count"`
													PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
													PossiblySensitive       bool          `json:"possibly_sensitive"`
													ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
													ProfileInterstitialType string        `json:"profile_interstitial_type"`
													ScreenName              string        `json:"screen_name"`
													StatusesCount           int           `json:"statuses_count"`
													TranslatorType          string        `json:"translator_type"`
													Verified                bool          `json:"verified"`
													WantRetweets            bool          `json:"want_retweets"`
													WithheldInCountries     []interface{} `json:"withheld_in_countries"`
												} `json:"legacy"`
											} `json:"result"`
										} `json:"user_results"`
										UserDisplayType string `json:"userDisplayType"`
									} `json:"itemContent"`
									ClientEventInfo struct {
										Component string `json:"component"`
										Element   string `json:"element"`
									} `json:"clientEventInfo"`
								} `json:"content"`
							} `json:"entries,omitempty"`
						} `json:"instructions"`
					} `json:"timeline"`
				} `json:"timeline"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

type Follower struct {
	Name string
	Id   string
}

type NewMessage struct {
	ConversationID    string `json:"conversation_id"`
	RecipientIds      bool   `json:"recipient_ids"`
	RequestID         string `json:"request_id"`
	Text              string `json:"text"`
	CardsPlatform     string `json:"cards_platform"`
	IncludeCards      int    `json:"include_cards"`
	IncludeQuoteCount bool   `json:"include_quote_count"`
	DmUsers           bool   `json:"dm_users"`
}
