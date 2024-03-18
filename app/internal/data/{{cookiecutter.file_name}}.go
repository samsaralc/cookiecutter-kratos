package data

import (
	"bird-ai-service/internal/biz"
	"bird-ecology/app/model/page"
	"bird-ecology/app/model/po"
	"bird-ecology/app/model/query"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type {{cookiecutter.service_name}}Repo struct {
	data *Data
	log  *log.Helper
}

func New{{cookiecutter.service_name}}Repo(data *Data, logger log.Logger) biz.{{cookiecutter.service_name}}Repo {
	return &{{cookiecutter.service_name}}Repo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "internal/data/{{cookiecutter.service_name}}Repo")),
	}
}

func (r *{{cookiecutter.service_name}}Repo) Page(ctx context.Context, g *query.{{cookiecutter.service_name}}) (*page.Model[po.{{cookiecutter.service_name}}], error) {
	if g != nil {
		g = &query.{{cookiecutter.service_name}}{}
	}
	db := r.data.db.WithContext(ctx).Model(&po.{{cookiecutter.service_name}}{})
	db.Where(&g.{{cookiecutter.service_name}})
	if g.StartTime != "" {
		parse, err := time.Parse(time.DateTime, g.StartTime)
		if err != nil {
			return nil, errors.New(500, "开始时间格式错误", err.Error())
		}
		db.Where("created_at >= ?", parse)
	}
	if g.EndTime != "" {
		parse, err := time.Parse(time.DateTime, g.EndTime)
		if err != nil {
			return nil, errors.New(500, "结束时间格式错误", err.Error())
		}
		db.Where("created_at <= ?", parse)
	}
	return page.GetResult[po.{{cookiecutter.service_name}}](db, g.PageNum, g.PageSize)
}

func (r *{{cookiecutter.service_name}}Repo) List(ctx context.Context, g *query.{{cookiecutter.service_name}}) ([]*po.{{cookiecutter.service_name}}, error) {
	if g != nil {
		g = &query.{{cookiecutter.service_name}}{}
	}
	db := r.data.db.WithContext(ctx).Model(&po.{{cookiecutter.service_name}}{})
	db.Where(&g.{{cookiecutter.service_name}})
	if g.StartTime != "" {
		parse, err := time.Parse(time.DateTime, g.StartTime)
		if err != nil {
			return nil, errors.New(500, "开始时间格式错误", err.Error())
		}
		db.Where("created_at >= ?", parse)
	}
	if g.EndTime != "" {
		parse, err := time.Parse(time.DateTime, g.EndTime)
		if err != nil {
			return nil, errors.New(500, "结束时间格式错误", err.Error())
		}
		db.Where("created_at <= ?", parse)
	}
	var res []*po.{{cookiecutter.service_name}}
	return res, db.Find(&res).Error
}

func (r *{{cookiecutter.service_name}}Repo) Detail(ctx context.Context, g *po.{{cookiecutter.service_name}}) (*po.{{cookiecutter.service_name}}, error) {
	return BaseMapper[*po.{{cookiecutter.service_name}}]{}.Detail(ctx, g)
}

func (r *{{cookiecutter.service_name}}Repo) Insert(ctx context.Context, g *po.{{cookiecutter.service_name}}) (*po.{{cookiecutter.service_name}}, error) {
	return g, BaseMapper[*po.{{cookiecutter.service_name}}]{}.Insert(ctx, g)
}

func (r *{{cookiecutter.service_name}}Repo) Update(ctx context.Context, g *po.{{cookiecutter.service_name}}) error {
	return BaseMapper[*po.{{cookiecutter.service_name}}]{}.Update(ctx, g)
}

func (r *{{cookiecutter.service_name}}Repo) Delete(ctx context.Context, g *po.{{cookiecutter.service_name}}) error {
	return BaseMapper[*po.{{cookiecutter.service_name}}]{}.Delete(ctx, g)
}
