package routes

import (
	// packages
	"IGI_API/internal/controllers"
	"IGI_API/internal/middlewares"

	// modules
	"net/http"
)

func InitRoutes() {
	// middleware arr to be used in all routes
	commonMiddlewares := []middlewares.Middleware{
		middlewares.ErrorMiddleware,    // Global Error Handling
		middlewares.CORSMiddleware,     // CORS Headers
		middlewares.SecurityMiddleware, // Security Headers
		middlewares.HeaderMiddleware,   // Content-Type Header
	}

	// IGI_API ROOT
	http.HandleFunc("/", middlewares.ChainMiddleware(
		func(resp http.ResponseWriter, req *http.Request) {
			resp.Write([]byte(`{"message":"Welcome to IGI_API"}`))
		},
		commonMiddlewares...,
	))

	// OPERATION: SEARCH --  /search?name="keyword"&source=[people|planet]&page=int&limit=int&sort=asc|desc&sortby=name|height|mass
	http.HandleFunc("/search", middlewares.ChainMiddleware(
		controllers.Search,
		commonMiddlewares...,
	))

	// OPERATION: PEOPLE-GetAll --  /people
	http.HandleFunc("/people", middlewares.ChainMiddleware(
		controllers.GetPeopleAll,
		commonMiddlewares...,
	))
	// OPERATION: PEOPLE-GetByID --  /people/:id
	http.HandleFunc("/people/", middlewares.ChainMiddleware(
		controllers.GetPeopleById,
		commonMiddlewares...,
	))

	// OPERATION: PLANET-GetAll --  /planet
	http.HandleFunc("/planet", middlewares.ChainMiddleware(
		controllers.GetPlanetAll,
		commonMiddlewares...,
	))
	// OPERATION: PLANET-GetByID --  /planet/:id
	http.HandleFunc("/planet/", middlewares.ChainMiddleware(
		controllers.GetPlanetById,
		commonMiddlewares...,
	))
}
