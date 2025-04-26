"use client"

import { useState, useEffect } from "react";
import { CurrentClimateData } from "@/interfaces";

export default function ClimateDate() {
  const [currentClimateData, setCurrentClimateData] = useState<CurrentClimateData | null>()

  useEffect(() => {

  }, [])

  return (
    <section className=""></section>
  )
}