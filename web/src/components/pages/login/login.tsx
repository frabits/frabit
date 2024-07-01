'use client'
import React, { useState } from 'react'
import { Switch } from '@headlessui/react'
import eye from '@/assets/eye.svg';
import eyeSlash from '@/assets/eye-slash.svg';
import Image from 'next/image'
import Link from 'next/Link'

import { useRouter } from 'next/navigation'

import Sso from './sso'
import Route from '@/lib/route';

const Resource = [
    { name: 'Docs', href: Route.DOCS },
    { name: 'Pricing', href: Route.PRICING },
    { name: 'Support', href: Route.SUPPORT  },
    { name: 'Community', href: Route.COMMUNITY },
    { name: 'Upgrade', href: Route.DOWNLOADS },
    { name: 'Contact', href: Route.CONTACTS },
    { name: 'Privacy', href: Route.PRIVACY },
];

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

const  Login = () => {
    const [login, setLogin] = useState(false)
    const [showPassword, setShowPassword] = useState(false)
    const [enableLdap, setEnableLdap] = useState(false)
    const version = "Ultimate-V2.2.3"
    const router = useRouter()
    return (
        <>
            <div>
                <div className="px-96 py-10">
                    <div className="px-40">
                        <div className="rounded-3xl bg-gradient-to-r from-cyan-900 via-zinc-800 to-cyan-900  p-10">
                            <div className="mx-auto max-w-2xl text-center">
                                <div className="row-span-1">
                                    <img className="justify-center items-center place-self-center " src="/images/logo.svg"/>
                                </div>
                            </div>
                            <form action="#" method="POST" className=" mx-auto max-w-xl">
                                <div className="grid grid-cols-1 gap-y-3 sm:grid-cols-2">
                                    <div className="sm:col-span-2">
                                        <label htmlFor="username" className="block text-sm font-semibold leading-6 text-cyan-50">
                                            Email or username
                                        </label>
                                        <div className="mt-2.5">
                                            <input
                                                type="text"
                                                name="username"
                                                id="username"
                                                autoComplete="organization"
                                                placeholder="email or username"
                                                className="w-full bg-cyan-50 rounded-md px-3.5 py-2 text-slate-800"
                                                required
                                            />
                                        </div>
                                    </div>
                                    <div className="sm:col-span-2">
                                        <label htmlFor="password" className="block text-sm font-semibold leading-6 text-cyan-50">
                                            Password
                                        </label>
                                        <div className="mt-2.5">
                                            <div className="row-span-1 rounded-md flex pr-2 items-center bg-cyan-50 gap-1 justify-center place-self-center">
                                                <input
                                                    type={showPassword ? "text" : "password"}
                                                    name="password"
                                                    id="password"
                                                    autoComplete="password"
                                                    placeholder="password"
                                                    className="w-full bg-cyan-50 rounded-l-md pl-3.5 py-2 text-slate-800"
                                                    required
                                                />
                                                <button
                                                    type="button"
                                                    className="w-8 h-8 p-1 rounded-md hover:bg-slate-200"
                                                    aria-label="Show password"
                                                    onClick={()=>{setShowPassword(!showPassword)}}
                                                >
                                                    <Image
                                                        src={showPassword ? eyeSlash:eye}
                                                        alt="Show password"
                                                        width={30}
                                                        height={30}
                                                    />
                                                </button>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div className="mt-10">
                                    <button
                                        type="button"
                                        onClick={() => router.push('/dashboard')}
                                        className={classNames(login?"animate-ping": "","block w-full rounded-md bg-cyan-500 hover:bg-cyan-600 px-3.5 py-2.5 text-center text-sm font-semibold text-white shadow-sm focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600")}
                                    >
                                        Sign In
                                    </button>
                                </div>
                                <div className="flex flex-cols-2 justify-between items-center px-5">
                                    <label className="flex content-center items-center place-items-center items-stretch gap-2">
                                        <input type="checkbox" className="w-4 h-4 accent-cyan-500" checked={enableLdap} onClick={()=>{setEnableLdap(!enableLdap)}}/>
                                        <p className="items-center text-center row-span-1 text-sm font-semibold text-slate-400">LDAP</p>
                                    </label>
                                    <button
                                        type="button"
                                        onClick={() => router.push('/reset-password')}
                                        className="block  rounded-md py-2.5 hover:underline underline-offset-4 text-center text-sm font-semibold text-slate-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                                    >
                                        Forget your password?
                                    </button>
                                </div>
                            </form>
                            <Sso hasSSO={true}/>
                            <div className="flex border-t-2 border-slate-400 justify-center place-self-center items-center pt-5 mt-5 gap-5">
                                <p className="text-center text-sm font-semibold text-slate-400">{version}</p>
                                <div className="flex items-center m-2">
                                    <select
                                        className="appearance-none  bg-cyan-900 m-2 ">
                                        <option>English</option>
                                        <option>简体中文</option>
                                    </select>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="gap-x-5 justify-center place-self-center items-center pl-10 pr-10 pb-1 py-2">
                    {/** office resoure */}
                    <div className="flex justify-center place-self-center gap-x-9">
                        {Resource.map(({name,href},idx) =>(
                            <div className="flex-row items-start justify-start" key={idx}>
                                <div className="flex flex-cols-2 justify-center place-self-center items-center">
                                    <a  target="_blank"
                                        href={href}>
                                    <p className="w-full text-10 font-bold pb-2 hover:underline underline-offset-4 leading-none tracking-wider text-slate-400">{name}</p>
                                </a>

                                </div>
                            </div>
                        ))}
                    </div>
                    {/** copyrights */}
                    <div className="flex my-5 justify-center place-self-center gap-x-9">
                        <p className="whitespace-nowrap border-t-2 border-slate-400 p-2 text-14 font-medium leading-none tracking-tight text-slate-400">
                            Copyright © {new Date().getFullYear()} Frabit Labs. All rights reserved.
                        </p>
                    </div>
                </div>
            </div>
        </>
    );
};

export default Login;