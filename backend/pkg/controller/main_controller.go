package controller

import "github.com/SzymekN/Car-rental-app/pkg/auth"

// nie wiem co to ma robić ale to chyba było do przechowywania wszystkich HandlerObjects i generycznej rejestracji
type MainController struct {
	jwtH     *auth.JWTHandler
	handlers []BasicHandler
}

func (mc *MainController) RegisterAllRoutes() {
	mc.jwtH.RegisterRoutes()
	for _, handler := range mc.handlers {
		handler.RegisterRoutes()
	}
}
