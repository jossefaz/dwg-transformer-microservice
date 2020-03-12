import logging
from functools import reduce  # forward compatibility for Python 3
import operator


class Store:
    __instance = None

    def __init__(self):
        if not Store.__instance:
            self.store = {}
            self.logger = logging.getLogger(__name__)

    @classmethod
    def get_instance(cls):
        if cls.__instance is None:
            cls.__instance = Store()
        return cls.__instance

    def get_result(self, path_string):
        path = path_string.split('.') if isinstance(path_string, str) else path_string
        try:
            result = reduce(operator.getitem, path, self.store)
        except KeyError as e:
            self.logger.error(
                "The result path '{}' was not found in store, check mispelling in the LINKED_TO key and be sur to be connected to a result name path".format(
                    path_string))
            raise
        return result

    def set_result(self, path_string, value):
        path = path_string.split('.')
        self.get_result(path[:-1])[path[-1]] = value


if __name__ == '__main__':
    pass
