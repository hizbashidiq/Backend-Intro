# import numpy as np

# Python data types: dicts, list, tuples, set
# dicts {"key":item, "key2":item2, ...}
# list [item1, item2, ...]
# tuples (item1, item2, ...)
# set {item1, item2, ...}

def FtoC(temp):
    return (temp-32)*5/9
temp1 = 77
temp2 = 95
temp3 = 50

# for i in range(3):
#     # print(KtoC(f"temp{i+1}"))
#     # print(f"temp{i+1}")
#     pass

i = 5
print(f"hi{i}")

print(globals()['temp1'])

print(FtoC(temp1))
print(FtoC(temp2))
print(FtoC(temp3))

# x = dir(np)
# print(x)

import platform

x = dir(platform)
print(x)

import numpy as np
x = dir(np)
print(x)

# shortcut
# ctrl+backtick : switching between code and terminal
# ctrl+c : copy a line
# ctrl+shift+k : delete a line
# ctrl+d : select a word
# ctrl+home : to the first line
# ctrl+end : to the last line
# ctrl+] : to indent a line


# File handling
# open()-> [r, a, w, x] [t, b]
# [read, append, write, create] [text, binary]
# r+, a+, w+ -> read + write, append + read, write + read

f = open("README.md")

# the default is rt so the above are exactly like
# f = open("README.md", "rt")

print(f.read())
f.close()

with open("README.md") as f:
    print(f.read())

with open("text.txt") as f:
    # print first x character
    print(f.read(3))
    print(f.readline())
    print(f.readline())

with open("text.txt") as f:
    for i in f:
        print(i)

with open("text.txt", "r") as f:
    print(f.readlines())

with open("text.txt", "a+") as f:
    # f.write(f"{len(f.readlines())+1} Hewwo")
    print("s")
    print(f.readline())

with open("text.txt", "a+") as f:
    print(len(f.readlines()))
    f.seek(0)
    print(len(f.readlines()))
    print(len(f.readlines()))
    f.seek(0)
    f.write(f"\n{len(f.readlines())+1} Hewwo")

# for big file it's risky to use readlines to count line, instead you can do it like this
count = 0
with open("text.txt", "r") as f:
    for _ in f:
        count += 1

print(count)

# count = 0
with open("text.txt", "r") as f:
    n_line = sum(1 for _ in f)

print(n_line)

# check a file before deleting a file
import os
if os.path.exists("abcd.txt"):
    os.remove("abcd.txt")
else:
    print("The file does not exist!")


# try -> check error
# except -> handle error
# else -> execute code if there's no error, usually normal logic outside of possible error command
# finally -> execute code anyway regardless, usually to close something i.e. connection, file
# raise -> intentionally trigger an exception, usually using if logic if ... then raise

# os.remove("abcd.txt")
# with open("abcd.txt") as f:
#     print("hi")

try:
    f = open("abcd.txt")
except:
    print("Idk, something error")

try:
    f = open("abcd.txt")
except FileNotFoundError:
    print("File Not Found!")
except:
    print("HEHE BOI")
else:
    print(f.readlines())
finally:
    # try:
    #     f.close()
    f.close()

try:
    f = open("abcd.txt")
except Exception as e:
    print(f"Error Type: {type(e).__name__}")
    print(f"Message: {e}")

# Never use except Exception in real code. Only use it to discover error type

try:
    f = open("nonexistent.txt", "r")  # file doesn’t exist
    # do something
except FileNotFoundError:
    print("File Not Found!")
finally:
    f.close()
    print("Hewwo")

# try:
#     f = open("nonexistent.txt", "r")  # file doesn’t exist
#     # do something
# finally:
#     # but this raise an error hmmm
#     # try:
#     #     f.close()
#     # except Exception as e:
#     #     print(type(e), e)
#     pass

# OOP
# DRY -> Don't Repeat Yourself
# Class and object
# you can named self other than "self", but first parameter is always a self
class MyClass:
    x = 5 # Class property
    def __init__(self, name, age=20):
        self.name = name # Instance property
        self.age = age

    def greet(self):
        print(f"Hi, my name is {self.name} and I'm {self.age} years old!")
    
    def happy_birthday(self):
        self.age += 1
        print(f"Happy birthday {self.name}!")
        print(f"You're now {self.age} years old.")
    
    # __str__ method is a special method that controls what's returned if you use print(object)
    def __str__(self):
        return f"{self.name} - {self.age}"



z = MyClass("Eric", 7)
print(z.x)
print(z.name)
print(z.age)
z.greet()
z.age = 15
z.greet()
z.city = "Bandung"
print(z.city)
z.happy_birthday()
print(z)
# del z to delete object z

# To make a child class use a parent class as parameter
class Student(MyClass):
    # When you add __init__ in child class, it no longer inherit the parent's __init__
    def __init__(self, name, age=12, year=2019):

        # To keep the parent's __init__ after add __init__ in child class, call it
        # MyClass.__init__(self, name, age)
        # super function will make child class to inherit all parent's methods and properties
        super().__init__(name, age)
        self.year = year

    # If you add same name method in child class, it'll override the parents method
    def greet(self):
        print(f"Hi, my name is {self.name} and I'm graduated in {self.year}!")

c = Student("Mike", 19)
c.greet()

# len() function is an example of polymorphism, 
# same name function that can be used on a lot of different type of object
# the greet method above also a type of polymorphism

# Encapsulation is about protecting a data inside a class
# in python use __ double underscore as prefix to make a data private. Example: self.__age
# private data can only accessed and modified using getter or setter
# single underscore _ mean it's meant to be protected like Java for example. but in python you still can acccess it
# python only have private and public I suppose
# lol you can bruteforce to use private methods or property in python
# instead z.age, use z._MyClass__age


import json
jsonstring = '{"name": "Eric", "age": 18, "year" : 2015}'

python_dict = json.loads(jsonstring)
print(python_dict["name"])

python_dict = {
    "name" : "Bill",
    "age" : 15,
    "year" : 2020
}
jsonstring = json.dumps(python_dict)
print(jsonstring)

print(json.dumps(["banana", "apple"]))

# use parameter indent to make it more readable
print(json.dumps(python_dict, indent=4))
# you can also use sort_keys=True to sort it based on key

# pip freeze > requirements.txt
# to create requirements.txt easily
# to setup in new environment, you can just 
# pip install -r requirements.txt