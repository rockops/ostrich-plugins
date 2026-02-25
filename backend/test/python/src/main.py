from fastapi import FastAPI
import random

app = FastAPI()

@app.get("/random")
def get_random():
    """
    Returns a random integer between 1 and 100.
    """
    return {"random_integer": random.randint(1, 100)}

if __name__ == "__main__":
    import uvicorn
    import os
    port = int(os.environ.get("PORT", 8000))
    uvicorn.run(app, host="0.0.0.0", port=port)
