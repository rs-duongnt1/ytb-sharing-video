import React from "react";
import VideoList from "../components/VideoList";
import { useVideoListQuery } from "../services/video";
import { useSelector } from "react-redux";
import { RootState } from "../store";

const VideosContainer = () => {
    useVideoListQuery();
    const videos = useSelector((state: RootState) => state.video.videos);
    return <div>
        <VideoList videos={videos} />
    </div>
}

export default VideosContainer;