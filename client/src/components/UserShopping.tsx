'use client'
import {useEffect, useState} from "react";
//components
import axios from "axios";
import Loading from "@/components/Loading";

export default function UserShopping() {
    const [cartStatus, setCartStatus] = useState<any[] | null>(null)
    const [lockCart, setLockCart] = useState<any[]>([])
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        const token = localStorage.getItem("token");
        const fetchCartStatus = async () => {
            try {
                setLoading(true);
                const response = await axios.get(`${process.env.NEXT_PUBLIC_URL}/client/cart`, {headers: {Authorization: token}});
                setCartStatus(response.data.data)
                setLoading(false);
            } catch (e) {
                console.log(e)
            } finally {
                setLoading(false);
            }
        }

        const fetchLockCart = async () => {
            try {
                setLoading(true);
                const response = await axios.get(`${process.env.NEXT_PUBLIC_URL}/client/lockCart`, {headers: {Authorization: token}});
                setLockCart(response.data.data)
                setLoading(false);
            } catch (e) {
                console.log(e)
            } finally {
                setLoading(false);
            }
        }

        fetchCartStatus()
        fetchLockCart()
    }, []);


    const getCartStatus = (text: string) => {
        switch (text) {
            case "active":
                return "فعال";
            case "locked":
                return "قفل شده"
            case "blocked" :
                return "مسدود شده"
        }
    }

    if (loading) return <Loading/>;

    return (
        <div className="p-4 ">
            {/* Shopping Cart Status */}
            <div className="bg-gray-100 p-4 rounded-xl mb-6">
                <p className="font-bold text-center mb-2">وضعیت سبد های خرید</p>
                <div className="grid grid-cols-2 gap-2">
                    {cartStatus?.map((status: any, index) => (
                        <div key={index} className="flex justify-between items-center bg-gray-200 p-3 rounded-lg">
                            <span className="text-gray-700 font-bold">{getCartStatus(status.cart_status)}</span>
                            <span
                                className="bg-gray-300 w-8 h-8 flex items-center justify-center rounded-full font-bold text-gray-700">
                {status.cart_number}
              </span>
                        </div>
                    ))}
                </div>
            </div>

            {/* Discount Codes Section */}
            <div className="bg-gray-100 p-4 rounded-xl">
                <p className="font-bold text-center mb-2">تاریخچه خرید</p>
                {lockCart?.length > 0 ? lockCart?.map((cart, index) => (
                    <div key={index} className="bg-white p-4 my-2 rounded-lg shadow-md relative">
                        <p className="text-gray-700 text-sm">مجموع ارزش سبد خرید</p>
                        <p className="font-bold">{cart.total_price}</p>
                        <p className="text-gray-700 text-sm mt-2">محصولات سبد خرید</p>
                        {cart.products.map((product: any, index: any) => (
                            <div key={index} className="flex flex-col gap-1 bg-gray-200 p-8 rounded-lg mt-2 relative">
                                <p className="text-sm font-bold">برند: {product.product.brand}</p>
                                <p className="text-sm">مدل: {product.product.model}</p>
                                <p className="text-sm">تعداد: {product.quantity}</p>
                                <p className="text-sm">قیمت: {product.cart_price}</p>
                                <span
                                    className="absolute top-2 right-2 bg-gray-300 w-6 h-6 flex items-center justify-center rounded-full font-bold text-gray-700">
                  {index + 1}
                </span>
                            </div>
                        ))}
                    </div>
                )) : <p className='font-bold text-lg text-center w-full py-10'>خریدی ثبت نشده است</p>}
            </div>
        </div>
    );
}