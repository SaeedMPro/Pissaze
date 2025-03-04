import {createSlice , PayloadAction} from "@reduxjs/toolkit";

interface State {
    CompatDialog : boolean
}

const initialState: State = {
    CompatDialog: false,
}

const modalSlice = createSlice({
    name : "modal",
    initialState,
    reducers : {
        updateCompatDialog : (state, action : PayloadAction<boolean>) => {
            state.CompatDialog = action.payload;
        }
    }
})

export const {updateCompatDialog} = modalSlice.actions;
export default modalSlice.reducer