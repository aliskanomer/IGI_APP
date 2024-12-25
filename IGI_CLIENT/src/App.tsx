import React, { Suspense } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

// ROUTES
const People = React.lazy(() => import("./pages/people/PeoplePage"));
const Planet = React.lazy(() => import("./pages/planet/PlanetPage"));
const Search = React.lazy(() => import("./pages/search/SearchPage"));
const NotFound = React.lazy(() => import("./pages/notFound/NotFound"));

// ERROR HANDLING
import ErrorBoundary from "./components/errorBoundry/ErrorBoundry";

// COMPONENTS
import Menu from "./components/common/menu/Menu";
import Loader from "./components/common/loader/Loader";

// CONTEXT
import { SearchProvider } from "./context/searchContext";
import "./assets/styles/styles.scss";

const App: React.FC = () => {
  return (
    <Router>
      <ErrorBoundary>
        <SearchProvider>
          <Menu />
          <Suspense fallback={<Loader/>}>
            <Routes>
              <Route path="/" element={<Search />} />
              <Route path="/people" element={<People />} />
              <Route path="/planet" element={<Planet />} />
              <Route path="*" element={<NotFound />} />
            </Routes>
          </Suspense>
        </SearchProvider>
      </ErrorBoundary>
    </Router>
  );
};

export default App;
