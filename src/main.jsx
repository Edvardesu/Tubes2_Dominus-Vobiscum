import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import ErrorPage from "./pages/404.jsx";
import ProductsPage from "./pages/Index/products";

const router = createBrowserRouter([
  {
    path: "/",
    element: <ProductsPage />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/register",
    element: <RegisterPage />,
  },
  {
    path: "/kasir",
    element: <KasirPage />,
  },
  {
    path: "/tentangKami",
    element: <TentangKami />,
  },
  {
    path: "/aktivitas",
    element: <Aktivitas />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
