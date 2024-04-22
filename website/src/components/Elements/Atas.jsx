import { useState } from "react";

const CucakRowo = (props) => {
  // const [selectedImage, setSelectedImage] = useState(null);
  // const [selectedImagesArray] = useState([]);
  // const [toggleState, setToggleSetState] = useState(0);
  // const { imageList, setImageList, timeElapsed, setTimeElapsed } = props;

  const [search, setSearch] = useState("");
  const [results, setResults] = useState([]);
  const [searchinfo, setSearchInfo] = useState({});

  const [results, setResults] = useState([]);

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

  return (
    <div className="py-20 w-full min-h-screen flex flex-col justify-between bg-blue-950">
      <div className="flex flex-col">
        <div className="overflow-x-auto">
          <p className="text-5xl mb-4 font-bold text-center text-white">
            WIKIRACE SOLVER
          </p>
          <p className="text-2xl font-bold mb-10 text-center text-yellow-400">
            DOMINVS VOBISCVM
          </p>
          <div className="flex flex-col mt-32 items-center justify-center">
            <div>
              <form className="search-box" onSubmit={handleSearch}>
                {/* <label
                  htmlFor="start_page"
                  className="text-center text-xl block mb-2 font-medium text-white"
                >
                  Start Page
                </label> */}
                <input
                  type="search"
                  placeholder="Input start page here ..."
                  value={search}
                  onChange={(e) => setSearch(e.target.value)}
                  // id="start_page"
                  className="mb-8 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-80 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  // required
                />
              </form>
              {/* <p className="text-center text-white">Search result</p> */}
              {searchinfo.totalhits ? (
                <p className="text-center text-white">
                  Search results: {searchinfo.totalhits}
                </p>
              ) : (
                ""
              )}
              <div className="results">
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
              </div>

              <label
                htmlFor="final_page"
                className="text-center text-xl block mb-2 font-medium text-white"
              >
                Final Page
              </label>
              
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CucakRowo;

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
