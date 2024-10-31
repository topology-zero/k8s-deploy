package deploy

import (
	"bytes"
	"encoding/json"
	"text/template"

	"k8s-deploy/model"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Add 添加部署
func Add(ctx *svc.ServiceContext, req *types.DeployAddRequest) error {
	projectModel := query.ProjectModel
	k8sTemplateModel := query.K8sTemplateModel
	deployModel := query.DeployModel

	projectInfo, _ := projectModel.WithContext(ctx).Where(projectModel.ID.Eq(req.ProjectID)).First()
	if projectInfo == nil {
		return errors.New("项目不存在")
	}

	templateInfo, _ := k8sTemplateModel.WithContext(ctx).Where(k8sTemplateModel.ID.Eq(req.TemplateID)).First()
	if templateInfo == nil {
		return errors.New("模板不存在")
	}

	localParams := make(map[string]string)
	for _, v := range req.Params {
		localParams[v.Name] = v.Value
	}

	t := template.Must(template.New("templateByte").Parse(templateInfo.Content))
	buffer := new(bytes.Buffer)
	err := t.Execute(buffer, localParams)
	if err != nil {
		ctx.Log.Error("%+v", errors.WithStack(err))
		return err
	}

	params, _ := json.Marshal(req.Params)
	project, _ := json.Marshal(projectInfo)
	templateByte, _ := json.Marshal(templateInfo)

	err = deployModel.WithContext(ctx).Create(&model.DeployModel{
		Name:          req.Name,
		ProjectID:     req.ProjectID,
		Project:       string(project),
		TemplateID:    req.TemplateID,
		Template:      string(templateByte),
		TemplateParse: buffer.String(),
		Params:        string(params),
	})

	if err != nil {
		ctx.Log.Error("%+v", errors.WithStack(err))
	}
	return err
}
