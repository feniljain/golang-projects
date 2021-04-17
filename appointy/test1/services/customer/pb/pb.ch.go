package pb

import (
	context "context"
	x "database/sql"

	sqrl "github.com/elgris/sqrl"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	chaku_globals "go.saastack.io/chaku/chaku-globals"
	driver "go.saastack.io/chaku/driver"
	sql "go.saastack.io/chaku/driver/pgsql"
	errors "go.saastack.io/chaku/errors"
)

var objectTableMap = chaku_globals.ObjectTable{
	"customer": {
		"id": "customer",
	},
}

func (m *Customer) PackageName() string {
	return "appointy_customer_v1_v1"
}

func (m *Customer) TableOfObject(f, s string) string {
	return objectTableMap[f][s]
}

func (m *Customer) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Customer) ObjectName() string {
	return "customer"
}

func (m *Customer) Fields() []string {
	return []string{
		"id",
	}
}

func (m *Customer) IsObject(field string) bool {
	switch field {
	default:
		return false
	}
}

func (m *Customer) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *Customer) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Customer) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Customer) New(field string) error {
	switch field {
	case "id":
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *Customer) Type(field string) string {
	switch field {
	case "id":
		return "string"
	default:
		return ""
	}
}

func (_ *Customer) GetEmptyObject() (m *Customer) {
	m = &Customer{}
	return
}

func (m *Customer) GetPrefix() string {
	return "cus"
}

func (m *Customer) GetID() string {
	return m.Id
}

func (m *Customer) SetID(id string) {
	m.Id = id
}

func (m *Customer) IsRoot() bool {
	return true
}

func (m *Customer) IsFlatObject(f string) bool {
	return false
}

func (m *Customer) NoOfParents(d driver.Descriptor) int {
	switch d.ObjectName() {
	}
	return 0
}

type CustomerStore struct {
	d      driver.Driver
	withTx bool
	tx     driver.Transaction

	limitMultiplier int
}

func (s CustomerStore) Execute(ctx context.Context, query string, args ...interface{}) error {
	if s.withTx {
		return s.tx.Execute(ctx, query, args...)
	}
	return s.d.Execute(ctx, query, args...)
}

func (s CustomerStore) QueryRows(ctx context.Context, query string, scanners []string, args ...interface{}) (driver.Result, error) {
	if s.withTx {
		return s.tx.QueryRows(ctx, query, scanners, args...)
	}
	return s.d.QueryRows(ctx, query, scanners, args...)
}

func NewCustomerStore(d driver.Driver) CustomerStore {
	return CustomerStore{d: d, limitMultiplier: 1}
}

func NewPostgresCustomerStore(db *x.DB, usr driver.IUserInfo) CustomerStore {
	return CustomerStore{
		d:               &sql.Sql{DB: db, UserInfo: usr, Placeholder: sqrl.Dollar},
		limitMultiplier: 1,
	}
}

type CustomerTx struct {
	CustomerStore
}

func (s CustomerStore) BeginTx(ctx context.Context) (*CustomerTx, error) {
	tx, err := s.d.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	return &CustomerTx{
		CustomerStore: CustomerStore{
			d:      s.d,
			withTx: true,
			tx:     tx,
		},
	}, nil
}

func (tx *CustomerTx) Commit(ctx context.Context) error {
	return tx.tx.Commit(ctx)
}

func (tx *CustomerTx) RollBack(ctx context.Context) error {
	return tx.tx.RollBack(ctx)
}

func (s CustomerStore) CreateCustomerPGStore(ctx context.Context) error {
	const queries = `
CREATE SCHEMA IF NOT EXISTS appointy_customer_v1_v1;
CREATE TABLE IF NOT EXISTS  appointy_customer_v1_v1.customer( id text DEFAULT ''::text , parent text DEFAULT ''::text , is_deleted boolean DEFAULT false, deleted_by text DEFAULT ''::text, deleted_on timestamp without time zone DEFAULT '0001-01-01 00:00:00'::timestamp without time zone, updated_by text DEFAULT ''::text, updated_on timestamp without time zone DEFAULT '0001-01-01 00:00:00'::timestamp without time zone, created_by text DEFAULT ''::text, created_on timestamp without time zone DEFAULT '0001-01-01 00:00:00'::timestamp without time zone, field_variable_mask text DEFAULT ''::text, PRIMARY KEY (id, parent)); 
CREATE TABLE IF NOT EXISTS  appointy_customer_v1_v1.customer_parent( id text DEFAULT ''::text , parent text DEFAULT ''::text ); 
`
	if err := s.d.Execute(ctx, queries); err != nil {
		return err
	}
	return nil
}

func (s CustomerStore) CreateCustomers(ctx context.Context, list ...*Customer) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	if s.withTx {
		return s.tx.Insert(ctx, vv, &Customer{}, &Customer{}, "", []string{})
	}
	return s.d.Insert(ctx, vv, &Customer{}, &Customer{}, "", []string{})
}

