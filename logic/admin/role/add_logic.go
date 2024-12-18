package role

import (
	"encoding/json"
	"strconv"

	"k8s-deploy/model"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Add 添加角色
func Add(ctx *svc.ServiceContext, req *types.RoleAddRequest) error {
	roleModel := query.AdminRoleModel
	roleInfo, _ := roleModel.WithContext(ctx).Where(roleModel.Name.Eq(req.Name)).First()
	if roleInfo != nil {
		return errors.New("该角色已存在")
	}

	marshal, _ := json.Marshal(req.Auth)
	authStr := string(marshal)

	defer model.Enforcer.LoadPolicy()

	err := query.Q.Transaction(func(tx *query.Query) error {
		saveRole := model.AdminRoleModel{
			Name: req.Name,
			Auth: authStr,
		}
		err := tx.AdminRoleModel.WithContext(ctx).Create(&saveRole)
		if err != nil {
			return err
		}

		var rules []*model.AdminCasbinRuleModel

		authModels, _ := tx.AdminAuthModel.WithContext(ctx).Where(tx.AdminAuthModel.ID.In(req.Auth...)).Find()
		for _, a := range authModels {
			if a.IsMenu == 1 {
				continue
			}

			rules = append(rules, &model.AdminCasbinRuleModel{
				Ptype: "p",
				V0:    "role:" + strconv.Itoa(saveRole.ID),
				V1:    a.API,
				V2:    a.Action,
			})
		}

		return tx.AdminCasbinRuleModel.WithContext(ctx).CreateInBatches(rules, 100)
	})
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
