package biz

import (
	"bird-ecology/app/model/page"
	"bird-ecology/app/model/po"
	"bird-ecology/app/model/query"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jiushengTech/common/utils/snowflakeutil"
)

type {{cookiecutter.service_name}}Repo interface {
	Page(context.Context, *query.{{cookiecutter.service_name}}) (*page.Model[po.{{cookiecutter.service_name}}], error)
	List(context.Context, *query.{{cookiecutter.service_name}}) ([]*po.{{cookiecutter.service_name}}, error)
	Detail(context.Context, *po.{{cookiecutter.service_name}}) (*po.{{cookiecutter.service_name}}, error)
	Insert(context.Context, *po.{{cookiecutter.service_name}}) (*po.{{cookiecutter.service_name}}, error)
	Update(context.Context, *po.{{cookiecutter.service_name}}) error
	Delete(context.Context, *po.{{cookiecutter.service_name}}) error
}

type {{cookiecutter.service_name}}UseCase struct {
	repo {{cookiecutter.service_name}}Repo
	log  *log.Helper
}

func New{{cookiecutter.service_name}}UseCase(
	{{cookiecutter.service_name}}Repo {{cookiecutter.service_name}}Repo,
	logger log.Logger) *{{cookiecutter.service_name}}UseCase {
	return &{{cookiecutter.service_name}}UseCase{
		repo: {{cookiecutter.service_name}}Repo,
		log:  log.NewHelper(log.With(logger, "module", "internal/biz/{{cookiecutter.file_name}}")),
	}
}

func (uc *{{cookiecutter.service_name}}UseCase) Page(ctx context.Context, g *query.{{cookiecutter.service_name}}) (*page.Model[po.{{cookiecutter.service_name}}], error) {
	res, err := uc.repo.Page(ctx, g)
	if err != nil {
		uc.log.Errorf("Model error: %s req: %+v", err, g)
		return nil, err
	}
	return res, err
}

func (uc *{{cookiecutter.service_name}}UseCase) List(ctx context.Context, g *query.{{cookiecutter.service_name}}) ([]*po.{{cookiecutter.service_name}}, error) {
	res, err := uc.repo.List(ctx, g)
	if err != nil {
		uc.log.Errorf("List error: %s req: %+v", err, g)
		return nil, err
	}
	return res, err
}

func (uc *{{cookiecutter.service_name}}UseCase) Detail(ctx context.Context, id uint64) (*po.{{cookiecutter.service_name}}, error) {
	res, err := uc.repo.Detail(ctx, &po.{{cookiecutter.service_name}}{ID: id})
	if err != nil {
		uc.log.Errorf("Detail error: %s req: %+v", err, id)
		return nil, err
	}
	return res, err
}

func (uc *{{cookiecutter.service_name}}UseCase) Save(ctx context.Context, g *po.{{cookiecutter.service_name}}) (*po.{{cookiecutter.service_name}}, error) {
	g.ID = snowflakeutil.GetId()
	res, err := uc.repo.Insert(ctx, g)
	if err != nil {
		uc.log.Errorf("Save error: %s, req: %+v", err, g)
		return nil, err
	}
	return res, err
}

func (uc *{{cookiecutter.service_name}}UseCase) Modify(ctx context.Context, g *po.{{cookiecutter.service_name}}) error {
	err := uc.repo.Update(ctx, g)
	if err != nil {
		uc.log.Errorf("ModifyById error: %s ,req: %+v", err, g)
		return err
	}
	return err
}

func (uc *{{cookiecutter.service_name}}UseCase) Remove(ctx context.Context, id uint64) error {
	err := uc.repo.Delete(ctx, &po.{{cookiecutter.service_name}}{ID: id})
	if err != nil {
		uc.log.Errorf("RemoveById error: %s, req: %+v", err, id)
		return err
	}
	return err
}
