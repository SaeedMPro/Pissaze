'use client'
import Image from "next/image";
import axios from "axios";
import {useState} from "react";
import {useRouter} from "next/navigation";
//icon
import {User} from 'lucide-react';
//utils
import { toast } from "sonner"
//utils
import Spinner from "@/components/Spinner";


export default function Login(){
    const [phoneNumber, setPhoneNumber] = useState<string>('');
    const [loading , setLoading] = useState<boolean>(false);
    const router = useRouter();

    const handleLogin = async () => {
        try {
            setLoading(true);
            const response = await axios.post("http://localhost:8082/api/login/", {"phone_number": phoneNumber});
            if(response.status === 200){
                setPhoneNumber("");
                localStorage.setItem('token', response.data.data.token);
                localStorage.setItem('is_vip', response.data.data.is_vip.toString());
                toast.success("وررود با موفقیت انجام شد.")
                setLoading(false);
                router.push("/");
            }
        }catch(e){
            console.log(e)
            toast.error("حساب کاربری با این شماره وجود ندارد.")
        }finally {
            setLoading(false);
        }
    }

    return <div className='w-full h-screen relative bg-[#2148C0]' style={{backgroundImage: `url('/images/BG.webp')` , backgroundRepeat: 'no-repeat' , backgroundSize:'cover'}}>
        <div className='size-full flex items-center justify-center'>
            <div className='flex flex-col items-center justify-center gap-10'>
                <Image src='/images/icon.svg' alt='logo' width={120} height={120}/>
                <div className='flex items-center flex-col justify-center gap-10'>
                    <div className='flex items-center border h-10 border-white w-[300px] rounded-md px-2'>
                        <User className='text-white'/>
                        <input type='text'
                               onChange={(e) => {
                                   setPhoneNumber(e.target.value)
                               }}
                               placeholder='شماره تلفن'
                               className='w-full text-white px-2 py-1 outline-none'/>
                    </div>
                    <button
                        disabled={loading}
                        onClick={() => handleLogin()}
                        className='bg-white text-[#2148C0] text-xl shadow-2xl rounded-md h-10 w-full flex items-center justify-center hover:scale-105 duration-300 active:scale-95 '>{loading ? <Spinner/> : "ورود"}</button>
                </div>
            </div>
        </div>
    </div>
}