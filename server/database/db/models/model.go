package models

// Represents the entities models as collections or tables
type CollectionList struct {
	User                string `json:"user"`
	VerificationRequest string `json:"verification_request"`
	Session             string `json:"session"`
	Env                 string `json:"env"`
}

var (
	// Prefix for table or collection names
	Prefix = "gate_guardian_"

	// Collections or tables available for gate_guardian in the database except for Gorm
	Collection = CollectionList{
		User:                Prefix + "users",
		VerificationRequest: Prefix + "verification_requests",
		Session:             Prefix + "sessions",
		Env:                 Prefix + "envs",
	}
)
