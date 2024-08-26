package auth

import (
	"os"
	"recipiary/models"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Init() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// https://try.supertokens.com is for demo purposes. Replace this with the address of your core instance (sign up on supertokens.com), or self host a core.
			ConnectionURI: os.Getenv("SUPERTOKENS_CONNECTION_URI"),
			APIKey:        os.Getenv("SUPERTOKENS_API_KEY"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "recipiary",
			APIDomain:       "http://localhost:8080",
			WebsiteDomain:   "http://localhost:5173",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdparty.Init(&tpmodels.TypeInput{
				Override: &tpmodels.OverrideStruct{
					Functions: func(originalImplementation tpmodels.RecipeInterface) tpmodels.RecipeInterface {
						// create a copy of the originalImplementation
						originalSignInUp := *originalImplementation.SignInUp

						// override the sign in up function
						(*originalImplementation.SignInUp) = func(thirdPartyID string, thirdPartyUserID string, email string, oAuthTokens map[string]interface{}, rawUserInfoFromProvider tpmodels.TypeRawUserInfoFromProvider, tenantId string, userContext *map[string]interface{}) (tpmodels.SignInUpResponse, error) {
							// First we call the original implementation of SignInUp.
							response, err := originalSignInUp(thirdPartyID, thirdPartyUserID, email, oAuthTokens, rawUserInfoFromProvider, tenantId, userContext)
							if err != nil {
								return tpmodels.SignInUpResponse{}, err
							}
							if response.OK != nil {
								// sign in / up was successful

								// user object contains the ID and email of the user
								user := response.OK.User

								if response.OK.CreatedNewUser {
									account := models.Account{UserID: user.ID, UserEmail: user.Email, Provider: thirdPartyID}
									models.DB.Create(&account)
								}
							}
							return response, nil
						}

						return originalImplementation
					},
				},
				SignInAndUpFeature: tpmodels.TypeInputSignInAndUp{
					Providers: []tpmodels.ProviderInput{
						// We have provided you with development keys which you can use for testing.
						// IMPORTANT: Please replace them with your own OAuth keys for production use.
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "google",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID:     "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
										ClientSecret: "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",
									},
								},
							},
						},
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "github",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID:     "467101b197249757c71f",
										ClientSecret: "e97051221f4b6426e8fe8d51486396703012f5bd",
									},
								},
							},
						},
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "apple",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID: "4398792-io.supertokens.example.service",
										AdditionalConfig: map[string]interface{}{
											"keyId":      "7M48Y4RYDL",
											"privateKey": "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgu8gXs+XYkqXD6Ala9Sf/iJXzhbwcoG5dMh1OonpdJUmgCgYIKoZIzj0DAQehRANCAASfrvlFbFCYqn3I2zeknYXLwtH30JuOKestDbSfZYxZNMqhF/OzdZFTV0zc5u5s3eN+oCWbnvl0hM+9IW0UlkdA\n-----END PRIVATE KEY-----",
											"teamId":     "YWQCXGJRJL",
										},
									},
								},
							},
						},
					},
				},
			}),
			session.Init(nil), // initializes session features
			dashboard.Init(nil),
		},
	})
	if err != nil {
		panic(err.Error())
	}
}