func (s CustomerStore) DeleteCustomer(ctx context.Context, cond CustomerCondition) error {
	if s.withTx {
		return s.tx.Delete(ctx, cond.customerCondToDriverCustomerCond(s.d), &Customer{}, &Customer{})
	}
	return s.d.Delete(ctx, cond.customerCondToDriverCustomerCond(s.d), &Customer{}, &Customer{})
}

func (s CustomerStore) UpdateCustomer(ctx context.Context, req *Customer, fields []string, cond CustomerCondition) error {
	if s.withTx {
		return s.tx.Update(ctx, cond.customerCondToDriverCustomerCond(s.d), req, &Customer{}, fields...)
	}
	return s.d.Update(ctx, cond.customerCondToDriverCustomerCond(s.d), req, &Customer{}, fields...)
}

func (s CustomerStore) UpdateCustomerMetaInfo(ctx context.Context, list ...*driver.UpdateMetaInfoRequest) error {
	fn := s.d.UpdateMetaInfo
	if s.withTx {
		fn = s.tx.UpdateMetaInfo
	}
	return fn(ctx, &Customer{}, &Customer{}, list...)
}

func (s CustomerStore) GetCustomer(ctx context.Context, fields []string, cond CustomerCondition, opt ...getCustomersOption) (*Customer, error) {
	if len(fields) == 0 {
		fields = (&Customer{}).Fields()
	}
	m := MetaInfoForList{}
	listOpts := []listCustomersOption{
		&CursorBasedPagination{Limit: 1},
	}
	for _, o := range opt {
		t, _ := o.getValue()
		switch t {
		case driver.OptionType_MetaInfo:
			listOpts = append(listOpts, &m)
		}
	}
	objList, err := s.ListCustomers(ctx, fields, cond, listOpts...)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	for _, o := range opt {
		t, in := o.getValue()
		switch t {
		case driver.OptionType_MetaInfo:
			in.(*MetaInfo).UpdatedBy = m[0].UpdatedBy
			in.(*MetaInfo).CreatedBy = m[0].CreatedBy
			in.(*MetaInfo).DeletedBy = m[0].DeletedBy
			in.(*MetaInfo).UpdatedOn = m[0].UpdatedOn
			in.(*MetaInfo).CreatedOn = m[0].CreatedOn
			in.(*MetaInfo).DeletedOn = m[0].DeletedOn
			in.(*MetaInfo).IsDeleted = m[0].IsDeleted
		}
	}
	return objList[0], nil
}

