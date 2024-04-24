import { useState } from "react";

const BonusScrapping = (props) => {
  const { imageList, setImageList } = props;

  const [url, setUrl] = useState("");

  const sendDataToBE = async (sendData) => {
    const response = await fetch("http://localhost:8081/scrapping", {
      method: "POST",
      body: JSON.stringify(sendData),
    });

    const data = await response.json();
    console.log("dari scrapping ===========");
    console.log(data);
    // setImageList(data.name);
  };

  const handleScrapping = (e) => {
    console.log("called scraping");
    e.preventDefault();

    var sendData = {
      url: url,
    };

    sendDataToBE(sendData);
  };

  return (
    <div className="w-fit-w-xs">
      <form
        className="justify-center shadow-md px-8 pt-6 pb-8 flex flex-row"
        style={{ backgroundColor: "#28293D" }}
        onSubmit={handleScrapping}
      >
        <div className="mb-4">
          <label
            className="block text-gray-700 text-sm font-bold mb-2 text-white"
            htmlFor="username"
          >
            Insert Link
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline w-64"
            id="username"
            type="text"
            placeholder="insert link here ..."
            onChange={(e) => setUrl(e.target.value)}
          />
        </div>
        <div className="flex items-center justify-between ml-12">
          <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-32">
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default BonusScrapping;
