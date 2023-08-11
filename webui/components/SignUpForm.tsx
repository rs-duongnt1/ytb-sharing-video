import * as React from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import { Box, Stack, Typography } from "@mui/material";
import { useForm } from "react-hook-form";
import { useAuth } from "../hooks/useAuth";
import LoadingButton from '@mui/lab/LoadingButton';
import { Link } from "react-router-dom";
import { UserRegisterPayload } from "../models/user";

interface Props {
    handleRegister: (data: UserRegisterPayload) => void;
    loading: boolean;
    errorMsg: string | undefined;
}

const SignUpForm: React.FC<Props> = ({ handleRegister, loading, errorMsg }) => {
    const { register, handleSubmit } = useForm({
        defaultValues: {
            email: '',
            password: '',
            confirmPassword: '',
        } as UserRegisterPayload,
    });
    return (
        <Box component="form" onSubmit={handleSubmit(handleRegister)}>
            <Typography variant="subtitle1">Register to new account</Typography>
            <Stack>
                <TextField
                    autoFocus
                    margin="dense"
                    required
                    label="Email Address"
                    type="email"
                    fullWidth
                    variant="standard"
                    {...register('email')}
                />
            </Stack>
            <Stack>
                <TextField
                    margin="dense"
                    required
                    label="Password"
                    type="password"
                    fullWidth
                    variant="standard"
                    {...register('password')}
                />
            </Stack>
            <Stack>
                <TextField
                    margin="dense"
                    required
                    label="Confirm Password"
                    type="password"
                    fullWidth
                    variant="standard"
                    {...register('confirmPassword')}
                />
            </Stack>
            {errorMsg && <Stack mt={4}>
                <Typography color="error" variant="subtitle2">{errorMsg}</Typography>
            </Stack>}
            <Stack justifyContent="flex-end" mt={3} mb={3} spacing={2}>
                <LoadingButton type="submit" variant="contained" loading={loading}>
                    Register
                </LoadingButton>
            </Stack>
            <Stack mb={2}>
                <Link to="/signin">{"Have an account? Sign In"}</Link>
            </Stack>
        </Box>
    );
}

export default SignUpForm;