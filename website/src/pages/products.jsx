import { Fragment, useEffect } from "react";
// import Navbar from "../../components/Elements/Navbar/Navbar";
import CucakRowo from "../components/Elements/Atas";
import { useState } from "react";

const ProductsPage = () => {
  // const slides = [img1, img2, img3];
  const [imageList, setImageList] = useState([]);
  const [timeElapsed, setTimeElapsed] = useState({});

  useEffect(() => {
    console.log("ini dari product.jsx ================");
    console.log(imageList);
  }, [imageList]);

  return (
    <div className="flex w-full" style={{ backgroundColor: "#28293D" }}>
      <Fragment>
        <div className="flex flex-initial justify-center items-center w-full">
          <div className="flex flex-col w-full">
            {/* <Navbar /> */}
            <CucakRowo
              imageList={imageList}
              setImageList={setImageList}
              timeElapsed={timeElapsed}
              setTimeElapsed={setTimeElapsed}
            />
          </div>
        </div>
      </Fragment>
    </div>
  );
};

export default ProductsPage;
