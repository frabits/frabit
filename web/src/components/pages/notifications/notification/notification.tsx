'use client'
import React, {useEffect, useState} from 'react'
import Image from 'next/image'

import { ChevronDownIcon } from '@heroicons/react/20/solid'
import { Switch } from '@headlessui/react'

import Route from '@/lib/route';
import {GithubLatestRelease} from '@/types/releaseGithub';
import eyeSlash from "@/assets/eye-slash.svg";
import eye from "@/assets/eye.svg";

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

const  Notification = ({id,channel,title,datetime}:{id:number,channel:string;title:string,datetime:string}) => {
    const notifyLink = "/notification/" + id
    const [Read, setRead] = useState(false)
    return (
        <>
            <div className="w-full bg-cyan-800 p-4">
                <div className="p-2 flex-cols-2 rounded-md">
                    <div className="col-span-1">
                        <button
                            type="button"
                            className="w-8 h-8 p-1 rounded-md hover:bg-slate-200"
                            aria-label="Show password"
                            onClick={()=>{setRead(!Read)}}
                        >
                            <Image
                                src={Read ? eye:eyeSlash}
                                alt="Set already read"
                                width={30}
                                height={30}
                            />
                        </button>
                    </div>
                    <div className="col-span-1">
                        <a className="text-cyan-500 hover:text-cyan-600 p-5" href={notifyLink}>
                            <div className={classNames(Read?"text-slate-500":"text-cyan-50","text-14 font-bold leading-none tracking-wider")}>
                                <div className="grid grid-cols-2">
                                    <p className="col-span-1">{channel}</p>
                                    <p className="col-span-1">{datetime}</p>
                                </div>
                                <div className="">{title}</div>
                            </div>
                        </a>
                    </div>
                </div>
            </div>
        </>
    );
};

export default Notification;