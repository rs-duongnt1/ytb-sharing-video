import { createApi } from '@reduxjs/toolkit/query/react'
import { baseQuery } from './baseQuery'
import { UserLoginPayload, UserRegisterPayload, UserResponse } from '../models/user'

export const authApi = createApi({
  reducerPath: 'authApi',
  baseQuery: baseQuery,
  endpoints: (builder) => ({
    login: builder.mutation<UserResponse, UserLoginPayload>({
      query: (user) => ({
        url: `v1/auth/login`,
        method: 'POST',
        body: user,
      })
    }),
    register: builder.mutation<UserResponse, UserRegisterPayload>({
      query: (user) => ({
        url: `v1/auth/register`,
        method: 'POST',
        body: user,
      })
    }),
    userInfo: builder.query<UserResponse, void>({
      query: () => ({
        url: `v1/user`,
      })
    }),
  }),
})

export const { useRegisterMutation, useLoginMutation, useUserInfoQuery } = authApi