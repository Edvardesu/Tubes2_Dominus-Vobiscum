import "./SearchResult.css";
export const SearchResult = ({ hasil }) => {
  return (
    <div
      className="search-result text-black"
      onClick={(e) => alert(`You selected ${hasil}!`)}
      // onClick={console.log(`You selected ${hasil}!`)}
    >
      {hasil}
    </div>
  );
};
