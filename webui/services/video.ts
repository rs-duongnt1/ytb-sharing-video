import { createApi } from "@reduxjs/toolkit/query/react";
import { baseQuery } from "./baseQuery";
import { IVideo } from "../models/video";

export const videoApi = createApi({
  reducerPath: "videoApi",
  baseQuery: baseQuery,
  endpoints: (builder) => ({
    videoList: builder.query<{ videos: IVideo[] }, void>({
      query: () => ({
        url: `v1/videos`,
      }),
    }),
  }),
});

export const { useVideoListQuery } = videoApi;
