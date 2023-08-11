import ThumbUpOffAltIcon from "@mui/icons-material/ThumbUpOffAlt";
import ThumbDownOffAltIcon from "@mui/icons-material/ThumbDownOffAlt";
import styled from "@emotion/styled";
import { Stack, Typography, IconButton, Link } from "@mui/material";
import { IVideo } from "../models/video";
import React, { useMemo } from "react";
interface VideoItemProps {
  video: IVideo;
}

const VideoItem: React.FC<VideoItemProps> = ({ video }) => {
  const videoEmbedUrl = useMemo(() => `https://www.youtube.com/embed/${video.videoId}`, [video.videoId])
  return (
    <VideoItemRoot>
      <VideoIframe
        src={videoEmbedUrl}
        title="YouTube video player"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
        allowFullScreen
      ></VideoIframe>
      <Stack direction="column" spacing={1}>
        <Typography variant="h5" component="h4" gutterBottom data-testid="title">
          {video.title}
        </Typography>
        <Typography mt={0} data-testid="share-by">Shared by <Link>{ video.user.email }</Link></Typography>
        <Stack>
          <Stack alignItems="center" direction="row" spacing={2}>
            <Stack alignItems="center" direction="row">
              <Typography>89</Typography>
              <IconButton>
                <ThumbUpOffAltIcon />
              </IconButton>
            </Stack>
            <Stack alignItems="center" direction="row" ml={1}>
              <Typography>89</Typography>
              <IconButton>
                <ThumbDownOffAltIcon />
              </IconButton>
            </Stack>
          </Stack>
        </Stack>
        <Typography>Description:</Typography>
        <VideoDescription data-testid="description">
          {video.description}
        </VideoDescription>
      </Stack>
    </VideoItemRoot>
  );
};

const VideoItemRoot = styled.div({
  display: "flex",
  gap: 30,
  flexDirection: "row",
  "@media (max-width: 600px)": {
    flexDirection: "column",
    gap: 10,
  },
});

const VideoIframe = styled.iframe({
  border: "none",
  width: "380px",
  height: "190px",
  "@media (max-width: 600px)": {
    width: "100%",
    height: "240px",
  },
});
const VideoTitle = styled(Typography)({
  overflow: 'hidden',
  display: '-webkit-box',
  WebkitLineClamp: 2,
  WebkitBoxOrient: 'vertical',
  lineClamp: "2",
})

const VideoDescription = styled(Typography)({
  overflow: 'hidden',
  display: '-webkit-box',
  WebkitLineClamp: "5",
  lineClamp: "5",
  WebkitBoxOrient: 'vertical'
})

export default VideoItem;
