import { videoReducer } from "../reducers/video";
import { authApi } from "../services/auth";
import { videoApi } from "../services/video";

export const reducerMapper = {
    [authApi.reducerPath]: authApi.reducer,
    [videoApi.reducerPath]: videoApi.reducer,
    video: videoReducer,
}

export const middlewareMapper = [
    authApi.middleware,
    videoApi.middleware,
]