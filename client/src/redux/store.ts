import {configureStore} from "@reduxjs/toolkit";
//reducer
import tabSlice from "@/redux/features/tabSlice";

export const store = configureStore({
    reducer: {
        tabs : tabSlice
    }
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;