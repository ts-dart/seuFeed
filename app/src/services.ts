import { PostsList } from "./interfaces"

export async function getAllPosts(): Promise<Array<PostsList>> {
  const response = await fetch("http://localhost:8080/getAllPosts", {
    method: "GET",
  })

  if (!response.ok) {
    throw new Error("Erro ao receber posts da api")
  }

  return await response.json()
}

export async function getPostsBySection(): Promise<Array<PostsList>> {
  const response = await fetch("http://localhost:8080/getPostsBySection", {
    method: "GET",
  })

  if (!response.ok) {
    throw new Error("Erro ao receber posts da api")
  }

  return await response.json()
}