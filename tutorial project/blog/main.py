from fastapi import FastAPI, Depends, status, Response, HTTPException
from typing import List
from .schemas import Blog
# from sqlalchemy.orm import Session
from . import models, schemas, hashing
# import models
# import schemas
# import hashing
from .database import engine, get_db
# from passlib.context import CryptContext
from .routers import blog, user, authentication

# python -m uvicorn blog.main:app --reload
# run in tutorial project repo


app = FastAPI()

# whenever there's new Base in models, it'll create new database for it
models.Base.metadata.create_all(engine)

app.include_router(authentication.router)
app.include_router(blog.router)
app.include_router(user.router)


# def get_db():
#     db = SessionLocal()
#     try:
#         yield db
#     finally:
#         db.close()



# @app.post("/blog", status_code=201)
# @app.post("/blog", status_code=status.HTTP_201_CREATED, tags=['blogs'])
# def create(request: Blog, db: Session = Depends(get_db)):
#     # new_blog = models.Blog(title=request.title, body=request.body)
#     new_blog = models.Blog(title=request.title, body=request.body, user_id=1)

#     db.add(new_blog)
#     db.commit()
#     db.refresh(new_blog)
#     return new_blog

# @app.get("/blog", response_model=List[schemas.ShowBlog] ,tags=['blogs'])
# def all(db: Session = Depends(get_db)):
#     blogs = db.query(models.Blog).all()
#     return blogs


# @app.get("/blog/{id}", status_code=200, response_model=schemas.ShowBlog ,tags=['blogs'])
# def show(id, response: Response, db: Session = Depends(get_db)):
#     blog = db.query(models.Blog).filter(models.Blog.id == id).first()
#     if not blog:
#         # response.status_code = status.HTTP_404_NOT_FOUND
#         # return {'detail': f'Blog with the id={id} is not available'}
#         raise HTTPException(status_code=status.HTTP_404_NOT_FOUND,
#                             detail=f'Blog with the id={id} is not available')
#     return blog

# # delete and update
# @app.delete("/blog/{id}", status_code=status.HTTP_204_NO_CONTENT ,tags=['blogs'])
# def delete_blog(id, db: Session = Depends(get_db)):
#     blog = db.query(models.Blog).filter(models.Blog.id == id)
#     if not blog:
#         raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail=f"Blog with id {id} is not found")
#     blog.delete(synchronize_session=False)
#     db.commit()
#     # return {f'Blog with id {id} and title {models.Blog[id].title} has been deleted'}
#     return 'Done'

# @app.put('/blog/{id}', status_code=status.HTTP_202_ACCEPTED ,tags=['blogs'])
# def update_blog(id, request: Blog, db: Session = Depends(get_db)):
    
#     # db.query(models.Blog).filter(models.Blog.id == id).\
#     # update({'title': 'updated title'})
#     # Because .update() in SQLAlchemy expects a dictionary of column:value pairs, not a Pydantic model (like request).
#     # so you need to add .dict()
#     blog = db.query(models.Blog).filter(models.Blog.id == id)
#     if not blog.first():
#         raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail=f"Blog with id {id} is not found")
#     blog.update(request.dict())
#     db.commit()

#     return f'Blog {id} Updated Succesfully'
#     # return request

# if __name__ == "__main__":
#     pass



# @app.post('/user', response_model=schemas.ShowUser, tags=['users'])
# def create_user(request: schemas.User, db: Session = Depends(get_db)):
#     # hashedPassword = pwd_cxt.hash(request.password)
#     # new_user = models.User(request)
#     new_user = models.User(name=request.name, email=request.email, password=hashing.Hash.encrypt(request.password))
#     db.add(new_user)
#     db.commit()
#     db.refresh(new_user)
#     return new_user

# # pydantic is called schemas while sqlalchemy is called models

# @app.get('/user/{id}', response_model=schemas.ShowUser, tags=['users'])
# def get_user(id:int, db: Session = Depends(get_db)):
#     # I don't get why he keep using first, I mean id is primary key so it shouldn't be an issue tbh
#     user = db.query(models.User).filter(models.User.id == id).first()
#     if not user:
#         raise HTTPException(status_code=status.HTTP_404_NOT_FOUND,
#                             detail= f"User with id {id} is not found")
#     return user

# # tags is for api documentation (?)
