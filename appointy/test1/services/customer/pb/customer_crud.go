package pb

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"go.saastack.io/chaku/errors"
	"go.saastack.io/idutil"
	"go.saastack.io/protos/types"
	"go.saastack.io/userinfo"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomersServiceCustomerServerCrud struct {
	store CustomerStore
	bloc  CustomersServiceCustomerServerBLoC
}

type CustomersServiceCustomerServerBLoC interface {
	CreateCustomerBLoC(context.Context, *CreateCustomerRequest) error

	GetCustomerBLoC(context.Context, *GetCustomerRequest) error

	UpdateCustomerBLoC(context.Context, *UpdateCustomerRequest) error

	DeleteCustomerBLoC(context.Context, *DeleteCustomerRequest) error

	BatchGetCustomerBLoC(context.Context, *BatchGetCustomerRequest) error

	ListCustomerBLoC(context.Context, *ListCustomerRequest) (CustomerCondition, error)
}

func NewCustomersServiceCustomerServerCrud(s CustomerStore, b CustomersServiceCustomerServerBLoC) *CustomersServiceCustomerServerCrud {
	return &CustomersServiceCustomerServerCrud{store: s, bloc: b}
}

func (s *CustomersServiceCustomerServerCrud) CreateCustomer(ctx context.Context, in *CreateCustomerRequest) (*Customer, error) {

	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.bloc.CreateCustomerBLoC(ctx, in)
	if err != nil {
		return nil, err
	}

	if idutil.GetPrefix(in.Customer.Id) != in.Customer.GetPrefix() {
		in.Customer.Id = in.Parent
	}

	ids, err := s.store.CreateCustomers(ctx, in.Customer)
	if err != nil {
		return nil, err
	}

	in.Customer.Id = ids[0]

	return in.GetCustomer(), nil
}

func (s *CustomersServiceCustomerServerCrud) UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest) (*Customer, error) {

	mask := s.GetViewMask(in.UpdateMask)
	if len(mask) == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot send empty update mask")
	}

	if err := in.GetCustomer().Validate(mask...); err != nil {
		return nil, err
	}

	err := s.bloc.UpdateCustomerBLoC(ctx, in)
	if err != nil {
		return nil, err
	}

	if err := s.store.UpdateCustomer(ctx,
		in.Customer, mask,
		CustomerIdEq{Id: in.Customer.Id},
	); err != nil {
		return nil, err
	}

	updatedCustomer, err := s.store.GetCustomer(ctx, []string{},
		CustomerIdEq{
			Id: in.GetCustomer().GetId(),
		},
	)
	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil
}

func (s *CustomersServiceCustomerServerCrud) GetCustomer(ctx context.Context, in *GetCustomerRequest) (*Customer, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.bloc.GetCustomerBLoC(ctx, in)
	if err != nil {
		return nil, err
	}

	mask := s.GetViewMask(in.ViewMask)

	res, err := s.store.GetCustomer(ctx, mask, CustomerIdEq{Id: in.Id})
	if err != nil {
		if err == errors.ErrNotFound {
			return nil, status.Error(codes.NotFound, "Customer not found")
		}
		return nil, err
	}

	return res, nil
}

func (s *CustomersServiceCustomerServerCrud) ListCustomer(ctx context.Context, in *ListCustomerRequest) (*ListCustomerResponse, error) {

	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	condition, err := s.bloc.ListCustomerBLoC(ctx, in)
	if err != nil {
		return nil, err
	}

	page, err := s.GetPagedCondition(ctx, in.First, in.After, in.Last, in.Before)
	if err != nil {
		return nil, err
	}

	mask := s.GetViewMask(in.ViewMask)

	return s.ListWithPagination(ctx, page, condition, mask)
}

func (s *CustomersServiceCustomerServerCrud) ListWithPagination(ctx context.Context, page *CursorBasedPagination, condition CustomerCondition, viewMask []string) (*ListCustomerResponse, error) {

	list, err := s.store.ListCustomers(ctx,
		viewMask,
		condition,
		page,
	)
	if err != nil {
		return nil, err
	}

	res := &ListCustomerResponse{
		PageInfo: &types.PageInfo{},
	}

	for _, it := range list {

		res.Nodes = append(res.Nodes, &CustomerNode{Position: it.Id, Node: it})
	}
	res.PageInfo.HasPrevious = page.HasPrevious
	res.PageInfo.HasNext = page.HasNext
	if len(list) > 0 {
		res.PageInfo.StartCursor = list[0].Id
		res.PageInfo.EndCursor = list[len(list)-1].Id
	}

	return res, nil

}

func (s *CustomersServiceCustomerServerCrud) GetPagedCondition(ctx context.Context, first uint32, after string, last uint32, before string) (*CursorBasedPagination, error) {

	page := &CursorBasedPagination{}
	flag := false

	if first != 0 {
		flag = true
		page = &CursorBasedPagination{
			Cursor:   idutil.GetId(after),
			Limit:    int(first),
			UpOrDown: false,
		}

	} else if last != 0 {
		flag = true
		page = &CursorBasedPagination{
			Cursor:   idutil.GetId(before),
			Limit:    int(last),
			UpOrDown: true,
		}
	}
	if !flag {
		return nil, status.Error(codes.InvalidArgument, "either after-first or before-last should be set in request")
	}

	return page, nil
}

func (s *CustomersServiceCustomerServerCrud) DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest) (*empty.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.bloc.DeleteCustomerBLoC(ctx, in)
	if err != nil {
		return nil, err
	}

	if err := s.store.DeleteCustomer(ctx, CustomerIdEq{Id: in.Id}); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *CustomersServiceCustomerServerCrud) BatchGetCustomer(ctx context.Context, in *BatchGetCustomerRequest) (*BatchGetCustomerResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.bloc.BatchGetCustomerBLoC(ctx, in)
	if err != nil {
		return nil, err
	}

	getIds := make([]string, 0, len(in.Ids))
	for _, id := range in.Ids {
		getIds = append(getIds, idutil.GetId(id))
	}

	mask := s.GetViewMask(in.ViewMask)

	list, err := s.store.ListCustomers(ctx, mask, CustomerIdIn{Id: getIds})
	if err != nil {
		return nil, err
	}

	resultMap := make(map[string]*Customer, 0)
	for i, it := range list {
		_ = i

		resultMap[it.Id] = it
	}

	isGrpc := userinfo.IsGrpcCall(ctx)

	result := make([]*Customer, 0, len(in.Ids))
	for _, id := range in.Ids {
		if resultMap[id] == nil && isGrpc {
			result = append(result, &Customer{})
			continue
		}
		result = append(result, resultMap[id])
	}

	return &BatchGetCustomerResponse{Customer: result}, nil
}

func (s *CustomersServiceCustomerServerCrud) GetViewMask(mask *field_mask.FieldMask) []string {
	if mask == nil || mask.GetPaths() == nil {
		return []string{}
	}
	return mask.GetPaths()
}
