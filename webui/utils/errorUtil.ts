import { SerializedError } from "@reduxjs/toolkit";
import { FetchBaseQueryError } from "@reduxjs/toolkit/dist/query";

export const getApiErrorMessage = (error: FetchBaseQueryError | SerializedError | undefined): string => {
    const err = error as FetchBaseQueryError;
    return (err?.data as { message: string })?.message || 'Unknow Error';
}