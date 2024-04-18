// import React from "react";
import PropTypes from "prop-types";

function Table(props) {
  return (
    <div>
      <table className="table text-center">
        <thead>
          <tr>
            <th className="p-3">id</th>
            <th className="p-3">Car</th>
            <th className="p-3">Color</th>
          </tr>
        </thead>
        <tbody>
          {props.users.map((user) => (
            <tr key={user.id}>
              <td>{user.id}</td>
              <td>{user.car}</td>
              <td>{user.color}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

Table.propTypes = {
  users: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number.isRequired,
      car: PropTypes.string.isRequired,
      color: PropTypes.string.isRequired,
    })
  ).isRequired,
};

export default Table;