func (s CustomerStore) ListCustomers(ctx context.Context, fields []string, cond CustomerCondition, opt ...listCustomersOption) ([]*Customer, error) {
	if len(fields) == 0 {
		fields = (&Customer{}).Fields()
	}
	var (
		res driver.Result
		err error
		m   driver.MetaInfo

		limit       = -1
		orderByList = make([]driver.OrderByType, 0, 5)
	)
	for _, o := range opt {
		t, in := o.getValue()
		switch t {
		case driver.OptionType_Pagination:
			page, ok := in.(*CursorBasedPagination)
			if page != nil && ok {
				if page.SetCustomerCondition == nil {
					page.SetCustomerCondition = defaultSetCustomerCondition
				}
				cond = page.SetCustomerCondition(page.UpOrDown, page.Cursor, cond)
				limit = page.Limit + 1
				if len(orderByList) == 0 {
					orderByList = append(orderByList, driver.OrderByType{
						Field:     "id",
						Ascending: !page.UpOrDown,
					})
				}
			}
		case driver.OptionType_MetaInfo:
			s.d.MetaInfoRequested(&ctx, &m)
		case driver.OptionType_OrderBy:
			by, ok := in.(OrderBy)
			if ok && len(by.Bys) != 0 {
				orderByList = by.Bys
			}
		}
	}
	if len(orderByList) == 0 {
		orderByList = append(orderByList, driver.OrderByType{
			Field:     "id",
			Ascending: true,
		})
	}
	ctx = driver.SetOrderBy(ctx, orderByList...)
	if limit > 0 {
		ctx = driver.SetListLimit(ctx, limit)
	}

	if s.withTx {
		res, err = s.tx.Get(ctx, cond.customerCondToDriverCustomerCond(s.d), &Customer{}, &Customer{}, fields...)
	} else {
		res, err = s.d.Get(ctx, cond.customerCondToDriverCustomerCond(s.d), &Customer{}, &Customer{}, fields...)
	}
	if err != nil {
		return nil, err
	}
	defer res.Close()

	mp := map[string]struct{}{}
	list := make([]*Customer, 0, 1000)
	infoMap := make(map[string]*driver.MetaInfo, 0)

	for res.Next(ctx) && limit != 0 {
		obj := &Customer{}
		if err := res.Scan(ctx, obj); err != nil {
			return nil, err
		}
		for _, o := range opt {
			t, _ := o.getValue()
			switch t {
			case driver.OptionType_MetaInfo:
				infoMap[obj.Id] = &driver.MetaInfo{
					UpdatedBy: m.UpdatedBy,
					CreatedBy: m.CreatedBy,
					DeletedBy: m.DeletedBy,
					UpdatedOn: m.UpdatedOn,
					CreatedOn: m.CreatedOn,
					DeletedOn: m.DeletedOn,
					IsDeleted: m.IsDeleted,
				}
				break
			}
		}
		list = append(list, obj)
		if _, ok := mp[obj.Id]; !ok {
			limit--
			mp[obj.Id] = struct{}{}
		}
	}
	if err := res.Close(); err != nil {
		return nil, err
	}

	list = MapperCustomer(list)
	meta := &MetaInfoForList{}

	for _, o := range opt {
		t, in := o.getValue()
		switch t {
		case driver.OptionType_Pagination:
			page, ok := in.(*CursorBasedPagination)
			if page != nil && ok {
				if len(list) <= page.Limit {
					page.HasNext = false
					page.HasPrevious = false
				} else {
					list = list[:page.Limit]
					if page.UpOrDown {
						page.HasPrevious = true
					} else {
						page.HasNext = true
					}
				}
			}
		case driver.OptionType_MetaInfo:
			meta = in.(*MetaInfoForList)
		}
	}
	for _, l := range list {
		*meta = append(*meta, infoMap[l.Id])
	}
	return list, nil
}

func (s CustomerStore) CountCustomers(ctx context.Context, cond CustomerCondition) (int, error) {
	cntFn := s.d.Count
	if s.withTx {
		cntFn = s.tx.Count
	}
	return cntFn(ctx, cond.customerCondToDriverCustomerCond(s.d), &Customer{}, &Customer{})
}

type getCustomersOption interface {
	getOptCustomers() // method of no significant use
	getValue() (driver.OptionType, interface{})
}

func (*MetaInfo) getOptCustomers() { // method of no significant use
}

type listCustomersOption interface {
	listOptCustomers() // method of no significant use
	getValue() (driver.OptionType, interface{})
}

func (*MetaInfoForList) listOptCustomers() {
}

func (OrderBy) listOptCustomers() {
}

func (*CursorBasedPagination) listOptCustomers() {
}

func defaultSetCustomerCondition(upOrDown bool, cursor string, cond CustomerCondition) CustomerCondition {
	if upOrDown {
		if cursor != "" {
			return CustomerAnd{cond, CustomerIdLt{cursor}}
		}
		return cond
	}
	if cursor != "" {
		return CustomerAnd{cond, CustomerIdGt{cursor}}
	}
	return cond
}

