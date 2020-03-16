
import csv
from functools import partial


import pyproj
import shapely.wkt as wktload
from shapely.ops import transform
from shapely.geometry import Polygon, LineString, MultiPolygon

from utils import ogr2ogr
from utils.projections import *
from utils.path import *
import os


def get_transformer(from_crs, to_crs):
    project = partial(pyproj.transform, from_crs, to_crs)
    return project

def check_polygons_intersect(polygon, polygon_list) :
    for p in polygon_list :
        if polygon.intersects(p) :
            return True
    return False


def checkPolygonInExtentList(polygon, extent, filtertype=None):
    check = ((polygon.centroid.x > extent[0] and polygon.centroid.x < extent[2]) and (polygon.centroid.y > extent[1] and polygon.centroid.y < extent[3]) )
    if check:
        return True
    return False


def dxf_to_geojson(dxf):
    basename = os.path.basename(dxf).split('.')[0]
    outfile = "{}.json".format(basename)
    ogr2ogr.main(["", "-f", "GeoJson", outfile, dxf])
    return outfile


def get_jerusalem_border():
    project = get_transformer(projections['wgs84'], projections['israel'])
    jerusalem_polygon = []
    work_file = os.path.join(GetParentDir(os.path.dirname(__file__)), 'resource/jer_border.csv')
    polylist = list(csv.reader(open(work_file, 'r'), delimiter='|'))
    for geom in polylist[0]:
        polygon = wktload.loads(geom)
        converted = transform(project, polygon)
        jerusalem_polygon.append(converted)
    return jerusalem_polygon

def is_polygon(poly) :
    return poly.boundary.is_closed and poly.is_valid

def convert_geom_to_polygon(coord) :
    try:
        poly = Polygon(coord)
        return poly
    except:
        return False


def inside_polygon(polygon, feature) :


    return True

def InsideJer(geomObj) :
    if "coordinates" in geomObj :
        polygon = convert_geom_to_polygon(geomObj['coordinates'])
        if polygon :
            if is_polygon(polygon) :
                jerusalem = get_jerusalem_border()
                if check_polygons_intersect(polygon, jerusalem) :
                    return True
    return False






    return True

