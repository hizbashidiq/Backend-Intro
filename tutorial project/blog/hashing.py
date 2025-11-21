from passlib.context import CryptContext

pwd_cxt = CryptContext(schemes=["pbkdf2_sha256"],deprecated="auto")

class Hash():
    def encrypt(password: str):
        return pwd_cxt.hash(password)
    def verify(plainPassword, hashedPassword):
        return pwd_cxt.verify(plainPassword, hashedPassword)