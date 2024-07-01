'use client'
import ArrowIcon from '@/svgs/arrow.inline.svg';
import React, {useEffect,useState} from 'react'
import Image from "next/image";

import {MENU} from "@/lib/menus";

const Sidebar = () => {
    const [display, setDisplay] = useState(true);
    return (
        <div className="fixed h-full md:flex md:flex-shrink-0 pl-10 text-white">
            <div className="flex flex-col w-52">
                <div className="flex-1 flex flex-col py-0 overflow-y-auto">
                    {MENU.sidebar.map(({ name, items}, idx) => (
                        <div className="py-4" key={idx}>
                            <span>{name}</span>
                            {items.map((item, idx) => (
                                <div className="pl-2" key={idx}>
                                    <div className="flex h-4 py-3 items-stretch gap-1">
                                        {item.icon}
                                        <span className="items-center text-xs text-cyan-50">{item.name}</span>
                                    </div>
                                </div>
                            ))}
                        </div>
                    ))}
                </div>
                <div className="bg-cyan-900">Ultimate-V2.2.1</div>
            </div>

           </div>
    );
};

export default Sidebar;