'use client'
import {useEffect, useRef, useState} from "react";
//redux
import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "@/redux/store";
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
import {X} from 'lucide-react';
//components
import axios from "axios";
import Image from "next/image";
import {SkeletonBox} from "@/utils/Skeleton";

export default function CompatibilityDialog() {
    const dispatch = useDispatch<AppDispatch>();
    const [products, setProducts] = useState<any[]>([]);
    const {categoryChoice, productsChoice} = useSelector(
        (state: RootState) => state.compatible
    );
    const token = localStorage.getItem("token");
    const [loading, setLoading] = useState<boolean>(true);
    const compatDialog = useSelector((state: RootState) => state.modal.CompatDialog)
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
    }, [compatDialog, dispatch]);

    useEffect(() => {
        const sendCompatible = async () => {
            setLoading(true);
            try {
                const res = await axios.post(`${process.env.NEXT_PUBLIC_URL}/product/compatible`, {
                    "filter_list_category": categoryChoice,
                    'product_id': productsChoice
                }, {headers: {Authorization: token},})
                if (res.data.data == null) {
                    setProducts([])
                } else {
                    setProducts(res.data.data);
                }
                setLoading(false);
            } catch (e) {
                console.log(e)
                setProducts([])
            } finally {
                setLoading(false);
            }
        }
        sendCompatible();
    }, [compatDialog]);


    return <Dialog open={compatDialog}>
        <DialogContent className='[&>button]:hidden p-0 overflow-hidden  rounded-md' ref={modalRef}>
            <DialogHeader>
                <DialogTitle
                    className=' w-full h-10 bg-[#1D3A5B] flex items-center justify-center text-white relative py-4'>
                    <button
                        onClick={() => dispatch(updateCompatDialog(false))}
                        className='absolute right-2 top-1/2 -translate-y-1/2 p-[1px]  bg-[#F44336] rounded-full text-white active:scale-90 duration-300 cursor-pointer'>
                        <X/></button>
                    <span>لیست سازگاری ها</span>
                </DialogTitle>
                <VisuallyHidden.Root>
                    <DialogDescription>X</DialogDescription>
                </VisuallyHidden.Root>
                <div className='flex items-center justify-center flex-col gap-3 max-h-[450px] overflow-auto'>
                    {
                        loading
                            ? Array.from({length : 2}).map((_,index)=>(
                                <div key={index} className='w-full bg-[#D9D9D9]  flex justify-between items-center relative px-10 py-4 text-sm'>
                                    <div className='flex gap-14 items-center'>
                                        <SkeletonBox className='w-10 h-10'/>
                                        <ul className='flex flex-col gap-2'>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>برند:</span>
                                                <SkeletonBox className='w-10 h-5'/>
                                            </li>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>مدل:</span>
                                                <SkeletonBox className='w-10 h-5'/>
                                            </li>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>دسته بندی:</span>
                                                <SkeletonBox className='w-10 h-5'/>
                                            </li>
                                        </ul>
                                    </div>
                                    <ul className='flex flex-col item-center gap-2'>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>موجودی:</span>
                                               <SkeletonBox className='w-10 h-5'/>
                                            </li>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>قیمت:</span>
                                                <SkeletonBox className='w-10 h-5'/>
                                            </li>
                                    </ul>
                                    <SkeletonBox className='absolute top-2 right-2  w-6 h-6 rounded-full '/>


                                </div>
                            ))
                            : products.length > 0 ? products.map((product: any, index) => (
                                <div key={product?.id}
                                     className='w-full bg-[#D9D9D9]  flex justify-between items-center relative px-10 py-4 text-sm'>
                                    <div className='flex gap-14 items-center'>
                                        <Image src='/images/picture.png' alt='picture' width='40' height='40'/>
                                        <ul>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>برند:</span>
                                                <span className='font-bold'>{product?.brand}</span>
                                            </li>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>مدل:</span>
                                                <span className='font-bold'>{product?.model}</span>
                                            </li>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>دسته بندی:</span>
                                                <span className='font-bold'>{product?.category}</span>
                                            </li>
                                        </ul>
                                    </div>
                                    <div className='flex flex-col item-center'>
                                        <ul>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>موجودی:</span>
                                                <span className='font-bold'>{product?.stock_count}</span>
                                            </li>
                                            <li className='flex items-center gap-3'>
                                                <span className='font-bold'>قیمت:</span>
                                                <span className='font-bold'>{product?.current_price} <span
                                                    className='text-sm text-gray-700'>تومان</span></span>
                                            </li>

                                        </ul>
                                    </div>
                                    <span
                                        className="absolute top-2 right-2 bg-white w-6 h-6 flex items-center justify-center rounded-full font-bold text-gray-700">{index + 1}</span>


                                </div>
                            )) : <p className='font-bold text-base py-5'>هیچ سازگاری وجود ندارد</p>
                    }
                </div>


            </DialogHeader>
        </DialogContent>
    </Dialog>

}