package data

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"log"
	"net/smtp"
	"time"
)

var SampleSecretKey = []byte("SecretYouShouldHide")

func CreateJwt(email string, role string) (string, error) {
	claims := &Claims{
		Role:  role,
		Email: email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 1200).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SampleSecretKey)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return tokenString, nil
}
func ValidateJwt(tokenString string) error {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SampleSecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return err
		}
		return err
	}
	// NOTE > read below
	if !token.Valid {
		return err
	}
	return nil
}

func SendAccountActivationEmail(providedEmail string) (string, error) {
	const accountActivationPath = "https://localhost:8081/auth/activate/"
	// Sender data
	//from := "twittertest282@outlook.com"

	from := "twittertest282@gmail.com"

	//password := "HaL4WI5p7m*8W(o)"
	//password := "wxxlhqljvcrvsefg"
	password := "xfxkdwgmyomnadua"

	// Receiver email
	to := []string{
		providedEmail,
	}

	// smtp server config
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	//smtpHost := "smtp.office365.com"
	// smtpHost := "smtp-mail.outlook.com"
	//smtpPort := "587"

	// Generating Activation uuid
	activationUUID := generateActivationUUID()

	// Text
	body := "Follow the verification link to activate your Twitterclone account: " + accountActivationPath + activationUUID
	stringMsg :=
		"From: " + from + "\n" +
			"To: " + to[0] + "\n" +
			"Subject: Twitter clone account activation\n\n" +
			body

	message := []byte(stringMsg)

	// Email Sender Auth
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println("\n\nPUKLO\n\n", err)
		return "", err
	}
	log.Println("\n\nUSPESNO SLANJE MEJLA\n\n")
	return activationUUID, nil
}
func generateActivationUUID() string {
	// *TODO: Generate uuid
	//requestUUID := uuid.NewUUID()
	requestUUID := uuid.New().String()
	return requestUUID
}
func GetFromClaims(token string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SampleSecretKey, nil
	})
	return claims, err
}

//func SaveAccountActivationRequest(activationUUID string, email string) error {
//	SaveActivationRequest()
//}

/*
	NOTE:  VALID = JWT Library functionality:

	func (c StandardClaims) Valid() error {
		vErr := new(ValidationError)
		now := TimeFunc().Unix()
		if !c.VerifyExpiresAt(now, false) {
			delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
			vErr.Inner = fmt.Errorf("%s by %s", ErrTokenExpired, delta)
			vErr.Errors |= ValidationErrorExpired
		}
*/
