import React from "react";
import Sidebar from "./Sidebar.tsx";
import Navbar from "./Navbar.tsx";

type LayoutProps = {
    children: React.ReactNode;
}

export default function Layout(props: LayoutProps) {
    const {children} = props;
    return <React.Fragment>
        <div className="layout flex flex-row items-start">
            <div className="w-64 h-screen flex-shrink-0">
                <Sidebar />
            </div>
            <div className="flex-grow">
                <Navbar />
                {children}
            </div>
        </div>
    </React.Fragment>
}