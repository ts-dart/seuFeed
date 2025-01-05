#classe para representar uma noticia
class News:

  def __init__(self):
    self._news_href_link = None
    self._news_img_src = None
    self._news_category = None
    self._news_title = None

  def __dict__(self):
    return {
      'newsHrefLink': self._news_href_link,
      'newsImgSrc': self._news_img_src,
      'newsCategory': self._news_category,
      'newsTitle': self._news_title
    }

  @property
  def news_href_link(self):
    return self._news_href_link
  
  @news_href_link.setter
  def news_href_link(self, news_href_link):
    if not news_href_link: 
      self._news_href_link = None
    else:
      self._news_href_link = news_href_link

  @property
  def news_img_src(self):
    return self._news_img_src
  
  @news_img_src.setter
  def news_img_src(self, news_img_src):
    if not news_img_src:
      self._news_img_src = None
    else:
      self._news_img_src = news_img_src

  @property
  def news_category(self):
    return self._news_category
  
  @news_category.setter
  def news_category(self, news_category):
    if not news_category:
      self._news_category = None
    else:
      self._news_category = news_category

  @property
  def news_title(self):
    return self._news_title
  
  @news_title.setter
  def news_title(self, news_title):
    if not news_title:
      self._news_title = None
    else:
      self._news_title = news_title