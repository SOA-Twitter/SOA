package data

type ActivationRequest struct {
	ActivationUUID string `json:"activationUUID" `
	Email          string `json:"email" `
}

type RecoveryRequest struct {
	RecoveryUUID string `json:"recoveryUUID" `
	Email        string `json:"email" `
}
