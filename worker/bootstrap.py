from shapely.geometry import Point
import sys
import json

from utils.io import oprint, eprint
from worker.utils.path import file_exists
from worker.registry.check_geofile import REGISTRY


def main():
    exist = file_exists(sys.argv[1])
    if exist:
        with open(sys.argv[1], 'r') as f:
            try:
                checks = sys.argv[2].split()
                for check in checks:
                    REGISTRY[check](sys.argv[1])
            except Exception as e:
                eprint("cannot convert args to json object :", str(e), sys.argv[2])
            print("OPENED")


if __name__ == "__main__":
    main()
