package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/service"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func AddEnvironmentConditionToStepExample() {
	var (
		// Declare working variables
		octopusURL       string   = "https://your_octopus_url"
		apiKey           string   = "API-YOUR_API_KEY"
		spaceID          string   = "space-id"
		projectName      string   = "project-name"
		environmentNames []string = []string{"Development", "Test"}
		stepName         string   = "Run a script"
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

	environmentIDs := []string{}

	for _, environmentName := range environmentNames {
		environments, err := client.Environments.GetByName(environmentName)
		if err != nil {
			// TODO: handle error
		}

		environmentIDs = append(environmentIDs, environments[0].GetID())
	}

	projects, err := client.Projects.Get(service.ProjectsQuery{
		Name: projectName,
	})
	if err != nil {
		// TODO: handle error
	}

	// sub-optimal; iterate through collection
	project := projects.Items[0]

	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)
	if err != nil {
		// TODO: handle error
	}

	for _, step := range deploymentProcess.Steps {
		if step.Name == stepName {
			for _, action := range step.Actions {
				for _, environmentID := range environmentIDs {
					action.Environments = append(action.Environments, environmentID)
				}
			}
		}
	}

	_, err = client.DeploymentProcesses.Update(deploymentProcess)
	if err != nil {
		// TODO: handle error
	}
}
