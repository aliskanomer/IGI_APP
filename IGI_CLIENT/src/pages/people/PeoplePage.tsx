/**
 *    Responsibilty: When routed, automatically fetches Planet data by @function PeopleService.getAll(page) and renders it by cards
 *                   Results are paginated by 15(default limit on calls) and can be navigated by Next and Prev buttons
 *                   Each pagination creates a new request.
 *                   Follows the default page layout. See _layout.scss for more details
 */

import React, { useEffect, useState } from "react";
import { GetAllResponse, ListItem } from "../../api/types";
import PeopleSVG from "../../assets/people.svg";

import PeopleService from "../../api/services/peopleService";
import Card from "../../components/common/card/Card";
import Loader from "../../components/common/loader/Loader";

const PeoplePage: React.FC = () => {
  // api manager states
  const [people, setPeople] = useState<GetAllResponse>();
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<any | null>(null);

  // pagination states
  const [page, setPage] = useState<number>(1);

  const fetchAllPeople = async (_page: number) => {
    // Reset states before fetching data
    setLoading(true);
    setError(null);
    setPeople(undefined);

    try {
      const response: GetAllResponse = await PeopleService.getAll(_page);
      setPeople(response); // set state with fetched data
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
    fetchAllPeople(page);
  }, [page]);

  return (
    <div className="row page">
      <div className="col page-info">
        <img src={PeopleSVG}></img>
        <h2>Inter-Galactic Index</h2>
        <h1>People</h1>
        <p>
          Republic registered people from different species are listed in here.
          If you are looking for someone specific please use search. May the
          force be with all of us.
        </p>
      </div>

      {loading && <Loader cover />}

      {error && (
        <p>Something went wrong. Please check console for more details.</p>
      )}

      {!loading && !error && people && (
        <div className="col page-content">
          <div className="row page-content--list">
            {people?.response?.results.map((person: ListItem) => {
              return (
                <Card
                  key={person?.uid}
                  uid={person?.uid}
                  name={person?.name}
                  resource="people"
                />
              );
            })}
          </div>
          <div className="row page-content--pagination">
            <button onClick={() => setPage(page - 1)} disabled={page <= 1}>
              Prev
            </button>
            <p>
              {page} / {people?.response?.total_pages}
            </p>
            <button
              onClick={() => setPage(page + 1)}
              disabled={page >= (people?.response?.total_pages || 1)}
            >
              Next
            </button>
          </div>
        </div>
      )}
    </div>
  );
};

export default PeoplePage;
