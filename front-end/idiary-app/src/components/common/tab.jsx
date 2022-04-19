import React from "react";
import { NavLink } from "react-router-dom";

const Tab = () => {
  return (
    <ul className="nav nav-tabs justify-content-center">
      <li className="nav-item">
        <NavLink
          className="nav-link active"
          aria-current="page"
          to="/idiary/myDiary/personalDiary"
        >
          My Diary
        </NavLink>
      </li>
      <li className="nav-item">
        <NavLink className="nav-link" to="/idiary/myDiary/friendsDiary">
          Friends
        </NavLink>
      </li>
    </ul>
  );
};

export default Tab;
