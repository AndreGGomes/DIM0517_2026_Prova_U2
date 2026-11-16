from fastapi import FastAPI

app = FastAPI()


@app.get("/hello")
def hello():
    return {"message": "Hello, World!"}


@app.get("/healthy")
def healthy():
    return {"status": "healthy"}

@app.get("/hi")
def hello():
    return {"message": "hi!"}

@app.get("/hi-errado")
def hello():
    return {"message": "hi-erradooooo!"}