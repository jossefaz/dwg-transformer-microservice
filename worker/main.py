from shapely.geometry import Point
import sys
import json
from worker.utils.io import eprint, oprint

def main():

    with open(sys.argv[1], 'r') as f :
        try :
            j = json.load(sys.argv[2])
            oprint(json.dumps(j))
        except Exception as e :
            eprint("cannot convert args to json object :", str(e), sys.argv[2])
        print("OPENED")


if __name__ == "__main__" :
    main()
