package tcr

import (
	"fmt"

	"github.com/cnrancher/tcr-access-control/pkg/utils"
	"github.com/sirupsen/logrus"
	tcrapi "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcr/v20190924"
)

var (
	Client *tcrapi.Client

	initialized bool
)

// Init initializes the Clinet of TCR
func Init() error {
	if initialized && Client != nil {
		return nil
	}
	logrus.Debugf("Start init TCR API")

	var err error
	Client, err = tcrapi.NewClient(
		utils.Credential,
		utils.Config.Region,
		utils.ClientProfile,
	)
	if err != nil {
		return fmt.Errorf("tcrapi.NewClient: %w", err)
	}

	initialized = true
	logrus.Debugf("Finished init TCR API")
	return nil
}

func DescribeExternalEndpointStatus() (*tcrapi.DescribeExternalEndpointStatusResponse, error) {
	request := tcrapi.NewDescribeExternalEndpointStatusRequest()
	request.RegistryId = &utils.Config.RegistryID

	return Client.DescribeExternalEndpointStatus(request)
}

func DescribeSecurityPolicies() (*tcrapi.DescribeSecurityPoliciesResponse, error) {
	request := tcrapi.NewDescribeSecurityPoliciesRequest()
	request.RegistryId = &utils.Config.RegistryID

	return Client.DescribeSecurityPolicies(request)
}

func CreateSecurityPolicy(
	cidr, description string,
) (*tcrapi.CreateSecurityPolicyResponse, error) {
	request := tcrapi.NewCreateSecurityPolicyRequest()
	request.RegistryId = &utils.Config.RegistryID
	request.CidrBlock = &cidr
	request.Description = &description
	return Client.CreateSecurityPolicy(request)
}

func DeleteSecurityPolicy(
	index int64, cidr, version string,
) (*tcrapi.DeleteSecurityPolicyResponse, error) {
	request := tcrapi.NewDeleteSecurityPolicyRequest()
	request.RegistryId = &utils.Config.RegistryID
	request.PolicyIndex = &index
	request.PolicyVersion = &version
	request.CidrBlock = &cidr
	return Client.DeleteSecurityPolicy(request)
}
