# HTTP -> Hypertext Transfer Protocol. protocol to transmitting hypermedia documents such as HTML
# what's hypertext?
# HTTP method -> GET, POST, PATCH, PUT, DELETE, etc
# HTTP request header
# HTTP response usually contain an HTTP code status, HTTP response header, optional HTTP body
# HTTP code status
# 1xx informational
# 2xx Success
# 3xx Redirection
# 4xx Client Error
# 5xx Server Error
# XHR?
# HTTP is a stateless protocol meaning they don't hold data between two request
# tho cookies kinda complicate this a bit
# 4 kind of HTTP headers context-wise:
# General Header->applied on both request response but not affecting the database body
# Request Header->information about the fetched request by the clients
# Response Header->location of the source that has been requested by the clients
# Entity Header->contains information about the body of the resource like mime type
# From what I see there's hundreds HTTP header
# TIL that cookies and cache are different thing
# from what I understand cookies are server side so it often manage thing like login session (?) and often
# store sensitive information
# while cache more like client side so it's stored in browser so you can open link that you've visited faster next
# time
# cookies are identity thing while cache are performance thing i guess
# so cookies also have some type: session cookies, preference cookies (like theme=dark),
# analytics cookies, and security cookies.

import requests

url = "https://pokeapi.co/api/v2/pokemon/pikachu"
# url = "https://w3schools.com/python/demopage.htm"

try:
    # response = requests.get(url, timeout=5) # wait max 5 seconds
    # response = requests.head(url, timeout=5) # wait max 5 seconds
    response = requests.options(url, timeout=5) # wait max 5 seconds


# url = "https://randomuser.me/api/"
# data = requests.get(url).json

# except Exception as e:
    # print(f"{type(e)} : {e}")
except requests.exceptions.ReadTimeout:
    print("The URL doesn't response")
else:
    print(type(response))
    print(response.status_code)
    print(response.headers)
    print(response.headers["access-control-allow-methods"])
    print(response.headers.get("allow"))
    # print(response.headers["content-type"])
    # print(response.headers["content-length"])
    # print(response.json())
    # print(response.text)
    # print(response.json)
    