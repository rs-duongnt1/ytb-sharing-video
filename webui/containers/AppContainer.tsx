import React, { useEffect } from "react";
import { CssBaseline, Link, Typography } from "@mui/material";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import VideosContainer from "./VideosContainer";
import MainLayout from "../layouts/MainLayout";
import ShareContainer from "./ShareContainer";
import SignInContainer from "./SignInContainer";
import SignUpContainer from "./SignUpContainer";
import { useUserInfoQuery } from "../services/auth";
import { ToastContainer, toast } from "react-toastify";

import "react-toastify/dist/ReactToastify.css";
import { useWebSocket } from "../providers/WebSocketProvider";
import { useDispatch } from "react-redux";
import { addVideo } from "../reducers/video";
const makeRoutes = () =>
  createBrowserRouter([
    {
      path: "/",
      element: <MainLayout />,
      children: [
        {
          path: "/",
          element: <VideosContainer />,
        },
        {
          path: "/share",
          element: <ShareContainer />,
        },
      ],
    },
    {
      path: "/signin",
      element: <SignInContainer />,
    },
    {
      path: "/signup",
      element: <SignUpContainer />,
    },
  ]);

const AppContainer: React.FC = () => {
  const dispatch = useDispatch();
  useUserInfoQuery();
  const { data } = useWebSocket();
  useEffect(() => {
    if (data && !data.error) {
      const video = data.data;
      dispatch(addVideo(video));
      toast(
        <div>
          <Typography color="blue">{video.user.email}</Typography>
          <Typography>{video.title}</Typography>
          <Link href={video.url} target="_blank">
            {video.url}
          </Link>
        </div>
      );
    } else {
      if (data?.message) {
        toast(data.message, {
          type: "error",
        });
      }
    }
  }, [data]);
  return (
    <div>
      <CssBaseline />
      <RouterProvider router={makeRoutes()}></RouterProvider>
      <ToastContainer />
    </div>
  );
};

export default AppContainer;
