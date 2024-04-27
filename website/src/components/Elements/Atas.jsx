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

  const handleSubmit = async (event) => {
    event.preventDefault();

    const requestData = {
      start: start,
      destination: destination,
    };

    const response = await fetch("http://localhost:8080/uploadids", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(requestData),
    });

    const data = await response.json();
    console.log(data); // You can do something with the response data, like updating state
  };

  // const handleImageUpload = async () => {
  //   if (selectedImage) {
  //     const formData = new FormData();
  //     formData.append("file", selectedImage);

  //     try {
  //       const response = await fetch("http://localhost:8080/upload", {
  //         method: "POST",
  //         body: formData,
  //       });

  //       if (response.ok) {
  //         const result = await response.json();
  //         alert(
  //           result.messageType === "S"
  //             ? "Image uploaded successfully!"
  //             : `Error: ${result.message}`
  //         );
  //       } else {
  //         alert("Error uploading image");
  //       }
  //     } catch (error) {
  //       console.error("Error:", error);
  //       alert("Error uploading image");
  //     }
  //   } else {
  //     alert("Please select an image to upload");
  //   }
  // };

  // const handleSearch = async (e) => {
  //   e.preventDefault();
  //   if (search === "") return;

  //   const endpoint = `https://en.wikipedia.org/w/api.php?action=query&list=search&prop=info&inprop=url&utf8=&format=json&origin=*&srlimit=20&srsearch=${search}`;

  //   const response = await fetch(endpoint);

  //   // console.log(response);

  //   if (!response.ok) {
  //     throw Error(response.statusText);
  //   }

  //   const json = await response.json();
  //   console.log(json);

  //   setResults(json.query.search);
  //   setSearchInfo(json.query.searchinfo);
  // };

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
          <div className="flex flex-row mt-32 items-center justify-center gap-4 ">
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
