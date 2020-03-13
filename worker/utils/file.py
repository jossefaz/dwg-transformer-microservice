import os
from pathlib import Path
import pandas as pd

from utils.geometry import dxf_to_geojson


def check_key_pair(searched_key, searched_val, dict_to_check, starting_key=None, return_dict=False):

    if starting_key :
        if starting_key not in dict_to_check :
            print("the requested strarting key in not in the target dict")
            return False
        dict_to_check = dict_to_check[starting_key]
    for key, val in dict_to_check.items():
        if key == searched_key and val == searched_val:
            return dict_to_check if return_dict else True
        if isinstance(val, dict):
            if check_key_pair(searched_key, searched_val, val):
                return dict_to_check[key] if return_dict else True


def border_exists(geojson_dict):

    return check_key_pair("Layer", "Border", geojson_dict, "features", True)["geometry"]


def rm_file(path):
    try:
        os.remove(path)
    except:
        try:
            Path.unlink(path)
        except Exception as e:
            print(str(e))


def dict_from_geojson(jsonfile):
    try:
        converted_json = pd.read_json(jsonfile).to_dict()
        return converted_json
    except Exception as e:
        print(str(e))
    return False


def json_from_file(filepath):
    if filepath.lower().endswith(('.json')):
        return filepath
    if filepath.lower().endswith(('.dxf')):
        try:
            dxf = dxf_to_geojson(filepath)
            return dxf
        except Exception as e:
            print(str(e))
            return False
