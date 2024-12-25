/**
 * @type {Object} APIResponse - Generic IGI_API response model.Regardles of status, IGI_API responses will have this model.
 * @property {number} status - HTTP status code of the response
 * @property {string} message - Message from the API server
 * @property {T} response - Response object of the API
 * @template T - Generic type of the response object see derived types
 */
export interface APIResponse<T> {
  status: number;
  message: string;
  response: T;
}

/**
 * @type {Object} ErrorInfo - Error response model.Nested in the APIResponse.response object when response is not 100 200 300.
 * @property {string} code - Error code
 * @property {string} details - Error details
 */
export interface ErrorInfo {
  code: string;
  details: string;
}

/**
 * @type {Object} ListItem - List item model
 * @property {string} name - Name of the item
 * @property {string} uid - Unique identifier of the item
 * @property {string} url - URL of the item
 * */
export interface ListItem {
  name: string;
  uid: string;
  url: string;
}

/**
 * @type {Object} ListInfo - List info model
 * @property {string} message - Message from the SWAPI server
 * @property {string | null} next - URL of the next page
 * @property {string | null} previous - URL of the previous page
 * @property {number} total_pages - Total number of pages
 * @property {number} total_records - Total number of records
 * @property {ListItem[]} results - List of items
 */
export interface ListInfo {
  message: string;
  next: string | null;
  previous: string | null;
  total_pages: number;
  total_records: number;
  results: ListItem[];
}

/**
 * @type {Object} PersonProperties - Properties of a person
 */
export interface PersonProperties {
  height: string;
  mass: string;
  hair_color: string;
  skin_color: string;
  eye_color: string;
  birth_year: string;
  gender: string;
  created: string;
  edited: string;
  name: string;
  homeworld: string;
  url: string;
}

/**
 * @typedef {Object} Person - A person record
 * @property {string} description - Description of the person
 * @property {PersonProperties} properties - Properties of the person
 * @property {string} uid - Unique identifier of record person
 */
export interface Person {
  description: string;
  properties: PersonProperties;
  uid: string;
} // A person record

/**
 * @typedef {Object} PlanetProperties - Properties of a planet
 */
export interface PlanetProperties {
  diameter: string;
  rotation_period: string;
  orbital_period: string;
  gravity: string;
  population: string;
  climate: string;
  terrain: string;
  surface_water: string;
  created: string;
  edited: string;
  name: string;
  url: string;
}

/**
 * @typedef {Object} Planet - A planet record
 * @property {string} description - Description of the planet
 * @property {PlanetProperties} properties - Properties of the planet
 * @property {string} uid - Unique identifier of record planet
 */
export interface Planet {
  description: string;
  properties: PlanetProperties;
  uid: string;
}

/**
 * @typedef {Object} SearchQuery - Search query params model
 * @property {string} keyword - Keyword to search
 * @property {"people" | "planets"} source - Source to search
 * @property {number} [limit] - Limit of the search results default 15
 * @property {number} [page] - Page number of the search results default 1
 */
export interface SearchQuery {
  keyword: string;
  source: "people" | "planets" | "people,planets";
  limit?: number;
  page?: number;
  sortType?: "name" | "created";
  sortOrder?: "asc" | "desc";
} // Search query params model

// -- Derived types -- //
export type GetAllResponse = APIResponse<ListInfo>; // Both :/people and :/planets from SWAPI returns same model.

export type PlanetByIDResponse = APIResponse<{
  message: string;
  result: Planet;
}>; // Response model of :/planet/{id}

export type PeopleByIDResponse = APIResponse<{
  message: string;
  result: Person;
}>; // Response model of :/people/{id}

export type SearchResponse = APIResponse<{
  // Response model of :/search
  dataSource: string;
  hitCount: {
    totalPagePeople: number;
    totalPagePlanet: number;
    total: number;
  };
  results: {
    people: Person[];
    planets: Planet[];
  };
}>;
