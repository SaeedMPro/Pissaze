//components
import Header from "@/components/Header";
import Tab from "@/components/Tabs";
import TabContent from "@/components/TabsContenet";

export default function Home() {
    return <div className="w-full flex flex-col">
        <Header />
        <Tab/>
        <TabContent/>
    </div>
}
