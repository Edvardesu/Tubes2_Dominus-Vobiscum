import { useEffect, useState } from "react";
import SearchIcon from "@mui/icons-material/Search";
import CloseIcon from "@mui/icons-material/Close";
import "./BarSearch.css";

const SearchBarGenah = () => {
  const [search, setSearch] = useState("");
  const [searchData, setSearchData] = useState([]);
  const [selectedItem, setSelectedItem] = useState(-1);
  const [loading, setLoading] = useState(false);

  const handleChange = (e) => {
    setSearch(e.target.value);
  };

  const handleClose = () => {
    setSearch("");
    setSearchData([]);
    setSelectedItem(-1);
  };

  const handleKeyDown = (e) => {
    if (e.key === "ArrowUp" && selectedItem > 0) {
      setSelectedItem(prev => prev - 1);
    } else if (e.key === "ArrowDown" && selectedItem < searchData.length - 1) {
      setSelectedItem(prev => prev + 1);
    } else if (e.key === "Enter" && selectedItem >= 0) {
      setSearch(searchData[selectedItem].title); // Update search to the selected item title
      setSearchData([]); // Hide the suggestions
      const url = `https://en.wikipedia.org/?curid=${searchData[selectedItem].pageid}`;
      window.open(url);
    }
  };

  useEffect(() => {
    if (search !== "" && !loading) {
      setLoading(true);
      const url = `https://en.wikipedia.org/w/api.php?action=query&list=search&prop=info&inprop=url&utf8=&format=json&origin=*&srlimit=20&srsearch=${search}`;
      fetch(url)
        .then(res => res.json())
        .then(data => {
          if (data.query) {
            setSearchData(data.query.search);
          } else {
            setSearchData([]);
          }
        })
        .catch(error => {
          console.error('Error fetching search data:', error);
        })
        .finally(() => {
          setLoading(false);
        });
    } else {
      setSearchData([]);
    }
  }, [search]);

  return (
    <section className="search_section h-screen">
      <div className="search_input_div">
        <input
          type="text"
          className="search_input"
          placeholder="Search..."
          autoComplete="off"
          onChange={handleChange}
          value={search}
          onKeyDown={handleKeyDown}
        />
        <div className="search_icon">
          {search === "" ? <SearchIcon /> : <CloseIcon onClick={handleClose} />}
        </div>
      </div>
      {searchData.length > 0 && (
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
                  setSearch(item.title); // Set search to the clicked item's title
                  setSearchData([]); // Immediately hide the suggestions
                  setSelectedItem(-1); // Reset selected item index
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

export default SearchBarGenah;
