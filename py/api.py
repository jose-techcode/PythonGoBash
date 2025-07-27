from fastapi import FastAPI, HTTPException
import uvicorn
from pydantic import BaseModel

application = FastAPI()

class Item(BaseModel):
    id: int
    name: str

items = {}

# POST

@application.post("/create")
def post_item(item: Item):
    if item.id in items:
        raise HTTPException(status_code=400, detail="Item com esse ID já existe!")
    items[item.id] = item.name
    return {"Mensagem": "item criado!", "item": item}

# PUT

@application.put("/update")
def put_item(item: Item):
    if item.id not in items:
        raise HTTPException(status_code=404, detail="Item com esse ID não existe!")
    items[item.id] = item.name
    return {"Mensagem": "item atualizado!", "item": item}

# DELETE

@application.delete("/delete")
def delete_item(item: Item):
    if item.id not in items:
        raise HTTPException(status_code=404, detail="Item com esse ID não existe!")
    deleted = items.pop(item.id)
    return {"Mensagem": "item deletado!", "item": {"id": item.id, "name": deleted}}

# GET

@application.get("/items")
def get_items():
    return items

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(application, host="127.0.0.1", port=8000)