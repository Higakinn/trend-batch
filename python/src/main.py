import os
import requests
from datetime import datetime

URL = os.getenv("DISCORD_WEBHOOK")

def main():
  print("python batch start!!")
  
  today = datetime.today()
  ym = f"{today.year}{str(today.month).rjust(2, '0')}"
  result = requests.get(f"https://connpass.com/api/v1/event/?keyword=python&order=2&count=1&ym={ym}").json()

  headers = {'Content-Type': 'application/json'}
  event_url = ""
  for event in result["events"]:
      print(f'{event["started_at"]}:{event["title"]}')
      print(event["event_url"])
      event_url=event["event_url"]
  result = requests.post(URL,data={'content': f"{today}\n{event_url}"})

  print(result)
  print("python batch end!!")

if __name__ == '__main__':
  main()
