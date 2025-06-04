"use client"

import { useState, useEffect } from "react";
import { CurrentClimateData } from "@/interfaces";
import { getClimateDate } from "@/services";
import { Sun, Moon} from "lucide-react"

export default function ClimateDate() {
  const [currentClimateData, setCurrentClimateData] = useState<CurrentClimateData | null>();
  const [hour, setHour] = useState<string>("");

  const getTime = () => {
    const now = new Date();
    const hour = now.getHours();
    const minutes = new Date().getMinutes();
    setHour(hour+":"+minutes);
  }

  const getTimeIcon = () => {
    const isDay = Number(hour.split(":")[0]) >= 6 && Number(hour.split(":")[0]) < 18;
  
    return isDay ? (
      <Sun className="text-yellow-400 w-20 h-20" />
    ) : (
      <Moon className="text-blue-500 w-20 h-20" />
    );
  }

  useEffect(() => {
    getTime();
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          getClimateDate(String(position.coords.latitude), String(position.coords.longitude))
            .then((data) => {
              if (data != null) {
                setCurrentClimateData(data);
              }
            })
        },
        (error) => {
          console.error("Erro ao obter localização:", error);
        }
      );
    } else {
      console.error("Geolocalização não é suportada pelo navegador.");
    }
  }, [])

  return (
    <section className="">
      {currentClimateData 
        ? <div className="flex flex-col items-end space-y-2">
            <div className="flex flex-row">
              {getTimeIcon()}
              <h1 className="text-black ml-2 text-[50px]">{Number.parseInt(currentClimateData.temperature_2m)}°C</h1>
            </div>
            <p className="text-black">{hour}</p>
            <p className="text-black">Vento: {currentClimateData.wind_speed_10m}km/h</p>
        </div> 
        : <div>
            <h1>Nao foi possivel obter informacoes climaticas</h1>
        </div>}
    </section>
  )
}