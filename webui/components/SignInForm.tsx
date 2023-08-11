import * as React from "react";
import TextField from "@mui/material/TextField";
import { Box, Stack, Typography } from "@mui/material";
import { useForm } from "react-hook-form";
import LoadingButton from "@mui/lab/LoadingButton";
import { Link } from "react-router-dom";
import { UserLoginPayload } from "../models/user";

interface Props {
  handleLogin: (data: UserLoginPayload) => void;
  loading: boolean;
}

const SignInForm: React.FC<Props> = ({ handleLogin, loading }) => {
  const { register, handleSubmit } = useForm({
    defaultValues: {
      email: "",
      password: "",
    } as UserLoginPayload,
  });
  return (
    <Box component="form" onSubmit={handleSubmit(handleLogin)}>
      <Typography variant="subtitle1">Login</Typography>
      <Stack>
        <TextField
          data-testid="email"
          autoFocus
          margin="dense"
          required
          label="Email Address"
          type="email"
          fullWidth
          variant="standard"
          {...register("email")}
        />
      </Stack>
      <Stack>
        <TextField
          data-testid="password"
          margin="dense"
          required
          label="Password"
          type="password"
          fullWidth
          variant="standard"
          {...register("password")}
        />
      </Stack>
      <Stack justifyContent="flex-end" mt={3} mb={3} spacing={2}>
        <LoadingButton
          type="submit"
          variant="contained"
          loading={loading}
          data-testid="button"
        >
          Login
        </LoadingButton>
      </Stack>
      <Stack mb={2}>
        <Link data-test="signup" to="/signup">
          {"Don't have an account? Sign Up"}
        </Link>
      </Stack>
    </Box>
  );
};

export default SignInForm;
