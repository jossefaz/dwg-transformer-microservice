from pathlib import Path
import os
def GetParentDir(path) :
    path = Path(path)
    return path.parent

def safeOpen(path) :
    Path(path).mkdir(parents=True, exist_ok=True)

def list_files(startpath):
    for root, dirs, files in os.walk(startpath):
        level = root.replace(startpath, '').count(os.sep)
        indent = ' ' * 4 * (level)
        print('{}{}/'.format(indent, os.path.basename(root)))
