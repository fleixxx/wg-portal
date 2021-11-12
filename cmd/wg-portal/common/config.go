package common

import (
	"os"

	"github.com/h44z/wg-portal/internal/persistence"
	"github.com/h44z/wg-portal/internal/portal"
	"gopkg.in/yaml.v3"
)

type OauthFields struct {
	UserIdentifier string `yaml:"user_identifier"`
	Email          string `yaml:"email"`
	Firstname      string `yaml:"firstname"`
	Lastname       string `yaml:"lastname"`
	Phone          string `yaml:"phone"`
	Department     string `yaml:"department"`
	IsAdmin        string `yaml:"is_admin"`
}

type LdapAuthProvider struct {
}

type OpenIDConnectProvider struct {
	// ProviderName is an internal name that is used to distinguish oauth endpoints. It must not contain spaces or special characters.
	ProviderName string `yaml:"provider_name"`

	// DisplayName is shown to the user on the login page. If it is empty, ProviderName will be displayed.
	DisplayName string `yaml:"display_name"`

	BaseUrl string `yaml:"base_url"`

	// ClientID is the application's ID.
	ClientID string `yaml:"client_id"`

	// ClientSecret is the application's secret.
	ClientSecret string `yaml:"client_secret"`

	ExtraScopes []string `yaml:"extra_scopes"`

	FieldMap OauthFields `yaml:"field_map"`

	RegistrationEnabled bool `yaml:"registration_enabled"`
}

type OAuthProvider struct {
	// ProviderName is an internal name that is used to distinguish oauth endpoints. It must not contain spaces or special characters.
	ProviderName string `yaml:"provider_name"`

	// DisplayName is shown to the user on the login page. If it is empty, ProviderName will be displayed.
	DisplayName string `yaml:"display_name"`

	BaseUrl string `yaml:"base_url"`

	// ClientID is the application's ID.
	ClientID string `yaml:"client_id"`

	// ClientSecret is the application's secret.
	ClientSecret string `yaml:"client_secret"`

	AuthURL     string `yaml:"auth_url"`
	TokenURL    string `yaml:"token_url"`
	UserInfoURL string `yaml:"user_info_url"`

	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	RedirectURL string `yaml:"redirect_url"`

	// Scope specifies optional requested permissions.
	Scopes []string `yaml:"scopes"`

	// Fielmap contains
	FieldMap OauthFields `yaml:"field_map"`

	// If RegistrationEnabled is set to true, wg-portal will create new users that do not exist in the database.
	RegistrationEnabled bool `yaml:"registration_enabled"`
}

type Config struct {
	Core struct {
		GinDebug bool   `yaml:"ginDebug"`
		LogLevel string `yaml:"logLevel"`

		ListeningAddress string `yaml:"listeningAddress"`
		SessionSecret    string `yaml:"sessionSecret"`

		ExternalUrl string `yaml:"externalUrl"`
		Title       string `yaml:"title"`
		CompanyName string `yaml:"company"`

		// TODO: check...
		AdminUser     string `yaml:"adminUser"` // must be an email address
		AdminPassword string `yaml:"adminPass"`

		EditableKeys            bool   `yaml:"editableKeys"`
		CreateDefaultPeer       bool   `yaml:"createDefaultPeer"`
		SelfProvisioningAllowed bool   `yaml:"selfProvisioning"`
		LdapEnabled             bool   `yaml:"ldapEnabled"`
		LogoUrl                 string `yaml:"logoUrl"`
	} `yaml:"core"`

	Auth struct {
		OpenIDConnect []OpenIDConnectProvider `yaml:"openIdCconnect"`
		OAuth         []OAuthProvider         `yaml:"oauth"`
		Ldap          []LdapAuthProvider      `yaml:"ldap"`
	} `yaml:"auth"`

	Mail     portal.MailConfig          `yaml:"email"`
	Database persistence.DatabaseConfig `yaml:"database"`
}

func LoadConfigFile(cfg interface{}, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}