import { useEffect, useState, useCallback } from "react";
import SearchIcon from "@mui/icons-material/Search";
import CloseIcon from "@mui/icons-material/Close";
import "./BarSearch.css";

const debounce = (func, delay) => {
  let debounceTimer;
  return function () {
    const context = this;
    const args = arguments;
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => func.apply(context, args), delay);
  };
};

const SearchBarGenahStart = ({ onSelect }) => {
  const [start, setStart] = useState("");
  const [search, setSearch] = useState("");
  const [searchData, setSearchData] = useState([]);
  const [selectedItem, setSelectedItem] = useState(-1);
  const [loading, setLoading] = useState(false);

  // const handleChange = (e) => {
  //   setSearch(e.target.value);
  // };

  const handleChange = (e) => {
    setSearch(e.target.value);
    setStart(e.target.value);
  };

  const handleClose = () => {
    setStart("");
    setSearchData([]);
    setSelectedItem(-1);
  };

  const handleKeyDown = (e) => {
    if (e.key === "ArrowUp" && selectedItem > 0) {
      setSelectedItem((prev) => prev - 1);
    } else if (e.key === "ArrowDown" && selectedItem < searchData.length - 1) {
      setSelectedItem((prev) => prev + 1);
    } else if (e.key === "Enter" && selectedItem >= 0) {
      setStart(searchData[selectedItem].title);
      setSearchData([]);
      const url = `https://en.wikipedia.org/?curid=${searchData[selectedItem].pageid}`;
      window.open(url);
      // sendSearchToServer(search);
    }
  };

  // Function to send search to the BFS server
  // const sendSearchToServer = useCallback(debounce((searchTerm) => {
  //   const url = `http://localhost:8080/startBFS`; // Adjust URL as per your server setup
  //   setLoading(true);
  //   fetch(url, {
  //     method: 'POST',
  //     headers: {
  //       'Content-Type': 'application/json'
  //     },
  //     body: JSON.stringify({ start: searchTerm, end: "desired end node here" }) // Adjust 'end' as needed
  //   })
  //   .then(res => res.json())
  //   .then(data => {
  //     setSearchData(data.paths); // Assuming data.paths contains the search results
  //   })
  //   .catch(error => {
  //     console.error('Error fetching search data:', error);
  //     setSearchData([]);
  //   })
  //   .finally(() => {
  //     setLoading(false);
  //   });
  // }, 500), []);

  const fetchSearchData = useCallback(
    debounce((input) => {
      const url = `https://en.wikipedia.org/w/api.php?action=query&list=search&prop=info&inprop=url&utf8=&format=json&origin=*&srlimit=20&srsearch=${input}`;
      setLoading(true);
      fetch(url)
        .then((res) => res.json())
        .then((data) => {
          // Ensure we safely access `data.query.search`
          if (data.query && data.query.search) {
            setSearchData(data.query.search);
          } else {
            setSearchData([]);
          }
        })
        .catch((error) => {
          console.error("Error fetching search data:", error);
          setSearchData([]);
        })
        .finally(() => {
          setLoading(false);
        });
    }, 500),
    []
  );

  useEffect(() => {
    if (start !== "") {
      fetchSearchData(start);
    } else {
      setSearchData([]);
    }
  }, [start, fetchSearchData]);

  return (
    <section className="search_section">
      <div className="search_input_div">
        <input
          type="text"
          className="search_input"
          placeholder="Search..."
          autoComplete="off"
          onChange={handleChange}
          value={start}
          onKeyDown={handleKeyDown}
        />
        <div className="search_icon">
          {start === "" ? <SearchIcon /> : <CloseIcon onClick={handleClose} />}
        </div>
      </div>
      {searchData && searchData.length > 0 && (
        <div className="search_result">
          {searchData.map((item, index) => {
            return (
              <div
                key={item.pageid}
                className={
                  selectedItem === index
                    ? "search_suggestion_line active"
                    : "search_suggestion_line"
                }
                onClick={() => {
                  onSelect(item.title)
                  setStart(item.title);
                  setSearchData([]);
                  setSelectedItem(-1);
                }}
              >
                {item.title}
              </div>
            );
          })}
        </div>
      )}
    </section>
  );
};

export default SearchBarGenahStart;
