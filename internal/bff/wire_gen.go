// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bff

import (
	"github.com/ecodeclub/webook/internal/bff/internal/web"
	"github.com/ecodeclub/webook/internal/cases"
	"github.com/ecodeclub/webook/internal/interactive"
	baguwen "github.com/ecodeclub/webook/internal/question"
)

// Injectors from wire.go:

func InitModule(intrModule *interactive.Module, caseModule *cases.Module, queModule *baguwen.Module) (*Module, error) {
	service := intrModule.Svc
	serviceService := caseModule.Svc
	caseSetService := caseModule.SetSvc
	service2 := queModule.Svc
	questionSetService := queModule.SetSvc
	examineService := queModule.ExamSvc
	handler := web.NewHandler(service, serviceService, caseSetService, service2, questionSetService, examineService)
	module := &Module{
		Hdl: handler,
	}
	return module, nil
}

// wire.go:

type Handler = web.Handler
