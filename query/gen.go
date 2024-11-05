// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                    = new(Query)
	AdminAuthModel       *adminAuthModel
	AdminCasbinRuleModel *adminCasbinRuleModel
	AdminRoleModel       *adminRoleModel
	AdminUserModel       *adminUserModel
	DeployLogModel       *deployLogModel
	DeployModel          *deployModel
	K8sTemplateModel     *k8sTemplateModel
	ProjectModel         *projectModel
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AdminAuthModel = &Q.AdminAuthModel
	AdminCasbinRuleModel = &Q.AdminCasbinRuleModel
	AdminRoleModel = &Q.AdminRoleModel
	AdminUserModel = &Q.AdminUserModel
	DeployLogModel = &Q.DeployLogModel
	DeployModel = &Q.DeployModel
	K8sTemplateModel = &Q.K8sTemplateModel
	ProjectModel = &Q.ProjectModel
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                   db,
		AdminAuthModel:       newAdminAuthModel(db, opts...),
		AdminCasbinRuleModel: newAdminCasbinRuleModel(db, opts...),
		AdminRoleModel:       newAdminRoleModel(db, opts...),
		AdminUserModel:       newAdminUserModel(db, opts...),
		DeployLogModel:       newDeployLogModel(db, opts...),
		DeployModel:          newDeployModel(db, opts...),
		K8sTemplateModel:     newK8sTemplateModel(db, opts...),
		ProjectModel:         newProjectModel(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AdminAuthModel       adminAuthModel
	AdminCasbinRuleModel adminCasbinRuleModel
	AdminRoleModel       adminRoleModel
	AdminUserModel       adminUserModel
	DeployLogModel       deployLogModel
	DeployModel          deployModel
	K8sTemplateModel     k8sTemplateModel
	ProjectModel         projectModel
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		AdminAuthModel:       q.AdminAuthModel.clone(db),
		AdminCasbinRuleModel: q.AdminCasbinRuleModel.clone(db),
		AdminRoleModel:       q.AdminRoleModel.clone(db),
		AdminUserModel:       q.AdminUserModel.clone(db),
		DeployLogModel:       q.DeployLogModel.clone(db),
		DeployModel:          q.DeployModel.clone(db),
		K8sTemplateModel:     q.K8sTemplateModel.clone(db),
		ProjectModel:         q.ProjectModel.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		AdminAuthModel:       q.AdminAuthModel.replaceDB(db),
		AdminCasbinRuleModel: q.AdminCasbinRuleModel.replaceDB(db),
		AdminRoleModel:       q.AdminRoleModel.replaceDB(db),
		AdminUserModel:       q.AdminUserModel.replaceDB(db),
		DeployLogModel:       q.DeployLogModel.replaceDB(db),
		DeployModel:          q.DeployModel.replaceDB(db),
		K8sTemplateModel:     q.K8sTemplateModel.replaceDB(db),
		ProjectModel:         q.ProjectModel.replaceDB(db),
	}
}

type queryCtx struct {
	AdminAuthModel       IAdminAuthModelDo
	AdminCasbinRuleModel IAdminCasbinRuleModelDo
	AdminRoleModel       IAdminRoleModelDo
	AdminUserModel       IAdminUserModelDo
	DeployLogModel       IDeployLogModelDo
	DeployModel          IDeployModelDo
	K8sTemplateModel     IK8sTemplateModelDo
	ProjectModel         IProjectModelDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AdminAuthModel:       q.AdminAuthModel.WithContext(ctx),
		AdminCasbinRuleModel: q.AdminCasbinRuleModel.WithContext(ctx),
		AdminRoleModel:       q.AdminRoleModel.WithContext(ctx),
		AdminUserModel:       q.AdminUserModel.WithContext(ctx),
		DeployLogModel:       q.DeployLogModel.WithContext(ctx),
		DeployModel:          q.DeployModel.WithContext(ctx),
		K8sTemplateModel:     q.K8sTemplateModel.WithContext(ctx),
		ProjectModel:         q.ProjectModel.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
