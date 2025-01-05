package com.api.services.interfaces;

public interface NewsInterface {
  String getNewsHrefLink();
  String getNewsImgSrc();
  String getNewsCategory();
  String getNewsTitle();

  void setNewsHrefLink(String newsHrefLink);
  void setNewsImgSrc(String newsImgSrc);
  void setNewsCategory(String newsCategory);
  void setNewsTitle(String newsTitle);
}
