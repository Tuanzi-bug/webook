// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package baguwen

import (
	"sync"

	"github.com/ecodeclub/ecache"
	"github.com/ecodeclub/mq-api"
	"github.com/ecodeclub/webook/internal/question/internal/event"
	"github.com/ecodeclub/webook/internal/question/internal/repository"
	"github.com/ecodeclub/webook/internal/question/internal/repository/cache"
	"github.com/ecodeclub/webook/internal/question/internal/repository/dao"
	"github.com/ecodeclub/webook/internal/question/internal/service"
	"github.com/ecodeclub/webook/internal/question/internal/web"
	"github.com/ego-component/egorm"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitModule(db *gorm.DB, ec ecache.Cache, q mq.MQ) (*Module, error) {
	questionDAO := InitQuestionDAO(db)
	questionCache := cache.NewQuestionECache(ec)
	repositoryRepository := repository.NewCacheRepository(questionDAO, questionCache)
	syncEventProducer := initSyncEventProducer(q)
	serviceService := service.NewService(repositoryRepository, syncEventProducer)
	handler := web.NewHandler(serviceService)
	questionSetDAO := InitQuestionSetDAO(db)
	questionSetRepository := repository.NewQuestionSetRepository(questionSetDAO)
	questionSetService := service.NewQuestionSetService(questionSetRepository, syncEventProducer)
	questionSetHandler, err := web.NewQuestionSetHandler(questionSetService)
	if err != nil {
		return nil, err
	}
	module := &Module{
		Svc:   serviceService,
		Hdl:   handler,
		QsHdl: questionSetHandler,
	}
	return module, nil
}

// wire.go:

var daoOnce = sync.Once{}

func InitTableOnce(db *gorm.DB) {
	daoOnce.Do(func() {
		err := dao.InitTables(db)
		if err != nil {
			panic(err)
		}
	})
}

func initSyncEventProducer(q mq.MQ) event.SyncEventProducer {
	producer, err := event.NewSyncEventProducer(q)
	if err != nil {
		panic(err)
	}
	return producer
}

func InitQuestionDAO(db *egorm.Component) dao.QuestionDAO {
	InitTableOnce(db)
	return dao.NewGORMQuestionDAO(db)
}

func InitQuestionSetDAO(db *egorm.Component) dao.QuestionSetDAO {
	InitTableOnce(db)
	return dao.NewGORMQuestionSetDAO(db)
}
