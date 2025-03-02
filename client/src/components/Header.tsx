'use client'
import {useRouter} from "next/navigation";
//icon
import { LogOut } from 'lucide-react';

export default function Header(){
    const router = useRouter();
    return <div className='w-full bg-[#244BC5] shadow-2xl  px-5 py-3 '>
        <div className='w-full flex items-center justify-center h-full relative'><p
            className='text-white cursor-pointer font-bold text-2xl' onClick={()=>router.push('/')}>سامانه پیساز</p>
            <button className='text-2xl cursor-pointer text-white absolute left-0'><LogOut className='hover:scale-105 active:scale-95 duration-300'/></button>
        </div>
    </div>
}