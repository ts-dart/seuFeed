"use client"

import { useState, useEffect } from "react";
import { CurrentClimateData } from "@/interfaces";
import { getClimateDate } from "@/services";
import { Sun, Moon} from "lucide-react"

export default function ClimateDate() {
  const [currentClimateData, setCurrentClimateData] = useState<CurrentClimateData | null>()

  const getTime = (dateString: string) => {
    const date = new Date(dateString);
    return date.getHours();
  }

  const getTimeIcon = (dateString: string) => {
    const hour = getTime(dateString);
  
    // Consideramos dia das 6h às 18h
    const isDay = hour >= 6 && hour < 18;
  
    return isDay ? (
      <Sun className="text-yellow-400 w-6 h-6" />
    ) : (
      <Moon className="text-blue-500 w-6 h-6" />
    );
  }

  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          getClimateDate(String(position.coords.latitude), String(position.coords.longitude))
            .then((data) => {
              if (data != null) {
                setCurrentClimateData(data)
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
        ? <div>
            {getTimeIcon(currentClimateData.time)}
            <h1>{currentClimateData.temperature2m}</h1>
            <p>{currentClimateData.windSpeed10m}</p>
            <p>{getTime(currentClimateData.time)}</p>
        </div> 
        : <div>
            <h1>Nao foi possivel obter informacoes climaticas</h1>
        </div>}
    </section>
  )
}