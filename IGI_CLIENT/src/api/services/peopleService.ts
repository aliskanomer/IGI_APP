import IgiClient from "../index";
import { GetAllResponse, PeopleByIDResponse } from "../types";

/**
 * PeopleService class
 * @class PeopleService - Singleton service class for people
 * @method getAll - HTTP_GET: /people
 * @method getById - HTTP_GET: /people/{id}
 * @returns PeopleService instance
 */
class PeopleService {
  private static instance: PeopleService;

  // define and store endpoints
  private endpoints: {
    root: string;
    getById: (id: string) => string;
  };
  private constructor() {
    this.endpoints = {
      root: "/people",
      getById: (id: string) => `/people/${id}`,
    };
  }

  // Singleton pattern
  public static getInstance(): PeopleService {
    if (!PeopleService.instance) {
      PeopleService.instance = new PeopleService();
    }
    return PeopleService.instance;
  }

  /**
   * HTTP_GET: /people
   * @param page - number (def1)
   * @param limit - number (def15)
   * @returns GetAllResponse type
   */
  public async getAll(
    page: number = 1,
    limit: number = 15
  ): Promise<GetAllResponse> {
    // runtime validation of query params
    if ((page && page <= 0) || (limit && limit <= 0)) {
      throw new Error("Page and/or Limit must be bigger than zero!");
    }

    // make HTTP request
    const response = await IgiClient.get<GetAllResponse>(this.endpoints.root, {
      params: { page, limit },
    });
    return response.data;
  }

  /**
   * HTTP_GET: /people/{id}
   * @param id - string
   * @returns PeopleByIDResponse type
   */
  public async getById(id: string): Promise<PeopleByIDResponse> {
    // make HTTP request
    const response = await IgiClient.get<PeopleByIDResponse>(
      this.endpoints.getById(id)
    );
    return response.data;
  }
}
// Export a singleton instance of the service class
export default PeopleService.getInstance();
