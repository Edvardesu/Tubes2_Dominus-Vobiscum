import { useEffect, useState } from "react";
import SearchIcon from "@mui/icons-material/Search";
import CloseIcon from "@mui/icons-material/Close";
// import data from "../books.json"
import "./BarSearch.css";

const BarSearch = () => {
  const [search, setSearch] = useState("");
  const [searchData, setSearchData] = useState([]);
  const [selectedItem, setSelectedItem] = useState(-1);

  const handleChange = (e) => {
    setSearch(e.target.value);
  };

  const handleClose = () => {
    setSearch("");
    setSearchData([]);
    setSelectedItem(-1);
  };

  const handleKeyDown = (e) => {
    if (selectedItem < searchData.length) {
      if (e.key === "ArrowUp" && selectedItem > 0) {
        setSelectedItem((prev) => prev - 1);
      } else if (
        e.key === "ArrowDown" &&
        selectedItem < searchData.length - 1
      ) {
        setSelectedItem((prev) => prev + 1);
      } else if (e.key === "Enter" && selectedItem >= 0) {
        window.open(searchData[selectedItem].show.url);
      }
    } else {
      setSelectedItem(-1);
    }
  };

  useEffect(() => {
    if (search !== "") {
      fetch(`http://api.tvmaze.com/search/shows?q=${search}`)
        .then((res) => res.json())
        .then((data) => setSearchData(data));

      // data.filter
    }
  }, [search]);

  // useEffect(() => {
  //   if (search !== "") {
  //     fetch(
  //       `https://en.wikipedia.org/w/api.php?action=query&list=search&prop=info&inprop=url&utf8=&format=json&origin=*&srlimit=20&srsearch=${search}`
  //     )
  //       .then((res) => res.json())
  //       .then((data) => setSearchData(data));

  //     // data.filter
  //   }
  // }, [search]);

  return (
    <section className="search_section bg-black h-screen">
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
      {/* <div className="search_result">
        {searchData.map((data) => {
          return <p>{data.show.name}</p>;
        })}
      </div> */}

      <div className="search_result">
        {searchData.map((data, index) => {
          return (
            <a
              href={data.show.url}
              key={index}
              target="_blank"
              className={
                selectedItem === index
                  ? "search_suggestion_line active"
                  : "search_suggestion_line"
              }
            >
              {data.show.name}
            </a>
          );
        })}
      </div>
    </section>
  );
};

export default BarSearch;
