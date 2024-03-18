package service

import (
	"bird-ai-service/internal/biz"
	pb "bird-ecology/api/bird_ai_srv/bird"
	"bird-ecology/app/model/po"
	"bird-ecology/app/model/query"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type {{cookiecutter.service_name}}Service struct {
	pb.Unimplemented{{cookiecutter.service_name}}Server
	uc  *biz.{{cookiecutter.service_name}}UseCase
	log *log.Helper
}

func New{{cookiecutter.service_name}}Service(
	uc *biz.{{cookiecutter.service_name}}UseCase,
	logger log.Logger,
) *{{cookiecutter.service_name}}Service {
	return &{{cookiecutter.service_name}}Service{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/{{cookiecutter.file_name}}")),
	}
}

func (s *{{cookiecutter.service_name}}Service) Create{{cookiecutter.service_name}}(ctx context.Context, req *pb.Create{{cookiecutter.service_name}}Request) (*pb.Create{{cookiecutter.service_name}}Reply, error) {
	to := &po.{{cookiecutter.service_name}}{}
	err := copier.Copy(&to, &req)
	if err != nil {
		s.log.Errorf("Create{{cookiecutter.service_name}} Copy1 error: %s", err)
		return &pb.Create{{cookiecutter.service_name}}Reply{}, err
	}
	res, err := s.uc.Save(ctx, to)
	if err != nil {
		return &pb.Create{{cookiecutter.service_name}}Reply{}, err
	}
	reply := &pb.Create{{cookiecutter.service_name}}Reply{}
	err = copier.Copy(reply, res)
	if err != nil {
		s.log.Errorf("Create{{cookiecutter.service_name}} Copy2 error: %s", err)
		return &pb.Create{{cookiecutter.service_name}}Reply{}, err
	}
	return reply, nil
}
func (s *{{cookiecutter.service_name}}Service) Update{{cookiecutter.service_name}}(ctx context.Context, req *pb.Update{{cookiecutter.service_name}}Request) (*pb.Update{{cookiecutter.service_name}}Reply, error) {
	to := &po.{{cookiecutter.service_name}}{}
	err := copier.Copy(&to, &req)
	if err != nil {
		s.log.Errorf("Update{{cookiecutter.service_name}} Copy error: %s", err)
		return &pb.Update{{cookiecutter.service_name}}Reply{}, err
	}
	err = s.uc.Modify(ctx, to)
	return &pb.Update{{cookiecutter.service_name}}Reply{}, err
}
func (s *{{cookiecutter.service_name}}Service) Delete{{cookiecutter.service_name}}(ctx context.Context, req *pb.Delete{{cookiecutter.service_name}}Request) (*pb.Delete{{cookiecutter.service_name}}Reply, error) {
	err := s.uc.Remove(ctx, req.GetId())
	return &pb.Delete{{cookiecutter.service_name}}Reply{}, err
}
func (s *{{cookiecutter.service_name}}Service) Get{{cookiecutter.service_name}}(ctx context.Context, req *pb.Get{{cookiecutter.service_name}}Request) (*pb.Get{{cookiecutter.service_name}}Reply, error) {
	res, err := s.uc.Detail(ctx, req.GetId())
	if err != nil {
		s.log.Errorf("Get{{cookiecutter.service_name}} Copy1 error: %s", err)
		return &pb.Get{{cookiecutter.service_name}}Reply{}, err
	}
	reply := &pb.Get{{cookiecutter.service_name}}Reply{}
	err = copier.Copy(reply, res)
	if err != nil {
		s.log.Errorf("Get{{cookiecutter.service_name}}Reply Copy2 error: %s", err)
		return &pb.Get{{cookiecutter.service_name}}Reply{}, err
	}
	return reply, nil
}
func (s *{{cookiecutter.service_name}}Service) List{{cookiecutter.service_name}}(ctx context.Context, req *pb.List{{cookiecutter.service_name}}Request) (*pb.List{{cookiecutter.service_name}}Reply, error) {
	to := &query.{{cookiecutter.service_name}}{}
	err := copier.Copy(&to, req)
	if err != nil {
		s.log.Errorf("List{{cookiecutter.service_name}} Copy1 error: %s", err)
		return &pb.List{{cookiecutter.service_name}}Reply{}, err
	}
	res, err := s.uc.List(ctx, to)
	if err != nil {
		return &pb.List{{cookiecutter.service_name}}Reply{}, err
	}
	reply := &pb.List{{cookiecutter.service_name}}Reply{}
	err = copier.Copy(&reply.CameraVideoClips, res)
	if err != nil {
		s.log.Errorf("List{{cookiecutter.service_name}}Reply Copy2 error: %s", err)
		return &pb.List{{cookiecutter.service_name}}Reply{}, err
	}
	return reply, nil
}

func (s *{{cookiecutter.service_name}}Service) Page{{cookiecutter.service_name}}(ctx context.Context, req *pb.Page{{cookiecutter.service_name}}Request) (*pb.Page{{cookiecutter.service_name}}Reply, error) {
	return &pb.Page{{cookiecutter.service_name}}Reply{}, nil
}


