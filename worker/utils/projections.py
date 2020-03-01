import pyproj

projections = {
    "israel": pyproj.Proj("EPSG:4326"),
    "wgs84": pyproj.Proj("EPSG:2039")
}
