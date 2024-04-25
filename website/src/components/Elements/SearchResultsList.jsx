import "./SearchResultsList.css";
import { SearchResult } from "./SearchResult";

export const SearchResultsList = ({ results }) => {
  return (
    <div className="results">
      {results.map((result, i) => {
        {
          /* const url = `https://en.wikipedia.org/?curid=${result.pageid}`; */
        }

        return (
          <>
            <div className="result" key={i}>
              <h3 className="text-white">{result.title}</h3>
              {/* <a
                className="text-white"
                href={url}
                target="_blank"
                rel="noreferrer"
              >
                Read more
              </a> */}
            </div>
          </>
        );
      })}
    </div>
  );
};

// export const SearchResultsList = ({ results }) => {
//   return (
//     <div className="results-list">
//       {results.map((result, id) => {
//         return <SearchResult result={result.name} key={id} />;
//       })}
//     </div>
//   );
// };
