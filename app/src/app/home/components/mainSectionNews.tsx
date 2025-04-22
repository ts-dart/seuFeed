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
    <section>
      {postsList?.map((post:Post) => (
          <div>
            <Image src={post.post_img_src} width={150} height={150} alt="Imagem de destaque da noticia"/>
            <h1>{post.post_text}</h1>
            <Image src={post.font_img_src} width={30} height={30} alt="Logo do pagina fonte"/>
            <p>{post.font}</p>
          </div>
      ))}
    </section>
  )
}