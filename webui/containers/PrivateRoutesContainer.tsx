import { Route, Router, RouterProvider, Routes } from "react-router";
import PrivateRoute from "../components/PrivateRoute";
import React from "react";
import { createBrowserRouter } from "react-router-dom";
import SignInForm from "../components/SignInForm";

interface Props {
    isLoggedIn: boolean;
}


const PrivateRoutesContainer: React.FC<Props> = ({ isLoggedIn }) => {
    return <PrivateRoute isLoggedIn={isLoggedIn}></PrivateRoute>
}

export default PrivateRoutesContainer;