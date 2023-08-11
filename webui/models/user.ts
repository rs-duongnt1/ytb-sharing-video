export interface IUser {
    id: string;
    email: string;
    password: string;
}

export type UserRegisterPayload = Pick<IUser, "email" | "password"> & {
    confirmPassword: string;
}

export type UserLoginPayload = Pick<IUser, "email" | "password">

export type UserResponse = Pick<IUser, "id" | "email" | "password"> & {
    token: string;
}