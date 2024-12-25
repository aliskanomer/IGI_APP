# Inter-Galactic Index CLIENT &#128640; &#11088; &#127759;

> v1.0.0 - Author: Ömer Faruk Alışkan - 25.12.24

This is documents covers the styling and structure of the IGI Client Application. It is a React Typescript application with WebPack bundling. 

---
---
---

## &#128187; Get on Board!

Before diving deep into the architectural aspects; lets run over the local test options. Please be aware that; IGI Client needs IGI API serve fully.

### Run on your local machine (Dev)

1. Open root directory on a terminal and run `npm install` to download packages.
2. If packets installed run `npm start` to run application on your local machine.
3. Application should be up an running at `http://localhos:3000`

### Run on your local machine (Prod/Static-server)

1. Open root directory on a terminal and run `npm install` to download packages.
2. If packets installed run `npm build` generate `/dist` folder.
3. Run `http-server dist` to run prod build in static HTTP server. Link to the server should be provided in logs after a succesfull run 

### Run on your local machine (Docker)

To run IGI Client in a docker container 
```
docker run -p 3000:80 \
  -e IGI_API_DEV=http://localhost:8080 \
  igi_client
``` 
in project directory to succesfully launch docker container.

---
---
---

## &#128194; Architecture

Just like IGI API does, IGI Client also has a moduler and segmented architecture. Components that returns a JSX elements follow a event-driven structure, while class-based typescript files follows the similar patterns to the IGI_API

```
| SRC
| --- api
|       |--- services
|       |       |---peopleService
|       |       |---planetService
|       |       |---searchService
|       |       |---IGIClient(index)
|       |       |---types
| --- assets
| --- components
|       |--- common
|       |--- errorBoundry
|       |--- searchFilters
|--- context
|--- pages
|       |--- search
|       |--- planet
|       |--- people
|       |--- notFouns
|--- App
```

Just like any React application, there is a public html file that acts as root node of the DOM and index binds App to the DOM and everything... There is no need to go in detail about each component of a application with a this base of UI but there are parts that needs to be described. API Layer is the base of the application layers because it holds the types that used troughout all of the application.  

## &#128233; API Layer (Service Layer)

Backbone of the whole application. Service layer makes the API connection to retrieve data from server. There is no need describe all services since each service holds JS documentation within its code. but index file which creates a IGI Client instance holds a special place.

*API Layer classes must follow singleton structure to prevent memory leaks*

### &#128268; IGI Client (Axios Client)

Is actually an axios singleton instance with interperters to make HTTP request. Every other service like `peopleService`, `planetService` or `searchService` uses a IGI Client instance to fetch data. Currently (since there is only GET request defined as endpoints) IGI Client can only make HTTP GET requests. 

This class also sets the base rules of HTTP request within the application like header settings or timeour 
> All requests has 10sc time out limit in IGI Client


## &#127760; Routing and Pages

Routing is provided by `ReactRoute` in IGI Client. Pages within the application are the representetives of the IGI_API endpoints and design by itself (Please see styling) requires each page to make an API call to show some data set. Both `/planet` and `/people` routes in IGI Application mounts with an API call to fill their content. Therefore pages usually has request management states like loading,error and data. Depending on the page's special requirements these states might be extended by others. 


### &#127912; Presenting Data and Details on Pages
Design structure in this manner is simple. Each page has two section: Meta & Data. Meta can really be just meta; an image and a some text like `/people` and `/planet` but for `/search` meta actually means a form. 

Data section uses common `Loaders` when mounted till page makes the request by service layer and fills the data. Once fetched data will be presented in `Cards` which again is a common component that changes it's style and content based on the invoker method and provided props.

There is not much of a difference between how `/planet` and `/people` pages works. On mount both makes their respective resource's `getAll()` call by service and fill the data to the cards. `Cards` hold the `uid` of their element. This `uid` prop is passed to an common `Overlay` component when any card is clicked. Overlay than makes the `getByID(uid)` calls by service layer to fetch the detail of given subject. 

> Due to Typescript there is lots of mapping and altering in this part's business.Check Overlay.tsx and Card.tsx for more details

> Managing the opening and closing of the Overlay requires two way binding between the Card and the Overlay it opens. Check Overlay.tsx and Card.tsx for more details

### &#128218; Context & Search & Filters 

Setting up a valid search query requires multiple validations by itself, mixing this with React component structure and best practices can easily cause lots of prop-drilling and two-way bindings which makes the code harder to read. 

This why IGI Client Application has a `SearchContext` that wraps up the whole application and provides values such as `query`, `validateQuery`,`resetQuery`, `updateQuery`. Currently this context only holds query but not the error states because there is not much of an error user can encounter due to restrictions. Regardless, when there is validation error; it is logged on console. 

### &#129309; Working with context

Components such as `SearchFilters` or `SearchPage` has deep bond with the context. Any input item on filters reads and sets the query by `context.updateQuery`. This prevents passing states and methods between multiple layers of component and makes code easier to read. 

