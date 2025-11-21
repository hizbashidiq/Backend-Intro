from fastapi import APIRouter
from .. import schemas, database, models
from sqlalchemy.orm import Session
from fastapi import APIRouter, Depends, status, HTTPException
from ..hashing import Hash

router = APIRouter(tags=['Users'], prefix="/user")

get_db = database.get_db

@router.post('/', response_model=schemas.ShowUser)
def create_user(request: schemas.User, db: Session = Depends(get_db)):
    # hashedPassword = pwd_cxt.hash(request.password)
    # new_user = models.User(request)
    new_user = models.User(name=request.name, email=request.email, password=Hash.encrypt(request.password))
    db.add(new_user)
    db.commit()
    db.refresh(new_user)
    return new_user

# pydantic is called schemas while sqlalchemy is called models

@router.get('/{id}', response_model=schemas.ShowUser)
def get_user(id:int, db: Session = Depends(get_db)):
    # I don't get why he keep using first, I mean id is primary key so it shouldn't be an issue tbh
    user = db.query(models.User).filter(models.User.id == id).first()
    if not user:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND,
                            detail= f"User with id {id} is not found")
    return user

# tags is for api documentation (?)