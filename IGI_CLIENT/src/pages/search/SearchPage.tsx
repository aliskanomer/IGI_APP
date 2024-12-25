/**
 * Responsibility: Run @function SearchService.search(query) with the query from the context and display the results.
 *                 By default Input field, default filters and search button. First click to search button initiates call and displays the results.
 *                 Each pagination click will update the page number in the context and trigger a new search.
 *                 Search results are displayed in two sections; people and planet by using Card component.
 *                 As a page, violets the default page layout and applies it's own layout. See SearchPage.scss for more details.
 */

import React, { useEffect } from "react";
import { useLocation } from "react-router-dom";
import { useSearch } from "../../context/searchContext";
import { Person, Planet } from "../../api/types";

import SearchService from "../../api/services/searchService";
import SearchFilters from "../../components/searchFilters/SearchFilters";

import Card from "../../components/common/card/Card";
import "./SearchPage.scss";
import Loader from "../../components/common/loader/Loader";

const SearchPage: React.FC = () => {
  // form context
  const { query, updateQuery, validateQuery } = useSearch();
  const location = useLocation();

  // api manager states
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState<any | null>(null);

  const [matchedPeople, setMatchedPeople] = React.useState<Person[]>([]);
  const [matchedPlanets, setMatchedPlanets] = React.useState<Planet[]>([]);
  const [searchMeta, setSearchMeta] = React.useState<any>();

  // pagination state
  const [page, setPage] = React.useState<number>(query.page || 1);
  const [paginated, setPaginated] = React.useState<boolean>(false); // *

  const handleSearch = async () => {
    // validate query abort when invalid
    const validationError = validateQuery();
    if (validationError) {
      setError(validationError);
      // reset states
      setMatchedPeople([]);
      setMatchedPlanets([]);
      setSearchMeta(undefined);
      return;
    }

    // start searching process
    setLoading(true);
    try {
      const result = await SearchService.search(query);
      // parse data to related states for rendering
      setMatchedPeople(result.response.results.people);
      setMatchedPlanets(result.response.results.planets);
      // parse search meta for rendering
      setSearchMeta({
        total: result.response.hitCount.total,
        // single paginator means bigger array gets to be the max limit for pagination
        totalPages:
          result.response.hitCount.totalPagePeople >
          result.response.hitCount.totalPagePlanet
            ? result.response.hitCount.totalPagePeople
            : result.response.hitCount.totalPagePlanet,
      });
      // terminate loading
      setLoading(false);
    } catch (error) {
      // log & set error
      console.error(error);
      setError(error);
      // reset states
      setMatchedPeople([]);
      setMatchedPlanets([]);
      setSearchMeta(undefined);
      // terminate loading
      setLoading(false);
    }
  };

  const onPaginate = (type: "prev" | "next", e: any) => {
    e.stopPropagation();
    if (type === "prev") {
      setPage(page - 1); // decrement page
      updateQuery("page", page - 1); // update query
    }
    if (type === "next") {
      setPage(page + 1); // increment page
      updateQuery("page", page + 1); // update query
    }
    setPaginated(true); // set paginated flag
  };

  useEffect(() => {
    //*
    if (paginated) {
      handleSearch();
      setPaginated(false);
    }
  }, [paginated]);

  useEffect(() => {
    // reset local states when location changes
    setMatchedPeople([]);
    setMatchedPlanets([]);
    setSearchMeta(undefined); // important: pages main render states relies on meta. **
  }, [location]);

  return (
    <div id="searchPage" className="col">
      {!searchMeta && (
        <h1>
          Inter-Galactic <br /> Index
        </h1>
      )}

      {/* Input & Filters & Pagination */}
      <div id="searchForm" className="col">
        {/* Input */}
        <input
          type="text"
          placeholder="Enter keyword"
          value={query.keyword}
          onChange={(e) => updateQuery("keyword", e.target.value)}
          maxLength={100}
          className="input"
        />
        {/* Filters */}
        <SearchFilters />
        {/* Actions */}
        <button
          onClick={handleSearch}
          className="btn btn-secondary"
          disabled={loading || query.keyword.length === 0}
        >
          Search
        </button>
        {/* Pagination */}
        {searchMeta && (
          <div className="row" id="pagination">
            <button
              onClick={(e) => onPaginate("prev", e)}
              disabled={page === 1}
            >
              Prev
            </button>
            <p>
              {page} / {searchMeta?.totalPages}
            </p>
            <button
              onClick={(e) => onPaginate("next", e)}
              disabled={page >= searchMeta?.totalPages}
            >
              Next
            </button>
          </div>
        )}
      </div>

      {/* Loading */}
      {loading && <Loader />}

      {/* Error */}
      {error && (
        <p>Something went wrong. Please check console for more details.</p>
      )}

      {/* Results */}
      {searchMeta && query?.keyword && (
        <div className="search-results">
          {/* People Results */}
          {matchedPeople?.length > 0 && (
            <div className="col search-results--container">
              <div className="col search-results--info">
                <h2>People</h2>
                <p>
                  Displaying people containing the keyword in their name field.
                </p>
              </div>
              <div className="row search-results--content">
                {matchedPeople?.map((person) => (
                  <Card
                    key={person.uid}
                    uid={person.uid}
                    name={person.properties.name}
                    resource="people"
                    details={{
                      year: person.properties.birth_year,
                      gender: person.properties.gender,
                      eyeColor: person.properties.eye_color,
                      hairColor: person.properties.hair_color,
                      skinColor: person.properties.skin_color,
                    }}
                  />
                ))}
              </div>
            </div>
          )}

          {/* Planet Results */}
          {matchedPlanets?.length > 0 && (
            <div className="col search-results--container">
              <div className="col search-results--info">
                <h2> Planets </h2>
                <p>
                  Displaying planets containg the keyword in their name field.
                </p>
              </div>
              <div className="row search-results--content">
                {matchedPlanets.map((planet) => (
                  <Card
                    key={planet.uid}
                    uid={planet.uid}
                    name={planet.properties.name}
                    resource="planet"
                    details={{
                      climate: planet.properties.climate,
                      terrain: planet.properties.terrain,
                      rotationPeriod: planet.properties.rotation_period,
                      orbitalPeriod: planet.properties.orbital_period,
                    }}
                  />
                ))}
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
};

export default SearchPage;

// * A note from the author: paginated boolean flag ensures that there is no call when page mounted and page number is correct when pagination is clicked.
// * react life-cycles can be tricky sometimes especially when dealing with async operations.
// * there is a possibility (not a small one) that the page number on the context remains the old value right after the moment the next or prev button clicked
// * If handleSearch() is directly called on the onPaginate function, and the life-cycle problem occurs, the page number will be wrong on the api call.
// * UI will change and say page 2/3 for example but on the api call page will be 1 and the results will be the same.
// * this is a simple solution to prevent this.

// ** About locational reset: Search represent itself in two ways; page when there is an active search and main page when there is no search.
// ** When location changes, the main page should be displayed. Query by SearchContext resets but the searchMeta does not and it should be reset manually.
// ** Using the menu as an escape; is an expected user behaviour in these kind of applications.
