'use client'
//redux
import {useDispatch} from "react-redux";
import {AppDispatch} from "@/redux/store";
import {updateCompatDialog} from "@/redux/features/modalSlice";
//components
import CompatibilityDialog from "@/components/CompatibilityDialog";

export default function Explore() {
    const dispatch = useDispatch<AppDispatch>();
    return <>
        <div className='w-full bg-muted rounded-b-xl flex items-center justify-center h-12 relative shadow-xl'>
            <button
                onClick={() => dispatch(updateCompatDialog(true))}
                className='bg-[#4771F1] w-64 h-[47.57px] rounded-3xl text-white absolute -bottom-6 hover:bg-white hover:text-[#4771F1] font-bold shadow duration-300 active:scale-95 cursor-pointer'>سازگاریاب
            </button>
        </div>
        <CompatibilityDialog/>
    </>
}