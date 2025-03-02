import {Tabs, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {SubTabsDB} from "@/data/db";

export default function Profile() {
    return <Tabs defaultValue="user_information" className="w-full">
        <TabsList className='w-full flex gap-4 flex-row-reverse rounded-b-2xl py-7 shadow-xl'>
            {
                SubTabsDB.map((item) => (
                    <TabsTrigger key={item.id} value={item.value}
                                 className=' cursor-pointer flex items-center justify-center data-[state=active]:bg-[#4771F1] data-[state=active]:text-white'
                    >{item.name}</TabsTrigger>
                ))
            }
        </TabsList>
    </Tabs>
}