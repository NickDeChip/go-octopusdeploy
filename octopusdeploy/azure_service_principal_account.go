package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureServicePrincipalAccount represents an Azure service principal account.
type AzureServicePrincipalAccount struct {
	ApplicationID           *uuid.UUID      `validate:"required"`
	ApplicationPassword     *SensitiveValue `validate:"required"`
	AuthenticationEndpoint  string          `validate:"required_with=AzureEnvironment,omitempty,uri"`
	AzureEnvironment        string          `validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	ResourceManagerEndpoint string          `validate:"required_with=AzureEnvironment,omitempty,uri"`
	SubscriptionID          *uuid.UUID      `validate:"required"`
	TenantID                *uuid.UUID      `validate:"required"`

	accounts.account
}

// NewAzureServicePrincipalAccount creates and initializes an Azure service
// principal account.
func NewAzureServicePrincipalAccount(name string, subscriptionID uuid.UUID, tenantID uuid.UUID, applicationID uuid.UUID, applicationPassword *SensitiveValue, options ...func(*AzureServicePrincipalAccount)) (*AzureServicePrincipalAccount, error) {
	if IsEmpty(name) {
		return nil, CreateRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	if applicationPassword == nil {
		return nil, CreateRequiredParameterIsEmptyOrNilError(ParameterApplicationPassword)
	}

	account := AzureServicePrincipalAccount{
		account: *accounts.newAccount(name, accounts.AccountType("AzureServicePrincipal")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.AccountType = accounts.AccountType("AzureServicePrincipal")
	account.ApplicationID = &applicationID
	account.ApplicationPassword = applicationPassword
	account.ID = services.emptyString
	account.ModifiedBy = services.emptyString
	account.ModifiedOn = nil
	account.Name = name
	account.SubscriptionID = &subscriptionID
	account.TenantID = &tenantID

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this Azure service principal account and
// returns an error if invalid.
func (a *AzureServicePrincipalAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
