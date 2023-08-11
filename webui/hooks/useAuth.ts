import { useDispatch, useSelector } from "react-redux";
import {
  useRegisterMutation,
  useLoginMutation,
  useUserInfoQuery,
} from "../services/auth";
import { useLocalStorage } from "./useLocalstorage";
import { useEffect, useState } from "react";
import { UserLoginPayload, UserRegisterPayload } from "../models/user";
import { useNavigate } from "react-router";
import { getApiErrorMessage } from "../utils/errorUtil";
import { toast } from "react-toastify";

export const useAuth = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [errorMsg, setErrorMsg] = useState<string>();
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [
    register,
    {
      data: registerData,
      isError: isRegisterError,
      error: registerError,
      isSuccess: isRegisterSuccess,
    },
  ] = useRegisterMutation();
  const [login, { isError, data, error, isSuccess }] = useLoginMutation();
  const { refetch } = useUserInfoQuery();
  useEffect(() => {
    if (isError || isRegisterError) {
      const msg = getApiErrorMessage(error || registerError);
      toast(msg, {
        type: "error",
      });
    }
  }, [isError, registerError]);

  useEffect(() => {
    if (isSuccess || isRegisterSuccess) {
      if (data?.token || registerData?.token) {
        localStorage.setItem(
          "t",
          (data?.token || registerData?.token) as string
        );
        refetch();
      }
      toast(isRegisterSuccess ? "Register successful" : "Login successful", {
        type: "success",
      });
      navigate("/");
    }
  }, [isSuccess, isRegisterSuccess]);

  const handleRegister = (data: UserRegisterPayload) => {
    register(data);
  };

  const handleLogin = (data: UserLoginPayload) => {
    login(data);
  };

  const handleLogout = () => {
    localStorage.removeItem("t");
    refetch();
  };

  return {
    loading,
    errorMsg,
    handleRegister,
    handleLogin,
    handleLogout,
  };
};
