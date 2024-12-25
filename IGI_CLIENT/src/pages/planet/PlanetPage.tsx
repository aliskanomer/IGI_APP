/**
 *    Responsibilty: When routed, automatically fetches Planet data by @function PlanetService.getAll(page) and renders it by cards
 *                   Results are paginated by 15(default limit on calls) and can be navigated by Next and Prev buttons
 *                   Each pagination creates a new request.
 *                   Follows the default page layout. See _layout.scss for more details
 */

import React, { useEffect, useState } from "react";
import { GetAllResponse, ListItem } from "../../api/types";
import PlanetSVG from "../../assets/planet.svg";

import PlanetService from "../../api/services/planetService";
import Card from "../../components/common/card/Card";
import Loader from "../../components/common/loader/Loader";

const PlanetPage: React.FC = () => {
  // api manager states
  const [planets, setPlanets] = useState<GetAllResponse>();
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<any | null>(null);

  // pagination states
  const [page, setPage] = useState<number>(1);

  const fetchAllPlanets = async (_page: number) => {
    // Reset states before fetching data
    setLoading(true);
    setError(null);
    setPlanets(undefined);

    try {
      const response: GetAllResponse = await PlanetService.getAll(_page);
      setPlanets(response); // set state with fetched data
    } catch (err: any) {
      console.error(err); // log error for debugging
      setError(err); // service error occured
    } finally {
      // Reset loading state
      setLoading(false);
    }
  };

  // Fetch data on mount
  useEffect(() => {
    fetchAllPlanets(page);
  }, [page]);

  return (
    <div className="row page">
      <div className="col page-info">
        <img src={PlanetSVG}></img>
        <h2>Inter-Galactic Index</h2>
        <h1>Planet</h1>
        <p>
          Here lies the list of planets of known galaxy for you. If you have a
          special place in mind; please use search to directly address it.
        </p>
      </div>
      {loading && <Loader cover />}

      {error && (
        <p>Something went wrong. Please check console for more details.</p>
      )}

      {!loading && !error && planets && (
        <div className="col page-content">
          <div className="row page-content--list">
            {planets?.response?.results.map((planet: ListItem) => {
              return (
                <Card
                  key={planet?.uid}
                  uid={planet?.uid}
                  name={planet?.name}
                  resource="planet"
                />
              );
            })}
          </div>
          <div className="row page-content--pagination">
            <button onClick={() => setPage(page - 1)} disabled={page <= 1}>
              Prev
            </button>
            <p>
              {page} / {planets?.response?.total_pages}
            </p>
            <button
              onClick={() => setPage(page + 1)}
              disabled={page >= (planets?.response?.total_pages || 1)}
            >
              Next
            </button>
          </div>
        </div>
      )}
    </div>
  );
};

export default PlanetPage;
