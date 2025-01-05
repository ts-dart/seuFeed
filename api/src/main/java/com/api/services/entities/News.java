package com.api.services.entities;

import com.api.services.interfaces.NewsInterface;

public class News implements NewsInterface {

  private String newsHrefLink;
  private String newsImgSrc;
  private String newsCategory;
  private String newsTitle;

  //getters
  public String getNewsHrefLink() {
    return this.newsHrefLink;
  }

  public String getNewsImgSrc() {
    return this.newsImgSrc;
  }

  public String getNewsCategory() {
    return this.newsCategory;
  }

  public String getNewsTitle() {
    return this.newsTitle;
  }

  //setters
  public void setNewsHrefLink(String newsHrefLink) {
    this.newsHrefLink = newsHrefLink;
  }

  public void setNewsImgSrc(String newsImgSrc) {
    this.newsImgSrc = newsImgSrc;
  }

  public void setNewsCategory(String newsCategory) {
    this.newsCategory = newsCategory;
  }

  public void setNewsTitle(String newsTitle) {
    this.newsTitle = newsTitle;
  }

}