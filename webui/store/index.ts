import { configureStore } from '@reduxjs/toolkit'
import { middlewareMapper, reducerMapper } from './mapper'

export const store = configureStore({
  reducer: {
    ...reducerMapper,
  },
  middleware: (getDefaultMiddleware) =>
  getDefaultMiddleware().concat(...middlewareMapper),
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch