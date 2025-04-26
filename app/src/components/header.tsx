"use client"

export default function Header() {
  return (
    <header className="flex flex-row fixed top-0 left-0 w-full z-50 h-[50px] shadow-md bg-[#181A1B] border-b-[1px] border-[#101112]">
      <div className="bg-[#db290b] h-full w-[6%] flex items-center justify-center ml-40">
        <h1 className="text-base text-white font-light">SeuFeed</h1>
      </div>
      <nav className="flex items-center gap-[3px] ml-8">
        <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Página inicial</button>
        <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Esportes</button>
        <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Ciência e tecnologia</button>
        <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Negócios</button>
        <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Saúde</button>
        <button className="text-xs text-[#bfbdbd] px-2 py-1 cursor-pointer hover:underline">Entretenimento</button>
      </nav>
    </header>
  )
}