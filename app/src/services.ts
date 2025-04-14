import { PostsList, CurrentClimateData } from "./interfaces"

export async function getAllPosts(): Promise<Array<PostsList>> {
  const response = await fetch("http://localhost:8080/AllPosts", {
    method: "GET",
  })

  if (!response.ok) {
    throw new Error("Erro ao receber posts da api")
  }

  return await response.json()
}

export async function getPostsBySection(): Promise<Array<PostsList>> {
  const response = await fetch("http://localhost:8080/PostsBySection", {
    method: "GET",
  })

  if (!response.ok) {
    throw new Error("Erro ao receber posts da api")
  }

  return await response.json()
}

export async function getClimateDate(): Promise<CurrentClimateData> {
  const response = await fetch("http://localhost:8080/ClimateData", {
    method: "GET",
  })

  if (!response.ok) {
    throw new Error("Erro ao receber posts da api")
  }

  return await response.json()
}