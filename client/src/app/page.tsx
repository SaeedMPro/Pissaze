import Image from "next/image";
//icon
import {User} from 'lucide-react';

export default function Home() {
    return <div className='w-full h-screen relative bg-[#2148C0]' style={{backgroundImage: `url('/images/BG.webp')` , backgroundRepeat: 'no-repeat' ,backgroundSize:'cover'}}>
        <div className='size-full flex items-center justify-center'>
            <div className='flex flex-col items-center justify-center gap-10'>
                <Image src='/images/icon.svg' alt='logo' width={120} height={120}/>
                <div className='flex items-center flex-col justify-center gap-10'>
                    <div className='flex items-center border border-white w-[300px] rounded-md px-2'>
                        <User className='text-white'/>
                        <input type='number' placeholder='شماره تلفن' className='w-full text-white px-2 py-1 outline-none'/>
                    </div>
                    <button className='bg-white text-[#2148C0] text-xl shadow-2xl rounded-md h-10 w-full flex items-center justify-center hover:scale-105 duration-300 active:scale-95 '>ورود</button>
                </div>
            </div>
        </div>
    </div>
}
