import sys
def eprint(*args, **kwargs):
    print(*args, file=sys.stderr, **kwargs)

def oprint(*args, **kwargs):
    print(*args, file=sys.stdout, **kwargs)