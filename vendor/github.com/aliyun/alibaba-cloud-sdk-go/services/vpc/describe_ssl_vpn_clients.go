package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSslVpnClients invokes the vpc.DescribeSslVpnClients API synchronously
func (client *Client) DescribeSslVpnClients(request *DescribeSslVpnClientsRequest) (response *DescribeSslVpnClientsResponse, err error) {
	response = CreateDescribeSslVpnClientsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSslVpnClientsWithChan invokes the vpc.DescribeSslVpnClients API asynchronously
func (client *Client) DescribeSslVpnClientsWithChan(request *DescribeSslVpnClientsRequest) (<-chan *DescribeSslVpnClientsResponse, <-chan error) {
	responseChan := make(chan *DescribeSslVpnClientsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSslVpnClients(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeSslVpnClientsWithCallback invokes the vpc.DescribeSslVpnClients API asynchronously
func (client *Client) DescribeSslVpnClientsWithCallback(request *DescribeSslVpnClientsRequest, callback func(response *DescribeSslVpnClientsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSslVpnClientsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSslVpnClients(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeSslVpnClientsRequest is the request struct for api DescribeSslVpnClients
type DescribeSslVpnClientsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string           `position:"Query" name:"VpnGatewayId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSslVpnClientsResponse is the response struct for api DescribeSslVpnClients
type DescribeSslVpnClientsResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	VpnGatewayId   string       `json:"VpnGatewayId" xml:"VpnGatewayId"`
	PageNumber     int          `json:"PageNumber" xml:"PageNumber"`
	PageSize       int          `json:"PageSize" xml:"PageSize"`
	TotalCount     int          `json:"TotalCount" xml:"TotalCount"`
	RegionId       string       `json:"RegionId" xml:"RegionId"`
	ClientInfoList []ClientInfo `json:"ClientInfoList" xml:"ClientInfoList"`
}

// CreateDescribeSslVpnClientsRequest creates a request to invoke DescribeSslVpnClients API
func CreateDescribeSslVpnClientsRequest() (request *DescribeSslVpnClientsRequest) {
	request = &DescribeSslVpnClientsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeSslVpnClients", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSslVpnClientsResponse creates a response to parse from DescribeSslVpnClients response
func CreateDescribeSslVpnClientsResponse() (response *DescribeSslVpnClientsResponse) {
	response = &DescribeSslVpnClientsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
