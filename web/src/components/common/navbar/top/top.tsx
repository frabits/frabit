"use client"
import { Fragment } from 'react'
import { Disclosure, Popover,Menu, Transition } from '@headlessui/react'
import {CloudArrowUpIcon, LockClosedIcon,DocumentIcon, ServerIcon,CircleStackIcon,UsersIcon,UserGroupIcon} from "@heroicons/react/20/solid";
import {MENU} from "@/lib/menus";
import Notification from "@/components/pages/notifications/notification"
import Workspace from "@/components/pages/workspace";
import { Bars3Icon, BellIcon, XMarkIcon } from '@heroicons/react/24/outline'
import Search from "@/components/common/search";
import { useRouter } from 'next/navigation'

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

const user = {
    name: 'Frabit',
    email: 'admin@frabit.com',
    imageUrl: '/assets/profile.jpg',
}

const userNavigation = [
    { name: 'Profile', href: '#',isExternal:false },
    { name: 'Frabit Home', href: 'https://www.frabit.tech',isExternal:true },
    { name: 'Docs', href: 'https://www.frabit.tech/docs',isExternal:true },
    { name: 'Sign out', href: '/',isExternal:false },
]

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

export default function Top() {
    const router = useRouter()
    const mesg_num = 2
    return (
        <>
            <div className="min-h-full">
                <Disclosure as="nav" className="bg-gradient-to-r from-cyan-900 to-zinc-800">
                    {({ open }) => (
                        <>
                            <div className="mx-auto w-full px-4">
                                <div className="flex h-16 items-center justify-between">
                                    <div className="flex items-center">
                                        <a href="/dashboard" className="-m-1.5 p-1.5">
                                            <div className="flex-shrink-0">
                                                <img
                                                    className="h-16 w-full"
                                                    src="/images/logo-full.svg"
                                                    alt="Your Company"
                                                />
                                            </div>
                                        </a>
                                        <div className="flex text-gray-300">
                                            <Workspace/>
                                        </div>
                                        <div className="hidden md:block">
                                            <div className="ml-10 flex items-baseline space-x-4">
                                                {MENU.navbar.map((item) => (
                                                    <a
                                                        key={item.name}
                                                        href={item.href}
                                                        className={classNames(
                                                            item.current
                                                                ? 'bg-cyan-500 text-cyan-50'
                                                                : 'text-gray-300 hover:bg-cyan-600 hover:text-white',
                                                            'rounded-md px-3 py-2 text-sm font-medium'
                                                        )}
                                                        aria-current={item.current ? 'page' : undefined}
                                                    >
                                                        {item.name}
                                                    </a>
                                                ))}
                                            </div>
                                        </div>
                                    </div>
                                    <Search/>
                                    <div className="">
                                        <div className="ml-4 flex items-center md:ml-6">
                                            <Popover className="relative ">
                                                {({ open }) => (
                                                    <>
                                                        <Popover.Button
                                                            className="group inline-flex items-center rounded-md px-3 py-2 text-base font-medium text-white focus:outline-none focus-visible:ring-2 focus-visible:ring-white/75"
                                                        >
                                                            <div className={classNames(mesg_num == 0 ? "hidden":"","absolute top-0 right-0 bg-cyan-600 flex justify-center place-items-center items-center rounded-full w-4 h-4")}>
                                                                <span className={classNames(mesg_num == 1 ? "":"hidden", "text-center items-center text-sm text-slate-50")}>1</span>
                                                                <span className={classNames(mesg_num == 2 ? "":"hidden", "text-center items-center text-sm text-slate-50")}>2</span>
                                                                <span className={classNames(mesg_num == 3 ? "":"hidden", "text-center place-self-center text-sm text-slate-50")}>3</span>
                                                                <span className={classNames(mesg_num >3 ? "":"hidden", "text-center items-center text-sm text-slate-50")}>3+</span>
                                                            </div>
                                                            <BellIcon className="h-8 w-8" aria-hidden="true" />
                                                        </Popover.Button>
                                                        <Transition
                                                            as={Fragment}
                                                            enter="transition ease-out duration-100"
                                                            enterFrom="transform opacity-0 scale-95"
                                                            enterTo="transform opacity-100 scale-100"
                                                            leave="transition ease-in duration-75"
                                                            leaveFrom="transform opacity-100 scale-100"
                                                            leaveTo="transform opacity-0 scale-95"
                                                        >
                                                            <Popover.Panel className="absolute right-0 z-10 mt-2  origin-top-right rounded-md bg-cyan-800 py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                                                                <div className="relative overflow-y-auto h-[32rem] flex-cols gap-8 bg-white p-7 lg:grid-cols-2">
                                                                    {notifications.map(({id,channel,title,datetime}, idx) => (
                                                                        <div className="w-48 flex gap-6" key={idx}>
                                                                            <Notification id={id} channel={channel} title={title} datetime={datetime}/>
                                                                        </div>
                                                                    ))}
                                                                </div>
                                                            </Popover.Panel>
                                                        </Transition>
                                                    </>
                                                )}
                                            </Popover>

                                            {/* Profile dropdown */}
                                            <Menu as="div" className="relative ml-3">
                                                <div>
                                                    <Menu.Button className="relative flex max-w-xs items-center rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
                                                        <span className="absolute -inset-1.5" />
                                                        <span className="sr-only">Open user menu</span>
                                                        <img className="h-8 w-8 rounded-full" src={user.imageUrl} alt="" />
                                                    </Menu.Button>
                                                </div>
                                                <Transition
                                                    as={Fragment}
                                                    enter="transition ease-out duration-100"
                                                    enterFrom="transform opacity-0 scale-95"
                                                    enterTo="transform opacity-100 scale-100"
                                                    leave="transition ease-in duration-75"
                                                    leaveFrom="transform opacity-100 scale-100"
                                                    leaveTo="transform opacity-0 scale-95"
                                                >
                                                    <Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                                                        {userNavigation.map((item) => (
                                                            <Menu.Item key={item.name}>
                                                                {({ active }) => (
                                                                    <a
                                                                        href={item.href}
                                                                        target={item.isExternal ? 'blank' : undefined}
                                                                        className={classNames(
                                                                            active ? 'bg-gray-100' : '',
                                                                            'block px-4 py-2 text-sm text-gray-700'
                                                                        )}
                                                                    >
                                                                        {item.name}
                                                                    </a>
                                                                )}
                                                            </Menu.Item>
                                                        ))}
                                                    </Menu.Items>
                                                </Transition>
                                            </Menu>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </>
                    )}
                </Disclosure>
            </div>
        </>
    )
}
