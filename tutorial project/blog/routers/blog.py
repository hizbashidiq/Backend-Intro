from typing import List
from fastapi import APIRouter, Depends, status, HTTPException
from .. import schemas, database, models, oauth2
from sqlalchemy.orm import Session


router = APIRouter(tags=['Blogs'], prefix="/blog")

get_db = database.get_db

@router.get("/", response_model=List[schemas.ShowBlog])
def all(db: Session = Depends(get_db), current_user: schemas.User = Depends(oauth2.get_current_user)):
    blogs = db.query(models.Blog).all()
    return blogs

@router.post("/", status_code=status.HTTP_201_CREATED)
def create(request: schemas.Blog, db: Session = Depends(get_db), current_user: schemas.User = Depends(oauth2.get_current_user)):
    # new_blog = models.Blog(title=request.title, body=request.body)
    new_blog = models.Blog(title=request.title, body=request.body, user_id=1)

    db.add(new_blog)
    db.commit()
    db.refresh(new_blog)
    return new_blog

@router.get("/{id}", status_code=200, response_model=schemas.ShowBlog)
def show(id, db: Session = Depends(get_db), current_user: schemas.User = Depends(oauth2.get_current_user)):
    blog = db.query(models.Blog).filter(models.Blog.id == id).first()
    if not blog:
        # response.status_code = status.HTTP_404_NOT_FOUND
        # return {'detail': f'Blog with the id={id} is not available'}
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND,
                            detail=f'Blog with the id={id} is not available')
    return blog

# delete and update
@router.delete("/{id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_blog(id, db: Session = Depends(get_db), current_user: schemas.User = Depends(oauth2.get_current_user)):
    blog = db.query(models.Blog).filter(models.Blog.id == id)
    if not blog:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail=f"Blog with id {id} is not found")
    blog.delete(synchronize_session=False)
    db.commit()
    # return {f'Blog with id {id} and title {models.Blog[id].title} has been deleted'}
    return 'Done'

@router.put('/{id}', status_code=status.HTTP_202_ACCEPTED)
def update_blog(id, request: schemas.Blog, db: Session = Depends(get_db), current_user: schemas.User = Depends(oauth2.get_current_user)):
    
    # db.query(models.Blog).filter(models.Blog.id == id).\
    # update({'title': 'updated title'})
    # Because .update() in SQLAlchemy expects a dictionary of column:value pairs, not a Pydantic model (like request).
    # so you need to add .dict()
    blog = db.query(models.Blog).filter(models.Blog.id == id)
    if not blog.first():
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail=f"Blog with id {id} is not found")
    blog.update(request.dict())
    db.commit()

    return f'Blog {id} Updated Succesfully'
    # return request
