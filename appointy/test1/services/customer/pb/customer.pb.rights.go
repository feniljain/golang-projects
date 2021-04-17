// Code generated by protoc-gen-defaults. DO NOT EDIT.

package pb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	rights "go.saastack.io/deployment/right"
	modulePB "go.saastack.io/modulerole/pb"
	rightspb "go.saastack.io/right/pb"
	"go.saastack.io/userinfo"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	CUSTOMERS_CREATE_CUSTOMER_PARENT = "/Customers/{parent}**/.CreateCustomer"

	CUSTOMERS_GET_CUSTOMER_ID = "/Customers/{id}**/.GetCustomer"

	CUSTOMERS_DELETE_CUSTOMER_ID = "/Customers/{id}**/.DeleteCustomer"

	CUSTOMERS_UPDATE_CUSTOMER_CUSTOMER_ID = "/Customers/{customer.id}**/.UpdateCustomer"

	CUSTOMERS_LIST_CUSTOMER_PARENT = "/Customers/{parent}**/.ListCustomer"

	CUSTOMERS_BATCH_GET_CUSTOMER_IDS = "/Customers/{ids}**/.BatchGetCustomer"
)

var CustomersModuleName = "Customers"
var CustomersModulePattern = "/Customers/{parent}**/.*"

var CustomersResourcePaths = []*rightspb.RightsResource{

	{
		Name:        `CreateCustomer`,
		Description: ` CreateCustomer creates new customer.\n`,
		Resource:    `/Customers/{parent}**/.CreateCustomer`,
	},

	{
		Name:        `GetCustomer`,
		Description: ` GetCustomer returns the customer by its unique id.\n`,
		Resource:    `/Customers/{id}**/.GetCustomer`,
	},

	{
		Name:        `DeleteCustomer`,
		Description: ` DeleteCustomer will delete the customer from the system by Id.\n This will be a soft delete from the system\n`,
		Resource:    `/Customers/{id}**/.DeleteCustomer`,
	},

	{
		Name:        `UpdateCustomer`,
		Description: ` UpdateCustomer will update the customer identified by its customer id.\n Update Customer uses Field Mask to update specific properties of customer object\n`,
		Resource:    `/Customers/{customer.id}**/.UpdateCustomer`,
	},

	{
		Name:        `ListCustomer`,
		Description: ` ListCustomer lists all the Customer(s)\n`,
		Resource:    `/Customers/{parent}**/.ListCustomer`,
	},

	{
		Name:        `BatchGetCustomer`,
		Description: ` Gets all the Customer(s) by their ids\n`,
		Resource:    `/Customers/{ids}**/.BatchGetCustomer`,
	},
}

type rightsCustomersServer struct {
	CustomersSrv     CustomersServer
	moduleRoleServer modulePB.ModuleRoleServiceServer
	rightsCli        rightspb.RightValidatorsClient
	db               *sql.DB
}

func NewRightsCustomersServer(db *sql.DB,
	mrs modulePB.ModuleRoleServiceServer,
	c rightspb.RightValidatorsClient,

	config rights.ModuleRoleConfig,
	in struct {
		fx.In
		S CustomersServer `name:"public"`
	},
) struct {
	fx.Out
	S CustomersServer `name:"public"`
} {
	srv := &rightsCustomersServer{
		db:               db,
		rightsCli:        c,
		CustomersSrv:     in.S,
		moduleRoleServer: mrs,
	}

	if err := srv.RegisterModuleRoles(); err != nil {
		panic(err)
	}

	return struct {
		fx.Out
		S CustomersServer `name:"public"`
	}{
		S: srv,
	}
}

