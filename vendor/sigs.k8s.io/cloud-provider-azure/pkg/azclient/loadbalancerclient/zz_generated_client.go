// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by client-gen. DO NOT EDIT.
package loadbalancerclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/utils"
)

type Client struct {
	*armnetwork.LoadBalancersClient
	subscriptionID string
	tracer         tracing.Tracer
}

func New(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (Interface, error) {
	if options == nil {
		options = utils.GetDefaultOption()
	}
	tr := options.TracingProvider.NewTracer(utils.ModuleName, utils.ModuleVersion)

	client, err := armnetwork.NewLoadBalancersClient(subscriptionID, credential, options)
	if err != nil {
		return nil, err
	}
	return &Client{
		LoadBalancersClient: client,
		subscriptionID:      subscriptionID,
		tracer:              tr,
	}, nil
}

const GetOperationName = "LoadBalancersClient.Get"

// Get gets the LoadBalancer
func (client *Client) Get(ctx context.Context, resourceGroupName string, resourceName string, expand *string) (result *armnetwork.LoadBalancer, rerr error) {
	var ops *armnetwork.LoadBalancersClientGetOptions
	if expand != nil {
		ops = &armnetwork.LoadBalancersClientGetOptions{Expand: expand}
	}
	ctx = utils.ContextWithClientName(ctx, "LoadBalancersClient")
	ctx = utils.ContextWithRequestMethod(ctx, "Get")
	ctx = utils.ContextWithResourceGroupName(ctx, resourceGroupName)
	ctx = utils.ContextWithSubscriptionID(ctx, client.subscriptionID)
	ctx, endSpan := runtime.StartSpan(ctx, GetOperationName, client.tracer, nil)
	defer endSpan(rerr)
	resp, err := client.LoadBalancersClient.Get(ctx, resourceGroupName, resourceName, ops)
	if err != nil {
		return nil, err
	}
	//handle statuscode
	return &resp.LoadBalancer, nil
}

const CreateOrUpdateOperationName = "LoadBalancersClient.Create"

// CreateOrUpdate creates or updates a LoadBalancer.
func (client *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, resource armnetwork.LoadBalancer) (result *armnetwork.LoadBalancer, err error) {
	ctx = utils.ContextWithClientName(ctx, "LoadBalancersClient")
	ctx = utils.ContextWithRequestMethod(ctx, "CreateOrUpdate")
	ctx = utils.ContextWithResourceGroupName(ctx, resourceGroupName)
	ctx = utils.ContextWithSubscriptionID(ctx, client.subscriptionID)
	ctx, endSpan := runtime.StartSpan(ctx, CreateOrUpdateOperationName, client.tracer, nil)
	defer endSpan(err)
	resp, err := utils.NewPollerWrapper(client.LoadBalancersClient.BeginCreateOrUpdate(ctx, resourceGroupName, resourceName, resource, nil)).WaitforPollerResp(ctx)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return &resp.LoadBalancer, nil
	}
	return nil, nil
}

const DeleteOperationName = "LoadBalancersClient.Delete"

// Delete deletes a LoadBalancer by name.
func (client *Client) Delete(ctx context.Context, resourceGroupName string, resourceName string) (err error) {
	ctx = utils.ContextWithClientName(ctx, "LoadBalancersClient")
	ctx = utils.ContextWithRequestMethod(ctx, "Delete")
	ctx = utils.ContextWithResourceGroupName(ctx, resourceGroupName)
	ctx = utils.ContextWithSubscriptionID(ctx, client.subscriptionID)
	ctx, endSpan := runtime.StartSpan(ctx, DeleteOperationName, client.tracer, nil)
	defer endSpan(err)
	_, err = utils.NewPollerWrapper(client.BeginDelete(ctx, resourceGroupName, resourceName, nil)).WaitforPollerResp(ctx)
	return err
}

const ListOperationName = "LoadBalancersClient.List"

// List gets a list of LoadBalancer in the resource group.
func (client *Client) List(ctx context.Context, resourceGroupName string) (result []*armnetwork.LoadBalancer, rerr error) {
	ctx = utils.ContextWithClientName(ctx, "LoadBalancersClient")
	ctx = utils.ContextWithRequestMethod(ctx, "List")
	ctx = utils.ContextWithResourceGroupName(ctx, resourceGroupName)
	ctx = utils.ContextWithSubscriptionID(ctx, client.subscriptionID)
	ctx, endSpan := runtime.StartSpan(ctx, ListOperationName, client.tracer, nil)
	defer endSpan(rerr)
	pager := client.LoadBalancersClient.NewListPager(resourceGroupName, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, nextResult.Value...)
	}
	return result, nil
}
