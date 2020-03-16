import sys
import json

from utils.io import oprint
from utils.path import file_exists
from utils.file import json_from_file, dict_from_geojson, border_exists

from registry.check_geofile import REGISTRY

from utils.file import rm_file
from Store.main import Store


def main():
    exist = file_exists(sys.argv[1])
    mainStore = Store.get_instance()
    if exist:
        geojson = json_from_file(sys.argv[1])
        if geojson :
            geojson_dict = dict_from_geojson(geojson)
            if geojson_dict :
                checks = sys.argv[2].split()
                target = geojson_dict
                for check in checks:
                    if REGISTRY[check]["return"] :
                        target = REGISTRY[check]["func"](target)
                        mainStore.set_result(check, int(bool(target)))
                    else :
                        res = REGISTRY[check]["func"](target)
                        mainStore.set_result(check, int(bool(res)))
            else :
                raise RuntimeError("cannot convert geojson to dict (error in json loading) : {}".format(geojson))
            rm_file(geojson)
            oprint(json.dumps(mainStore.store))
        else :
            raise FileNotFoundError("cannot convert file to geojson : {}".format(sys.argv[1]))
    else :
        raise FileNotFoundError("cannot find input file : {}".format(sys.argv[1]))

if __name__ == "__main__":
    main()
