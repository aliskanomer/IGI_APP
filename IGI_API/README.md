# Inter-Galactic Index API &#128640; &#11088; &#127759;

> v1.0.0 - Author: Ömer Faruk Alışkan - 25.12.24

This documentation covers the general architectural structure of the IGI API. It covers the general structure, components, constraints and rule set of the application with some detailed business clearences.

---
---
---

## &#128187; Get on Board!

Before diving deep into the architectural aspects; lets run over the local test options. IGI API can run standalone and supports Docker & Swagger

### Run on your local machine (Standalone)

1. Open root directory on a terminal and run `go run ./cmd/main.go`
2. Application should be up and running on `localhost:8080`
3. URL for the local server is provided in application logs

### Run on your local machine (Docker)

This application requires env variables to operate succesfully. Please run

```
docker run -p 8080:8080 -e DOCKER_ENV=true -e BASE_URL=http://localhost:8080 -e SWAPI_BASE_URL=https://www.swapi.tech/api/ -e SWAGGER_URL=http://localhost:8080/docs/ igi_api
```

in project directory to succesfully launch docker container.

### How to update documentation (Swagger)

1. Open rrot directory on a terminal and run `swag init --generalInfo ./cmd/main.go --output ./docs --parseDependency --parseInternal --parseDepth=2`
2. This should generate swagger files by checking the root directory and depth 3 files(eg. controllers)
3. URL for the server documentation ( Swagger UI ) is provided in application logs

---
---
---

## &#128194; Architecture

