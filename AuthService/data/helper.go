package data

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"log"
	"net/smtp"
	"os"
	"time"
)

var SampleSecretKey = []byte("SecretYouShouldHide")

func CreateJwt(email string, role string, username string) (string, error) {
	claims := &Claims{
		Role:     role,
		Email:    email,
		Username: username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 120).Unix(),
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

func SendEmail(providedEmail string, intention string) (string, error) {
	const accountActivationPath = "https://localhost:8081/auth/activate/"
	const accountRecoveryPath = "https://localhost:4200/recover-account"
	// Sender data
	from := os.Getenv("MAIL_ADDRESS")

	password := os.Getenv("MAIL_PASSWORD")

	// Receiver email
	to := []string{
		providedEmail,
	}

	// smtp server config
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	address := smtpHost + ":" + smtpPort
	activationUUID := generateActivationUUID()
	var subject string
	var body string

	if intention == "activation" {
		subject = "Twitter clone account activation"
		body = "Follow the verification link to activate your Twitterclone account: \n" + accountActivationPath + activationUUID
	} else if intention == "recovery" {
		subject = "Twitter clone password recovery"
		body = "To reset your password, copy the given code & then follow the recovery link: \n" + activationUUID + "\n" + accountRecoveryPath
	}
	// Text
	stringMsg :=
		"From: " + from + "\n" +
			"To: " + to[0] + "\n" +
			"Subject: " + subject + "\n\n" +
			body

	message := []byte(stringMsg)

	// Email Sender Auth
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		log.Println("Error sending mail", err)
		return "", err
	}
	log.Println("Mail successfully sent")
	return activationUUID, nil
}

func generateActivationUUID() string {
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
