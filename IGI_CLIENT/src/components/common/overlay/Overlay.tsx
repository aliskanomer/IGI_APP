/**
 *       Responsbility : Opens an overlay on top of DOM to display detailed information about a resource
 *                       Depending on resource type invokes @function PlanetService.getById or @function PeopleService.getById
 *                       Relies on two-way binding to be opened and closed
 *                       @function propKeyFormatter formats the property key (eye_color -> Eye Color | name -> Name)
 *                       @function propValueFormatter formats the property value for special cases (dates and urls)
 */

import React, { JSX, useEffect, useState } from "react";
import { PlanetByIDResponse, PeopleByIDResponse } from "../../../api/types";

import peopleService from "../../../api/services/peopleService";
import planetService from "../../../api/services/planetService";

import "./Overlay.scss";
import Loader from "../loader/Loader";

interface OverlayProps {
  name: string;
  uid: string;
  resource: "people" | "planet";
  open: boolean;
  onClose: (e: any) => void;
}

const Overlay: React.FC<OverlayProps> = ({
  name,
  uid,
  resource,
  open,
  onClose,
}) => {
  const [data, setData] = useState<PlanetByIDResponse | PeopleByIDResponse>();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(false);

  const fetchData = async () => {
    setLoading(true);
    try {
      if (resource === "people") {
        const response = await peopleService.getById(uid);
        setLoading(false);
        setData(response);
      } else {
        const response = await planetService.getById(uid);
        setLoading(false);
        setData(response);
      }
    } catch (error) {
      setLoading(false);
      setError(true);
    }
  };

  // Fetch data on mount
  useEffect(() => {
    if (open) {
      fetchData();
    } else {
      setData(undefined);
      setLoading(true);
      setError(false);
    }
  }, [uid, resource, open]);

  // format the property key (eye_color -> Eye Color | name -> Name)
  const propKeyFormatter = (key: string): string => {
    return key
      .split("_")
      .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
      .join(" ");
  };

  // format the property value for special cases (dates and urls)
  const propValueFormatter = (
    key: string,
    value: any
  ): JSX.Element | string => {
    // safe-fail for unexpected values
    if (typeof value === "string") {
      // URL value to clickable link
      if (value.startsWith("http://") || value.startsWith("https://")) {
        return (
          <a
            href={value}
            target="_blank"
            rel="noopener noreferrer"
            className="property-link"
          >
            {propKeyFormatter(key)}
          </a>
        );
      }

      // Date string value to formatted date
      if (value.match(/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}/)) {
        const date = new Date(value);
        if (!isNaN(date.getTime())) {
          return `${date.getFullYear()}.${String(date.getMonth() + 1).padStart(
            2,
            "0"
          )}.${String(date.getDate()).padStart(2, "0")} - ${String(
            date.getHours()
          ).padStart(2, "0")}:${String(date.getMinutes()).padStart(2, "0")}`;
        }
      }
    }
    return value;
  };

  if (!open) return null;

  return (
    <div id={`overlay-${uid}`} className="overlay">
      <div className="overlay-close">
        <button
          id={`close-overlay-${uid}`}
          className="btn btn-primary"
          onClick={onClose}
        >
          X
        </button>
      </div>

      <div id={`overlay-content-${uid}`} className="overlay-container">
        <div className="overlay-title">
          <h1>{name}</h1>
        </div>
        {loading && <Loader color="primary" cover />}
        {error && <h2> Oh No!</h2>}
        <div className="overlay-content">
          {data?.response?.result?.properties &&
            Object.entries(data.response.result.properties).map(
              ([key, value]) => (
                <div key={key} className="table-row">
                  <p className="table-col">{propKeyFormatter(key)}</p>
                  <p className="table-col">{propValueFormatter(key, value)}</p>
                </div>
              )
            )}
        </div>
      </div>
    </div>
  );
};

export default Overlay;
