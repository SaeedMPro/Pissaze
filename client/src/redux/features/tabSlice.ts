import {createSlice , PayloadAction} from "@reduxjs/toolkit";

interface Tabs {
    tab : string
    subTab : string
}

const initialState : Tabs = {
    tab : 'profile',
    subTab:'UserInformation'

}

const tabsSlice = createSlice({
    name : 'tabs',
    initialState,
    reducers : {
        updateTab(state , action : PayloadAction<string>){
            state.tab = action.payload
        },
        updateSubTab(state , action : PayloadAction<string>){
            state.subTab = action.payload
        }

    }
})

export const {updateTab , updateSubTab} = tabsSlice.actions
export default tabsSlice.reducer