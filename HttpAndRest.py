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


# API -> a set of rules or protocols that enabled software applications to communicate with each other
# i.e. exchange data, features, and functionality
# type of API
# Web API -> (majority of APIs)
# data API/database API, operating system API, remote API used to define how applications on different device interact
# type of Web API
# Open API or public API -> anyone can access
# partner API -> for partner business
# internal API -> for specific organization
# composite API -> combine multiple data or service, allow programmer to access several endpoints in single call
# useful in microservice.
# SaaS -> software as a service 

# API Protocol
# SOAP (Simple Object Access Protocol) -> a lightweight XML-based messaging protocol specification
# RPC (Remote Procedure Call) -> protocol that provide high level communications paradigm used in OS
# XML-RPC, JSON-RPC, gRPC, GraphQL
# Websocket -> enable bidirectional communication between client and server. Once the connection established,
# it allows for continous exchange make it great for real time communication



# SOAP is more strict and secure and built into many development tools than REST.
# While REST support more than XML such as plaintext, HTML, JSON, etc.

# REST API is the base layer of web communication, so learn it and master it. 
# After you master it, you can learn Websocket and GraphQL
# REST API -> Representational State Transfer Application Programming Interface
# tipically in JSON format

# 6 REST design principles (architectural constraints)
# 1. Unifrom Interface -> all API requests for the same resource should look the same regrardless of the client
# 2. Client-Server Decoupling -> client and server should be completely independent. The only information client
# so based on chatGPT, it's only mean that client doesn't care about how the server store or process data.
# should know is URI (Uniform Resource Identifier). and client can't modify the server other than passing it to
# the requested data via HTTP
# 3. Stateless -> not storing any data from a session basically, so it need to be done in one process (?)
# 4. Cacheability -> When possible, data should be cacheable to increase performance.
# 5. Layered System Architecture -> calls and responses go through different layers. Do not assume client and server
# connect directly.
# 6. Code on Demand (Optional) -> usually send static data. but in certain cases, responses can also contain 
# executable code, in this case the code should only run on demand

# REST API Best Practice
# 1. Using OpenAPI Specification (OAS)
# 2. Securing through use hashing algorithms for password and HTTPS for secure data transmission


# curl command (Client URL Request Library)
# https://curl.se/docs/tutorial.html
# curl usually used for quick testing or debugging API


# swagger -> API documentation framework
# for documenting REST API
# provides HTML view of the API documentation with JSON support and detailed infromation oh the HTTP methods

# Best practice to learn how to interact with API is always by look at the API documentation

# Postman -> testing and call API tools
# query parameters: optional and required
# query parameters can only be known by reading documentation since there's no rule to it
# query parameters is lowercase sensitive
# Path variables automatically replaced by Postman to it's value and so path variable name doesn't matter
# not like query params
# Most of post endpoint need some way of authentication (?)
# important things in postman: console, random variable, send shortcut(ctrl+enter), 
# postman isn't a tool for performance testing, security testing, 
# when I do GET method, and send a data, I get an error, so they don't just ignore the data that been sent?
