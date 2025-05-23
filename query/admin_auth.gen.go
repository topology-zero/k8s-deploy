// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"k8s-deploy/model"
)

func newAdminAuthModel(db *gorm.DB, opts ...gen.DOOption) adminAuthModel {
	_adminAuthModel := adminAuthModel{}

	_adminAuthModel.adminAuthModelDo.UseDB(db, opts...)
	_adminAuthModel.adminAuthModelDo.UseModel(&model.AdminAuthModel{})

	tableName := _adminAuthModel.adminAuthModelDo.TableName()
	_adminAuthModel.ALL = field.NewAsterisk(tableName)
	_adminAuthModel.ID = field.NewInt(tableName, "id")
	_adminAuthModel.Pid = field.NewInt(tableName, "pid")
	_adminAuthModel.Name = field.NewString(tableName, "name")
	_adminAuthModel.Key = field.NewString(tableName, "key")
	_adminAuthModel.IsMenu = field.NewInt(tableName, "is_menu")
	_adminAuthModel.API = field.NewString(tableName, "api")
	_adminAuthModel.Action = field.NewString(tableName, "action")
	_adminAuthModel.CreateTime = field.NewTime(tableName, "create_time")
	_adminAuthModel.UpdateTime = field.NewTime(tableName, "update_time")
	_adminAuthModel.DeleteTime = field.NewField(tableName, "delete_time")

	_adminAuthModel.fillFieldMap()

	return _adminAuthModel
}

// adminAuthModel 权限
type adminAuthModel struct {
	adminAuthModelDo adminAuthModelDo

	ALL        field.Asterisk
	ID         field.Int
	Pid        field.Int    // 上级ID
	Name       field.String // 节点名
	Key        field.String // 权限标识
	IsMenu     field.Int    // 是否是菜单栏 0：否 1：是
	API        field.String // 接口
	Action     field.String // 操作方法
	CreateTime field.Time
	UpdateTime field.Time
	DeleteTime field.Field

	fieldMap map[string]field.Expr
}

func (a adminAuthModel) Table(newTableName string) *adminAuthModel {
	a.adminAuthModelDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminAuthModel) As(alias string) *adminAuthModel {
	a.adminAuthModelDo.DO = *(a.adminAuthModelDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminAuthModel) updateTableName(table string) *adminAuthModel {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt(table, "id")
	a.Pid = field.NewInt(table, "pid")
	a.Name = field.NewString(table, "name")
	a.Key = field.NewString(table, "key")
	a.IsMenu = field.NewInt(table, "is_menu")
	a.API = field.NewString(table, "api")
	a.Action = field.NewString(table, "action")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.DeleteTime = field.NewField(table, "delete_time")

	a.fillFieldMap()

	return a
}

func (a *adminAuthModel) WithContext(ctx context.Context) IAdminAuthModelDo {
	return a.adminAuthModelDo.WithContext(ctx)
}

func (a adminAuthModel) TableName() string { return a.adminAuthModelDo.TableName() }

func (a adminAuthModel) Alias() string { return a.adminAuthModelDo.Alias() }

func (a adminAuthModel) Columns(cols ...field.Expr) gen.Columns {
	return a.adminAuthModelDo.Columns(cols...)
}

func (a *adminAuthModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminAuthModel) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["pid"] = a.Pid
	a.fieldMap["name"] = a.Name
	a.fieldMap["key"] = a.Key
	a.fieldMap["is_menu"] = a.IsMenu
	a.fieldMap["api"] = a.API
	a.fieldMap["action"] = a.Action
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["delete_time"] = a.DeleteTime
}

func (a adminAuthModel) clone(db *gorm.DB) adminAuthModel {
	a.adminAuthModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminAuthModel) replaceDB(db *gorm.DB) adminAuthModel {
	a.adminAuthModelDo.ReplaceDB(db)
	return a
}

type adminAuthModelDo struct{ gen.DO }

