'use client'
import {useEffect, useState} from "react";
import axios from "axios";
import Spinner from "@/components/Spinner";
import Image from "next/image";

export default function UserInformation() {
    const [userInfo, setUserInfo] = useState<any | null>(null)
    const [loading, setLoading] = useState<boolean>(true);
    const token = localStorage.getItem("token");
    const isVip = localStorage.getItem("is_vip");

    useEffect(() => {
        const fetchUserInformation = async () => {
            try {
                setLoading(true);
                const response = await axios.get('http://localhost:8082/api/client/', {headers: {Authorization: token}});
                setUserInfo(response.data.data);
                setLoading(false);
            } catch (e) {
                console.log(e)
            } finally {
                setLoading(false);
            }
        }
        fetchUserInformation()
    }, []);

    if (loading) <Spinner/>;

    if(!isVip) return <div className='w-full h-full p-8'>
        <div className='size-full flex flex-col  gap-8'>
            <div className='flex w-full items-start justify-between'>
                <div className='relative border px-10 py-3 bg-transparent rounded-lg  flex items-center '>
                    <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>زمان ثبت
                        نام:</p>
                    <p className='text-lg font-bold'>{userInfo?.first_name}</p>
                </div>
                <Image src='/images/avatar.png' width='128' height='128' alt='user avatar'/>
                <div className='relative border px-10 py-3 bg-transparent rounded-md flex items-center justify-center'>
                    <p className='text-slate-600 text-xs absolute -top-[11px] right-1 bg-transparent px-1'>موجودی کیف پول شما
                    </p>
                    <p className='text-lg font-bold'>{userInfo?.wallet_balance} <span className='text-sm text-gray-700'>تومان</span></p>
                </div>
            </div>
            <div className='flex items-center justify-center flex-col gap-5 w-full'>
                <div
                    className='relative border px-10 py-3 bg-transparent rounded-md min-w-[200px] flex items-center justify-center'>
                    <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>نام و نام خانوادگی :
                    </p>
                    <p className='text-lg font-bold'>{userInfo?.first_name} {userInfo?.last_name}</p>
                </div>
                <div className='flex items-center w-full justify-center gap-8'>
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md min-w-[350px]  flex items-center justify-center'>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>زمان باقی
                            کد معرفی شما
                        </p>
                        <p className='text-lg font-bold'>{userInfo?.referral_code}</p>
                    </div>
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md min-w-[350px] flex items-center justify-center'>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>زمان باقی
                            تعداد افرادی که تاکنون معرفی کرده اید
                        </p>
                        <p className='text-lg font-bold'>{userInfo?.number_of_refererred}نفر</p>
                    </div>
                </div>
                <div className='w-full flex flex-col gap-5 py-8 px-6 border   rounded-2xl relative '>
                    <p className='text-slate-600 text-sm absolute -top-[11px] right-2 bg-white px-1'>آدرس ها :
                    </p>
                    {
                        userInfo?.addresses && userInfo?.addresses.map((address:any, index:any) => (
                            <div key={index} className='flex items-center justify-start gap-16 font-bold bg-[#D9D9D9] rounded-md px-5 py-3'>
                                <p className='font-bold text-2xl'>{index+1}</p>
                                <div className='felx flex-col items-start'>
                                    <p>{address.province}</p>
                                    <p>{address.remain_address}</p>
                                </div>
                            </div>
                        ))
                    }

                </div>


            </div>

        </div>
    </div>

    return <div className='w-full h-full p-8'>
        <div className='size-full flex flex-col  gap-8'>
            <div className='flex w-full items-start justify-between'>
                <div className='relative border px-10 py-3 bg-transparent rounded-lg  flex items-center '>
                    <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>زمان ثبت
                        نام:</p>
                    <p className='text-lg font-bold'>{userInfo?.client.first_name}</p>
                </div>
                <Image src='/images/avatar.png' width='128' height='128' alt='user avatar'/>
                <div className='relative border px-10 py-3 bg-transparent rounded-md flex items-center justify-center'>
                    <p className='text-slate-600 text-xs absolute -top-[11px] right-1 bg-transparent px-1'>موجودی کیف پول شما
                     </p>
                    <p className='text-lg font-bold'>{userInfo?.client.wallet_balance} <span className='text-sm text-gray-700'>تومان</span></p>
                </div>
            </div>
            <div className='flex items-center justify-center flex-col gap-5 w-full'>
                <div
                    className='relative border px-10 py-3 bg-transparent rounded-md min-w-[200px] flex items-center justify-center'>
                    <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>نام و نام خانوادگی :
                    </p>
                    {
                        isVip && <Image src='/images/vip.svg' alt='vip' width={30} height={30} className='absolute -top-[15px] -left-3 '/>
                    }
                    <p className='text-lg font-bold'>{userInfo?.client.first_name} {userInfo?.client.last_name}</p>
                </div>
                <div
                    className='relative border px-10 py-3 bg-transparent rounded-md min-w-[230px] flex items-center justify-center'>
                    <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>زمان باقی مانده تا اتمام اشتراک ویژه
                    </p>
                    <p className='text-lg font-bold'>{userInfo?.expiration_time}<span
                        className='text-sm text-gray-700'>روز</span></p>
                </div>

                <div
                    className='relative border px-10 py-3 bg-transparent rounded-md min-w-[230px] flex items-center justify-center'>
                    <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>میزان سود شما از اشتراک ویژه در یک ماه اخیر
                    </p>
                    <p className='text-lg font-bold'>{userInfo?.month_profit}<span
                        className='text-sm text-gray-700'>تومان</span></p>
                </div>
                <div className='flex items-center w-full justify-center gap-8'>
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md min-w-[350px]  flex items-center justify-center'>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>زمان باقی
                            کد معرفی شما
                        </p>
                        <p className='text-lg font-bold'>{userInfo?.client.referral_code}</p>
                    </div>
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md min-w-[350px] flex items-center justify-center'>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-transparent px-1'>
                            تعداد افرادی که تاکنون معرفی کرده اید
                        </p>
                        <p className='text-lg font-bold'> {userInfo?.client.number_of_referred}نفر  </p>
                    </div>
                </div>
                <div className='w-full flex flex-col gap-5 py-8 px-6 border   rounded-2xl relative '>
                    <p className='text-slate-600 text-sm absolute -top-[11px] right-2 bg-white px-1'>آدرس ها :
                    </p>
                        {
                        userInfo?.client.addresses && userInfo?.client.addresses.map((address:any, index:any) => (
                            <div key={index} className='flex items-center justify-start gap-16 font-bold bg-[#D9D9D9] rounded-md px-5 py-3'>
                                <p className='font-bold text-2xl'>{index+1}</p>
                                <div className='felx flex-col items-start'>
                                    <p>{address.province}</p>
                                    <p>{address.remain_address}</p>
                                </div>
                            </div>
                        ))
                    }

                </div>


            </div>

        </div>
    </div>
}