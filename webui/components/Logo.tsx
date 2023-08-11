import styled from "@emotion/styled";
import { ReactComponent as HomeIcon } from "../assets/icons/home.svg";
import { Stack } from "@mui/material";
import { useNavigate } from "react-router-dom";
import React from "react";

const Logo = () => {
  const navigate = useNavigate();
  return (
    <LogoRoot
      alignItems="center"
      justifyContent="center"
      direction="row"
      spacing={2}
      onClick={() => navigate("/")}
    >
      <HomeIcon />
      <LogoTitle>FUNNY MOVIE</LogoTitle>
    </LogoRoot>
  );
};

const LogoRoot = styled(Stack)({
  cursor: "pointer",
});

const LogoTitle = styled.h2({
  marginTop: "0 !important",
  fontSize: "2.0rem",
  background: " linear-gradient(to right, #30CFD0 0%, #330867 100%)",
  backgroundClip: "text",
  WebkitTextFillColor: "transparent",
  "@media (max-width: 600px)": {
    fontSize: "1.4rem",
  }
});

export default Logo;
