'use client'
import {useEffect, useState} from "react";
import axios from "axios";
import Spinner from "@/components/Spinner";

export default function UserDiscount() {
    const [loading, setLoading] = useState<boolean>(true);
    const token = localStorage.getItem("token");
    const [userDiscount, setUserDiscount] = useState<any | null>(null)

    useEffect(() => {
        const fetchUserDiscount = async () => {
            try {
                setLoading(true);
                const response = await axios.get('http://localhost:8082/api/client/discountCode', {headers: {Authorization: token}});
                setUserDiscount(response.data.data);
                console.log(response);
                setLoading(false);
            } catch (e) {
                console.log(e)
            } finally {
                setLoading(false);
            }
        }
        fetchUserDiscount()
    }, []);

    if (loading) <Spinner/>;

    return <div className="p-8">
        {/* Counter */}
        <div
            className='relative border px-10 py-3 bg-transparent rounded-md min-w-[230px] flex items-center justify-center'>
            <p className='text-slate-600 text-sm absolute -top-[11px] right-12 bg-white px-1 '>
                تعداد کدهای تخفیف هدیه گرفته شده از سیستم معرفی            </p>
            <p className='text-lg font-bold'>{userDiscount?.number_of_discount_code}</p>
        </div>

        {/* List of Discount Codes */}
        <div className="mt-6 bg-gray-100 p-4 rounded-xl">
            <p className="font-bold text-center mb-2">کد تخفیف‌های شخصی با کمتر از یک هفته مهلت</p>
            {userDiscount?.discount_code?.map((item:any,index:number) => (
                <div key={index} className="relative bg-white p-8 my-2 rounded-lg shadow-md flex flex-col gap-2">
                    <button
                        className="absolute left-4 top-4 text-gray-500"
                        onClick={() => navigator.clipboard.writeText(item.code)}
                    >
                        📋
                    </button>
                    <p className="text-gray-700 text-sm">سریال کد</p>
                    <p className="font-bold">{item.code}</p>
                    <p className="text-gray-700 text-sm">زمان انقضا</p>
                    <p className="text-gray-600">{item.expiration_time}</p>
                    <p className="text-gray-700 text-sm">مقدار</p>
                    <p className="font-bold text-lg">{item.amount}</p>
                    <span
                        className="absolute top-2 right-2 bg-gray-300 w-6 h-6 flex items-center justify-center rounded-full font-bold text-gray-700">
              {index+1}
            </span>
                </div>
            ))}
        </div>
    </div>
}