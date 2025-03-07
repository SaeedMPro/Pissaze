import {configureStore} from "@reduxjs/toolkit";
//reducer
import tabSlice from "@/redux/features/tabSlice";
import modalSlice from "@/redux/features/modalSlice";
import compatibleSlice from "@/redux/features/campatible"

export const store = configureStore({
    reducer: {
        tabs : tabSlice,
        modal : modalSlice,
        compatible : compatibleSlice
    }
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;