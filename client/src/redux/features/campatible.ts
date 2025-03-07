import {createSlice , PayloadAction} from "@reduxjs/toolkit";

interface State {
    categoryChoice:string[]
    productsChoice:number[]
}

const initialState: State = {
    categoryChoice: [],
    productsChoice:[]
}

const compatibleSlice = createSlice({
    name : "compatible",
    initialState,
    reducers : {
        updateCategory : (state, action : PayloadAction<string[]>) => {
            state.categoryChoice=action.payload
        },
        updateProducts : (state, action : PayloadAction<number[]>) => {
            state.productsChoice=action.payload
        }
    }
})

export const {updateCategory , updateProducts} = compatibleSlice.actions;
export default compatibleSlice.reducer