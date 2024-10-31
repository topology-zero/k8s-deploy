package common

import (
	"encoding/json"
	"net/url"
	"os"
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

// ProjectParams 项目参数
func ProjectParams(ctx *svc.ServiceContext, req *types.PathID) (resp []types.CommonProjectParamsResponse, err error) {
	projectModel := query.ProjectModel
	projectInfo, _ := projectModel.WithContext(ctx).Where(projectModel.ID.Eq(req.ID)).First()

	var tag []string
	if projectInfo.UseTag == 0 {
		tag = []string{"latest"}
	} else {
		tag, err = getTags(ctx, projectInfo)
		if err != nil {
			return nil, err
		}
	}

	namespace, err := getNamespace(ctx)
	if err != nil {
		return nil, err
	}

	tagVal := ""
	if len(tag) > 0 {
		tagVal = tag[0]
	}

	resp = append(resp,
		types.CommonProjectParamsResponse{
			Name:    "namespace",
			Options: namespace,
		},
		types.CommonProjectParamsResponse{
			Name:    "tag",
			Value:   tagVal,
			Options: tag,
		},
		types.CommonProjectParamsResponse{
			Name:  "subset",
			Value: strings.ReplaceAll(tagVal, ".", ""),
		})

	var params []types.CommonProjectParamsResponse
	_ = json.Unmarshal([]byte(projectInfo.Params), &params)
	resp = append(resp, params...)

	return
}

func getTags(ctx *svc.ServiceContext, projectInfo *model.ProjectModel) ([]string, error) {
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
