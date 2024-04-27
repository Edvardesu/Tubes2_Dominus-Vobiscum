import { useState } from "react";
import styled from "styled-components";
import "./bg.css";
import SearchBarGenahStart from "./SearchBarGenahStart";
import SearchBarGenahDest from "./SearchBarGenahDest";

function clickMe() {
  alert("You clicked me!");
}

const theme = {
  blue: {
    default: "#3f51b5",
    hover: "#283593",
  },
  pink: {
    default: "#e91e63",
    hover: "#ad1457",
  },
};

const Button = styled.button`
  background-color: ${(props) => theme[props.theme].default};
  color: white;
  padding: 5px 15px;
  border-radius: 5px;
  outline: 0;
  border: 0;
  text-transform: uppercase;
  margin: 10px 0px;
  cursor: pointer;
  box-shadow: 0px 2px 2px lightgray;
  transition: ease background-color 250ms;
  &:hover {
    background-color: ${(props) => theme[props.theme].hover};
  }
  &:disabled {
    cursor: default;
    opacity: 0.7;
  }
`;

Button.defaultProps = {
  theme: "blue",
};

const CucakRowo = (props) => {
  const [search, setSearch] = useState("");
  // const [results, setResults] = useState([]);
  const [searchinfo, setSearchInfo] = useState({});
  const [results, setResults] = useState([]);
  const [inputStart, setInputStart] = useState("");
  const [inputEnd, setInputEnd] = useState("");
  const [start, setStart] = useState("");
  const [destination, setDestination] = useState("");
  const [responseOutput, setResponseOutput] = useState(""); // State to store the backend response
  const [singlePath, setSinglePath] = useState(false); // false as the default value

  const handleSubmit = async (event) => {
    event.preventDefault();

    const requestData = {
      start: start,
      destination: destination,
      single_path: singlePath, // Add this line to include the toggle state in the request
    };

    const response = await fetch("http://localhost:8080/uploadids", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(requestData),
    });

    const data = await response.json();
    setResponseOutput(JSON.stringify(data)); // Store the response data as a string in state
    console.log(data);
  };

  return (
    <div className="py-20 w-full h-full flex flex-col justify-between">
      <div className="background">
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>

        <div className="flex flex-col mt-20 overflow-x-auto ">
          <p className="text-6xl mb-4 font-bold text-center text-white relative">
            WIKIRACE SOLVER
          </p>
          <p className="text-3xl font-bold mb-10 text-center text-yellow-400 relative">
            DOMINVS VOBISCVM
          </p>
          <div className="mt-4 relative">
            <label
              htmlFor="single_path_toggle"
              className="text-center text-xl block mb-2 font-medium text-white relative"
            >
              Single Path Only
            </label>
            <div className="text-center">
              <input
                id="single_path_toggle"
                type="checkbox"
                checked={singlePath}
                onChange={() => setSinglePath(!singlePath)}
                className="w-6 h-6 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 focus:ring-2"
              />
            </div>
          </div>
          <div className="flex flex-row mt-16 items-center justify-center gap-4 ">
            {/* GOES HERE */}
            <form onSubmit={handleSubmit}>
              <div className="">
                <label
                  htmlFor="start_page"
                  className="text-center text-xl block mb-2 font-medium text-white relative"
                >
                  Start Page
                </label>
                <div className="h-full text-black">
                  <div className="search-bar-container">
                    <SearchBarGenahStart
                      onSelect={(title) => setStart(title)}
                    />
                  </div>
                </div>
              </div>

              <div className="">
                <label
                  htmlFor="final_page"
                  className="text-center text-xl block mb-2 font-medium text-white relative"
                >
                  Final Page
                </label>
                <div className="h-full text-black">
                  <div className="search-bar-container">
                    <SearchBarGenahDest
                      onSelect={(title) => setDestination(title)}
                    />
                  </div>
                </div>
              </div>

              <div className="mt-8 text-center items-center justify-center gap-4 relative">
                <button
                  type="submit"
                  className="text-3xl bg-red-500 hover:bg-red-700 text-white font-semibold py-8 px-20 rounded-full relative border border-white hover:border-transparent"
                >
                  SIKATTT !!!
                </button>
              </div>
            </form>
          </div>
          {responseOutput && (
            <div className="response-output text-center mt-16 text-xl text-white relative">
              {responseOutput}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default CucakRowo;

{
  /* <label>
                Start:
                <input
                  type="text"
                  value={start}
                  onChange={(e) => setStart(e.target.value)}
                />
              </label>
              <label>
                Destination:
                <input
                  type="text"
                  value={destination}
                  onChange={(e) => setDestination(e.target.value)}
                />
              </label> */
}
{
  /* <div className="bg-white text-center">
                <button type="submit">Submit</button>
              </div> */
}

{
  /* <div className="">
              <label
                htmlFor="start_page"
                className="text-center text-xl block mb-2 font-medium text-white relative"
              >
                Start Page
              </label>
              <div className="h-full text-black">
                <div className="search-bar-container">
                  <SearchBarGenahStart />
                </div>
              </div>
            </div>

            <div className="">
              <label
                htmlFor="final_page"
                className="text-center text-xl block mb-2 font-medium text-white relative"
              >
                Final Page
              </label>
              <div className="h-full text-black">
                <div className="search-bar-container">
                  <SearchBarGenahDest />
                </div>
              </div>
            </div> */
}

{
  /* <div className="flex flex-row">
                  <a
                    href="#"
                    className="inline-block text-xl align-middle px-6 py-2 mx-2 leading-none border rounded-lg text-black border-white hover:border-transparent hover:text-white hover:bg-yellow-500 bg-white font-semibold"
                    onClick={handleColorSearch}
                  >
                    Color
                  </a>
                  <a
                    href="#"
                    className="inline-block text-xl align-middle px-6 py-2 leading-none border rounded-lg text-black border-white hover:border-transparent hover:text-white hover:bg-yellow-500 bg-white font-semibold"
                  >
                    Texture
                  </a>
                </div> */
}

{
  /* <form className="search-box" onSubmit={handleSearch}>
                <label
                  htmlFor="start_page"
                  className="text-center text-xl block mb-2 font-medium text-white"
                >
                  Start Page
                </label>
                <input
                  type="search"
                  placeholder="Input start page here ..."
                  value={search}
                  onChange={(e) => setSearch(e.target.value)}
                  // id="start_page"
                  className="mb-8 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-80 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  // required
                />
              </form> */
}
{
  /* <p className="text-center text-white">Search result</p> */
}
{
  /* {searchinfo.totalhits ? (
                <p className="text-center text-white">
                  Search results: {searchinfo.totalhits}
                </p>
              ) : (
                ""
              )} */
}
{
  /* <div className="results">
                {results.map((result, i) => {
                  const url = `https://en.wikipedia.org/?curid=${result.pageid}`;

                  return (
                    <>
                      <div className="result" key={i}>
                        <h3 className="text-white">{result.title}</h3>
                        <p
                          dangerouslySetInnerHTML={{ __html: result.snippet }}
                          className="text-white"
                        ></p>
                        <a
                          className="text-white"
                          href={url}
                          target="_blank"
                          rel="noreferrer"
                        >
                          Read more
                        </a>
                      </div>
                    </>
                  );
                })}
              </div> */
}