- Context listens the active route and resets it's query anytime active route changes. This prevents memory leaks and also helps user experience. (Many users use the navigation menu as an escape mechanism)
- Context holds the default values and safe fail algorithms to prevent unnecesarry calls and unexpected errors
- UI elements are selected and manipulated in a way to ensure search safety as well. For an example resource selection in `SearchFilters` comes as a checkbox group but it is impossible to set resource as none by UI elements.
- Even if everything goes wrong; service layer always validates and inserts default values for non-valid parameters before making any HTTP Request.
- To ensure search only happens by user events; only way to make a `searchService.search()` call is by clicking the search button on `SearchPage` which is disabled when the query is empty. `SearchFilter` is a seperated component because of that. Ensuring the query validation and keyword validation will not become spagetti code
> PS: Seperating the results from the page might even be good idea for better readability but i have some concerns on that manner as well(like two way binding neccesity)

### &#128209; Results & Pagination

As other pages does, `SearchPage` displays the results with `Card` component that opens an `Overlay` to show more information about the record. There are minor differences due to [SWAPI](https://swapi.tech/) but idea is the same.

> Search queries on SWAPI returns full resource model array as response, meanwhile list calls like people and planet returns more plain ones

And just like other pages and list data, by the IGI API principle; results are returned in paginated data sets. For other pages pagination is not even an issue since they only serve one resource at a time.

But search serves multiple resources grouped under different arrays. This makes pagination a bit more complex and React life-cycle makes it even more....

#### &#128208; Pagination Logic on Search

 There is only one pagination button on search page that paginates both `people` and `planet` response. This can create issue if those resources has different amount of segments available(Which is the case most of the time) To cover this, algorithm runs on logic below

- &#127919; After a succesfull request, `searchPage` saves the segmented resource data on its local state by two seperate array for each of the resources. Each rescourde is rendered with cards as seperated lists, if there is any hit for that resource.

- &#127919; When data fethced `hitCount` transferred into a local state called `searchMeta`. This structure has a map as below:
    ```JS
        searchMeta = {
            totalHits: 0,
            totalPages: 0
        }
    ```
   -  `totalHits` are read from the response but `totalPages` decides depending on the which resource array has more total page count. Because you can easily measure 10cm with a 1m ruler but other way around will be harder.

        - Imagine you have `people = [1,2,3]` and `planet= [1,2]`
        - Max page count in this case will be decided by `people[]` and its 3
        - When call for `/search/?...page=3&limit=1` response will have data for `people` but since there is not a third segment for `planet`, on last call UI will recieve a `people[]` but `planet` will be `null` 

- &#127919; Decided max page count then setted to a local state. Why? Because of React! Everytime next or prev buttons clicked, application must make another fetch to collect the new set. React Life-cycle works mysterious ways when it comes to this. 
    - If page count is read and written from the context; there is high possiblity for the page count to update on `nextTick()` this causes the same search query to happen again, which can affect the pagination in unexpected ways.

- &#127919; Each time pagination buttons clicked, both context and local states are updated before update both prev and next are checked to prevent negative values and overflows. 

- &#127919; Since each pagination means a new call, `SearchPage` have to invoke the `searchService` whenever page value is set. Problem is page value is always set. It's default value is 1, because of this: `SearchPage` holds another booelan flag called `paginated`. This state is by default false and only setted true when any pagination button clicked; it is also resetted when location changes(mentioned escape mechanism). Instead of listening the page and invoking the service; `searchPage` checks this flag to invoke a new call.
    - And this call will get it's data from cache unless the first call happened more than 15 minutes ago.Because cached items are checked by query and resource and pagination changes neither of those values


## &#128165; Error Handling

1. `errorBoundry` component wraps up the whole application to catch any unHandled error. This will prevent the app from crashing and inform the user that there is an error.

2. `notFound` page handles the unrecognized route errors. 

3. `IGI Client (Axios Client)` ensures that every `Promise` either resolved or rejected so that the call stack is empty once HTTP request is finalized. 

Other then this there is not much of an error handling in the UI application. 


## &#127912; Styling & Assets

It would a wrong to say application has a design system; but it will also be wrong to say that it has not &#128513; . IGI Client Application uses `SASS` for stying due to it's capacity. General structure is setted on the `/assets/styles` yet each component can have it's own stylesheet. Those stylesheet usually uses and extends common stylings defined in `/assets/styles`

```
| assets
| --- styles
|       |---_common            -> I/O element generalization (buttons,inputs)
|       |---_layout            -> Layout class declerations (row,col,page)
|       |---_typography        -> Generalization of typeface (h1,h2...,p,a..)
|       |---_variables         -> Common variable declerations (colors,font)
|       |---styles             -> global styling

```

Modular structural approach applies the stylings as well. Modules named with underscore are `@use`d in `styles.scss` and then impoerted on the root level. So there is no need to import styles again and again for multiple components. 

As an example we can take `peoplePage` and `planetPage`. Both pages have the same layout and instead of defining them individually; `_layout.scss` contains the `.page` styling that is applied on both of them. 

General coding practice on styling is BEM a-like approach with identifier nested like `.page`,`.page-header`,`.page-list--item`. 

For more information about styling and to see the ideation of it you can visit [Figma](https://www.figma.com/design/osKWlMEf4lXEz2gcIpmkZM/Galactic-Index?node-id=3-722&t=vB8INFnvwdEtTN1m-1)

>PS: I am not sure about the links expiration date. Please get in contact with me if there is an issue.

----

If you have any other questions or problem please don't hesitate to get in contact:

**Support:** Ömer Faruk Alışkan - omerfarukaliskan@icloud.com








