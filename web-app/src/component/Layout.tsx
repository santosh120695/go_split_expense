import React from "react";
import Sidebar from "./Sidebar.tsx";
import Navbar from "./Navbar.tsx";
import { useSidebarStore } from "../store/useStore.ts";

type LayoutProps = {
    children: React.ReactNode;
}

export default function Layout(props: LayoutProps) {
    const { children } = props;
    const { isOpen, toggle } = useSidebarStore();

    return (
        <React.Fragment>
            <div className="layout flex flex-row items-start">
                <div className="md:w-64 h-screen flex-shrink-0 md:block">
                    <Sidebar />
                </div>
                {isOpen && (
                    <div
                        className="fixed inset-0 bg-black opacity-50 z-40 md:hidden"
                        onClick={toggle}
                    ></div>
                )}
                <div className="flex-grow">
                    <Navbar />
                    {children}
                </div>
            </div>
        </React.Fragment>
    );
}