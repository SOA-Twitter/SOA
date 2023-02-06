package handlers

import (
	"ProfileService/data"
	"ProfileService/proto/auth"
	"ProfileService/proto/profile"
	social "ProfileService/proto/social"
	"context"
	"github.com/sony/gobreaker"
	"log"
	"net/http"
	"time"
)

type ProfileHandler struct {
	profile.UnimplementedProfileServiceServer
	l    *log.Logger
	repo *data.ProfileRepo
	as   auth.AuthServiceClient
	ss   social.SocialServiceClient
}

func NewProfileHandler(l *log.Logger, repo *data.ProfileRepo, as auth.AuthServiceClient, ss social.SocialServiceClient) *ProfileHandler {
	return &ProfileHandler{
		l:    l,
		repo: repo,
		as:   as,
		ss:   ss,
	}
}
func (pr *ProfileHandler) Register(ctx context.Context, r *profile.ProfileRegisterRequest) (*profile.ProfileRegisterResponse, error) {

	cb := pr.CircuitBreaker()

	pr.l.Println("Profile service - Register")
	user := &data.User{
		Username:       r.Username,
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		Email:          r.Email,
		Gender:         r.Gender,
		Country:        r.Country,
		Age:            int(r.Age),
		CompanyName:    r.CompanyName,
		CompanyWebsite: r.CompanyWebsite,
		Private:        true,
		Role:           r.Role,
	}

	// C BREAKER

	//bodyBytes, errCB := cb.Execute(func() (interface{}, error) {
	_, errCB := cb.Execute(func() (interface{}, error) {

		resp, err6 := pr.ss.RegUser(context.Background(), &social.RegUserRequest{
			Username: r.Username,
		})
		if err6 != nil {
			return nil, err6
		}
		// ?
		pr.l.Println("response of rpc to social service: ", resp)
		return resp, nil
	})
	if errCB != nil {
		pr.l.Println("\n------------", errCB, "\n--------------")
		// ?
		return nil, errCB
		//continue
	}

	errMongoReg := pr.repo.Register(user)
	if errMongoReg != nil {
		pr.l.Println("Error inserting user into db")
		return nil, errMongoReg
	}
	//pr.l.Println("Body bytes: " + string(bodyBytes.([]byte)))

	// C BREAKER
	return &profile.ProfileRegisterResponse{}, nil
}

func (pr *ProfileHandler) GetUserProfile(ctx context.Context, r *profile.UserProfRequest) (*profile.UserProfResponse, error) {
	pr.l.Println("Profile service - Get User Profile")

	user, err1 := pr.repo.GetByUsername(r.Username)
	if err1 != nil {
		pr.l.Println("Cannot find user")
		return nil, err1
	}

	return &profile.UserProfResponse{
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Gender:         user.Gender,
		Country:        user.Country,
		Age:            int32(user.Age),
		CompanyName:    user.CompanyName,
		CompanyWebsite: user.CompanyWebsite,
		Private:        user.Private,
		Role:           user.Role,
	}, nil

}

func (pr *ProfileHandler) ManagePrivacy(ctx context.Context, r *profile.ManagePrivacyRequest) (*profile.ManagePrivacyResponse, error) {
	pr.l.Println("Profile service - Manage privacy")
	claims, err := data.GetFromClaims(r.Token)
	if err != nil {
		pr.l.Println("Error getting claims")
		return &profile.ManagePrivacyResponse{
			Status: http.StatusNotFound,
		}, err
	}

	err1 := pr.repo.ChangePrivacy(claims.Username, r.Privacy)

	if err1 != nil {
		pr.l.Println("Cannot update account privacy for not-found user " + claims.Username)
		return nil, err1
	}
	return &profile.ManagePrivacyResponse{
		Privacy: r.Privacy,
		Status:  http.StatusOK,
	}, nil
}

func (pr *ProfileHandler) CircuitBreaker() *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(
		gobreaker.Settings{
			// circuit breaker's name
			Name: "cb",
			// maximum number of requests allowed to pass through when the circuit breaker is half-open
			MaxRequests: 1,
			// period of the open state, after which the state of the circuit breaker becomes half-open
			// if Timeout is 0, the timeout value of CircuitBreaker is set to 60 seconds
			Timeout: 2 * time.Second,
			// cyclic period of the closed state for the circuit breaker to clear the internal Counts
			// if Interval is 0, CircuitBreaker doesn't clear the internal Counts during the closed state
			Interval: 0,
			// called whenever a request fails in the closed state
			// the circuit breaker will come into the open state if this function returns true
			// type Counts struct {
			//    Requests             uint32
			//    TotalSuccesses       uint32
			//    TotalFailures        uint32
			//    ConsecutiveSuccesses uint32
			//    ConsecutiveFailures  uint32
			// }
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				return counts.ConsecutiveFailures > 0
			},
			// called whenever the state of the circuit breaker changes
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				pr.l.Printf("Circuit Breaker '%s' changed from '%s' to '%s'\n", name, from, to)
			},
		},
	)
}
