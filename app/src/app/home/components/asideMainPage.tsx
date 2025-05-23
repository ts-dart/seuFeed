import ClimateDate from "./climateDate";

export default function AsideMainPage() {
  return(
    <aside className="w-[40%] ml-0 mt-[60px] px-4 py-4">
      <div id="clima">
        <div className="px-1 py-1 border-b-[1px] border-[#3C3C3C] mb-[10px]">
          <h1 className="text-black text-left text-base">Tempo</h1>
        </div>
        <ClimateDate/>
      </div>
    </aside>
  );
}