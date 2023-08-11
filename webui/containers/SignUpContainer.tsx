import React from "react";
import { Stack } from "@mui/material";
import SignUpForm from "../components/SignUpForm";
import { useAuth } from "../hooks/useAuth";

const SignUpContainer = () => {
  const { handleRegister, loading, errorMsg } = useAuth();
  return (
    <Stack
      sx={{
        width: "100%",
        height: "100vh",
        alignItems: "center",
        justifyContent: "center",
        backgroundColor: "#d6d6d6",
      }}
    >
      <Stack
        sx={{
          width: "100%",
          maxWidth: "400px",
          backgroundColor: "#fff",
          padding: "15px 30px",
          borderRadius: "10px",
        }}
      >
        <SignUpForm
          handleRegister={handleRegister}
          loading={loading}
          errorMsg={errorMsg}
        />
      </Stack>
    </Stack>
  );
};

export default SignUpContainer;
