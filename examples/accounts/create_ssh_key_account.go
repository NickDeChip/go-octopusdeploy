package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateSSHKeyAccountExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// account values
		name           string                        = "SSH Key Account"
		privateKeyFile *octopusdeploy.SensitiveValue = octopusdeploy.NewSensitiveValue("private-key")
		username       string                        = "account-username"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	sshKeyAccount, err := accounts.NewSSHKeyAccount(name, username, privateKeyFile)
	if err != nil {
		_ = fmt.Errorf("error creating SSH key account: %v", err)
		return
	}

	// create account
	createdAccount, err := client.Accounts.Add(sshKeyAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access SSH key-specific fields
	sshKeyAccount = createdAccount.(*accounts.SSHKeyAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", sshKeyAccount.GetID())
}
