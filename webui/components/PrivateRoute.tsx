import React from "react";
import { Route, Navigate, RouteProps, Outlet } from "react-router-dom";

type PrivateRouteProps = {
    isLoggedIn: boolean;
}

function PrivateRoute({ isLoggedIn }: PrivateRouteProps) {
    return (
        isLoggedIn ? (
            <Outlet />
        ) : (
            /* istanbul ignore next */
            <Navigate to="/signin" replace={true} />
        )
    );
}

export default PrivateRoute;
