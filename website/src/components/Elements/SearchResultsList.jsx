import { SearchResult } from "./SearchResult";
import "./SearchResultsList.css";

export const SearchResultsList = ({ hasils }) => {
  return (
    <div className="results-list text-black">
      {hasils.map((hasil, id) => {
        return <SearchResult result={hasil.name} key={id} />;
      })}
    </div>
  );
};