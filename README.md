# Inter-Galactic Index (IGI) &#128640; &#11088; &#127759;

> v1.0.0 - Author: Ömer Faruk Alışkan - 25.12.24
> v1.0.1 - Author: Ömer Faruk Alışkan - 31.12.24

Welcome! Inter-Galactic Index is a product that helps you find Star Wars universe characters and planets. It uses [SWAPI](https://swapi.tech/) to collect data, and represents it with a fresh user interface and a proxy backend service. 

This document is here to help you to get on board with the project. So let's start


## &#128681; Let's Start

First things first goal of this documentation is to guide you about ***"How to"*** questions. Both API and Client directories has their own documentation files (Check out the respective directory's README.md) to inform about the application structure and business flows. 

>Therefore if you are looking for ***"Oh how does that search work?"*** you are in the wrong place. You should check Client and API documentation for that.

### &#127939; How To Run Application

IGI_APP is a dockerized application that runs on `PORT:6969` in order for IGI_APP to serve on your local machine:

1. Make sure Docker Desktop is up and running.
2. Open a terminal and locate to the application directory 
3. Run `docker-compose up --build` to create docker containers to be built and ready to serve
4. If everything went well your terminal should be locked with docker and you can open `http://localhost:6969/` to see the client application 

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

IGI Application consumes `People` and `Planet` resources and has a derived `Search` source to perform it's job.

### &#129489; People

People registered in [SWAPI](https://swapi.tech/) is readable by IGI APP and its modules. You can get all people registered or detailed information of a single person. You can also search people by their name field. 

### &#129680; Planets

Planets registered in [SWAPI](https://swapi.tech/) is readable by IGI APP and its modules. You can get all planets registered or detailed information of a single one. You can also search planets by their name field. 

### &#128270; Search

Even though [SWAPI](https://swapi.tech/) does not provides an combined search algorithm, IGI APP does. With IGI API you are able to run a query on both people and planet resources at the same time. Each resource hits will be represented in 15 piece set, within themselves.

That concludes it actually. If you dont have [Docker](https://www.docker.com/) or [Node](https://nodejs.org/tr/download/package-manager) or [Go](https://go.dev/doc/install) in your local machine you might have some issues with local running but downloading docker and setting up the Docker Desktop as it's suggested in the documentation might be the easiest fix for that.

Please continue to listed directories to gain more insight about technical aspect of this project

- .../IGIAPP/IGI_CLIENT/README.md -> Client & Design Documentation
- .../IGIAPP/IGI_API/README.md -> API & Business Documentation

## &#128051; Some helpful Docker compose commands for possible debugging
- Build from scratch    `docker-compose up --build`
- Run on background     `docker-compose up -d`
- See logs              `docker-compose logs -f`
- See running services  `docker-compose ps`
- Kill container        `docker-compose down`
- Clean Up (!Dangerous) `docker-compose down --volumes`


## &#129302; Possible Improvements

Below technologies and features listed would be important to implement for real prod build. 

#### Backend
1. Unit tests should be integrated using standart `testing` library for go. 
1. Message broker integration for asynchronous communication between services. It would be important for scaling and handling concurrent requests. `RabbitMQ` or `Kafka` might be a good choice.
2. Integrate the API with a separate cache service for performance improvements. This would reduce the cost of search. `Redis` might be a good option to choose.
3. SWAPI has rate limiting; since IGI_API is a proxy API, it should obey the dependencies of SWAPI. Therefore, integrating rate limiting using an API Gateway such as `Kong` might be good for better service quality and performance. A gateway would also be set up for throttling to prevent malicious requests.
4. Logging and monitoring could be improved in different ways depending on the lifespan of the product and business requests. If a deep understanding of API usage is required, mechanisms with dashboards can be implemented using `OpenSearch, Kibana` or any other monitoring and logging tools. `Prometheus` could be a good choice for metric tracking.
5. Improve `IgiClient.ts` class to type secure conditions by integrating dynamic type insertion.

#### Frontend
1. Integrate a minimizer for the production build for faster service time. `Terser` might be used.
2. Integrate `ESLint` and `SonarLint`  and `Prettier` to enforce coding rules and proper linting.
3. Extend a `/public` directory in the root for SEO integrations.
    - Create `/public/sitemap.xml`, `/public/manifest.json`, and `/public/robots.txt` files to represent the website and activate findability by search engines.
    - Enhance `/public/index.html` `<head/>` with Structured Data string for better SEO outcomes.
    - Cross check `<tag/>` usage to ensure SEO optimized information presentation and organization.
4. Implement animations for more smooth experience. 
5. Implement a design system for more accurate interface.
6. Implement full responsiveness to prevent unexpected renderings. Currently flex layout covers a lot but better experience lies in details.

#### Operations
Stages should be visible by automation server like `Jenkins`. This requires an whole devops planning. Currently there are 2 stages for each appplication; which are build and serve. In order to create a CI/CD pipeline these stages could be improved by adding `Testing` and `Linting` stages. For test; backend could run unit tests. For linting; `Sonarqube` could be good choice.

Listed improvements are crucial before a real production release for a proffessional application.

## Support

If you have any other questions or problem please don't hesitate to get in contact:

**Support:** Ömer Faruk Alışkan - omerfarukaliskan@icloud.com
