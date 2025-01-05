package com.api.controllers;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.api.controllers.interfaces.NewsControllerInterface;
import com.api.services.NewsService;
import com.api.services.interfaces.NewsInterface;


@RestController
@RequestMapping("/news")
public class NewsController implements NewsControllerInterface {

  private NewsService newsService;
  
  @Autowired
  public NewsController(NewsService newsService) {
    this.newsService = newsService;
  }

  @GetMapping
  public ResponseEntity<Iterable<NewsInterface>> GetNews() {
    Iterable<NewsInterface> news = newsService.GetNews();
    return ResponseEntity.ok(news);
  }
  
}