'use client'
import {useState, useEffect} from "react";
//redux
import {useDispatch, useSelector} from "react-redux";
import {AppDispatch, RootState} from "@/redux/store";
import {updateCompatDialog} from "@/redux/features/modalSlice";
import {updateProducts, updateCategory} from "@/redux/features/campatible";
//components
import CompatibilityDialog from "@/components/CompatibilityDialog";
import axios from "axios";
import Image from "next/image";
import Select from 'react-select'
import makeAnimated from 'react-select/animated';
import {SkeletonBox} from "@/utils/Skeleton";


export default function Explore() {
    const dispatch = useDispatch<AppDispatch>();
    const [products, setProducts] = useState<any[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const token = localStorage.getItem("token");
    const isVip = localStorage.getItem("is_vip");
    const animatedComponents = makeAnimated();
    const {productsChoice} = useSelector((state: RootState) => state.compatible)


    const options = [
        {value: 'Motherboard', label: 'Motherboard'},
        {value: 'CPU', label: 'CPU'},
        {value: 'GPU', label: 'GPU'},
        {value: 'RAM Stick', label: 'RAM Stick'},
        {value: 'Cooler', label: 'Cooler'},
        {value: 'Power Supply', label: 'Power Supply'},
        {value: 'Case', label: 'Case'},
        {value: 'SSD', label: 'SSD'},
        {value: 'HDD', label: 'HDD'},
    ]


    // Product selection handler
    const handleProductSelect = (productId: number) => {
        const isSelected = productsChoice.includes(productId);
        const updatedProducts = isSelected
            ? productsChoice.filter(id => id !== productId)
            : [...productsChoice, productId];
        dispatch(updateProducts(updatedProducts));
    };

    const handleCategorySelect = (selectedOptions: any) => {
        const selectedValues = selectedOptions.map((option: any) => option.value);
        dispatch(updateCategory(selectedValues));
    };

    useEffect(() => {
        const fetchProduct = async () => {
            try {
                setLoading(true);
                const response = await axios.get(`${process.env.NEXT_PUBLIC_URL}/product/list`, {headers: {Authorization: token}});
                setProducts(response.data.data);
                setLoading(false);
            } catch (e) {
                console.log(e)
            } finally {
                setLoading(false);
            }
        }
        fetchProduct()
    }, []);

    return <>
        <div className='size-full flex flex-col'>
            <div className='w-full bg-muted rounded-b-xl flex items-center justify-center h-12 relative shadow-xl'>
                {isVip && <button
                    onClick={() => dispatch(updateCompatDialog(true))}
                    className='bg-[#4771F1] w-64 h-[47.57px] rounded-3xl text-white absolute -bottom-6 hover:bg-white hover:text-[#4771F1] font-bold shadow duration-300 active:scale-95 cursor-pointer'>سازگاریاب
                </button>}
            </div>
            <div className='flex flex-col w-full gap-5 my-10 px-10 py-4'>
                {
                    isVip && <Select
                        options={options}
                        isMulti
                        className="basic-multi-select"
                        classNamePrefix="select"
                        components={animatedComponents}
                        closeMenuOnSelect={false}
                        placeholder='فیلتر'
                        onChange={handleCategorySelect}
                        isSearchable
                        noOptionsMessage={() => "دسته بندی یافت نشد"}
                    />
                }
                {
                    loading
                        ? Array.from({length : 4}).map((_,index)=>(
                            <div key={index} className='w-full bg-[#D9D9D9]  flex justify-between items-center relative px-10 py-4 text-sm'>
                                <div className='flex gap-14 items-center'>
                                    <SkeletonBox className='w-20 h-20'/>
                                    <ul className='flex flex-col gap-2'>
                                        <li className='flex items-center gap-3'>
                                            <span className='font-bold'>برند:</span>
                                            <SkeletonBox className='w-20 h-5'/>
                                        </li>
                                        <li className='flex items-center gap-3'>
                                            <span className='font-bold'>مدل:</span>
                                            <SkeletonBox className='w-20 h-5'/>
                                        </li>
                                        <li className='flex items-center gap-3'>
                                            <span className='font-bold'>دسته بندی:</span>
                                            <SkeletonBox className='w-20 h-5'/>
                                        </li>
                                    </ul>
                                </div>
                                <ul className='flex flex-col item-center gap-2'>
                                    <li className='flex items-center gap-3'>
                                        <span className='font-bold'>موجودی:</span>
                                        <SkeletonBox className='w-20 h-5'/>
                                    </li>
                                    <li className='flex items-center gap-3'>
                                        <span className='font-bold'>قیمت:</span>
                                        <SkeletonBox className='w-20 h-5'/>
                                    </li>
                                </ul>
                                <SkeletonBox className='absolute top-2 right-2  w-6 h-6 rounded-full '/>


                            </div>
                        ))
                        : products.length > 0 && products.map((product: any, index) => (
                        <div key={product?.id}
                             className='w-full bg-[#D9D9D9] rounded-md flex justify-between items-center relative px-10 py-4'>
                            {isVip && (
                                <input
                                    type="checkbox"
                                    checked={productsChoice.includes(product.id)}
                                    onChange={() => handleProductSelect(product.id)}
                                    className="absolute left-2 top-2 w-5 h-5 cursor-pointer"
                                />
                            )}
                            <div className='flex gap-14 items-center'>
                                <Image src='/images/picture.png' alt='picture' width='80' height='80'/>
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
                    ))
                }
            </div>
        </div>

        <CompatibilityDialog/>
    </>
}