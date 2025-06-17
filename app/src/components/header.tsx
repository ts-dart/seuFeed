"use client"

import Link from "next/link"

export default function Header() {
  return (
    <header className="flex flex-row fixed top-0 left-0 w-full z-50 h-[50px] shadow-md bg-[#181A1B] border-b-[1px] border-[#101112]">
      <div className="bg-[#db290b] h-full w-[6%] flex items-center justify-center ml-40">
        <h1 className="text-base text-white font-light">SeuFeed</h1>
      </div>
      <nav className="flex items-center gap-[3px] ml-8">
        <Link href={"/home"}>
          <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Página inicial</button>
        </Link>
        <Link href={"/sports"}>
          <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Esportes</button>
        </Link>
        <Link href={"/scienceATech"}>
          <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Ciência e tecnologia</button>
        </Link>
        <Link href={"/business"}>
          <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Negócios</button>
        </Link>
        <Link href={"/health"}>
          <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Saúde</button>
        </Link>
        <Link href={"/entertainment"}>
          <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Entretenimento</button>
        </Link>
      </nav>
    </header>
  )
}