'use client'
import {Tabs, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {SubTabsDB} from "@/data/db";
import SubTabContent from "@/components/SubTabsContent";
import {useDispatch} from "react-redux";
import {AppDispatch} from "@/redux/store";
import {updateSubTab} from "@/redux/features/tabSlice";

export default function Profile() {
    const dispatch = useDispatch<AppDispatch>();
    return <>
        <Tabs defaultValue="UserInformation" className="w-full">
            <TabsList className='w-full flex gap-4 flex-row-reverse rounded-b-2xl py-7 shadow-xl'>
                {
                    SubTabsDB.map((item) => (
                        <TabsTrigger key={item.id} value={item.value}
                                     onClick={()=>dispatch(updateSubTab(item.value))}
                                     className=' cursor-pointer flex items-center justify-center data-[state=active]:bg-[#4771F1] data-[state=active]:text-white'
                        >{item.name}</TabsTrigger>
                    ))
                }
            </TabsList>
        </Tabs>
        <SubTabContent/>
    </>
}