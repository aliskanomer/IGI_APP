/**
 *    Responsibilty: Renders a 404 page when a route is not found. See routes.tsx for more details
 */

import React from "react";
import { Link } from "react-router-dom";

const NotFoundPage: React.FC = () => {
  return (
    <div className="col">
      <h1>What is that?</h1>
      <p>
        Sorry but the page you are looking for is not found in here. Are you
        sure you are in the right place?
      </p>
      <Link to="/">
        <strong>Back to Home</strong>
      </Link>
    </div>
  );
};

export default NotFoundPage;
