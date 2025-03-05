'use client'
import {useEffect , useRef} from "react";
//redux
import {useDispatch , useSelector} from "react-redux";
import {AppDispatch , RootState} from "@/redux/store";
import {updateCompatDialog} from "@/redux/features/modalSlice";
//shadCN
import {
    Dialog,
    DialogContent, DialogDescription,
    DialogHeader,
    DialogTitle,
} from "@/components/ui/dialog"
import * as VisuallyHidden from "@radix-ui/react-visually-hidden";
//icon
import { X } from 'lucide-react';




export default function CompatibilityDialog() {
    const dispatch = useDispatch<AppDispatch>();
    const compatDialog = useSelector((state:RootState)=>state.modal.CompatDialog)
    const modalRef = useRef<HTMLDivElement | null>(null);
    useEffect(() => {
        const removeIsOpen = (event: MouseEvent) => {
            // Check if the modal is open and whether the clicked element is outside the modal content
            if (compatDialog && modalRef.current && !modalRef.current.contains(event.target as Node)) {
                dispatch(updateCompatDialog(false));
            }
        };

        // Attach the event listener to the document instead of the window
        document.addEventListener("mousedown", removeIsOpen);

        return () => {
            document.removeEventListener("mousedown", removeIsOpen);
        };
    }, [compatDialog,dispatch]);


    return <Dialog open={compatDialog}>
        <DialogContent className='[&>button]:hidden p-0 rounded-md' ref={modalRef}>
            <DialogHeader>
                <DialogTitle className=' w-full h-10 bg-[#1D3A5B] flex items-center justify-center text-white relative py-4'>
                    <button
                        onClick={()=>dispatch(updateCompatDialog(false))}
                        className='absolute right-2 top-1/2 -translate-y-1/2 p-[1px]  bg-[#F44336] rounded-full text-white active:scale-90 duration-300 cursor-pointer'><X/></button>
                    <span>لیست سازگاری ها</span>
                </DialogTitle>
                <VisuallyHidden.Root>
                    <DialogDescription>X</DialogDescription>
                </VisuallyHidden.Root>


            </DialogHeader>
        </DialogContent>
    </Dialog>

}