type IAdminAuthModelDo interface {
	gen.SubQuery
	Debug() IAdminAuthModelDo
	WithContext(ctx context.Context) IAdminAuthModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminAuthModelDo
	WriteDB() IAdminAuthModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminAuthModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminAuthModelDo
	Not(conds ...gen.Condition) IAdminAuthModelDo
	Or(conds ...gen.Condition) IAdminAuthModelDo
	Select(conds ...field.Expr) IAdminAuthModelDo
	Where(conds ...gen.Condition) IAdminAuthModelDo
	Order(conds ...field.Expr) IAdminAuthModelDo
	Distinct(cols ...field.Expr) IAdminAuthModelDo
	Omit(cols ...field.Expr) IAdminAuthModelDo
	Join(table schema.Tabler, on ...field.Expr) IAdminAuthModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminAuthModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminAuthModelDo
	Group(cols ...field.Expr) IAdminAuthModelDo
	Having(conds ...gen.Condition) IAdminAuthModelDo
	Limit(limit int) IAdminAuthModelDo
	Offset(offset int) IAdminAuthModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminAuthModelDo
	Unscoped() IAdminAuthModelDo
	Create(values ...*model.AdminAuthModel) error
	CreateInBatches(values []*model.AdminAuthModel, batchSize int) error
	Save(values ...*model.AdminAuthModel) error
	First() (*model.AdminAuthModel, error)
	Take() (*model.AdminAuthModel, error)
	Last() (*model.AdminAuthModel, error)
	Find() ([]*model.AdminAuthModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminAuthModel, err error)
	FindInBatches(result *[]*model.AdminAuthModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminAuthModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminAuthModelDo
	Assign(attrs ...field.AssignExpr) IAdminAuthModelDo
	Joins(fields ...field.RelationField) IAdminAuthModelDo
	Preload(fields ...field.RelationField) IAdminAuthModelDo
	FirstOrInit() (*model.AdminAuthModel, error)
	FirstOrCreate() (*model.AdminAuthModel, error)
	FindByPage(offset int, limit int) (result []*model.AdminAuthModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminAuthModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminAuthModelDo) Debug() IAdminAuthModelDo {
	return a.withDO(a.DO.Debug())
}

func (a adminAuthModelDo) WithContext(ctx context.Context) IAdminAuthModelDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminAuthModelDo) ReadDB() IAdminAuthModelDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminAuthModelDo) WriteDB() IAdminAuthModelDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminAuthModelDo) Session(config *gorm.Session) IAdminAuthModelDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminAuthModelDo) Clauses(conds ...clause.Expression) IAdminAuthModelDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminAuthModelDo) Returning(value interface{}, columns ...string) IAdminAuthModelDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminAuthModelDo) Not(conds ...gen.Condition) IAdminAuthModelDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminAuthModelDo) Or(conds ...gen.Condition) IAdminAuthModelDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminAuthModelDo) Select(conds ...field.Expr) IAdminAuthModelDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminAuthModelDo) Where(conds ...gen.Condition) IAdminAuthModelDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminAuthModelDo) Order(conds ...field.Expr) IAdminAuthModelDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminAuthModelDo) Distinct(cols ...field.Expr) IAdminAuthModelDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminAuthModelDo) Omit(cols ...field.Expr) IAdminAuthModelDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminAuthModelDo) Join(table schema.Tabler, on ...field.Expr) IAdminAuthModelDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminAuthModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminAuthModelDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminAuthModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminAuthModelDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminAuthModelDo) Group(cols ...field.Expr) IAdminAuthModelDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminAuthModelDo) Having(conds ...gen.Condition) IAdminAuthModelDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminAuthModelDo) Limit(limit int) IAdminAuthModelDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminAuthModelDo) Offset(offset int) IAdminAuthModelDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminAuthModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminAuthModelDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminAuthModelDo) Unscoped() IAdminAuthModelDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminAuthModelDo) Create(values ...*model.AdminAuthModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminAuthModelDo) CreateInBatches(values []*model.AdminAuthModel, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminAuthModelDo) Save(values ...*model.AdminAuthModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminAuthModelDo) First() (*model.AdminAuthModel, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminAuthModel), nil
	}
}

func (a adminAuthModelDo) Take() (*model.AdminAuthModel, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminAuthModel), nil
	}
}

func (a adminAuthModelDo) Last() (*model.AdminAuthModel, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminAuthModel), nil
	}
}

func (a adminAuthModelDo) Find() ([]*model.AdminAuthModel, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminAuthModel), err
}

func (a adminAuthModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminAuthModel, err error) {
	buf := make([]*model.AdminAuthModel, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminAuthModelDo) FindInBatches(result *[]*model.AdminAuthModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminAuthModelDo) Attrs(attrs ...field.AssignExpr) IAdminAuthModelDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminAuthModelDo) Assign(attrs ...field.AssignExpr) IAdminAuthModelDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminAuthModelDo) Joins(fields ...field.RelationField) IAdminAuthModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminAuthModelDo) Preload(fields ...field.RelationField) IAdminAuthModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminAuthModelDo) FirstOrInit() (*model.AdminAuthModel, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminAuthModel), nil
	}
}

func (a adminAuthModelDo) FirstOrCreate() (*model.AdminAuthModel, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminAuthModel), nil
	}
}

func (a adminAuthModelDo) FindByPage(offset int, limit int) (result []*model.AdminAuthModel, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminAuthModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminAuthModelDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminAuthModelDo) Delete(models ...*model.AdminAuthModel) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminAuthModelDo) withDO(do gen.Dao) *adminAuthModelDo {
	a.DO = *do.(*gen.DO)
	return a
}
