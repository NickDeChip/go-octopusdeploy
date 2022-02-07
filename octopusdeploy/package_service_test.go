package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createPackageService(t *testing.T) *packageService {
	service := newPackageService(nil, TestURIPackages, TestURIPackageDeltaSignature, TestURIPackageDeltaUpload, TestURIPackageNotesList, TestURIPackagesBulk, TestURIPackageUpload)
	services.testNewService(t, service, TestURIPackages, ServicePackageService)
	return service
}

func TestPackageServiceGetByID(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(services.emptyString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(services.whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)
}

func TestPackageServiceNew(t *testing.T) {
	ServiceFunction := newPackageService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	ServiceName := ServicePackageService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string, string, string, string, string) *packageService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, TestURIPackageDeltaSignature, TestURIPackageDeltaUpload, TestURIPackageNotesList, TestURIPackagesBulk, TestURIPackageUpload)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestPackageServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", services.emptyString},
		{"Whitespace", services.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createPackageService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			require.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
			require.Nil(t, resource)

			err = service.DeleteByID(tc.parameter)
			require.Error(t, err)
			require.Equal(t, err, createInvalidParameterError(OperationDeleteByID, ParameterID))
		})
	}
}

func TestPackageServiceUpdateWithEmptyPackage(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	updatedPackage, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, updatedPackage)

	updatedPackage, err = service.Update(&Package{})
	require.Error(t, err)
	require.Nil(t, updatedPackage)
}
