export const SearchBar = () => {
  const [input, setInput] = useState("");
  const fetchData = (value) => {
    fetch("https://jsonplaceholder.typicode.com/users")
      .then((response) => response.json())
      .then((json) => {
        const results = json.filter((user) => {
          return (
            value &&
            user &&
            user.name &&
            user.name.toLowerCase().includes(value)
          );
        });
        setResults(results)
      });
  };

  const handleChange = (value) => {
    setInput(value);
    fetchData(value);
  };

  return (
    <input
      type="search"
      id="final_page"
      value={input}
      onChange={(e) => setInput(e.target.value)}
      className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-80 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
      placeholder="Input final page here ..."
      // required
      setResults={setResults}
    />
  );
};
