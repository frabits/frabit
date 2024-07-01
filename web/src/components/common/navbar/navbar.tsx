"use client"

import Top from './top'
import Second from './second'

import PROMO_DATA from "@/lib/promo-data";
import Banner from "@/components/common/banner";

export default function Navbar() {
    const topBanner = PROMO_DATA.LICENSE_INFO
    const licenseDesc = "Your license will be expired at: "
    const licenseExpired = "Your license already expired at: "
    const ExpireDate = "2024-05-28"

    topBanner.title = licenseExpired + ExpireDate
    topBanner.disabled = false
    return (
        <div className="">
            {topBanner && <Banner bannerText={topBanner.title} bannerUrl={topBanner.link} disable={topBanner.disabled} />}
            <Top/>
            <Second/>
        </div>
    )
}
