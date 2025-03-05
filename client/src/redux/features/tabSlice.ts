import {createSlice , PayloadAction} from "@reduxjs/toolkit";

interface Tabs {
    tab : string
}

const initialState : Tabs = {
    tab : 'profile'
}

const tabsSlice = createSlice({
    name : 'tabs',
    initialState,
    reducers : {
        updateTab(state , action : PayloadAction<string>){
            state.tab = action.payload
        }
    }
})

export const {updateTab} = tabsSlice.actions
export default tabsSlice.reducer