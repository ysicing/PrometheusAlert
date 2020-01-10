package teslamaxcompute

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

// ListUserQuotas invokes the teslamaxcompute.ListUserQuotas API synchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/listuserquotas.html
func (client *Client) ListUserQuotas(request *ListUserQuotasRequest) (response *ListUserQuotasResponse, err error) {
	response = CreateListUserQuotasResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserQuotasWithChan invokes the teslamaxcompute.ListUserQuotas API asynchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/listuserquotas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUserQuotasWithChan(request *ListUserQuotasRequest) (<-chan *ListUserQuotasResponse, <-chan error) {
	responseChan := make(chan *ListUserQuotasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserQuotas(request)
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

// ListUserQuotasWithCallback invokes the teslamaxcompute.ListUserQuotas API asynchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/listuserquotas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUserQuotasWithCallback(request *ListUserQuotasRequest, callback func(response *ListUserQuotasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserQuotasResponse
		var err error
		defer close(result)
		response, err = client.ListUserQuotas(request)
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

// ListUserQuotasRequest is the request struct for api ListUserQuotas
type ListUserQuotasRequest struct {
	*requests.RpcRequest
}

// ListUserQuotasResponse is the response struct for api ListUserQuotas
type ListUserQuotasResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      int    `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListUserQuotasRequest creates a request to invoke ListUserQuotas API
func CreateListUserQuotasRequest() (request *ListUserQuotasRequest) {
	request = &ListUserQuotasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "ListUserQuotas", "teslamaxcompute", "openAPI")
	return
}

// CreateListUserQuotasResponse creates a response to parse from ListUserQuotas response
func CreateListUserQuotasResponse() (response *ListUserQuotasResponse) {
	response = &ListUserQuotasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}