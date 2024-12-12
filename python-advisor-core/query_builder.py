
class QueryParamsBuilder:
    def __init__(self):
        self.params = {}

    def add(self, key, value):
        if value is not None:
            self.params[key] = value
        return self

    def build(self):
        return self.params
