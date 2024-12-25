import IgiClient from "../index";
import { GetAllResponse, PlanetByIDResponse } from "../types";

/**
 * PlanetService class
 * @class PlanetService - Singleton service class for planets
 * @method getAll - HTTP_GET: /planet
 * @method getById - HTTP_GET: /planet/{id}
 * @returns PlanetService instance
 */
class PlanetService {
  private static instance: PlanetService;

  // define and store endpoints
  private endpoints: {
    root: string;
    getById: (id: string) => string;
  };
  private constructor() {
    this.endpoints = {
      root: "/planet",
      getById: (id: string) => `/planet/${id}`,
    };
  }

  // Singleton pattern
  public static getInstance(): PlanetService {
    if (!PlanetService.instance) {
      PlanetService.instance = new PlanetService();
    }
    return PlanetService.instance;
  }

  /**
   * HTTP_GET : /planet
   * @param page - number (def1)
   * @param limit - number (def15)
   * @returns GetAllResponse type
   * */
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
    }); //*
    return response.data;
  }

  /**
   * HTTP_GET : /planet/{id}
   * @param id - ID of the planet
   * @returns PlanetByIDResponse type
   * */
  public async getById(id: string): Promise<PlanetByIDResponse> {
    const response = await IgiClient.get<PlanetByIDResponse>(
      this.endpoints.getById(id)
    );
    return response.data;
  }
}

// Export a singleton instance of the service class
export default PlanetService.getInstance();

// * Dont wrap up the awaits with try-catch. Let the caller handle the error. Also axios interceptor will catch the error and return a rejected promise in any case of error.Check documentation for more details.
