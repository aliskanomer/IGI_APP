import axios, { AxiosInstance, AxiosResponse } from "axios";
import { APIResponse, ErrorInfo } from "./types";

/**
 * IgiClient class
 * @class IgiClient - Singleton class for Axios instance to be used on all service layer classes
 * @returns AxiosInstance
 */
class IgiClient {
  private static instance: AxiosInstance;

  private constructor() {}

  // Singleton pattern
  public static getInstance(): AxiosInstance {
    // Create a singleton axios instance to perform HTTP requests
    if (!IgiClient.instance) {
      IgiClient.instance = axios.create({
        baseURL: process.env.IGI_API_DEV || "http://localhost:8080", // red URL from process or fallback to local dev server
        headers: {
          "Content-Type": "application/json",
        },
        timeout: 10000, // time limit for each request is 10 seconds
      });

      // Response interceptor to map API response to APIResponse model and log errors
      IgiClient.instance.interceptors.response.use(
        (response: AxiosResponse<APIResponse<any>>) => response, // *1
        (error) => {
          const errorResponse: APIResponse<ErrorInfo> = error.response
            ?.data || {
            status: error.response?.data?.status || 500,
            message: error.response?.data?.message || "Unknown error occurred.",
            response: {
              code: error.response?.data?.code || "UNKNOWN_ERROR",
              details:
                error.response?.data?.details ||
                "No additional error details available.",
            },
          };
          return Promise.reject(errorResponse);
        }
      );
    }

    return IgiClient.instance;
  }
}

// Export a singleton instance of the Axios instance
export default IgiClient.getInstance();

// *1 (response: AxiosResponse<APIResponse<ANY>>) => Implement dynamic type insertion to achive full type safety.
