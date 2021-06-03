package main

import (
	"go-dig/repository"
	"go-dig/usecases"
	"miidas/domain/admin/light_message_stats"
	"miidas/domain/company/section_matching"
	"miidas/domain/connect/talk"
	"miidas/usecase/admin/stats"

	"go.uber.org/dig"
)

var diContainer = newDiContainer()

func GetDiContainer() *DiContainer {
	return &diContainer
}

type DiContainer struct {
	Container *dig.Container
}

func newDiContainer() DiContainer {
	c := DiContainer{
		Container: dig.New(),
	}
	c.provide()
	return c
}

func (di *DiContainer) provide() {
	di.provideRepository()
	di.provideUsecase()
}

func (di *DiContainer) provideRepository() {
	di.Container.Provide(repository.NewUserRepository)
}

func (di *DiContainer) provideUsecase() {
	di.Container.Provide(usecases.NewGetUser)
}
