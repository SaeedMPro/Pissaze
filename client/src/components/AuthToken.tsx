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

    // Add this in your _app.tsx or main layout component
    useEffect(() => {
        const checkTokenExpiration = () => {
            const expiration = localStorage.getItem('token_expiration');
            if (expiration && new Date() > new Date(expiration)) {
                localStorage.removeItem('token');
                localStorage.removeItem('token_expiration');
                localStorage.removeItem('is_vip');
                router.push('/login');
            }
        };

        // Check on mount and every minute
        checkTokenExpiration();
        const interval = setInterval(checkTokenExpiration, 60000);
        return () => clearInterval(interval);
    }, []);


    return children


}