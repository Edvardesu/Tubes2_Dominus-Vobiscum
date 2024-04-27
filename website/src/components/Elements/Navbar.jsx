import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav
      className="z-10 w-full flex items-center justify-between flex-wrap p-6 sticky top-0 border-b-2 border-yellow-700 bg-black"
    >
      <div className="flex items-center flex-shrink-0 text-white">
        <span className="font-semibold text-3xl text-yellow-300 tracking-tight">DOMINVS VOBISCVM</span>
      </div>
      <div className="block lg:hidden">
        <button className="flex items-center px-3 py-2 border rounded text-teal-200 border-teal-400 hover:text-white hover:border-white">
          <svg
            className="fill-current h-3 w-3"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <title>Menu</title>
            <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
          </svg>
        </button>
      </div>
      <div className="w-full block lg:flex lg:items-center lg:w-auto text-white">
        <div className="text-xl mr-5 lg:flex-grow font-bold">
          {/* Replace anchor elements with Link components */}
          <Link to="/" className="block mb-4 lg:inline-block lg:mt-0 text-white hover:text-yellow-300 mr-10 no-underline">
            Main
          </Link>
          <Link to="/konsep" className="block mb-4 lg:inline-block lg:mt-0 text-white hover:text-yellow-300 mr-10 no-underline">
            Konsep Search Engine
          </Link>
          <Link to="/howToUse" className="block mt-4 lg:inline-block lg:mt-0 text-white hover:text-yellow-300 mr-10 no-underline">
            How to Use
          </Link>
          <Link to="/aboutUs" className="block mt-4 lg:inline-block lg:mt-0 text-white hover:text-yellow-300 mr-10 no-underline">
            About us
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;