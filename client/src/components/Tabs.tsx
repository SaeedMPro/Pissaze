'use client'
//shadCn
import {Tabs, TabsList, TabsTrigger} from "@/components/ui/tabs"
//DB
import {TabsDB} from "@/data/db"
//redux
import {useDispatch} from "react-redux";
import {AppDispatch} from "@/redux/store";
import {updateTab} from "@/redux/features/tabSlice";

export default function Tab() {
    const dispatch = useDispatch<AppDispatch>();
    return <Tabs defaultValue="profile" className="w-full">
        <TabsList className='w-full flex justify-between items-center gap-2 '>
            {
                TabsDB.map((item) => (
                    <TabsTrigger key={item.id} value={item.value} className='flex items-center justify-center w-full text-[#244BC5] cursor-pointer'
                                 onClick={() => dispatch(updateTab(item.value))}>{item.name}</TabsTrigger>
                ))
            }
        </TabsList>
    </Tabs>
}
