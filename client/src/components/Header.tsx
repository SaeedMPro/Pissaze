'use client'
import {useRouter} from "next/navigation";
//icon
import { LogOut } from 'lucide-react';
import {toast} from "sonner";

export default function Header(){
    const router = useRouter();

    const handleLogout = () => {
        localStorage.clear();
        router.push('/login');
        toast.success('خروج شما با موفقیت انجام شد.')
        window.location.reload();
    }
    return <div className='w-full bg-[#244BC5] shadow-2xl  px-5 py-3 '>
        <div className='w-full flex items-center justify-center h-full relative'><p
            className='text-white cursor-pointer font-bold text-2xl hover:scale-110 duration-500' onClick={()=>router.push('/')}>سامانه پیساز</p>
            <button
                onClick={()=>handleLogout()}
                className='text-2xl cursor-pointer text-white absolute left-0'><LogOut className='hover:scale-105 active:scale-95 duration-300'/></button>
        </div>
    </div>
}