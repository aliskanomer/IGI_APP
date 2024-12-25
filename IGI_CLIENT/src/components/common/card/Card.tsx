/**
 *    Responsibilty: Display masked data of resources in a list as a card. Each card contains a name, uid & resource, and details if available.
 *                   Cards open overlay to display further information when clicked.  (see Overlay.tsx)
 *                   In concept; search result cards provide detail but resource routes does not. For taoliring; details are masked in sentences.               
 */

import React, { useEffect } from "react";
import Overlay from "../overlay/Overlay";

import "./Card.scss";

interface CardProps {
  uid: string;
  name: string;
  resource: "people" | "planet";
  details?: {
    [key: string]: string;
  };
}

const Card: React.FC<CardProps> = ({ uid, name, resource, details }) => {
  const [text, setText] = React.useState<{
    subTitle: string;
    details: string;
  }>();
  const [openOverlay, setOpenOverlay] = React.useState<boolean>(false);

  // Card detail text setter
  useEffect(() => {
    if (resource === "people") {
      setText({
        subTitle: `${details?.year} - ${details?.gender}`,
        details: `${details?.eyeColor} eyes with ${details?.hairColor} hair and ${details?.skinColor} skin`,
      });
    }
    if (resource === "planet") {
      setText({
        subTitle: `${details?.climate} - ${details?.terrain}`,
        details: `${details?.rotationPeriod}h rotation period with  ${details?.orbitalPeriod}h orbital period`,
      });
    }
  }, [resource, details]);

  return (
    <div className="col card" key={uid} onClick={() => setOpenOverlay(true)}>
      <h3 className="card-title">{name}</h3>
      {details && (
        <div className="col card-details">
          <h4 className="card-details--subtitle">{text?.subTitle}</h4>
          <p className="card-details--text">{text?.details}</p>
        </div>
      )}
      {openOverlay && (
        <Overlay
          name={name}
          uid={uid}
          resource={resource}
          open={openOverlay}
          onClose={(e) => {
            e.stopPropagation();
            setOpenOverlay(false);
          }}
        />
      )}
    </div>
  );
};

export default Card;
