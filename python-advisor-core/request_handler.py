import requests
import time

class RequestHandler:
    def __init__(self, base_url, token, retries=2, delay=1):
        self.base_url = base_url
        self.token = token
        self.retries = retries
        self.delay = delay
        self.session = requests.Session()

    def make_request(self, method, endpoint, params=None, json_data=None, retries=None):
        retries = retries if retries is not None else self.retries
        full_url = f"{self.base_url}{endpoint}"

        try:
            if method == "GET":
                response = self.session.get(full_url, params=params)
            else:
                response = self.session.request(method, full_url, json=json_data)

            response.raise_for_status()
            return {"data": response.json(), "error": None}

        except requests.exceptions.RequestException as error:
            if retries > 0:
                time.sleep(self.delay)
                print(f"Re-trying in {self.delay}s... attempts left: {retries}")
                return self.make_request(method, endpoint, params, json_data, retries - 1)

            return {"data": None, "error": str(error)}
