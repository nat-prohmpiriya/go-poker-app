package main

import (
	"poker-app/handler"
	"poker-app/repository"
	"poker-app/service"
)

func main() {
	repo := repository.NewGameRepository()
	deckService := service.NewDeckService(repo)
	handService := service.NewHandService()
	gameService := service.NewGameService(handService)
	gameHandler := handler.NewGameHandler(deckService, gameService, repo)

	gameHandler.Run()
}
