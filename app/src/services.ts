import { Post, CurrentClimateData } from "./interfaces";

export async function getAllPosts(): Promise<Array<Post>> {
  const response = await fetch("http://localhost:8080/AllPosts", {
    method: "GET",
  })

  if (!response.ok) {
    throw new Error("Erro ao receber posts da api");
  }

  return await response.json();
}

export async function getPostsBySection(ft:string): Promise<Array<Post>> {
  const response = await fetch(`http://localhost:8080/PostsBySection?ft=${ft}`, {
    method: "GET",
  })

  if (!response.ok) {
    throw new Error("Erro ao receber posts da api");
  }

  return await response.json();
}

export async function getClimateDate(latitude: string | undefined, longitude: string | undefined): Promise<CurrentClimateData> {
  const response = await fetch(`http://localhost:8080/ClimateData?latitude=${latitude}&longitude=${longitude}`, {
    method: "GET",
  })

  if (!response.ok) {
    throw new Error("Erro ao receber posts da api");
  }

  return await response.json();
}