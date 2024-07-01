"use client"
import { useState } from 'react'
import { Tab } from '@headlessui/react'
import License from "@/components/pages/admin/subscription/license";
import Account from "@/components/pages/admin/subscription/account";
import subscription from "@/components/pages/admin/subscription/index";

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

const subscriptionMethod = [
    {
        id: 1,
        Name: 'Frabit Account',
        comp: <Account />,
    },
    {
        id: 2,
        Name: 'License',
        comp: <License />,
    },

]

const Subscription = () => {
    return (
        <div className="w-full justify-center items-center place-self-center max-w-md px-2 py-16 sm:px-0">
            <Tab.Group>
                <Tab.List className="flex space-x-1 rounded-xl bg-blue-900/20 p-1">
                    {subscriptionMethod.map((menhod,idx) => (
                        <Tab
                            key={idx}
                            className={({ selected }) =>
                                classNames(
                                    'w-full rounded-lg py-2.5 text-sm font-medium leading-5',
                                    'ring-white/60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2',
                                    selected
                                        ? 'bg-white text-blue-700 shadow'
                                        : 'text-blue-100 hover:bg-white/[0.12] hover:text-white'
                                )
                            }
                        >
                            {menhod.Name}
                        </Tab>
                    ))}
                </Tab.List>
                <Tab.Panels className="mt-2">
                    {subscriptionMethod.map((method, idx) => (
                        <Tab.Panel
                            key={idx}
                            className={classNames(
                                'rounded-xl bg-white p-3',
                                'ring-white/60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2'
                            )}
                        >
                            <div>
                                {method.comp}
                            </div>
                        </Tab.Panel>
                    ))}
                </Tab.Panels>
            </Tab.Group>
        </div>
    );
};

export default Subscription;

