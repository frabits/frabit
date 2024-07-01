import Route from '@/lib/route';

export default function NotFound() {
    return (
        <>
            <main className="grid min-h-full place-items-center bg-gradient-to-r from-cyan-900 via-zinc-800 to-cyan-900 px-6 py-24 sm:py-32 lg:px-8">
                <div className="text-center p-10">
                    <p className="text-base font-semibold text-cyan-50">404</p>
                    <h1 className="mt-4 text-3xl font-bold tracking-tight text-cyan-50 sm:text-5xl">Page not found</h1>
                    <p className="mt-6 text-base leading-7 text-cyan-50">Sorry, we couldn’t find the page you’re looking for.</p>
                    <div className="mt-10 flex items-center justify-center gap-x-6">
                        <a
                            href={Route.INDEX}
                            className="rounded-md bg-cyan-500 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-cyan-600 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-cyan-600"
                        >
                            Go back home
                        </a>
                        <a href={Route.CONTACTS} className="text-sm font-semibold text-cyan-50">
                            Contact support <span aria-hidden="true">&rarr;</span>
                        </a>
                    </div>
                </div>
            </main>
        </>
    )
}
