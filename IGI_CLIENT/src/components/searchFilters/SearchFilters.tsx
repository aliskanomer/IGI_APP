/**
 *        Responsbility: Displays a search filter menu with options to filter and sort search results
 *                       Contains a button and a form field. Acts as dropdown for search filter.
 *                       I/O element events - updates query state by @context SearchContext
 *                       Thus preventing prop drilling within the application
 */

import React from "react";
import { useSearch } from "../../context/searchContext";

import "./SearchFilters.scss";

const SearchFilters: React.FC = () => {
  const { query, updateQuery } = useSearch();
  const [showFilters, setShowFilters] = React.useState(false);

  const handleCheckboxChange = (value: "people" | "planets") => {
    const sources = query.source.split(",");
    if (sources.includes(value)) {
      if (sources.length > 1) {
        updateQuery(
          "source",
          sources.filter((s) => s !== value).join(",") as any
        );
      }
    } else {
      updateQuery("source", [...sources, value].join(",") as any);
    }
  };

  return (
    <div id="searchFilter" className="col filter">
      <button
        onClick={() => setShowFilters(!showFilters)}
        className={showFilters ? "btn btn-tertiary" : "btn btn-primary"}
      >
        Filters
      </button>

      {showFilters && (
        <div className="row filter-menu">
          {/* Source Checkboxes */}
          <div className="col filter-menu--group">
            <div className="col filter-menu--group_info">
              <h3>Resource</h3>
              <p>
                Set the resources to be checked for items that include the
                provided keyword.
              </p>
            </div>

            <label className="label">
              <input
                type="checkbox"
                className="input"
                checked={query.source.includes("people")}
                onChange={() => handleCheckboxChange("people")}
              />
              People
            </label>
            <label className="label">
              <input
                type="checkbox"
                className="input"
                checked={query.source.includes("planets")}
                onChange={() => handleCheckboxChange("planets")}
              />
              Planets
            </label>
          </div>

          {/* Sorting Options */}
          <div className="col filter-menu--group">
            <div className="col filter-menu--group_info">
              <h3>Sorting</h3>
              <p>
                Select a sort field and sorting order to see more organized
                results.
              </p>
            </div>

            <div className="row">
              <div className="col filter-menu--group">
                <h4>Type</h4>
                <label className="label">
                  <input
                    type="radio"
                    name="sortType"
                    value="name"
                    className="input"
                    checked={query.sortType === "name"}
                    onChange={() => updateQuery("sortType", "name")}
                  />
                  Name
                </label>
                <label>
                  <input
                    type="radio"
                    className="input"
                    name="sortType"
                    value="created"
                    checked={query.sortType === "created"}
                    onChange={() => updateQuery("sortType", "created")}
                  />
                  Created
                </label>
              </div>

              {/* SortOrder */}
              <div className="col filter-menu--group">
                <h4> Order </h4>
                <label className="label">
                  <input
                    type="radio"
                    name="sortOrder"
                    value="asc"
                    className="input"
                    checked={query.sortOrder === "asc"}
                    onChange={() => updateQuery("sortOrder", "asc")}
                  />
                  Ascending
                </label>
                <label className="label">
                  <input
                    type="radio"
                    name="sortOrder"
                    value="desc"
                    className="input"
                    checked={query.sortOrder === "desc"}
                    onChange={() => updateQuery("sortOrder", "desc")}
                  />
                  Descending
                </label>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default SearchFilters;
