import json

from time import sleep
from urlo import Urlo

news_list = []
index = 0

while(True):
  index += 1
  Urlo(news_list)

  with open('dados.json', 'w') as file:
    for data in news_list:
      json.dump(data.__dict__(), file)

  print(index, len(news_list), news_list)
  news_list.clear()
  sleep(300)

  