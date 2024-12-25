# Inter-Galactic Index (IGI) &#128640; &#11088; &#127759;

> v1.0.0 - Author: Ömer Faruk Alışkan - 25.12.24

Welcome! Inter-Galactic Index is a product that helps you find Star Wars universe characters and planets. It uses [SWAPI](https://swapi.tech/) to collect data, and represents it with a fresh user interface and a proxy backend service. 

This document is here to help you to get on board with the project. So let's start


## &#128681; Let's Start

First things first goal of this documentation is to guide you about ***"How to"*** questions. Both API and Client directories has their own documentation files (Check out the respective directory's README.md) to inform about the application structure and business flows. 

>Therefore if you are looking for ***"Oh how does that search work?"*** you are in the wrong place. You should check Client and API documentation for that.

### &#127939; How To Run Application

IGI_APP is a dockerized application that runs on `PORT:6969` in order for IGI_APP to serve on your local machine:

1. Open a terminal and locate to the application directory 
2. Run `docker-compose up --build` to create docker containers to be built and ready to serve
3. If everything went well your terminal should be locked with docker and you can open `http://localhost:6969/` to see the client application 

And Voila! May the force be with you in your seeking of information!

### &#9881; Proxy & Routes

IGI_API has a proxy to serve all of its modules in a single port. This proxy allows you to reach both backend and frontend application from single port

- Frontend Application : `http://localhost:6969/`
- Backend Application : `http://localhost:6969/api/`
- SWAGGER Documentation :  `http://localhost:6969/api/docs/`

> Please make sure that you have "/" at end of your URL's to reach to related points. There is minor issue with SWAGGER redirection. Insert `/index.html` at the end of the route if you come up with any problems. In any case `localhost:8080/docs` must return the SWAGGER documentation even when the Nginx redirection fails. 


### &#9995; Specs

Backend and Frontend modules of this application has the specs listed below as well as their own dependency arrays. 

| Module      | Technology Used           | Base Image        | Runtime Environment | Application Server   | Port  |
|-----------------|---------------------------|--------------------|----------------------|-----------------------|-------|
| Backend        | GoLang                    | golang:1.23.4      | alpine:latest       | net/http             | 8080  |
| Frontend       | React + TypeScript + Webpack | node:20           | nginx:alpine        | NGINX                | 3000  |


## &#129300; Features and Resources

IGI Application does not have it's own resources. There is no database connection on backend module, making it a proxy backend service. Therefore there is an huge dependency to [SWAPI](https://swapi.tech/). Yet, that does not mean there isn't any business model for the API. Details of business object and logics explained in the IGI_CLIENT's documentation

> Basically if they change anything this application needs to be updated in order to perform.

IGI Application consumes `People` and `Planet` resources and has a derrived `Search` source to perform it's job.

### &#129489; People

People registered in [SWAPI](https://swapi.tech/) is readable by IGI APP and its modules. You can get all people registered or detailed information of a single person. You can also search people by their name field. 

### &#129680; Planets

Planets registered in [SWAPI](https://swapi.tech/) is readable by IGI APP and its modules. You can get all planets registered or detailed information of a single one. You can also search planets by their name field. 

### &#128270; Search

Even though [SWAPI](https://swapi.tech/) does not provides an combined search algorithm, IGI APP does. With IGI API you are able to run a query on both people and planet resources at the same time. Each resource hits will be represented in 15 piece set, within themselves.

That concludes it actually. If you dont have [Docker](https://www.docker.com/) or [Node](https://nodejs.org/tr/download/package-manager) or [Go](https://go.dev/doc/install) in your local machine you might have some issues with local running but downloading docker and setting up the Docker Desktop as it's suggested in the documentation might be the easiest fix for that.

Please contuniue to listed directories to gain more insight about technical aspect of this project

- .../IGIAPP/IGI_CLIENT/README.md -> Client & Design Documentation
- .../IGIAPP/IGI_API/README.md -> API & Business Documentation

### &#128051; Some helpful Docker compose commands for possible debugging
- Build from scratch    `docker-compose up --build`
- Run on background     `docker-compose up -d`
- See logs              `docker-compose logs -f`
- See running services  `docker-compose ps`
- Kill container        `docker-compose down`
- Clean Up (!Dangerous) `docker-compose down --volumes`


If you have any other questions or problem please don't hesitate to get in contact:

**Support:** Ömer Faruk Alışkan - omerfarukaliskan@icloud.com
