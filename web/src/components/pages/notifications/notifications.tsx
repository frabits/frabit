'use client'
import React, { useState } from 'react'
import { ChevronDownIcon } from '@heroicons/react/20/solid'
import { Switch } from '@headlessui/react'
import Notification from "./notification";

import Route from '@/lib/route';

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

const notifications = [
    {id: 101234, channel: "Orders",title:"Just a demo",datetime:"2023/04/03 14:05:56"},
    {id: 101234, channel: "Orders",title:"Just a demo",datetime:"2023/04/03 14:05:56"},
    {id: 101234, channel: "Orders",title:"发布通知",datetime:"2023/04/03 14:05:56"},
    {id: 101234, channel: "Orders",title:"Just a demo",datetime:"2023/04/03 14:05:56"},
    {id: 101234, channel: "Orders",title:"Just a demo",datetime:"2023/04/03 14:05:56"},
    {id: 101234, channel: "Orders",title:"Just a demo",datetime:"2023/04/03 14:05:56"},
    {id: 101234, channel: "Orders",title:"Just a demo",datetime:"2023/04/03 14:05:56"},
    {id: 101234, channel: "Orders",title:"Just a demo",datetime:"2023/04/03 14:05:56"},
    ]
;

const  Notifications = () => {

    return (
        <>
        <div className="isolate px-6 py-24 sm:py-32 lg:px-8">
            <div className="mx-auto max-w-2xl text-center">
                <h2 className="text-3xl font-bold tracking-tight text-cyan-50 sm:text-4xl">My Notifications</h2>
                <p className="mt-2 text-lg leading-8 text-cyan-50">
                    Aute magna irure deserunt veniam aliqua magna enim voluptate.
                </p>
            </div>
            <div className="flex grid grid-cols place-items-center gap-2">
                {notifications.map(({id,channel,title,datetime}, idx) => (
                    <div className="w-full px-52 gap-x-8" key={idx}>
                        <Notification id={id} channel={channel} title={title} datetime={datetime}/>
                    </div>
                ))}
            </div>
        </div>
        </>
    );
};

export default Notifications;