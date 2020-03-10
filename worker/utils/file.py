import json
import geojson
import pandas as pd

from utils.geometry import dxf_to_geojson
def check_key_pair(searched_key, searched_val, dict_to_check):
    for key, val in dict_to_check.items() :
        if key == searched_key and val == searched_val :
            return True
        if isinstance(val, dict) :
            if check_key_pair(searched_key, searched_val, val) :
                return True




def border_exists(geojson_dict) :
    return check_key_pair("Layer", "Border", geojson_dict)




def dict_from_geojson(jsonfile) :
    try :
        converted_json = pd.read_json(jsonfile).to_dict()
        return converted_json
    except Exception as e:
        print(str(e))
    return False

def json_from_file(filepath) :
    if filepath.lower().endswith(('.json')) :
        return filepath
    if filepath.lower().endswith(('.dxf')) :
        try :
            dxf = dxf_to_geojson(filepath)
            return dxf
        except Exception as e :
            print(str(e))
            return False
