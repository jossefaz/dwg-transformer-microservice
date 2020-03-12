from utils.file import border_exists
from utils.geometry import InsideJer

REGISTRY = {
    "BorderExist": {
        "func": border_exists,
        "return": True
    },
    "InsideJer": {
        "func": InsideJer,
        "return": False
    }
}
