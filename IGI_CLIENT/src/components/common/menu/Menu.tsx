/**
 *      Responsiblity: Navigation of the application. Root is search page.
 */

import React from "react";
import { Link } from "react-router-dom";
import "./Menu.scss";

const Menu: React.FC = () => {
  return (
    <nav id="navMenu" className="row">
      <ul id="menu" className="row">
        <li>
          <Link to="/">Search</Link>
        </li>
        <li id="logo">
          <Link to="/">IGI</Link>
        </li>
        <li>
          <Link to="/planet">Planet</Link>
        </li>
        <li>
          <Link to="/people">People</Link>
        </li>
      </ul>
    </nav>
  );
};

export default Menu;
