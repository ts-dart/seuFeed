package com.api.controllers.interfaces;

import org.springframework.http.ResponseEntity;
import com.api.services.interfaces.NewsInterface;

public interface NewsControllerInterface {
  ResponseEntity<Iterable<NewsInterface>> GetNews();
}
