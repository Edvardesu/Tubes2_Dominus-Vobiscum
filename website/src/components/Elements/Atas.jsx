import { useState } from "react";

const CucakRowo = (props) => {
  const [selectedImage, setSelectedImage] = useState(null);
  const [selectedImagesArray] = useState([]);
  const [toggleState, setToggleSetState] = useState(0);
  const { imageList, setImageList, timeElapsed, setTimeElapsed } = props;

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
              <label
                htmlFor="start_page"
                className="text-center text-xl block mb-2 font-medium text-white"
              >
                Start Page
              </label>
              <input
                type="text"
                id="start_page"
                className="mb-8 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-80 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="Input start page here ..."
                required
              />

              <label
                htmlFor="final_page"
                className="text-center text-xl block mb-2 font-medium text-white"
              >
                Final Page
              </label>
              <input
                type="text"
                id="final_page"
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-80 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="Input final page here ..."
                required
              />
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
