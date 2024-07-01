const Footer = () => (
    <footer className="w-full border-t-2  border-slate-400">
        <div className="gap-x-5 items-center py-2">
            {/** copyrights */}
            <div className="flex justify-center place-self-center items-center gap-x-9">
                <p className="whitespace-nowrap text-14 font-medium leading-none tracking-tight text-cyan-50">
                    Copyright © {new Date().getFullYear()} Frabit Labs. All rights reserved.
                </p>
            </div>
        </div>
    </footer>
);

export default Footer;