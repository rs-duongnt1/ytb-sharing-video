import { fetchBaseQuery } from "@reduxjs/toolkit/dist/query";
import { Config } from "../config/config";
import { useLocalStorage } from "../hooks/useLocalstorage";
export const baseQuery = fetchBaseQuery({
  baseUrl: Config.apiUrl,
  prepareHeaders: (headers, { getState }) => {

    // If we have a token set in state, let's assume that we should be passing it.
    const t = localStorage.getItem('t');
    if (t) {
      headers.set("authorization", `Bearer ${t}`);
    }

    return headers;
  },
});
