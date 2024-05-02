# import urllib.request    

# url = "https://www.fragrantica.com/perfume/Tom-Ford/Tobacco-Vanille-1825.html"
# headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.4 Safari/605.1.15',
#        'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
#        'Accept-Charset': 'ISO-8859-1,utf-8;q=0.7,*;q=0.3',
#        'Accept-Encoding': 'none',
#        'Accept-Language': 'en-US,en;q=0.8',
#        'Connection': 'keep-alive'} 

# req=urllib.request.Request(url=url, headers=headers) #The assembled request
# http_handler = urllib.request.HTTPHandler(debuglevel=1)
# opener = urllib.request.build_opener(http_handler)
# response = opener.open(req)
# # data = response.read() # The data u need
# # print(response.status, response.headers)

import urllib.request

url = "https://www.fragrantica.com/perfume/Tom-Ford/Tobacco-Vanille-1825.html"
headers = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.4 Safari/605.1.15',
    'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
    'Accept-Charset': 'ISO-8859-1,utf-8;q=0.7,*;q=0.3',
    'Accept-Encoding': 'none',
    'Accept-Language': 'en-US,en;q=0.8',
    'Connection': 'keep-alive'
}

req = urllib.request.Request(url=url, headers=headers)

# HTTPHandler 인스턴스 생성 및 debuglevel 설정
http_handler = urllib.request.HTTPHandler(debuglevel=1)

# HTTPHandler를 사용하는 opener 생성
opener = urllib.request.build_opener(http_handler)

# URL 열기
response = opener.open(req)