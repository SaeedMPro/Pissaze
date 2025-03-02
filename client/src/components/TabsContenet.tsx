'use client';
// Redux
import {useSelector} from "react-redux";
import {RootState} from "@/redux/store";
// React
import React, {lazy, Suspense} from "react";
//components
const LazyProfile = lazy(() => import('@/components/Profile'));
const LazyExplore = lazy(() => import('@/components/Explore'));
import Loading from "@/components/Loading";

export default function TabContent() {
    const {tab} = useSelector((state: RootState) => state.tabs);

    // Tab-to-Component mapping
    const tabComponents: Record<string, React.ReactNode> = {
        profile: <LazyProfile/>, // or <LazyDesks />
        explore: <LazyExplore/>,
    };

    return (
        <Suspense fallback={<Loading/>}>
            <div className='w-full h-full'>
                {tabComponents[tab]}
            </div>
        </Suspense>
    );
}
