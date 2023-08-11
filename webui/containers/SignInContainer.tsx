import React from "react"
import SignInForm from "../components/SignInForm"
import { Box, Stack } from "@mui/material";
import { useAuth } from "../hooks/useAuth";

const SignInContainer = () => {
    const { handleLogin, loading, errorMsg } = useAuth();

    return <Stack
        sx={{

            width: '100%',
            height: '100vh',
            alignItems: 'center',
            justifyContent: 'center',
            backgroundColor: '#d6d6d6'
        }}>
        <Stack sx={{
            width: '100%',
            maxWidth: '400px',
            backgroundColor: '#fff',
            padding: '15px 30px',
            borderRadius: '10px',
        }}>
            <SignInForm loading={loading} errorMsg={errorMsg as string} handleLogin={handleLogin} />
        </Stack>
    </Stack>
}

export default SignInContainer;