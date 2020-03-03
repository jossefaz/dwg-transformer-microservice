from shapely.geometry import Point
import sys
import json
from utils.io import eprint, oprint

def main():

    with open(sys.argv[1], 'r') as f :
        try :
            j = sys.argv[2].split()
            oprint(j)
        except Exception as e :
            eprint("cannot convert args to json object :", str(e), sys.argv[2])
        print("OPENED")


if __name__ == "__main__" :
    main()
