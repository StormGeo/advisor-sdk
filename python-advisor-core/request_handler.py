import requests
import time

class RequestHandler:
    def __init__(self, base_url, token, retries, delay):
        self.base_url = base_url
        self.token = token
        self.retries = retries
        self.delay = delay
        self.session = requests.Session()

    def make_request(self, method, endpoint, params=None, json_data=None, retries=None):
        retries = retries if retries is not None else self.retries
        full_url = f"{self.base_url}{endpoint}"
        error_message = ''

        try:
            if method == "GET":
                response = self.session.get(full_url, params=params)
            elif method == "POST":
                response = self.session.post(full_url, params=params, json=json_data)
            else:
                response = self.session.request(method, full_url, json=json_data)

            if response.status_code != 200:
                error_message = response.json().get("error", response.text)
            
            response.raise_for_status()
            return {"data": response.json(), "error": None}

        except requests.exceptions.RequestException as error:
            if retries > 0:
                time.sleep(self.delay)
                print(f"Re-trying in {self.delay}s... attempts left: {retries}")
                return self.make_request(method, endpoint, params, json_data, retries - 1)

            return {"data": None, "error": error_message if error_message != '' else error}
