import { Fragment, useState } from "react";
import Navbar from "../../components/Elements/Navbar/Navbar";
import Ringkasan from "../../components/Elements/Contents/Products/Ringkasan";
// import img1 from "path-to-local-image.jpg";
// import img2 from "path-to-local-image.jpg";
// import img3 from "path-to-local-image.jpg";

const ProductsPage = () => {
  // const slides = [img1, img2, img3];
  return (
    <div className="flex w-full">
      <Fragment>
        {/* <Sidebar></Sidebar> */}

        {/* <div className="flex justify-end h-20 bg-blue-500 text-white items-center px-10">
          {email}
          <Button className="ml-5 bg-black" onClick={handleLogout}>
            Logout
          </Button>
        </div> */}
        {/* <div className="flex justify-end h-20 w-3/4 bg-black text-white items-center px-10">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Aliquam qui
          expedita unde praesentium, esse perferendis illo reprehenderit
          temporibus nam. Illo voluptatum ipsum blanditiis accusamus ut
        </div> */}
        {/* <div className="flex h-screen">
          <Navbar />
        </div> */}

        <div className="flex flex-initial justify-center items-center">
          {/* <Navbar /> */}
          <div className="flex flex-col ">
            <Navbar />
            <Ringkasan />
          </div>
          {/* <div className="w-1/4">
          <h1 className="text-3xl font-bold text-blue-600">Cart</h1>
        </div> */}
        </div>
      </Fragment>
    </div>
  );
};

export default ProductsPage;
