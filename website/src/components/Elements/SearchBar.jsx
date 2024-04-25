import { useState } from "react";
// import { FaSearch } from "react-icons/fa";

import "./SearchBar.css";

export const SearchBar = ({ setResults }) => {
  const [input, setInput] = useState("");

  const [search, setSearch] = useState("");
  // const [results, setResults] = useState([]);
  const [searchinfo, setSearchInfo] = useState({});

  const fetchData = (value) => {
    fetch("https://jsonplaceholder.typicode.com/users")
      .then((response) => response.json())
      .then((json) => {
        const results = json.filter((user) => {
          return (
            value &&
            user &&
            user.name &&
            user.name.toLowerCase().includes(value)
          );
        });
        setResults(results);
      });
  };

  const handleChange = (value) => {
    setInput(value);
    fetchData(value);
  };

  const handleSearch = async (e) => {
    e.preventDefault();
    if (search === "") return;

    const endpoint = `https://en.wikipedia.org/w/api.php?action=query&list=search&prop=info&inprop=url&utf8=&format=json&origin=*&srlimit=20&srsearch=${search}`;

    const response = await fetch(endpoint);

    // console.log(response);

    if (!response.ok) {
      throw Error(response.statusText);
    }

    const json = await response.json();
    console.log(json);

    setResults(json.query.search);
    setSearchInfo(json.query.searchinfo);
  };

  const handleSearchNyar = (value) => {
    fetch(
      `https://en.wikipedia.org/w/api.php?action=query&list=search&prop=info&inprop=url&utf8=&format=json&origin=*&srlimit=20&srsearch=${value}`
    )
      .then((response) => response.json())
      .then((json) => {
        const results = json.filter(() => {
          return value;
          // user &&
          // user.name &&
          // user.name.toLowerCase().includes(value)
        });
        setResults(results);
      });
  };

  const handleChangeNyar = (value) => {
    setSearch(value);
    handleSearch;
  };

  return (
    <div className="">
      {/* <FaSearch id="search-icon" /> */}
      {/* <input
        placeholder="Type to search..."
        value={input}
        onChange={(e) => handleChange(e.target.value)}
      /> */}
      <form className="search-box" onSubmit={handleSearch}>
        <input
          type="search"
          placeholder="Input start page here ..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          // onChange={handleSearch}
          // id="start_page"
          className="mb-8 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-80 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          // required
        />
      </form>
    </div>
  );
};
