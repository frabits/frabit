'use client'
import ArrowIcon from '@/svgs/arrow.inline.svg';
import React, {useEffect,useState} from 'react'

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

// NOTE: Height of this component is used in mobile sidebars components
const Banner = ({ bannerText, bannerUrl,disable }: { bannerText: string; bannerUrl: string,disable:boolean }) => {
    const [display, setDisplay] = useState("1");
    function handleClick() {
        setDisplay("0");
        window.sessionStorage.setItem('displayBanner', "0");
    }

    useEffect(() => {
        const prevValue = window.sessionStorage.getItem('displayBanner');

        if (prevValue) {
            setDisplay(prevValue);
            return;
        }
        //getReleaseVersion();
    }, []);
    return (
        <div className={classNames(display=="1"?"":"hidden","bg-cyan-500 pr-20 pl-20 relative z-20 text-white transition-colors duration-200")}>
            <div className="grid grid-cols-3 justify-items-center">
                <div className="flex items-center w-8"></div>
                <a className="group/link relative -z-10 mx-auto flex  w-full items-center justify-center px-4.5"
                   href={bannerUrl}
                >
                    <p className="flex flex-row justify-items-center items-center justify-center mb-2 mt-2 font-bold text-cyan-50">{bannerText}
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" className="w-5 h-5">
                            <path fill-rule="evenodd" d="M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z" clip-rule="evenodd"/>
                        </svg>
                    </p>
                </a>
                <div className="flex hidden items-center"></div>
                <div className={classNames(disable?"":"invisible","flex absolute inset-y-0 right-3 w-8 hover:border-cyan-600 rounded-xl")}>
                    <button onClick={handleClick}>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                             stroke="currentColor" className="w-6 h-6">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>
                        </svg>
                    </button>
                </div>
            </div>
        </div>
    );
};

export default Banner;
