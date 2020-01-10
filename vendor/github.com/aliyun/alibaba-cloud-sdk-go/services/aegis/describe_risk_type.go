package aegis

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

// DescribeRiskType invokes the aegis.DescribeRiskType API synchronously
// api document: https://help.aliyun.com/api/aegis/describerisktype.html
func (client *Client) DescribeRiskType(request *DescribeRiskTypeRequest) (response *DescribeRiskTypeResponse, err error) {
	response = CreateDescribeRiskTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRiskTypeWithChan invokes the aegis.DescribeRiskType API asynchronously
// api document: https://help.aliyun.com/api/aegis/describerisktype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRiskTypeWithChan(request *DescribeRiskTypeRequest) (<-chan *DescribeRiskTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeRiskTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRiskType(request)
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

// DescribeRiskTypeWithCallback invokes the aegis.DescribeRiskType API asynchronously
// api document: https://help.aliyun.com/api/aegis/describerisktype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRiskTypeWithCallback(request *DescribeRiskTypeRequest, callback func(response *DescribeRiskTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRiskTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeRiskType(request)
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

// DescribeRiskTypeRequest is the request struct for api DescribeRiskType
type DescribeRiskTypeRequest struct {
	*requests.RpcRequest
	SourceIp   string           `position:"Query" name:"SourceIp"`
	StrategyId requests.Integer `position:"Query" name:"StrategyId"`
	Lang       string           `position:"Query" name:"Lang"`
	Uuids      string           `position:"Query" name:"Uuids"`
}

// DescribeRiskTypeResponse is the response struct for api DescribeRiskType
type DescribeRiskTypeResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Count     int        `json:"Count" xml:"Count"`
	RiskTypes []RiskType `json:"RiskTypes" xml:"RiskTypes"`
}

// CreateDescribeRiskTypeRequest creates a request to invoke DescribeRiskType API
func CreateDescribeRiskTypeRequest() (request *DescribeRiskTypeRequest) {
	request = &DescribeRiskTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeRiskType", "vipaegis", "openAPI")
	return
}

// CreateDescribeRiskTypeResponse creates a response to parse from DescribeRiskType response
func CreateDescribeRiskTypeResponse() (response *DescribeRiskTypeResponse) {
	response = &DescribeRiskTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}