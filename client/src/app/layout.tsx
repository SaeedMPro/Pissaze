import React from "react";
import type {Metadata} from "next"
import "./globals.css";
import localFont from 'next/font/local'
import ReduxProvider from "@/redux/provider";
import {Toaster} from "@/components/ui/sonner"
import AuthToken from '@/components/AuthToken'


export const metadata: Metadata = {
    title: "Pissaze",
    description: "This is Pissaze, an online shopping platform",
    icons: [
        {
            rel: 'icon',
            type: 'image/webp',
            url: '/images/main-logo.webp',
            sizes: 'any',
        },
        // Fallback for Safari and older browsers
        {
            rel: 'icon',
            type: 'image/png',
            url: '/images/main-logo.png',
            sizes: 'any',
        },
    ]
};

const myFont = localFont({
    src: '../../public/fonts/YekanBakhFaNum-Regular.woff2',
})

export default function RootLayout({children}: Readonly<{ children: React.ReactNode; }>) {
    return (
        <html lang="fa" dir="rtl">
        <body className={myFont.className}>
        <ReduxProvider>
            <AuthToken>
                {children}
                <Toaster/>
            </AuthToken>
        </ReduxProvider>
        </body>
        </html>
    );
}