func (s *rightsCustomersServer) CreateCustomer(ctx context.Context, rightsvar *CreateCustomerRequest) (*Customer, error) {

	ResourcePathOR := make([]string, 0)
	ResourcePathAND := make([]string, 0)

	ResourcePathOR = append(ResourcePathOR,

		fmt.Sprintf("/Customers/%s**/.CreateCustomer",

			rightsvar.GetParent(),
		),
	)

	validations := map[string]bool{}

	res, err := s.rightsCli.IsValid(ctx, &rightspb.IsValidRequest{
		ResourcePathOr:       ResourcePathOR,
		ResourcePathAnd:      ResourcePathAND,
		UserId:               userinfo.FromContext(ctx).Id,
		ModuleName:           "Customers",
		AttributeValidations: validations,
		AllowParent:          false,
		AllowStaff:           false,
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.CustomersSrv.CreateCustomer(ctx, rightsvar)
}

func (s *rightsCustomersServer) GetCustomer(ctx context.Context, rightsvar *GetCustomerRequest) (*Customer, error) {

	ResourcePathOR := make([]string, 0)
	ResourcePathAND := make([]string, 0)

	ResourcePathOR = append(ResourcePathOR,

		fmt.Sprintf("/Customers/%s**/.GetCustomer",

			rightsvar.GetId(),
		),
	)

	validations := map[string]bool{}

	res, err := s.rightsCli.IsValid(ctx, &rightspb.IsValidRequest{
		ResourcePathOr:       ResourcePathOR,
		ResourcePathAnd:      ResourcePathAND,
		UserId:               userinfo.FromContext(ctx).Id,
		ModuleName:           "Customers",
		AttributeValidations: validations,
		AllowParent:          false,
		AllowStaff:           false,
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.CustomersSrv.GetCustomer(ctx, rightsvar)
}

func (s *rightsCustomersServer) DeleteCustomer(ctx context.Context, rightsvar *DeleteCustomerRequest) (*empty.Empty, error) {

	ResourcePathOR := make([]string, 0)
	ResourcePathAND := make([]string, 0)

	ResourcePathOR = append(ResourcePathOR,

		fmt.Sprintf("/Customers/%s**/.DeleteCustomer",

			rightsvar.GetId(),
		),
	)

	validations := map[string]bool{}

	res, err := s.rightsCli.IsValid(ctx, &rightspb.IsValidRequest{
		ResourcePathOr:       ResourcePathOR,
		ResourcePathAnd:      ResourcePathAND,
		UserId:               userinfo.FromContext(ctx).Id,
		ModuleName:           "Customers",
		AttributeValidations: validations,
		AllowParent:          false,
		AllowStaff:           false,
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.CustomersSrv.DeleteCustomer(ctx, rightsvar)
}

func (s *rightsCustomersServer) UpdateCustomer(ctx context.Context, rightsvar *UpdateCustomerRequest) (*Customer, error) {

	ResourcePathOR := make([]string, 0)
	ResourcePathAND := make([]string, 0)

	ResourcePathOR = append(ResourcePathOR,

		fmt.Sprintf("/Customers/%s**/.UpdateCustomer",

			rightsvar.GetCustomer().GetId(),
		),
	)

	validations := map[string]bool{}

	res, err := s.rightsCli.IsValid(ctx, &rightspb.IsValidRequest{
		ResourcePathOr:       ResourcePathOR,
		ResourcePathAnd:      ResourcePathAND,
		UserId:               userinfo.FromContext(ctx).Id,
		ModuleName:           "Customers",
		AttributeValidations: validations,
		AllowParent:          false,
		AllowStaff:           false,
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.CustomersSrv.UpdateCustomer(ctx, rightsvar)
}

func (s *rightsCustomersServer) ListCustomer(ctx context.Context, rightsvar *ListCustomerRequest) (*ListCustomerResponse, error) {

	ResourcePathOR := make([]string, 0)
	ResourcePathAND := make([]string, 0)

	ResourcePathOR = append(ResourcePathOR,

		fmt.Sprintf("/Customers/%s**/.ListCustomer",

			rightsvar.GetParent(),
		),
	)

	validations := map[string]bool{}

	res, err := s.rightsCli.IsValid(ctx, &rightspb.IsValidRequest{
		ResourcePathOr:       ResourcePathOR,
		ResourcePathAnd:      ResourcePathAND,
		UserId:               userinfo.FromContext(ctx).Id,
		ModuleName:           "Customers",
		AttributeValidations: validations,
		AllowParent:          false,
		AllowStaff:           false,
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.CustomersSrv.ListCustomer(ctx, rightsvar)
}

func (s *rightsCustomersServer) BatchGetCustomer(ctx context.Context, rightsvar *BatchGetCustomerRequest) (*BatchGetCustomerResponse, error) {

	ResourcePathOR := make([]string, 0)
	ResourcePathAND := make([]string, 0)

	for _, Ids := range rightsvar.GetIds() {

		ResourcePathAND = append(ResourcePathAND,

			fmt.Sprintf("/Customers/%s**/.BatchGetCustomer",

				Ids,
			),
		)

	}

	validations := map[string]bool{}

	res, err := s.rightsCli.IsValid(ctx, &rightspb.IsValidRequest{
		ResourcePathOr:       ResourcePathOR,
		ResourcePathAnd:      ResourcePathAND,
		UserId:               userinfo.FromContext(ctx).Id,
		ModuleName:           "Customers",
		AttributeValidations: validations,
		AllowParent:          false,
		AllowStaff:           false,
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.CustomersSrv.BatchGetCustomer(ctx, rightsvar)
}

// function for constructor
func (s *rightsCustomersServer) RegisterModuleRoles() error {

	if _, err := s.moduleRoleServer.RegisterModuleRoleInMemory(context.Background(), &modulePB.ModuleRoleList{
		List: []*modulePB.ModuleRole{},
	}); err != nil {
		return err
	}

	return nil
}
