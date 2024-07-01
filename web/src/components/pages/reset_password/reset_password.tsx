'use client'
import React, { useState,Fragment } from 'react'
import { Switch,Dialog,Transition } from '@headlessui/react'
import eye from '@/assets/eye.svg';
import eyeSlash from '@/assets/eye-slash.svg';
import Image from 'next/image'
import Link from 'next/Link'

import { useRouter } from 'next/navigation'

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

const  ResetPassword = () => {
    let [isOpen, setIsOpen] = useState(false)

    function closeModal() {
        setIsOpen(false)
    }

    function openModal() {
        setIsOpen(true)
    }

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
                                        <h1 className="text-center text-3xl font-semibold leading-6 text-cyan-50 pb-4">Reset Password</h1>
                                        <label htmlFor="username" className="block text-sm font-semibold leading-6 text-cyan-50">
                                            Enter your email address to receive a password reset link.
                                        </label>
                                        <Transition appear show={isOpen} as={Fragment}>
                                            <Dialog as="div" className="relative z-10 bg-cyan-900" onClose={closeModal}>
                                                <Transition.Child
                                                    as={Fragment}
                                                    enter="ease-out duration-300"
                                                    enterFrom="opacity-0"
                                                    enterTo="opacity-100"
                                                    leave="ease-in duration-200"
                                                    leaveFrom="opacity-100"
                                                    leaveTo="opacity-0"
                                                >
                                                    <div className="fixed inset-0 bg-black/25" />
                                                </Transition.Child>

                                                <div className="fixed inset-0 overflow-y-auto">
                                                    <div className="flex min-h-full items-center justify-center p-4 text-center">
                                                        <Transition.Child
                                                            as={Fragment}
                                                            enter="ease-out duration-300"
                                                            enterFrom="opacity-0 scale-95"
                                                            enterTo="opacity-100 scale-100"
                                                            leave="ease-in duration-200"
                                                            leaveFrom="opacity-100 scale-100"
                                                            leaveTo="opacity-0 scale-95"
                                                        >
                                                            <Dialog.Panel className="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all">
                                                                <Dialog.Title
                                                                    as="h3"
                                                                    className="text-lg font-medium leading-6 text-gray-900"
                                                                >
                                                                    Sending successful
                                                                </Dialog.Title>
                                                                <div className="mt-2">
                                                                    <p className="text-sm text-gray-500">
                                                                        An email with a reset link has been sent to the email address,You should receive is shortly.
                                                                    </p>
                                                                </div>

                                                                <div className="mt-4">
                                                                    <button
                                                                        type="button"
                                                                        className="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                                                                        onClick={closeModal}
                                                                    >
                                                                        <a href="/login">Back to sign in </a>
                                                                    </button>
                                                                </div>
                                                            </Dialog.Panel>
                                                        </Transition.Child>
                                                    </div>
                                                </div>
                                            </Dialog>
                                        </Transition>
                                        <div className="mt-2.5">
                                            <input
                                                type="email"
                                                name="email"
                                                id="username"
                                                autoComplete="organization"
                                                placeholder="your@domain.com"
                                                className="w-full bg-cyan-50 rounded-md px-3.5 py-2 text-slate-800"
                                                required
                                            />
                                        </div>
                                    </div>
                                </div>
                                <div className="mt-10">
                                    <button
                                        type="button"
                                        onClick={openModal}
                                        className="block w-full rounded-md bg-cyan-500 hover:bg-cyan-600 px-3.5 py-2.5 text-center text-sm font-semibold text-white shadow-sm focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                                    >
                                        Send
                                    </button>
                                </div>
                            </form>
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

export default ResetPassword;