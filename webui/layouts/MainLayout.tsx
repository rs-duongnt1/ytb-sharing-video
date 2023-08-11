import styled from "@emotion/styled";
import { Outlet } from "react-router-dom";
import Header from "../components/Header";
import { Container } from "@mui/material";
import React from "react";

const Root = styled.div({
  paddingBottom: "60px",
});

const MainLayout = () => {
  return (
    <Root>
      <Header />
      <Container>
        <Outlet />
      </Container>
    </Root>
  );
};

export default MainLayout;
