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

const ResponseOutput = styled.div`
  display: block; // Ensures the div is a block element, making children stack vertically
  background-color: white;
  color: black;
  padding: 20px;
  margin-top: 8px;
  text-align: center;

  li {
    display: block; // Makes each list item a block, ensuring they stack vertically
    margin-bottom: 4px; // Adds space between items
  }
`;

const ToggleContainer = styled.div`
  position: relative;
  width: 50px;
  height: 25px;
  background-color: ${(props) => (props.checked ? "#4CAF50" : "#ccc")};
  border-radius: 25px;
  transition: background-color 0.3s;
`;

const Toggle = styled.span`
  position: absolute;
  top: 2px;
  left: ${(props) => (props.checked ? "26px" : "2px")};
  width: 21px;
  height: 21px;
  border-radius: 50%;
  background-color: white;
  transition: left 0.3s;
`;

Button.defaultProps = {
  theme: "blue",
};

const CucakRowo = (props) => {
  const [method, setMethod] = useState("uploadbfs"); // State to control which method is used
  const [search, setSearch] = useState("");
  // const [results, setResults] = useState([]);
  const [searchinfo, setSearchInfo] = useState({});
  const [results, setResults] = useState([]);
  const [inputStart, setInputStart] = useState("");
  const [inputEnd, setInputEnd] = useState("");
  const [start, setStart] = useState("");
  const [destination, setDestination] = useState("");
  const [responseOutput, setResponseOutput] = useState({
    paths: [],
    extraData: {},
  });
  const [singlePath, setSinglePath] = useState(false); // false as the default value
  const [isLoading, setIsLoading] = useState(false); // New state for loading indicator

  const handleSubmit = async (event) => {
    event.preventDefault();

    // Clear previous results before fetching new data
    setResponseOutput({ paths: [], extraData: {} });
    setIsLoading(true); // Set loading to true before the request

    const requestData = {
      start: start,
      destination: destination,
      single_path: singlePath,
    };

    try {
      const response = await fetch(`http://localhost:8080/${method}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestData),
      });

      const data = await response.json();
      console.log("Received data:", data);

      if (data && Array.isArray(data.paths)) {
        setResponseOutput({
          paths: data.paths,
          extraData: {
            singlePath: data.single_path,
            totalLinks: data.total_links, // from IDS
            pathLength: data.path_length, // from IDS
            execTime: data.exec_time, // from IDS
            pathAmount: data.path_amount, // from IDS
          },
        });
      } else {
        setResponseOutput({ paths: [], extraData: {} });
      }
    } catch (error) {
      console.error("Failed to fetch paths:", error);
      setResponseOutput({ paths: [], extraData: {} });
    }
    setIsLoading(false); // Set loading to false after the fetch completes
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

        <div className="h-full flex flex-col overflow-x-auto">
          <div className="pt-40 bg-black bg-opacity-80 relative">
            <p className="text-6xl mb-4 font-bold text-center text-white relative">
              WIKIRACE SOLVER
            </p>
            <p className="text-3xl font-bold mb-10 text-center text-yellow-400 relative">
              DOMINVS VOBISCVM
            </p>
          </div>
          <div className="flex flex-row mt-16 items-center justify-center gap-4 ">
            {/* GOES HERE */}
            <form onSubmit={handleSubmit}>
              {/* Toggle for BFS and IDS */}
              <div className="text-center relative">
                <label className="text-white font-bold mx-4">
                  BFS
                  <input
                    type="radio"
                    name="method"
                    value="uploadbfs"
                    checked={method === "uploadbfs"}
                    onChange={() => setMethod("uploadbfs")}
                    className="mx-2"
                  />
                </label>
                <label className="text-white font-bold mx-4">
                  IDS
                  <input
                    type="radio"
                    name="method"
                    value="uploadids"
                    checked={method === "uploadids"}
                    onChange={() => setMethod("uploadids")}
                    className="mx-2"
                  />
                </label>
              </div>

              <div className="mt-4 relative my-8">
                <label
                  htmlFor="single_path_toggle"
                  className="text-center text-xl block mb-2 font-medium text-white relative"
                >
                  Single Path?
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
                  className="text-center text-xl block mb-2 font-medium text-white"
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
          <div className="mt-40 mb-20 mx-20 bg-black relative border-8 border-yellow-400 rounded-xl">
            <div className="pt-16 pb-8 bg-white text-center">
              <p className="text-black font-bold text-5xl">RESULTS</p>
              <div>
                {responseOutput.extraData.totalLinks && (
                  <div className="extra-info text-center text-xl text-black bg-white my-8 mx-8">
                    <p className="font-bold text-red-700">
                      Total Links : {responseOutput.extraData.totalLinks}
                    </p>
                    <p className="font-bold text-yellow-600">
                      Path Length : {responseOutput.extraData.pathLength}
                    </p>
                    <p className="font-bold text-green-700">
                      Execution Time : {responseOutput.extraData.execTime} ms
                    </p>
                    <p className="font-bold text-blue-700">
                      Path Amount : {responseOutput.extraData.pathAmount}
                    </p>
                  </div>
                )}
              </div>
            </div>
            <div>
              {responseOutput.paths.length > 0 ? (
                <div className="response-container items-center justify-center bg-blue-200 flex flex-row">
                  {responseOutput.paths.slice(0, 4).map(
                    (
                      pathList,
                      listIndex // Only display up to 4 paths
                    ) => (
                      <div
                        key={listIndex}
                        className="response-output text-center text-xl text-black bg-white my-8 px-8 py-20 mx-8 border-4 border-black rounded-3xl"
                      >
                        <ul>
                          {pathList.map((path, pathIndex) => (
                            <li
                              key={pathIndex}
                              className="mb-4 text-2xl font-bold"
                            >
                              {path}
                            </li>
                          ))}
                        </ul>
                      </div>
                    )
                  )}
                </div>
              ) : (
                <p className="text-center text-xl text-white">
                  No paths found.
                </p>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CucakRowo;
