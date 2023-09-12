package tcr

import (
	"fmt"

	"github.com/cnrancher/tcr-access-control/pkg/utils"
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

	var err error
	Client, err = tcrapi.NewClient(
		utils.Credential,
		utils.Region,
		utils.ClientProfile,
	)
	if err != nil {
		return fmt.Errorf("tcrapi.NewClient: %w", err)
	}

	initialized = true
	return nil
}

// DescribeExternalEndpointStatus
// 查询实例公网访问入口状态
func DescribeExternalEndpointStatus() (*tcrapi.DescribeExternalEndpointStatusResponse, error) {
	request := tcrapi.NewDescribeExternalEndpointStatusRequest()
	request.RegistryId = &utils.Config.InstanceID

	response, err := Client.DescribeExternalEndpointStatus(request)
	return response, err
}

func DescribeSecurityPolicies() (*tcrapi.DescribeSecurityPoliciesResponse, error) {
	request := tcrapi.NewDescribeSecurityPoliciesRequest()
	request.RegistryId = &utils.Config.InstanceID

	response, err := Client.DescribeSecurityPolicies(request)
	return response, err
}
