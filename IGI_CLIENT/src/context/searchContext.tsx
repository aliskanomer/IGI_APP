/**
 *    Responsibilty: Provides a @context SearchContext for search query management.
 *                   @function updateQuery(SearchQuery) - Updates the provided key value pair on context query
 *                   @function resetQuery() - Resets the query to default values. Invoked by location change to prevent stale data.
 *                   @function validateQuery() - Validates the query and returns the first error as @object {field,message} if any. Returns null if query is valid.
 */
import React, {
  createContext,
  useContext,
  ReactNode,
  useState,
  useEffect,
} from "react";
import { SearchQuery } from "../api/types";
import { useLocation } from "react-router-dom";

/**
 * @type {Object} SearchContextType - Search context model
 * @property {SearchQuery} query - Current search query state
 * @property {(field: keyof SearchQuery, value: SearchQuery[Key]) => void} updateQuery - Update query state
 * @property {() => void} resetQuery - Reset query to default values
 */
interface SearchContextType {
  query: SearchQuery;
  updateQuery: <Key extends keyof SearchQuery>(
    field: Key,
    value: SearchQuery[Key]
  ) => void;
  resetQuery: () => void;
  validateQuery: () => { field: keyof SearchQuery; message: string } | null;
}

// Default search query values
const defaultQuery: SearchQuery = {
  keyword: "",
  source: "people,planets",
  page: 1,
  limit: 15,
  sortType: "name",
  sortOrder: "asc",
};

// Context
const SearchContext = createContext<SearchContextType | undefined>(undefined);

// Provider
export const SearchProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [query, setQuery] = useState<SearchQuery>(defaultQuery);
  const location = useLocation(); // Get current location

  // Update query by providing field and value
  const updateQuery = <Key extends keyof SearchQuery>(
    field: Key,
    value: SearchQuery[Key]
  ) => {
    setQuery((prev) => ({
      ...prev,
      [field]: value,
    }));
  };

  const validateQuery = (): {
    field: keyof SearchQuery;
    message: string;
  } | null => {
    if (!query.keyword.trim() || query.keyword.length > 100) {
      return {
        field: "keyword",
        message: "Keyword must be between 1 and 100 characters.",
      };
    }
    if (
      !query.source ||
      !["people", "planets", "people,planets"].includes(query.source)
    ) {
      return { field: "source", message: "Invalid source." };
    }
    if (query.sortType && !["name", "created"].includes(query.sortType)) {
      return { field: "sortType", message: "Invalid sort type." };
    }
    if (query.sortOrder && !["asc", "desc"].includes(query.sortOrder)) {
      return { field: "sortOrder", message: "Invalid sort order." };
    }
    return null; // Validasyon başarılıysa null döner
  };

  // Reset query to default values
  const resetQuery = () => {
    setQuery(defaultQuery);
  };

  useEffect(() => {
    resetQuery(); // Reset query on location change
  }, [location]); 

  // Provide query, updateQuery and resetQuery to children for easy access to query on all levels
  return (
    <SearchContext.Provider
      value={{ query, updateQuery, resetQuery, validateQuery }}
    >
      {children}
    </SearchContext.Provider>
  );
};

// Custom hook to validate context usage and return context values
export const useSearch = () => {
  const context = useContext(SearchContext);
  if (!context) {
    throw new Error("useSearch must be used within a SearchProvider");
  }
  return context;
};
