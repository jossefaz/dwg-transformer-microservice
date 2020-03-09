import collections
import csv
from functools import partial

import pyproj
import shapely.wkt as wktload
from shapely.ops import transform
from shapely.geometry import Point, LineString, MultiPolygon

from utils import ogr2ogr
from utils.projections import *
from utils.path import *
import os


def get_transformer(from_crs, to_crs):
    project = partial(pyproj.transform, from_crs, to_crs)
    return project


def dxf_to_geojson(dxf):
    basename = os.path.basename(dxf).split('.')[0]
    outfile = "{}.json".format(basename)
    ogr2ogr.main(["", "-f", "GeoJson", outfile, dxf])
    return outfile


def get_jerusalem_border():
    project = get_transformer(projections['wgs84'], projections['israel'])
    jerusalem_polygon = []
    work_file = os.path.join(GetParentDir(os.path.dirname(__file__)), 'ressource/border2.csv')
    polylist = list(csv.reader(open(work_file, 'r'), delimiter='|'))
    for geom in polylist[0]:
        polygon = wktload.loads(geom)
        converted = transform(project, polygon)
        jerusalem_polygon.append(converted)
    return jerusalem_polygon

def line_is_closed(xy_list) :
    # xy_dict = {}
    # for xy in xy_list :
    #     if len(xy == 2) :
    #         if xy[0] in xy_dict :
    #             xy_dict[xy[0]]


    return True

def inside_polygon(polygon, feature) :
    return True

def InsideJer(xy_list) :
    return True

