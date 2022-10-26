package constants

var VERSION = "0.0.0"

const (
	// Env key for env variable ENV
	Env = "ENV"
	// EnvPath key for cli arg variable ENV_PATH
	EnvPath = "ENV_PATH"
	// GateGuardianURL key for env variable GATE_GUARDIAN_URL
	GateGuardianURL = "GATE_GUARDIAN_URL"
	// KeyPort key for env variable PORT
	Port = "PORT"
	// AccessTokenExpiryTime key for env variable ACCESS_TOKEN_EXPIRY_TIME
	AccessTokenExpiryTime = "ACCESS_TOKEN_EXPIRY_TIME"
	// AdminSecret key for env variable ADMIN_SECRET
	AdminSecret = "ADMIN_SECRET"
	// DatabaseType key for env variable DATABASE_TYPE
	DatabaseURL = "DATABASE_URL"
	// DatabaseName key for env variable DATABASE_NAME
	DatabaseName = "DATABASE_NAME"
	// DatabaseName key for env variable DATABASE_NAME
	DatabaseNameSpace = "DATABASE_NAMESPACE"
	// DatabaseUsername key for env variable DATABASE_USERNAME
	DatabaseUsername = "DATABASE_USERNAME"
	// DatabasePassword key for env variable DATABASE_PASSWORD
	DatabasePassword = "DATABASE_PASSWORD"
	// DatabasePort key for env variable DATABASE_PORT
	DatabasePort = "DATABASE_PORT"
	// DatabaseHost key for env variable DATABASE_HOST
	DatabaseHost = "DATABASE_HOST"
	// DatabaseCert key for env variable DATABASE_CERT
	DatabaseCert = "DATABASE_CERT"
	// DatabaseCertKey key for env variable DATABASE_KEY
	DatabaseCertKey = "DATABASE_CERT_KEY"
	// DatabaseCACert key for env variable DATABASE_CA_CERT
	DatabaseCACert = "DATABASE_CA_CERT"
	// SmtpHost key for env variable SMTP_HOST
	SmtpHost = "SMTP_HOST"
	// SmtpPort key for env variable SMTP_PORT
	SmtpPort = "SMTP_PORT"
	// SmtpUsername key for env variable SMTP_USERNAME
	SmtpUsername = "SMTP_USERNAME"
	// SmtpPassword key for env variable SMTP_PASSWORD
	SmtpPassword = "SMTP_PASSWORD"
	// SenderEmail key for env variable SENDER_EMAIL
	SenderEmail = "SENDER_EMAIL"
	// IsEmailServiceEnabled key for env variable IS_EMAIL_SERVICE_ENABLED
	IsEmailServiceEnabled = "IS_EMAIL_SERVICE_ENABLED"
	// AppCookieSecure key for env variable APP_COOKIE_SECURE
	AppCookieSecure = "APP_COOKIE_SECURE"
	// AdminCookieSecure key for env variable ADMIN_COOKIE_SECURE
	AdminCookieSecure = "ADMIN_COOKIE_SECURE"
	// JwtType key for env variable JWT_TYPE
	JwtType = "JWT_TYPE"
	// JwtSecret key for env variable JWT_SECRET
	JwtSecret = "JWT_SECRET"
	// JwtPrivateKey key for env variable JWT_PRIVATE_KEY
	JwtPrivateKey = "JWT_PRIVATE_KEY"
	// JwtPublicKey key for env variable JWT_PUBLIC_KEY
	JwtPublicKey = "JWT_PUBLIC_KEY"
	// AppURL key for env variable APP_URL
	AppURL = "APP_URL"
	// EnvKeyResetPasswordURL key for env variable RESET_PASSWORD_URL
	ResetPasswordURL = "RESET_PASSWORD_URL"
	// DefaultJwtRoleClaim key for env variable DEFAULT_JWT_ROLE_CLAIM
	DefaultJwtRoleClaim = "DEFAULT_JWT_ROLE_CLAIM"
	// JwtRoleClaim key for env variable JWT_ROLE_CLAIM
	JwtRoleClaim = "JWT_ROLE_CLAIM"
	// GoogleClientID key for env variable GOOGLE_CLIENT_ID
	GoogleClientID = "GOOGLE_CLIENT_ID"
	// GoogleClientSecret key for env variable GOOGLE_CLIENT_SECRET
	GoogleClientSecret = "GOOGLE_CLIENT_SECRET"
	// GithubClientID key for env variable GITHUB_CLIENT_ID
	GithubClientID = "GITHUB_CLIENT_ID"
	// GithubClientSecret key for env variable GITHUB_CLIENT_SECRET
	GithubClientSecret = "GITHUB_CLIENT_SECRET"
	// FacebookClientID key for env variable FACEBOOK_CLIENT_ID
	FacebookClientID = "FACEBOOK_CLIENT_ID"
	// FacebookClientSecret key for env variable FACEBOOK_CLIENT_SECRET
	FacebookClientSecret = "FACEBOOK_CLIENT_SECRET"
	// LinkedinClientID key for env variable LINKEDIN_CLIENT_ID
	LinkedInClientID = "LINKEDIN_CLIENT_ID"
	// LinkedinClientSecret key for env variable LINKEDIN_CLIENT_SECRET
	LinkedInClientSecret = "LINKEDIN_CLIENT_SECRET"
	// AppleClientID key for env variable APPLE_CLIENT_ID
	AppleClientID = "APPLE_CLIENT_ID"
	// AppleClientSecret key for env variable APPLE_CLIENT_SECRET
	AppleClientSecret = "APPLE_CLIENT_SECRET"
	// TwitterClientID key for env variable TWITTER_CLIENT_ID
	TwitterClientID = "TWITTER_CLIENT_ID"
	// TwitterClientSecret key for env variable TWITTER_CLIENT_SECRET
	TwitterClientSecret = "TWITTER_CLIENT_SECRET"
	// OrganizationName key for env variable ORGANIZATION_NAME
	OrganizationName = "ORGANIZATION_NAME"
	// OrganizationLogo key for env variable ORGANIZATION_LOGO
	OrganizationLogo = "ORGANIZATION_LOGO"
	// CustomAccessTokenScript key for env variable CUSTOM_ACCESS_TOKEN_SCRIPT
	CustomAccessTokenScript = "CUSTOM_ACCESS_TOKEN_SCRIPT"

	// Not Exposed Keys
	// ClientID key for env variable CLIENT_ID
	ClientID = "CLIENT_ID"
	// ClientSecret key for env variable CLIENT_SECRET
	ClientSecret = "CLIENT_SECRET"
	// EncryptionKey key for env variable ENCRYPTION_KEY
	EncryptionKey = "ENCRYPTION_KEY"
	// JWK key for env variable JWK
	JWK = "JWK"

	// Boolean variables
	// IsProd key for env variable IS_PROD
	IsProd = "IS_PROD"
	// DisableEmailVerification key for env variable DISABLE_EMAIL_VERIFICATION
	DisableEmailVerification = "DISABLE_EMAIL_VERIFICATION"
	// DisableBasicAuthentication key for env variable DISABLE_BASIC_AUTH
	DisableBasicAuthentication = "DISABLE_BASIC_AUTHENTICATION"
	// DisableMagicLinkLogin key for env variable DISABLE_MAGIC_LINK_LOGIN
	DisableMagicLinkLogin = "DISABLE_MAGIC_LINK_LOGIN"
	// DisableLoginPage key for env variable DISABLE_LOGIN_PAGE
	DisableLoginPage = "DISABLE_LOGIN_PAGE"
	// DisableSignUp key for env variable DISABLE_SIGN_UP
	DisableSignUp = "DISABLE_SIGN_UP"
	// DisableStrongPassword key for env variable DISABLE_STRONG_PASSWORD
	DisableStrongPassword = "DISABLE_STRONG_PASSWORD"
	// EnforceMultiFactorAuthentication is key for env variable ENFORCE_MULTI_FACTOR_AUTHENTICATION
	// If enforced and changed later on, existing user will have MFA but new user will not have MFA
	EnforceMultiFactorAuthentication = "ENFORCE_MULTI_FACTOR_AUTHENTICATION"
	// DisableMultiFactorAuthentication is key for env variable DISABLE_MULTI_FACTOR_AUTHENTICATION
	// this variable is used to completely disable multi factor authentication. It will have no effect on profile preference
	DisableMultiFactorAuthentication = "DISABLE_MULTI_FACTOR_AUTHENTICATION"

	// Slice variables
	// Roles key for env variable ROLES
	Roles = "ROLES"
	// ProtectedRoles key for env variable PROTECTED_ROLES
	ProtectedRoles = "PROTECTED_ROLES"
	// DefaultRoles key for env variable DEFAULT_ROLES
	DefaultRoles = "DEFAULT_ROLES"
	// AllowedOrigins key for env variable ALLOWED_ORIGINS
	AllowedOrigins = "ALLOWED_ORIGINS"
)
