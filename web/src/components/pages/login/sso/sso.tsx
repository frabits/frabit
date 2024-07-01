import Image from "next/image";

import { ChevronDownIcon, PhoneIcon, PlayCircleIcon,ChevronRightIcon } from '@heroicons/react/20/solid'
import Route from '@/lib/route';
import React from "react";

import githubIcon from '@/assets/social/github.inline.svg';
import microsoftIcon from '@/assets/social/microsoft.inline.svg';
import amazonIcon from '@/assets/social/amazon .inline.svg';
import googleIcon from '@/assets/social/google.inline.svg';
const socialAddrs = [
    {
        name: 'Github',
        href: Route.GITHUB,
        icon: githubIcon,
        enable:true,
    },
    {
        name: 'Amazon',
        href: Route.GITHUB,
        icon: amazonIcon,
        enable:false,
    },
    {
        name: 'Google',
        href: Route.GITHUB,
        icon: googleIcon,
        enable:true,
    },
    {
        name: 'Microsoft',
        href: Route.GITHUB,
        icon: microsoftIcon,
        enable:true,
    },
    {
        name: 'OIDC',
        href: Route.GITHUB,
        icon: microsoftIcon,
        enable:true,
    },
];

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}
// { bannerText, bannerUrl,disable }: { bannerText: string; bannerUrl: string,disable:boolean }
const Sso = ({ hasSSO = false }: { hasSSO: boolean }) =>{
    return (
        <div className={classNames(hasSSO?"":"hidden")}>
            <div className="grid grid-cols-1 justify-between row-span-1 w-full h-full">
                <div className="col-span-1 w-full text-sm font-semibold  place-self-center  leading-6 text-cyan-50">
                    <div className="col-span-1 w-full flex justify-evenly items-center gap-2">
                        <hr className="w-full border-t-2 border-slate-400 pr-4 bg-slate-400"></hr>
                        <p className="text-sm font-semibold text-center items-center leading-6 text-cyan-50">OR</p>
                        <hr className="w-full pl-4 bg-slate-400"></hr>
                    </div>
                    <p className="flex w-full text-xs font-semibold justify-center place-self-center items-center leading-6 text-slate-400">Sign in with below methods:</p>
                </div>
                <div className="col-span-1  p-2 w-full mx-auto flex justify-evenly items-center gap-1">
                    {socialAddrs.map(({ name, href, icon: Icon,enable }, idx) => (
                        <button className={classNames(enable?"bg-cyan-500 hover:bg-cyan-600":"cursor-not-allowed bg-slate-400","rounded-lg p3 h-9 w-20 h-10 ")} disabled={enable}>
                            <div className={classNames(enable?"hidden":"","flex items-stretch gap-1")}>
                                <Image className="flex place-self-center" alt="" width={14} height={14} src={Icon} />
                                <span className="items-center text-xs text-cyan-50">{name}</span>
                            </div>
                            <a className={classNames(enable?"":"hidden","rounded grid rounded-md place-items-center")}
                               key={idx}
                               href={href}
                               rel="noopener noreferrer"
                               target="_blank">
                                <div className="flex items-stretch gap-1">
                                    <Image className="flex place-self-center" alt="" width={14} height={14} src={Icon} />
                                    <span className="items-center text-xs text-cyan-50">{name}</span>
                                </div>
                            </a>
                        </button>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Sso;