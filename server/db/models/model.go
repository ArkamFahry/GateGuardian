package models

// Tables available for GateGuardian in the database
type CollectionList struct {
	User                string
	VerificationRequest string
	Session             string
	Env                 string
	Webhook             string
	WebhookLog          string
	EmailTemplate       string
	OTP                 string
	Metadata            string
	Permissions         string
}

var (
	// Prefix for table name
	Prefix = "gate_guardian_"
	// Tables in the database
	Collections = CollectionList{
		User:                Prefix + "users",
		VerificationRequest: Prefix + "verification_requests",
		Session:             Prefix + "sessions",
		Env:                 Prefix + "env",
		Webhook:             Prefix + "webhooks",
		WebhookLog:          Prefix + "webhook_logs",
		EmailTemplate:       Prefix + "email_templates",
		OTP:                 Prefix + "otps",
	}
)