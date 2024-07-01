"use client"
import React, {useState} from "react";
import { useRouter } from 'next/navigation'

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

const License = () => {
    const router = useRouter()
    const [apply, setApply] = useState(false)
    return (
        <div className="mt-2.5">
                            <textarea
                                name="message"
                                id="message"
                                rows={4}
                                className="block w-full rounded-md border-0 px-3.5 py-2 text-cyan-50 shadow-sm ring-1 ring-inset ring-cyan-50 placeholder:text-cyan-50 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                defaultValue="Paste or drop license here"
                            />
            <div className="mt-10">
                <button
                    type="button"
                    onClick={() => router.push('/dashboard')}
                    className={classNames(apply?"animate-ping": "","block w-full rounded-md bg-cyan-500 hover:bg-cyan-600 px-3.5 py-2.5 text-center text-sm font-semibold text-white shadow-sm focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600")}
                >
                    Apply
                </button>
            </div>
        </div>
    );
};

export default License;
