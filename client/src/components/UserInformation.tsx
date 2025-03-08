'use client'
import {useCallback, useEffect, useState} from "react";
//components
import axios from "axios";
import {SkeletonBox} from "@/utils/Skeleton";
import Image from "next/image";
import {jalaliMoment} from "@/utils/helper"
import Loading from "@/components/Loading";

export default function UserInformation() {
    const [userInfo, setUserInfo] = useState<any | null>(null)
    const [loading, setLoading] = useState<boolean>(true);
    const [isVip, setIsVip] = useState<boolean | null>(null);
    useEffect(() => {
        // Access localStorage inside useEffect to ensure client-side execution
        const token = localStorage.getItem("token");
        const isVipStatus = localStorage.getItem("is_vip") === "true"; // Add this conversion
        setIsVip(isVipStatus);
        const fetchUserInformation = async () => {
            try {
                setLoading(true);
                const response = await axios.get(`${process.env.NEXT_PUBLIC_URL}/client/`, {headers: {Authorization: token}});
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


    const getExpireDate = useCallback((expire: string) => {
        const expireDate = new Date(expire);
        const today = new Date();

        // Convert both dates to UTC midnight to ignore time differences
        const expireUtc = Date.UTC(expireDate.getUTCFullYear(), expireDate.getUTCMonth(), expireDate.getUTCDate());
        const todayUtc = Date.UTC(today.getUTCFullYear(), today.getUTCMonth(), today.getUTCDate());

        // Calculate difference in days
        const msPerDay = 24 * 60 * 60 * 1000;
        const diffDays = Math.floor((expireUtc - todayUtc) / msPerDay);

        return Math.max(diffDays, 0); // Never return negative days
    },[]);

    if (loading || userInfo == null || isVip == null) return <Loading/>;

    if (!isVip) {
        return <div className='w-full h-full p-8'>
            <div className='size-full flex flex-col  gap-8'>
                <div className='flex w-full items-start justify-between'>
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-lg  flex items-center  min-w-[300px] '>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1  px-1 bg-white'>زمان ثبت
                            نام:</p>
                        {loading ? <SkeletonBox className='w-full h-8'/> :
                            <p dir='ltr' className='text-lg font-bold'>{jalaliMoment(userInfo?.timestamp)}</p>}
                    </div>
                    {loading ? <SkeletonBox className='w-32 h-32 rounded-full'/> :
                        <Image src='/images/avatar.png' width='128' height='128' alt='user avatar'/>}
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md flex items-center justify-center min-w-[300px]'>
                        <p className='text-slate-600 text-xs absolute -top-[11px] right-1 bg-white px-1'>موجودی کیف پول
                            شما
                        </p>
                        {loading ? <SkeletonBox className='w-full h-8'/> :
                            <p className='text-lg font-bold'>{userInfo?.wallet_balance} <span
                                className='text-sm text-gray-700'>تومان</span></p>}
                    </div>
                </div>
                <div className='flex items-center  flex-col gap-5 w-full'>
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md w-full flex items-center justify-center'>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-white px-1'>نام و نام
                            خانوادگی
                            :
                        </p>
                        {loading ? <SkeletonBox className='w-full h-8'/> :
                            <p className='text-lg font-bold'>{userInfo?.first_name} {userInfo?.last_name}</p>}
                    </div>

                    <div className='flex items-center w-full gap-8'>
                        <div
                            className='relative border px-10 py-3 bg-transparent rounded-md w-full  flex items-center justify-center'>
                            <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-white px-1'>زمان باقی
                                کد معرفی شما
                            </p>
                            {loading ? <SkeletonBox className='w-full h-8'/> :
                                <p className='text-lg font-bold'>{userInfo?.referral_code}</p>}
                        </div>
                        <div
                            className='relative border px-10 py-3 bg-transparent rounded-md w-full flex items-center justify-center'>
                            <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-white px-1'>
                                تعداد افرادی که تاکنون معرفی کرده اید
                            </p>
                            {loading ? <SkeletonBox className='w-full h-8'/> :
                                <p className='text-lg font-bold'> {userInfo?.number_of_referred}<span
                                    className='text-sm text-gray-700 mr-1'>نفر</span></p>}
                        </div>
                    </div>
                    <div className='w-full flex flex-col gap-5 py-8 px-6 border   rounded-2xl relative '>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-2 bg-white px-1'>آدرس ها :
                        </p>
                        {
                            loading ? Array.from({length: 4}).map((_, index) => (
                                    <SkeletonBox className='w-full h-12' key={index}/>
                                ))
                                : userInfo?.addresses ? userInfo?.addresses.map((address: any, index: any) => (
                                    <div key={index}
                                         className='flex items-center justify-start gap-16 font-bold bg-[#D9D9D9] rounded-md px-5 py-3'>
                                        <p className='font-bold text-2xl'>{index + 1}</p>
                                        <div className='felx flex-col items-start'>
                                            <p>{address.province}</p>
                                            <p>{address.remain_address}</p>
                                        </div>
                                    </div>
                                )) : <p className='text-lg font-bold'>هیچ آدرسی برای این کاربر ثبت نشده است</p>
                        }

                    </div>
                </div>

            </div>
        </div>
    } else {
        return <div className='w-full h-full p-8'>
            <div className='size-full flex flex-col  gap-8'>
                <div className='flex w-full items-start justify-between'>
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-lg  flex items-center  min-w-[300px] '>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1  px-1 bg-white'>زمان ثبت
                            نام:</p>
                        {loading ? <SkeletonBox className='w-full h-8'/> :
                            <p dir='ltr' className='text-lg font-bold'>{jalaliMoment(userInfo?.client?.timestamp)}</p>}
                    </div>
                    {loading ? <SkeletonBox className='w-32 h-32 rounded-full'/> :
                        <Image src='/images/avatar.png' width='128' height='128' alt='user avatar'/>}
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md flex items-center justify-center min-w-[300px]'>
                        <p className='text-slate-600 text-xs absolute -top-[11px] right-1 bg-white px-1'>موجودی کیف پول
                            شما
                        </p>
                        {loading ? <SkeletonBox className='w-full h-8'/> :
                            <p className='text-lg font-bold'>{userInfo?.client?.wallet_balance} <span
                                className='text-sm text-gray-700'>تومان</span></p>}
                    </div>
                </div>
                <div className='flex items-center  flex-col gap-5 w-full'>
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md w-full flex items-center justify-center'>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-white px-1'>نام و نام
                            خانوادگی
                            :
                        </p>
                        <Image src='/images/vip.svg' alt='vip' width={30} height={30}
                               className='absolute -top-[15px] -left-3 '/>
                        {loading ? <SkeletonBox className='w-full h-8'/> :
                            <p className='text-lg font-bold'>{userInfo?.client?.first_name} {userInfo?.client?.last_name}</p>}
                    </div>
                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md w-full flex items-center justify-center'>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-white px-1'>زمان باقی مانده
                            تا
                            اتمام اشتراک ویژه
                        </p>
                        {loading ? <SkeletonBox className='w-full h-8'/> :
                            <p className='text-lg font-bold'>{getExpireDate(userInfo?.expiration_time)}<span
                                className='text-sm text-gray-700 mr-1'>روز</span></p>}
                    </div>

                    <div
                        className='relative border px-10 py-3 bg-transparent rounded-md w-full flex items-center justify-center'>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-white px-1'>میزان سود شما
                            از
                            اشتراک ویژه در یک ماه اخیر
                        </p>
                        {loading ? <SkeletonBox className='w-full h-8'/> :
                            <p className='text-lg font-bold'>{userInfo?.month_profit}<span
                                className='text-sm text-gray-700 mr-1'>تومان</span></p>}
                    </div>
                    <div className='flex items-center w-full gap-8'>
                        <div
                            className='relative border px-10 py-3 bg-transparent rounded-md w-full  flex items-center justify-center'>
                            <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-white px-1'>زمان باقی
                                کد معرفی شما
                            </p>
                            {loading ? <SkeletonBox className='w-full h-8'/> :
                                <p className='text-lg font-bold'>{userInfo?.client?.referral_code}</p>}
                        </div>
                        <div
                            className='relative border px-10 py-3 bg-transparent rounded-md w-full flex items-center justify-center'>
                            <p className='text-slate-600 text-sm absolute -top-[11px] right-1 bg-white px-1'>
                                تعداد افرادی که تاکنون معرفی کرده اید
                            </p>
                            {loading ? <SkeletonBox className='w-full h-8'/> :
                                <p className='text-lg font-bold'> {userInfo?.client?.number_of_referred}<span
                                    className='text-sm text-gray-700 mr-1'>نفر</span></p>}
                        </div>
                    </div>
                    <div className='w-full flex flex-col gap-5 py-8 px-6 border   rounded-2xl relative '>
                        <p className='text-slate-600 text-sm absolute -top-[11px] right-2 bg-white px-1'>آدرس ها :
                        </p>
                        {
                            loading ? Array.from({length: 4}).map((_, index) => (
                                    <SkeletonBox className='w-full h-12' key={index}/>
                                ))
                                : userInfo?.client?.addresses ? userInfo?.client?.addresses.map((address: any, index: any) => (
                                    <div key={index}
                                         className='flex items-center justify-start gap-16 font-bold bg-[#D9D9D9] rounded-md px-5 py-3'>
                                        <p className='font-bold text-2xl'>{index + 1}</p>
                                        <div className='felx flex-col items-start'>
                                            <p>{address.province}</p>
                                            <p>{address.remain_address}</p>
                                        </div>
                                    </div>
                                )) : <p className='text-lg font-bold'>هیچ آدرسی برای این کاربر ثبت نشده است</p>
                        }

                    </div>
                </div>

            </div>
        </div>

    }

}