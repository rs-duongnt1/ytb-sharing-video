import { createSlice } from "@reduxjs/toolkit";
import { videoApi } from "../services/video";
import { IVideo } from "../models/video";

const slice = createSlice({
  name: "videos",
  initialState: { videos: [] as IVideo[] },
  reducers: {
    addVideo(state, action) {
      console.log('add...', action.payload);
      state.videos = [action.payload, ...state.videos];
    },
  },
  extraReducers: (builder) => {
    builder.addMatcher(
      videoApi.endpoints.videoList.matchFulfilled,
      (state, { payload }) => {
        state.videos = payload.videos;
      }
    );
  },
});
export const { addVideo } = slice.actions;
export const videoReducer = slice.reducer;
