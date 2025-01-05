import os
import requests

from news import News
from parsel import Selector
from time import sleep
from dotenv import load_dotenv 

'''
- criar uma classe para representar urlo(), organizar melhor o codigo e refatorar a classe noticia
para que possa ser generica.
'''


class Urlo:

  def __init__(self, news_list):
    load_dotenv()
    self.news_list = news_list
    self.scraper()

  @property
  def get_news_list(self):
    return self.news_list

  #faz a requisicao
  def makes_requests(self, url):
    try:
      response = requests.get(url, timeout=1)
      sleep(10)

      return response.text if response.status_code == 200 else None
    except requests.exceptions.ReadTimeout:
      return None
    
  #faz a filtragem das informacoes necessarias e cria uma lista de objetos (cada objeto representa uma noticia)
  def scraper(self):
    response = self.makes_requests(os.getenv("URLO_ENV"))
    data = response if isinstance(response, str) else ''
    selector = Selector(text=data)

    news_href_link = selector.css('div.post a.post-img::attr(href)').getall()
    news_img_src = selector.css('div.post a.post-img img::attr(src)').getall()
    news_category = selector.css('div.post div.post-body div.post-meta a.post-category::text').getall()
    news_title = selector.css('div.post div.post-body h3.post-title a::text').getall()

    for i in range(len(selector.css('div.col-xs-12 div.col-md-3 div.post').getall())):
      news_index = News()
      news_index.news_href_link = news_href_link[i]
      news_index.news_img_src = news_img_src[i]
      news_index.news_category = news_category[i]
      news_index.news_title = news_title[i]

      self.news_list.append(news_index)


#implementacao posterior.
#def urlt():
#  data = makes_requests(os.getenv('URLT'))