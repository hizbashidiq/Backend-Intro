from pydantic import BaseModel
from typing import List
class Blog(BaseModel):
    title: str
    body: str

# class Blog(BlogBase):
    

class User(BaseModel):
    name:str
    email:str
    password:str

class ShowUser(BaseModel):
    name:str
    email:str
    blogs: List[Blog] = []


# class ShowBlog(Blog):
class ShowBlog(BaseModel):
    title: str
    body: str
    creator: ShowUser
    class Config():
        # orm_mode = True
        from_attributes = True

class Login(BaseModel):
    username: str
    password: str

class Token(BaseModel):
    access_token: str
    token_type: str


class TokenData(BaseModel):
    # username: str | None = None
    email: str | None = None
