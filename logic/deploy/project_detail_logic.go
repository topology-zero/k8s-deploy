package deploy

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"

	"k8s-deploy/model"
	"k8s-deploy/pkg/kubectl"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var k8sTemplateVar = regexp.MustCompile(`\{\{\s*\.(\w+)\s*}}`)

// ProjectDetail 项目详情
func ProjectDetail(ctx *svc.ServiceContext, req *types.PathID) (resp []types.DeployProjectDetailResponse, err error) {
	k8sTemplateModel := query.K8sTemplateModel
	projectModel := query.ProjectModel
	deployModel := query.DeployModel

	commonTemp, _ := k8sTemplateModel.WithContext(ctx).Order(k8sTemplateModel.ID.Desc()).Find()
	projectInfo, _ := projectModel.WithContext(ctx).Where(projectModel.ID.Eq(req.ID)).First()

	var projectTemp []types.NameAndValue
	_ = json.Unmarshal([]byte(projectInfo.Template), &projectTemp)

	for _, v := range projectTemp {
		resp = append(resp, types.DeployProjectDetailResponse{
			TemplateName:    fmt.Sprintf("[项目] - %s", v.Name),
			TemplateContent: v.Value,
		})
	}
	for _, v := range commonTemp {
		resp = append(resp, types.DeployProjectDetailResponse{
			TemplateName:    fmt.Sprintf("[通用] - %s", v.Name),
			TemplateContent: v.Content,
		})
	}

	var tag []string
	tagVal := ""
	if projectInfo.UseTag == 0 {
		tag = []string{"latest"}
	} else {
		tag, err = getGitTags(ctx, projectInfo)
		if err != nil {
			return nil, err
		}
	}

	if len(tag) > 0 {
		tagVal = tag[0]
	}

	namespace, err := getNamespace(ctx)
	if err != nil {
		return nil, err
	}

	for i, v := range resp {
		subMatch := k8sTemplateVar.FindAllStringSubmatch(v.TemplateContent, -1)
		tempVar := make(map[string]int)
		for _, vv := range subMatch {
			tempVar[vv[1]] = 1
		}

		if _, ok := tempVar["tag"]; ok {
			resp[i].Params = append(resp[i].Params, types.ProjectParams{
				Name:    "tag",
				Value:   tagVal,
				Options: tag,
			})
			delete(tempVar, "tag")
		}

		if _, ok := tempVar["namespace"]; ok {
			resp[i].Params = append(resp[i].Params, types.ProjectParams{
				Name:    "namespace",
				Options: namespace,
			})
			delete(tempVar, "namespace")
		}

		if _, ok := tempVar["subset"]; ok {
			resp[i].Params = append(resp[i].Params, types.ProjectParams{
				Name:  "subset",
				Value: strings.ReplaceAll(tagVal, ".", ""),
			})
			delete(tempVar, "subset")
		}

		var params []types.NameAndValue
		_ = json.Unmarshal([]byte(projectInfo.Params), &params)
		for _, param := range params {
			if _, ok := tempVar[param.Name]; ok {
				resp[i].Params = append(resp[i].Params, types.ProjectParams{
					Name:  param.Name,
					Value: param.Value,
				})
				delete(tempVar, param.Name)
			}
		}

		// 查找上次写入的值
		fingerprint := fmt.Sprintf("%x", md5.Sum([]byte(v.TemplateContent)))
		lastDeploy, _ := deployModel.WithContext(ctx).
			Where(deployModel.ProjectID.Eq(req.ID), deployModel.Fingerprint.Eq(fingerprint)).
			Order(deployModel.ID.Desc()).
			First()

		lastTempVar := make(map[string]string)

		if lastDeploy != nil {
			var params []types.NameAndValue
			_ = json.Unmarshal([]byte(lastDeploy.Params), &params)
			for _, vv := range params {
				lastTempVar[vv.Name] = vv.Value
			}

		}

		for k := range tempVar {
			newVale := ""
			if _, ok := lastTempVar[k]; ok {
				newVale = lastTempVar[k]
			}
			resp[i].Params = append(resp[i].Params, types.ProjectParams{
				Name:  k,
				Value: newVale,
			})
		}
	}

	return
}

func getGitTags(ctx *svc.ServiceContext, projectInfo *model.ProjectModel) ([]string, error) {
	urlParse, err := url.Parse(projectInfo.Git)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return nil, err
	}

	urlParse.User = url.UserPassword("joy_list", projectInfo.Token)

	split := strings.Split(projectInfo.Git, "/")

	dir := "./temp/git/" + split[len(split)-1]
	_, err = os.Stat(dir)
	var repository *git.Repository
	if err != nil && os.IsNotExist(err) {
		repository, err = git.PlainClone(dir, false, &git.CloneOptions{
			URL: urlParse.String(),
		})
	} else {
		repository, err = git.PlainOpen(dir)
	}

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return nil, err
	}

	worktree, err := repository.Worktree()
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return nil, err
	}

	err = worktree.Pull(&git.PullOptions{Force: true})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return nil, err
	}

	tags, err := repository.Tags()
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return nil, err
	}

	var newTags []string
	tags.ForEach(func(reference *plumbing.Reference) error {
		newTags = append(newTags, strings.TrimLeft(reference.Name().String(), "refs/tags/"))
		return nil
	})

	sort.Slice(newTags, func(i, j int) bool {
		return newTags[i] > newTags[j]
	})

	if len(newTags) > 6 {
		newTags = newTags[:6]
	}

	return newTags, nil
}

func getNamespace(ctx *svc.ServiceContext) ([]string, error) {
	ns, err := kubectl.K8sClient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		return nil, err
	}

	var namespace []string
	for _, v := range ns.Items {
		namespace = append(namespace, v.Name)
	}
	return namespace, nil
}
