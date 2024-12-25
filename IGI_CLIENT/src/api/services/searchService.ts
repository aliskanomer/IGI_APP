import IgiClient from "../index";
import { SearchResponse, SearchQuery } from "../types";

/**
 * SearchService class
 * @class SearchService - Singleton service class for search
 * @method search - HTTP_GET: /search
 * @returns SearchService instance
 */
class SearchService {
  private static instance: SearchService;

  // define and store endpoints
  private endpoints: {
    search: string;
  };
  private constructor() {
    this.endpoints = {
      search: "/search",
    };
  }

  // Singleton pattern
  public static getInstance(): SearchService {
    if (!SearchService.instance) {
      SearchService.instance = new SearchService();
    }
    return SearchService.instance;
  }

  /**
   * HTTP_GET: /search
   * @param query - SearchQuery
   * @returns SearchResponse type
   */
  public async search(query: SearchQuery): Promise<SearchResponse> {
    // runtime validation of query params *
    if (!query.keyword || query.keyword.trim() === "") {
      throw new Error("Keyword is required and cannot be empty!");
    }
    if (
      !query.source ||
      !["people", "planets", "people,planets"].includes(query.source)
    ) {
      throw new Error("Source must be either 'people' or 'planets'!");
    }
    if ((query.page && query.page <= 0) || (query.limit && query.limit <= 0)) {
      throw new Error("Page and/or Limit must be bigger than zero!");
    }
    if (query.sortType && !["name", "created"].includes(query.sortType)) {
      throw new Error("Invalid sort type. Must be either 'name' or 'created'.");
    }
    if (query.sortOrder && !["asc", "desc"].includes(query.sortOrder)) {
      throw new Error("Invalid sort order. Must be either 'asc' or 'desc'.");
    }

    // build URL with default or provided query params
    const params = new URLSearchParams();
    params.append("keyword", query?.keyword);
    if (query?.source === "people,planets") { // not the best but easiest way to handle %2C in our case
      params.append("source", "people,planets");
    }else{
      params.append("source", query?.source);
    }
    params.append("page", query.page?.toString() || "1");
    params.append("limit", query.limit?.toString() || "15");
    params.append("sortType", query?.sortType || "name");
    params.append("sortOrder", query?.sortOrder || "asc");
    const URL = `${this.endpoints.search}?${params.toString()}`;

    // make HTTP request
    const response = await IgiClient.get<SearchResponse>(URL);

    return response.data;
  }
}

// Export a singleton instance of the service class
export default SearchService.getInstance();

// * both components by elements and context, client invoker by itself and IGI_API runs parameter validations.
//  This is a good practice to avoid unnecessary API calls and to prevent errors.
