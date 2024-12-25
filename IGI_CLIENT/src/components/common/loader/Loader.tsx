/**
 *       Responsbility: Animate a loader with 3 dots when invoked in given color.
 *                      cover property can be used to make the loader take up the available space.
 */
import React from "react";
import "./Loader.scss";

interface LoaderProps {
  color?: "primary" | "secondary" | "tertiary";
  cover?: boolean; // rather the loader should take up the available space
}

const Loader: React.FC<LoaderProps> = ({ color = "secondary", cover = false }) => {
  return (
    <div className="loader-container loader-cover">
      {Array.from({ length: 3 }).map((_, index) => (
        <div key={index} className={`loader-dot loader-dot-${color}`} />
      ))}
    </div>
  );
};

export default Loader;
