import styled from "@emotion/styled";
import { Stack } from "@mui/material";
import VideoItem from "./VideoItem";
import React from "react";
import { IVideo } from "../models/video";


interface VideoListProps {
  videos?: IVideo[];
}

const VideoList: React.FC<VideoListProps> = ({ videos }) => {
  return (
    <VideoListRoot>
      <Stack direction="column" spacing={10}>
        {videos?.map((video) => (
          <VideoItem video={video} key={video.id}></VideoItem>
        ))}
      </Stack>
    </VideoListRoot>
  );
};

const VideoListRoot = styled.div({});

export default VideoList;
