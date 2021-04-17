package pb

import (
	"context"

	. "github.com/golang/protobuf/ptypes/empty"
	"go.saastack.io/protoc-gen-caw/convert"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
)

var (
	_ = Empty{}
	_ = convert.JsonB{}
)

const (
	CustomersCreateCustomerActivity   = "/appointy.customer.v1.v1.Customers/CreateCustomer"
	CustomersGetCustomerActivity      = "/appointy.customer.v1.v1.Customers/GetCustomer"
	CustomersDeleteCustomerActivity   = "/appointy.customer.v1.v1.Customers/DeleteCustomer"
	CustomersUpdateCustomerActivity   = "/appointy.customer.v1.v1.Customers/UpdateCustomer"
	CustomersListCustomerActivity     = "/appointy.customer.v1.v1.Customers/ListCustomer"
	CustomersBatchGetCustomerActivity = "/appointy.customer.v1.v1.Customers/BatchGetCustomer"
)

func RegisterCustomersActivities(cli CustomersClient) {
	activity.RegisterWithOptions(
		func(ctx context.Context, in *CreateCustomerRequest) (*Customer, error) {
			res, err := cli.CreateCustomer(ctx, in)
			return res, err
		},
		activity.RegisterOptions{Name: CustomersCreateCustomerActivity},
	)
	activity.RegisterWithOptions(
		func(ctx context.Context, in *GetCustomerRequest) (*Customer, error) {
			res, err := cli.GetCustomer(ctx, in)
			return res, err
		},
		activity.RegisterOptions{Name: CustomersGetCustomerActivity},
	)
	activity.RegisterWithOptions(
		func(ctx context.Context, in *DeleteCustomerRequest) (*Empty, error) {
			res, err := cli.DeleteCustomer(ctx, in)
			return res, err
		},
		activity.RegisterOptions{Name: CustomersDeleteCustomerActivity},
	)
	activity.RegisterWithOptions(
		func(ctx context.Context, in *UpdateCustomerRequest) (*Customer, error) {
			res, err := cli.UpdateCustomer(ctx, in)
			return res, err
		},
		activity.RegisterOptions{Name: CustomersUpdateCustomerActivity},
	)
	activity.RegisterWithOptions(
		func(ctx context.Context, in *ListCustomerRequest) (*ListCustomerResponse, error) {
			res, err := cli.ListCustomer(ctx, in)
			return res, err
		},
		activity.RegisterOptions{Name: CustomersListCustomerActivity},
	)
	activity.RegisterWithOptions(
		func(ctx context.Context, in *BatchGetCustomerRequest) (*BatchGetCustomerResponse, error) {
			res, err := cli.BatchGetCustomer(ctx, in)
			return res, err
		},
		activity.RegisterOptions{Name: CustomersBatchGetCustomerActivity},
	)
}

// CustomersActivitiesClient is a typesafe wrapper for CustomersActivities.
type CustomersActivitiesClient struct {
}

// NewCustomersActivitiesClient creates a new CustomersActivitiesClient.
func NewCustomersActivitiesClient(cli CustomersClient) CustomersActivitiesClient {
	RegisterCustomersActivities(cli)
	return CustomersActivitiesClient{}
}

func (ca *CustomersActivitiesClient) CreateCustomer(ctx workflow.Context, in *CreateCustomerRequest) (*Customer, error) {
	future := workflow.ExecuteActivity(ctx, CustomersCreateCustomerActivity, in)
	var result Customer
	if err := future.Get(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ca *CustomersActivitiesClient) GetCustomer(ctx workflow.Context, in *GetCustomerRequest) (*Customer, error) {
	future := workflow.ExecuteActivity(ctx, CustomersGetCustomerActivity, in)
	var result Customer
	if err := future.Get(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ca *CustomersActivitiesClient) DeleteCustomer(ctx workflow.Context, in *DeleteCustomerRequest) (*Empty, error) {
	future := workflow.ExecuteActivity(ctx, CustomersDeleteCustomerActivity, in)
	var result Empty
	if err := future.Get(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ca *CustomersActivitiesClient) UpdateCustomer(ctx workflow.Context, in *UpdateCustomerRequest) (*Customer, error) {
	future := workflow.ExecuteActivity(ctx, CustomersUpdateCustomerActivity, in)
	var result Customer
	if err := future.Get(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ca *CustomersActivitiesClient) ListCustomer(ctx workflow.Context, in *ListCustomerRequest) (*ListCustomerResponse, error) {
	future := workflow.ExecuteActivity(ctx, CustomersListCustomerActivity, in)
	var result ListCustomerResponse
	if err := future.Get(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (ca *CustomersActivitiesClient) BatchGetCustomer(ctx workflow.Context, in *BatchGetCustomerRequest) (*BatchGetCustomerResponse, error) {
	future := workflow.ExecuteActivity(ctx, CustomersBatchGetCustomerActivity, in)
	var result BatchGetCustomerResponse
	if err := future.Get(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
