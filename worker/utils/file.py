import json
import geojson
import pandas as pd

from utils.geometry import dxf_to_geojson
def check_key_pair(k, v, dict_to_check):
    return (k, v) in dict_to_check.viewitems()

def border_exists(geojson_dict) :

    return True

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
