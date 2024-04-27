import { Fragment, useEffect } from "react";
// import Navbar from "../../components/Elements/Navbar/Navbar";
import CucakRowo from "../components/Elements/Atas";
import { useState } from "react";
import Navbar from "../components/Elements/Navbar";
import "./../components/Elements/bg.css";

const AboutUs = () => {
  return (
    <div className="py-20 w-full h-full flex flex-col justify-between">
      <div className="background">
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>

        <div className="h-full flex flex-col overflow-x-auto">
          <Navbar />
          <div className="flex flex-initial justify-center items-center w-full">
          <div className="text-white justify-center ">JUDUL</div>
            <div className="flex flex-col w-full h-full">
            <p></p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default AboutUs;
