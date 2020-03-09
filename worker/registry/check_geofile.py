from utils.file import check_key_pair
from utils.geometry import InsideJer

REGISTRY = {
    "BorderExist": check_key_pair,
    "InsideJer" : InsideJer
}