from fastapi import FastAPI
from typing import Optional
from pydantic import BaseModel
import uvicorn
# the app name is important.
# uvicorn main:app --reload
# main means file name, app means the FastAPI instance, and reload means it's reload automatically if the file changes
app = FastAPI()


@app.get('/')
# function name doesn't matter?
def index():
    return {"data":{"name":"John"}}

@app.get("/about")
def about():
    return {"data":"about page"}

@app.get("/blog")
def blog(limit = 10, published : bool = True, sort: Optional[str] = None):
    if published:
        return {"data": f"list of {limit} published blogs"}
    else:
        return {"data": f"list of {limit} unpublished blogs"}

# fastapi read scripts line by line
# so the best practice to write specific endpoint before dynamic one (?)
@app.get("/blog/unpublished")
def unpublished():
    return {"data": "all unpublished blogs"}

@app.get("/blog/{id}")
def blog(id:int):
    return {"data": id}




@app.get("/blog/{id}/comments")
# if there's the curly bracket in app.get, fastapi will treat it as path parameter, other wise query parameter
# remember query parameter always put in the end of the link (?)
def comments(id:int, limit=10):
    return {"data":{"1","2", limit}}

# API doc
# FastAPI automatically create swagger UI and Redoc for you
# Just typing http://127.0.0.1:8000/docs or http://127.0.0.1:8000/redoc

class Blog(BaseModel):
    title : str
    body : str
    published : Optional[bool] = False

@app.post("/blog")
def create_a_blog(blog: Blog):
    return {"data": f"Blog with title {blog.title} is created"}

# if __name__ == "__main__":
    
#     # 127.0.0.1:8000 is a default btw (?)
#     uvicorn.run(app, host="127.0.0.1", port=8000)