type CustomerAnd []CustomerCondition

func (p CustomerAnd) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.customerCondToDriverCustomerCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type CustomerOr []CustomerCondition

func (p CustomerOr) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.customerCondToDriverCustomerCond(d))
	}
	return driver.Or{Conditioners: dc, Operator: d}
}

type CustomerParentEq struct {
	Parent string
}

func (c CustomerParentEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerFullParentEq struct {
	Parent string
}

func (c CustomerFullParentEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "fullParent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerParentNotEq struct {
	Parent string
}

func (c CustomerParentNotEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerFullParentNotEq struct {
	Parent string
}

func (c CustomerFullParentNotEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "fullParent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerParentLike struct {
	Parent string
}

func (c CustomerParentLike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Like{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerFullParentLike struct {
	Parent string
}

func (c CustomerFullParentLike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Like{Key: "fullParent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerParentILike struct {
	Parent string
}

func (c CustomerParentILike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.ILike{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerFullParentILike struct {
	Parent string
}

func (c CustomerFullParentILike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.ILike{Key: "fullParent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerParentIn struct {
	Parent []string
}

func (c CustomerParentIn) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerFullParentIn struct {
	Parent []string
}

func (c CustomerFullParentIn) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "fullParent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerParentNotIn struct {
	Parent []string
}

func (c CustomerParentNotIn) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerFullParentNotIn struct {
	Parent []string
}

func (c CustomerFullParentNotIn) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "fullParent", Value: c.Parent, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdEq struct {
	Id string
}

func (c CustomerIdEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdNotEq struct {
	Id string
}

func (c CustomerIdNotEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdGt struct {
	Id string
}

func (c CustomerIdGt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdLt struct {
	Id string
}

func (c CustomerIdLt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdGtOrEq struct {
	Id string
}

func (c CustomerIdGtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdLtOrEq struct {
	Id string
}

func (c CustomerIdLtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdLike struct {
	Id string
}

func (c CustomerIdLike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Like{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdILike struct {
	Id string
}

func (c CustomerIdILike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.ILike{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeleted struct {
	IsDeleted bool
}

func (c CustomerDeleted) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: c.IsDeleted, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedByEq struct {
	By string
}

func (c CustomerCreatedByEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c CustomerCreatedOnEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedByNotEq struct {
	By string
}

func (c CustomerCreatedByNotEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c CustomerCreatedOnNotEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedByGt struct {
	By string
}

func (c CustomerCreatedByGt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c CustomerCreatedOnGt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedByLt struct {
	By string
}

func (c CustomerCreatedByLt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c CustomerCreatedOnLt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedByGtOrEq struct {
	By string
}

func (c CustomerCreatedByGtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c CustomerCreatedOnGtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedByLtOrEq struct {
	By string
}

func (c CustomerCreatedByLtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c CustomerCreatedOnLtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedByLike struct {
	By string
}

func (c CustomerCreatedByLike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Like{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerCreatedByILike struct {
	By string
}

func (c CustomerCreatedByILike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.ILike{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedByEq struct {
	By string
}

func (c CustomerUpdatedByEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c CustomerUpdatedOnEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedByNotEq struct {
	By string
}

func (c CustomerUpdatedByNotEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c CustomerUpdatedOnNotEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedByGt struct {
	By string
}

func (c CustomerUpdatedByGt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c CustomerUpdatedOnGt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedByLt struct {
	By string
}

func (c CustomerUpdatedByLt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c CustomerUpdatedOnLt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedByGtOrEq struct {
	By string
}

func (c CustomerUpdatedByGtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c CustomerUpdatedOnGtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedByLtOrEq struct {
	By string
}

func (c CustomerUpdatedByLtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c CustomerUpdatedOnLtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedByLike struct {
	By string
}

func (c CustomerUpdatedByLike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Like{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerUpdatedByILike struct {
	By string
}

func (c CustomerUpdatedByILike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.ILike{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedByEq struct {
	By string
}

func (c CustomerDeletedByEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c CustomerDeletedOnEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedByNotEq struct {
	By string
}

func (c CustomerDeletedByNotEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c CustomerDeletedOnNotEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedByGt struct {
	By string
}

func (c CustomerDeletedByGt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c CustomerDeletedOnGt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedByLt struct {
	By string
}

func (c CustomerDeletedByLt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c CustomerDeletedOnLt) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedByGtOrEq struct {
	By string
}

func (c CustomerDeletedByGtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c CustomerDeletedOnGtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedByLtOrEq struct {
	By string
}

func (c CustomerDeletedByLtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c CustomerDeletedOnLtOrEq) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedByLike struct {
	By string
}

func (c CustomerDeletedByLike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.Like{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerDeletedByILike struct {
	By string
}

func (c CustomerDeletedByILike) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.ILike{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Customer{}, RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdIn struct {
	Id []string
}

func (c CustomerIdIn) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

type CustomerIdNotIn struct {
	Id []string
}

func (c CustomerIdNotIn) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &Customer{}, FieldMask: "id", RootDescriptor: &Customer{}, CurrentDescriptor: &Customer{}}
}

func (c TrueCondition) customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type customerMapperObject struct {
	id string
}

func (s *customerMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperCustomer(rows []*Customer) []*Customer {

	ids := make([]string, 0, len(rows))
	uniqueIDMap := map[string]bool{}
	for _, r := range rows {
		if uniqueIDMap[r.Id] {
			continue
		}
		uniqueIDMap[r.Id] = true
		ids = append(ids, r.Id)
	}

	combinedCustomerMappers := map[string]*customerMapperObject{}

	for _, rw := range rows {

		tempCustomer := &customerMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		tempCustomer.id = rw.Id

		if combinedCustomerMappers[tempCustomer.GetUniqueIdentifier()] == nil {
			combinedCustomerMappers[tempCustomer.GetUniqueIdentifier()] = tempCustomer
		}
	}

	combinedCustomers := make(map[string]*Customer, 0)

	for _, customer := range combinedCustomerMappers {
		tempCustomer := &Customer{}
		tempCustomer.Id = customer.id

		if tempCustomer.Id == "" {
			continue
		}

		combinedCustomers[tempCustomer.Id] = tempCustomer

	}
	list := make([]*Customer, 0, len(combinedCustomers))
	for _, i := range ids {
		list = append(list, combinedCustomers[i])
	}
	return list
}

func (m *Customer) IsUsedMultipleTimes(f string) bool {
	return false
}

type TrueCondition struct{}

type CustomerCondition interface {
	customerCondToDriverCustomerCond(d driver.Driver) driver.Conditioner
}

type CursorBasedPagination struct {
	// Set UpOrDown = true for getting list of data above Cursor-ID,
	// limited to 'limit' amount, when ordered by ID in Ascending order.
	// Set UpOrDown = false for getting list of data below Cursor-ID,
	// limited to 'limit' amount, when ordered by ID in Ascending order.
	Cursor   string
	Limit    int
	UpOrDown bool

	// All pagination-cursor condition functions for different objects
	// SetCustomerCondition will be used to set the condition parameter for
	// setting parameter based on UpOrDown value,
	// if null default IdGt or IdLt condition will be used.
	SetCustomerCondition func(upOrDown bool, cursor string, cond CustomerCondition) CustomerCondition

	// Response objects Items - will be updated and set after the list call
	HasNext     bool // Used in case of UpOrDown = false
	HasPrevious bool // Used in case of UpOrDown = true
}

func (p *CursorBasedPagination) getValue() (driver.OptionType, interface{}) {
	return driver.OptionType_Pagination, p
}

type MetaInfo driver.MetaInfo
type MetaInfoForList []*driver.MetaInfo

func (p *MetaInfo) getValue() (driver.OptionType, interface{}) {
	return driver.OptionType_MetaInfo, p
}

func (p *MetaInfoForList) getValue() (driver.OptionType, interface{}) {
	return driver.OptionType_MetaInfo, p
}

type OrderBy struct {
	Bys []driver.OrderByType
}

func (o OrderBy) getValue() (driver.OptionType, interface{}) {
	return driver.OptionType_OrderBy, o
}
