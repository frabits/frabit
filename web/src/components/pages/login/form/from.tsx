import Image from "next/image";

import { ChevronDownIcon, PhoneIcon, PlayCircleIcon,ChevronRightIcon } from '@heroicons/react/20/solid'
import Route from '@/lib/route';
import React from "react";

const From = () =>{
    return (
        <>
            <div className="col-start-2 rounded-3xl text-4xl items-center ">
                <Image
                    src='/images/page/main/illustration.svg'
                    alt=""
                    width={20}
                    height={20}
                    className="dark:bg-slate-800 p-3 md:w-full sm:-mt-2.5"
                />
            </div>
        </>
    );
};

export default From;