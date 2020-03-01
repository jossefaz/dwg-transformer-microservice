import csv
from functools import partial

import shapely.wkt as wktload
from shapely.ops import transform
from shapely.geometry import Point, LineString, MultiPolygon
from worker.utils.projections import *
from worker.utils.path import *
import os

def getTransformer(fromCRS, toCRS):
    project = partial(pyproj.transform,fromCRS, toCRS)
    return project

def getJerusalemBorder() :
    project = getTransformer(projections['wgs84'], projections['israel'])
    jerusalem_polygon = []
    work_file = os.path.join(GetParentDir(os.path.dirname(__file__)), 'ressource/border2.csv')
    polylist = list(csv.reader(open(work_file, 'r'), delimiter='|'))
    for geom in polylist[0]:
        polygon = wktload.loads(geom)
        converted = transform(project, polygon)
        jerusalem_polygon.append(converted)
    return jerusalem_polygon


