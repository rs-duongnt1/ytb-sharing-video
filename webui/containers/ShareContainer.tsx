import styled from "@emotion/styled";
import { Stack, TextField, Button, Typography, Link } from "@mui/material";
import React, { useEffect, useState } from "react";
import { WebSocketContext, useWebSocket } from "../providers/WebSocketProvider";
import { toast } from "react-toastify";

const ShareContainer = () => {
  const { data, sendMessage } = useWebSocket();
  const [url, setUrl] = useState("");
  const handleShare = () => {
    sendMessage("share", url);
  };

  useEffect(() => {
    // if (data) {
    //   const video = data;
    //   toast(
    //     <div>
    //       <Typography color="blue">{video.user.email}</Typography>
    //       <Typography>{video.title}</Typography>
    //       <Link href={video.url} target="_blank">
    //         {video.url}
    //       </Link>
    //     </div>
    //   );
    // }
  }, [data]);
  return (
    <SharePageRoot>
      <ShareFormLegend>Share a youtube movie</ShareFormLegend>
      <Stack direction="column">
        <Stack direction="column">
          <TextField
            placeholder="Youtube URL"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
          />
        </Stack>
        <Stack mt={3}>
          <Button variant="contained" fullWidth onClick={handleShare}>
            Share
          </Button>
        </Stack>
      </Stack>
    </SharePageRoot>
  );
};

const SharePageRoot = styled.fieldset({
  margin: "0 auto",
  maxWidth: "440px",
  border: "1px solid #DAE2ED",
  padding: "40px 30px",
  borderRadius: "3px",
});

const ShareFormLegend = styled.legend({
  fontSize: "1.6rem",
  padding: "0 10px",
  color: "#131313",
});

export default ShareContainer;
