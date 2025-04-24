"use client"

import { useState, useEffect } from "react";
import { Post } from "@/interfaces";
import { getPostsBySection } from "@/services";
import Image from "next/image";
import Loading from "@/app/loading";
import Error from "@/app/error";

export default function MainSectionNews() {
  const [postsList, setPostsList] = useState<Array<Post> | null>();
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<boolean>(false);

  useEffect(() => {
    setLoading(true)
    setError(false)
  
    const timeout = setTimeout(() => {
      setError(true)
      setLoading(false)
    }, 5000) // tempo máximo para esperar os dados
  
    getPostsBySection("main")
      .then((data) => {
        if (data.length > 0) {
          clearTimeout(timeout) // se os dados chegaram a tempo, cancela o timeout
          setPostsList(data)
          setLoading(false)
        }
      })
      .catch(() => {
        clearTimeout(timeout)
        setError(true)
        setLoading(false)
      })
  
    // cleanup se o componente for desmontado
    return () => clearTimeout(timeout)
  }, [])

  return (
    <section className="w-[65%] ml-0 mt-[50px] px-4 py-4">
      <div className="px-1 py-1 border-b-[1px] border-[#3C3C3C] mb-[10px]">
        <h1 className="text-left text-base">Principais notícias</h1>
      </div>
      {!loading 
        ? postsList?.map((post: Post) => (
            <a href={post.post_href_link} target="_blank" key={post.post_href_link}>
              <div className="flex bg-[#181A1B] rounded-[2px] px-1 py-1 border-[1px] border-[#101112] mb-[5px] w-full min-h-[150px]">
                {/* Imagem de destaque à esquerda */}
                <div className="w-[20%] px-1 py-1">
                  <Image
                    src={post.post_img_src}
                    width={100}
                    height={100}
                    alt="Imagem de destaque da noticia"
                    className="rounded-[2px] object-cover w-full h-full"
                  />
                </div>

                {/* Conteúdo à direita */}
                <div className="flex flex-col justify-start px-2 py-1 w-[80%]">
                  {/* Fonte */}
                  <div className="flex items-center gap-2 mb-2">
                    <Image
                      src={post.font_img_src}
                      width={16}
                      height={16}
                      alt="Logo do pagina fonte"
                      className="rounded-[2px] object-contain"
                    />
                    <p className="text-xs text-white">{post.font}</p>
                  </div>

                  {/* Título da notícia */}
                  <h1 className="text-left text-base font-bold text-white leading-tight mb-2">
                    {post.post_text}
                  </h1>

                  {/* Tempo */}
                  <h2 className="text-[#808080] text-xs mt-[3px]">00 minutos atrás</h2>
                </div>
              </div>
            </a>
      )) : <Loading />}
      {error ? <Error/> : null}
    </section>
  )
}