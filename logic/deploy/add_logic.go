package deploy

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
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
	deployModel := query.DeployModel

	projectInfo, _ := projectModel.WithContext(ctx).Where(projectModel.ID.Eq(req.ID)).First()
	if projectInfo == nil {
		return errors.New("项目不存在")
	}

	localParams := make(map[string]string)
	for _, v := range req.Params {
		localParams[v.Name] = v.Value
	}

	t := template.Must(template.New("templateByte").Parse(req.TemplateContent))
	buffer := new(bytes.Buffer)
	err := t.Execute(buffer, localParams)
	if err != nil {
		ctx.Log.Error("%+v", errors.WithStack(err))
		return err
	}

	params, _ := json.Marshal(req.Params)
	project, _ := json.Marshal(projectInfo)

	err = deployModel.WithContext(ctx).Create(&model.DeployModel{
		Name:            req.Name,
		ProjectID:       req.ID,
		Project:         string(project),
		Fingerprint:     fmt.Sprintf("%x", md5.Sum([]byte(req.TemplateContent))),
		TemplateName:    req.TemplateName,
		TemplateContent: req.TemplateContent,
		TemplateParse:   buffer.String(),
		Params:          string(params),
	})

	if err != nil {
		ctx.Log.Error("%+v", errors.WithStack(err))
	}
	return err
}
