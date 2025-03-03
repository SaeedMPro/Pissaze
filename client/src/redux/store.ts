import {configureStore} from "@reduxjs/toolkit";
//reducer
import tabSlice from "@/redux/features/tabSlice";
import modalSlice from "@/redux/features/modalSlice";

export const store = configureStore({
    reducer: {
        tabs : tabSlice,
        modal : modalSlice
    }
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;