IGI API follows a **Layered Architecture(N-Tier)** that follows the SOLID principles (as much as it can &#128522;). Here is the base folder structure to make everything a bit more easier to understand.

```
| API
| --- cmd
|       |--- main.go
| --- docs
| --- internal
        |--- models
        |--- routes
        |--- middlewares
        |--- controllers
        |--- services
        |--- utils
        |--- cache
```

This structure also contains other modules to enhance readabiliy and maintainence but there are 4 core layers that **every single request** must go through. Well there are actually 7 since each of those folders are considered as layers but we can small it down to 4

1. Server Layer
2. Presentation Layer
3. Service Layer
4. Data Layer

Those layers works together with other components to provide the information that is requested.

### &#9881; Server Layer

This layer includes the portions of the API that makes the API itself running

- **Entrance Point and Routes**
- Environment Config, Server Quality & Health Utilities (Env,Logger,Error)
- Some of the Middlewares (Header, Securtiy, CORS, Error)

are included in this layer. Before any communication between the client and the data provider starts, this layers activates itself to set up the communication channel right. This layer sets up the rules of the communication itself.

### &#128221; Presentation Layer

This layer is responsible of interpertion in the most basic terms. Modules of this layer are related to either caching and parsing the client request or configuring the service layer responses before they are passed to the client.

- **Controllers**
- Response or request related mappers and builder utilities (RespBuilder, Helper)

are included in this layer. Welcoming and bidding farewell to HTTP requests happens via this layer, therefore running validations, parsing and rebuilding data according to the request and ensuring securty of the communication is this layers responsibility.

### &#128279; Service Layer

This layer is the only layer within the application that has data source access. Service layer directly communicates with Data Layer to fetch (so far this API only supports HTTP-GET methods) data from [SWAPI](https://swapi.tech/).

- **Services**
- Response or request related parsers (Helper)

are included in this layer. Service layer has a unique rule to keep in mind:
**_Service structs must follow Singleton Pattern to prevent garbage instances_** This rule exists to prevent multiple instance creation for the same service. This layer is responsible to manipulate the data itself not by the request's need but the API's need. Since this backend application acts as a proxy; service layer holds a crucial value to hold and secure the contunium of the communication. Any error or success when communicating from [SWAPI](https://swapi.tech/) must be resolved here so that Presentation Layer always have an information to provide to the client. Unhandled threads or errors in this layer might cause drastic problems.

### &#128274; Data Layer

This is the layer where everything is based actually. Whole communication of this API relies on this layers integrity.

- Models
- Cache Modules

are included in this layer. This layer gives meaning to other layers as well providing them a framework that they can communicate with. Cache is in this layer because Search business can feed on both [SWAPI](https://swapi.tech/) or Application Cache memory. (Check Business part for more information about cache and search). This layer is responsible to validate and secure the data itself and/or it's model.

> Important thing to point out: Errors and Logs! API is a communication ruleset within itself as well as it's between two distinct parties. Logging mechanisms and Error Handling does not actually fit in to the any layer because they are the tools that used in every layer.

## &#129504; Business

Since there is not an exact business object (because this application does not have a direct manipulation right on the data it provides) there is not much to talk about the business part but the search. People and Planet resources serves the same sort of business: you can either read all data or a spesific one.

Search on the other hand has a bit more complex structure. [SWAPI](https://swapi.tech/) does not provides a [famous searching company that owns everything]-alike search endpoint. Each resource has its own query parameter to run search.

- People Resource => `/people/?name=muaddib`
- Planet Resource => `/planets/?name=arakis`

Can provide you the data you are searching for but there isnt a way to check out for a keyword within the whole data sturcture.

**IGI API** on the other hand has a "proxy" resource like that to provide information in a more simpler way. This is the reason why this application has a caching mechanism.

### Search & Caching

A Search in this application means running concurrent queries on desired resource set. Fancy sentence huh? It simply means: you select the resources you want to search within, backend runs query for them at the same time and gives you a response.

Well thats how every search kinda works so lets see the whole qury to understand. A very base search URL (with default parameters) does looks like this and return a JSON object like below:

```URL
/search/?keyword=attreidies&source="people,planets"&page=1&limit=15&sortBy="name"&sortOrder="asc"
```

```JSON
{
  "message": "string",
  "response": {
    "dataSource": "string",
    "hitCount": {
      "total": 0,
      "totalPagePeople": 0,
      "totalPagePlanet": 0
    },
    "results": {
      "people": [],
      "planets": []
    }
}
```

- &#10067; **Results** -> Currently resource can only be either `People | Planet | People,Planet`. Search algorithm will run the query for the source with the URL's provided above. Each matched data set will fill it's own array on the `results{}` object.

- &#10067; **Hit Count** -> It holds the metadata of the search response. When service layer provides the raw response data in their respective resource array; controller counts, paginates and sort the data before sending it as a response. By doing so controller also builds this metadata for client. This sorting and segmentation is needed to improve performance and keep an healthy API. Default segmentation keys are (1:page-15:limit).

  - Why? Let's say user only searched by the keyword "e". It is the most used character in english alphabet by far. There will be maybe hundreds of data that matches with that keyword. One way to prevent this kind of paylod would be adding UI restrictions but as long as route is out there; somebody might try to violete it or some bot might. Therefore there should be restriction on how much data is passed between two parties to prevent long response times and ease the backend applications workload a bit.
    > PS : UI also does prevent unnecesary calls by the hitCount meta data like disabling the pagination buttons or search button itself in some cases.

- &#10067; **Data Source** -> This is where the Cache mechanism gets it's curtain call; When any search query is made controllers first checks the cache to see if same search keyword on same resources has been made by anyone else within the 15 minute limit. - Why? Resources on [SWAPI](https://swapi.tech/) returns whole data at once for each query. Lets say you searched for "e" in people and there is 82 matches. [SWAPI](https://swapi.tech/) response will have 82 entries then. This creates a problem: - How? If the response containt too big of a data set, response [SWAPI](https://swapi.tech/) will be late, meaning IGI API will be late. Also there is API limitations on [SWAPI](https://swapi.tech/) like max limits and ratings.
  Thats why instead of making a new call on every query **IGI API saves each response set to it's Cache for 15 mins as key value pair, so that if the same user looks for the second data set of the same query(page=2) or two different person makes runs the same query there is no need to fetch data again we can just read it from cache**
  Cached items are stored in key-value pairs. Key holds the information of the _keyword_ and _resource_ parameters making the results pair bounded to their query. If there is no cahced response for the given query, than controller invokes the service to fetch the data from the source and then sorted but not paginated data is stored on cache. - Why sorted but not paginated? Well sorted because it does not make sense to sort each data-set(page-limit group in this case) within themselves. Whole set of response must be sorted before getting segmentated. Also if IGI API only holds the segment of the response for the provided query, then it has to make another call again to get another segment. > PS: This is an architectural choice, depending of the cost of caching it might make more sense to make a call for each query. In order to see which one is best, product requires lots of test from different user cases. But it is common practice to use cache in such scnearios and it would be better if the IGI API would use a seperate cacheing mechanism like Redis instead of server memory.

And here we are: this is how search works. When a request is made

1. Check cache to see if there is any data matches by the requested data
2. If not invoke service ro read data
   2.1 Sort raw data from service (each array on itself)
   2.2 Cache sorted raw data with its keyword and resource list as a key
3. If there is a cached data read it
4. Paginate the response and finish

## &#128230; Logging & Error Handling

IGI API has a modular structure that seperates concerns by their responsive businesses and use cases. Commonly used methodologies such as server logs are organized within the utilities to provide a more readable and maintainable code base.

A log in IGI API would propably look like this.

```
INF_STAT:0_OPS:ServerCONFG Environment file loaded!
```

```
type_STAT:code_OPS:operation message
```

Type can be, info(INF), error(ERR) or success(SCC). code is usually HTTP status code but in-server logs can also use 0 as code. operation is the invoker method that creates the log and the message contain detailed information about the log or error.

In server; there is a global error handling middleware that chained to each route. This middleware is there to catch any unresolved error or threat to provide an error message to the client no matter what. Each layer tries to prevent the errors propagating by if-else cases or safe-fails but if any thing wents south, ErrorMiddleware is there to provide an answer the to client.

---

If you have any other questions or problem please don't hesitate to get in contact:

**Support:** Ömer Faruk Alışkan - omerfarukaliskan@icloud.com
