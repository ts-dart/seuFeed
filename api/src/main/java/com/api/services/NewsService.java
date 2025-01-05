package com.api.services;

import org.springframework.stereotype.Service;
import com.api.services.interfaces.NewsInterface;
import com.api.services.interfaces.NewsServiceInterface;

//DESENVOLVER METODO PARA PEGAR AS NOTICIAS

@Service
public class NewsService implements NewsServiceInterface {
  
  public Iterable<NewsInterface> GetNews() {

  }
}
