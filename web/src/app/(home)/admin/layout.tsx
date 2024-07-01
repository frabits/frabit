import type { Metadata } from 'next'
import React from "react";

import Navbar from '@/components/common/navbar/navbar';
import Sidebar from '@/components/common/sidebar/admin/sidebar';
import Footer from '@/components/common/footer/footer';

import '@/styles/main.css';
import PROMO_DATA from "@/lib/promo-data";

export const metadata: Metadata = {
  title: 'Frabit | The next-gen database automatic platform',
  description: 'The next-gen database automatic platform',
}

export default function AdminLayout({ children,}: { children: React.ReactNode }) {
  return (
      <div className="relative h-screen overflow-hidden flex flex-col">
          <div className="h-full flex flex-col overflow-hidden">
              <div className="grid grid-cols-8 ">
                  <div className="sticky top-0  col-span-1 bg-cyan-900 col-span-1">
                      <Sidebar/>
                  </div>
                  <div className="col-span-7 bg-cyan-800 p-2">
                      <div className="overflow-y-auto h-[45rem] bg-gradient-to-r from-cyan-900 via-zinc-800 to-cyan-900">
                          {children}
                      </div>
                  </div>
              </div>
          </div>

      </div>
  )
}
