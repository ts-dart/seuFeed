'use client'

import { useState, useEffect } from "react";
import { Post } from "@/interfaces";
import { getPostsBySection } from "@/services";
import Image from "next/image";

export default function MainSectionNews() {
  const [postsList, setPostsList] = useState<Array<Post> | null>();

  useEffect(() => {
    getPostsBySection("main")
        .then((data) => {
          console.log(data)
          setPostsList(data)
        })
  }, [])

  return (
    <section className="w-[65%] ml-0 mt-[70px] px-4 py-4">
      <div className="px-1 py-1 border-b-[1px] border-[#3C3C3C] mb-[10px]">
        <h1 className="text-left text-xl">Notícias principais.</h1>
      </div>
      {postsList?.map((post: Post) => (
        <a href={post.post_href_link} target="_blank" key={post.post_href_link}>
          <div className="flex bg-[#181A1B] rounded-[2px] px-3 py-3 border-[1px] border-[#101112] mb-[5px] w-full min-h-[150px]">
            <div id="img-news" className="w-[20%] flex flex-col justify-between px-1 py-1">
              <Image
                src={post.post_img_src}
                width={150}
                height={100}
                alt="Imagem de destaque da noticia"
                className="rounded-[2px] object-cover w-full h-[100px]"
              />
              <div className="flex justify-start items-center gap-1 mt-2">
                <Image
                  src={post.font_img_src}
                  width={16}
                  height={16}
                  alt="Logo do pagina fonte"
                  className="rounded-[2px] object-contain"
                />
                <p className="text-xs text-white">{post.font}</p>
              </div>
            </div>

            <div className="flex flex-col justify-between px-2 py-2 w-[80%]">
              <h1
                id="title"
                className="text-left text-xl font-bold text-white leading-tight mb-2"
              >
                {post.post_text}
              </h1>
              <h2
                id="time"
                className="text-[#808080] text-sm mt-[3px]"
              >
                00 minutos atrás
              </h2>
            </div>
          </div>
        </a>
      ))}
    </section>
  )
}