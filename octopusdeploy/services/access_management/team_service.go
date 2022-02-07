package access_management

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/access_management"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

const teamsV1BasePath = "teams"

type teamService struct {
	client *octopusdeploy.AdminClient
	services.AdminService
	services.GetsByIDer[access_management.Team]
	services.ResourceQueryer[access_management.Teams]
	services.CanAddService[access_management.Team]
	services.CanUpdateService[access_management.Team]
	services.CanDeleteService[access_management.Team]
}

func NewTeamService(client *octopusdeploy.AdminClient) *teamService {
	teamService := &teamService{
		AdminService: services.NewAdminService(octopusdeploy.ServiceTeamService, client),
	}

	return teamService
}

func (s teamService) getPagedResponse(path string) ([]*access_management.Team, error) {
	resources := []*access_management.Team{}
	loadNextPage := true

	for loadNextPage {
		resp, err := octopusdeploy.ApiGetPaged(s.client, new(access_management.Teams))
		if err != nil {
			return resources, err
		}

		responseList := resp.(*access_management.Teams)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = octopusdeploy.LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Query returns a collection of teams based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s teamService) Query(teamsQuery octopusdeploy.TeamsQuery) (*access_management.Teams, error) {
	template, err := uritemplates.Parse(fmt.Sprintf("%s{?skip,take,ids,partialName,spaces,includeSystem}", s.BasePath))
	path, err := s.getURITemplate().Expand(teamsQuery)
	if err != nil {
		return &access_management.Teams{}, err
	}

	response, err := s.client.apiQuery(new(access_management.Teams), template)
	if err != nil {
		return &access_management.Teams{}, err
	}

	return response.(*access_management.Teams), nil
}

// GetByPartialName performs a lookup and returns teams with a matching partial
// name.
func (s teamService) GetByPartialName(name string) ([]*access_management.Team, error) {
	path, err := s.getByPartialNamePath(name)
	if err != nil {
		return []*access_management.Team{}, err
	}

	return s.getPagedResponse(path)
}

func (s teamService) GetScopedUserRoles(team access_management.Team, query octopusdeploy.SkipTakeQuery) (*octopusdeploy.ScopedUserRoles, error) {
	template, _ := uritemplates.Parse(team.Links["ScopedUserRoles"])
	path, _ := template.Expand(query)

	resp, err := s.client.apiGet(new(octopusdeploy.ScopedUserRoles), path)
	if err != nil {
		return nil, err
	}

	return resp.(*octopusdeploy.ScopedUserRoles), nil
}
