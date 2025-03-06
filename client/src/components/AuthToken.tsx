'use client'
import React from "react";
import {useEffect} from "react";
import {useRouter} from "next/navigation";

export default function AuthToken({children}: { children: React.ReactNode }) {
    const router = useRouter();

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (!token) router.push("/login");
    }, [router])


    return children


}