import styled from "@emotion/styled";
import { Button, Stack, Typography } from "@mui/material";
import Logo from "./Logo";
import { useNavigate } from "react-router-dom";
import React from "react";
import { useUserInfoQuery } from "../services/auth";
import { useAuth } from "../hooks/useAuth";
import IosShareIcon from "@mui/icons-material/IosShare";
import LogoutIcon from "@mui/icons-material/Logout";
import LoginIcon from "@mui/icons-material/Login";

const Header = () => {
  const navigate = useNavigate();
  const { handleLogout } = useAuth();
  const { data, isSuccess, isFetching } = useUserInfoQuery();
  return (
    <HeaderRoot>
      <HeaderWrapper>
        <Logo />
        <div>
          <Stack alignItems="center" spacing={10}>
            <Stack>
              <Stack spacing={5}>
                {isSuccess && (
                  <Stack direction="row" alignItems="center" spacing={2}>
                    <Username>{data?.email}</Username>
                    <Button
                      color="success"
                      variant="outlined"
                      onClick={() => navigate("/share")}
                    >
                      <ButtonShareText>Share a movie</ButtonShareText>
                      <ButtonShareIcon></ButtonShareIcon>
                    </Button>
                    <Button color="error" onClick={() => handleLogout()}>
                      <ButtonLogoutText>Logout</ButtonLogoutText>
                      <ButtonLogoutIcon></ButtonLogoutIcon>
                    </Button>
                  </Stack>
                )}
                {!isSuccess && !isFetching && (
                  <Button
                    variant="contained"
                    onClick={() => navigate("/signin")}
                  >
                    <ButtonLoginText>Login / Register</ButtonLoginText>
                    <ButtonLoginIcon></ButtonLoginIcon>
                  </Button>
                )}
              </Stack>
            </Stack>
          </Stack>
        </div>
      </HeaderWrapper>
    </HeaderRoot>
  );
};

const MOBILE_WIDTH = 768;

const HeaderRoot = styled.div({
  padding: "0 30px",
  marginBottom: 30,
});

const Username = styled(Typography)({
  [`@media (max-width: ${MOBILE_WIDTH}px)`]: {
    display: "none",
  },
});

const ButtonShareText = styled.span({
  [`@media (max-width: ${MOBILE_WIDTH}px)`]: {
    display: "none",
  },
});
const ButtonShareIcon = styled(IosShareIcon)({
  display: "none",
  [`@media (max-width: ${MOBILE_WIDTH}px)`]: {
    display: "block",
  },
});

const ButtonLogoutText = styled.span({
  [`@media (max-width: ${MOBILE_WIDTH}px)`]: {
    display: "none",
  },
});
const ButtonLogoutIcon = styled(LogoutIcon)({
  display: "none",
  [`@media (max-width: ${MOBILE_WIDTH}px)`]: {
    display: "block",
  },
});

const ButtonLoginText = styled.span({
  [`@media (max-width: ${MOBILE_WIDTH}px)`]: {
    display: "none",
  },
});

const ButtonLoginIcon = styled(LoginIcon)({
  display: "none",
  [`@media (max-width: ${MOBILE_WIDTH}px)`]: {
    display: "block",
  },
});

const HeaderWrapper = styled.div({
  display: "flex",
  justifyContent: "space-between",
  alignItems: "center",
  borderBottom: "1px solid #DAE2ED",
  padding: "12px 0",
});

export default Header;
