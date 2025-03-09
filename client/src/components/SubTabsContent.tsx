'use client';
// Redux
import {useSelector} from "react-redux";
import {RootState} from "@/redux/store";
// React
import React, {lazy, Suspense} from "react";
import Loading from "@/components/Loading";
//components
const LazyUserInformation = lazy(() => import('@/components/UserInformation'));
const LazyUserDiscount = lazy(() => import('@/components/UserDiscount'));
const LazyUserShopping = lazy(()=>import('@/components/UserShopping'))

export default function SubTabContent() {
    const subTab = useSelector((state: RootState) => state.tabs.subTab);

    // Tab-to-Component mapping
    const tabComponents: Record<string, React.ReactNode> = {
        UserInformation: <LazyUserInformation/>, // or <LazyDesks />
        UserShopping: <LazyUserShopping/>,
        UserDiscount: <LazyUserDiscount/>,
    };

    return (
        <Suspense fallback={<Loading/>}>
            <div className='w-full h-full'>
                {tabComponents[subTab]}
            </div>
        </Suspense>
    );
}
