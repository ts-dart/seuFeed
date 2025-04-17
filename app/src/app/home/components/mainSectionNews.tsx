'use client'

import { useState, useEffect } from "react";
import { Post } from "@/interfaces";
import { getPostsBySection } from "@/services";
import Image from "next/image";

export default function MainSectionNews() {
  const [postsList, setPostsList] = useState<Array<Post> | null>();

  useEffect(() => {
    getPostsBySection("main")
        .then((data) => setPostsList(data))
  }, [])

  return (
    <section>
      {postsList?.map((post:Post) => (
          <div>
            <Image src={post.fontImgSrc} width={150} height={150} alt="Imagem de destaque da noticia"/>
            <h1>{post.postText}</h1>
            <p>{post.fontImgSrc}</p>
          </div>
      ))}
    </section>
  )
}