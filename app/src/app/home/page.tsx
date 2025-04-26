import Header from "@/components/header";
import MainSectionNews from "@/app/home/components/mainSectionNews";
import ClimateDate from "./components/climateDate";
import AsideMainPage from "./components/asideMainPage";

export default function HomePage() {
  return (
    <>
      <Header/>
        <section className="flex flex-row">
            <MainSectionNews/>
            <AsideMainPage/>
        </section>
    </>
  )
}