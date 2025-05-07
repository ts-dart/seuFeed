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
  const [emphasis, setEmphasis] = useState<number | null>(0)

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
          setEmphasis(Math.floor(Math.random() * data.length))
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
    <section className="w-[55%] ml-0 mt-[60px] px-4 py-4">
      <div className="px-1 py-1 border-b-[1px] border-[#3C3C3C] mb-[10px]">
        <h1 className="text-black text-left text-base">Destaques</h1>
      </div>
      <div>
      {typeof emphasis === 'number' && postsList && postsList[emphasis] && (
        <div className="relative w-[100%] h-[300px] rounded-[6px] overflow-hidden rounded-[10px] px-4 py-4">
          {/* Imagem como plano de fundo */}
          <Image
            src={postsList[emphasis].post_img_src}
            alt="Imagem de destaque da noticia"
            fill
            className="object-cover z-0"
          />
        
          {/* Conteúdo sobreposto */}
          <div className="absolute bottom-0 left-0 w-full bg-gradient-to-t from-black/80 to-transparent px-4 py-4 z-10">
            <h1 className="text-white text-2xl font-bold mb-2">
              {postsList[emphasis].post_text}
            </h1>
            <div className="flex items-center gap-2">
              <Image
                src={postsList[emphasis].font_img_src}
                width={20}
                height={20}
                alt="Logo do pagina fonte"
                className="rounded-[2px]"
              />
              <p className="text-white text-sm">{postsList[emphasis].font}</p>
            </div>
          </div>
        </div>
      )}
      </div>
      {!loading 
        ? postsList?.map((post: Post) => (
            <a href={post.post_href_link} target="_blank" key={post.post_href_link}>
              <div className="flex bg-[#F9F9F9] rounded-[2px] px-1 py-1 border-[1px] border-[#F9F9F9] mb-[5px] w-full min-h-[150px]">
                {/* Imagem de destaque à esquerda */}
                <div className="w-[20%] px-1 py-1">
                  <Image
                    src={post.post_img_src}
                    width={100}
                    height={100}
                    alt="Imagem de destaque da noticia"
                    className="rounded-[10px] object-cover w-full h-full"
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
                    <p className="text-xs text-black">{post.font}</p>
                  </div>

                  {/* Título da notícia */}
                  <h1 className="text-left text-base font-bold text-black leading-tight mb-2">
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