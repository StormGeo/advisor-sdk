import os
import sys
from pathlib import Path

import pytest


PACKAGE_ROOT = Path(__file__).resolve().parents[1]
if str(PACKAGE_ROOT) not in sys.path:
    sys.path.insert(0, str(PACKAGE_ROOT))

from advisor_core import AdvisorCore  # noqa: E402


def _get_env_int(name, default):
    value = os.getenv(name)
    return int(value) if value else default


def _require_env(name):
    value = os.getenv(name)
    if not value:
        raise pytest.UsageError(f"Set {name} before running the route tests.")
    return value


@pytest.fixture(scope="session")
def advisor():
    token = _require_env("ADVISOR_TOKEN")
    advisor = AdvisorCore(token, retries=1, delay=0)
    advisor.setHeaderAccept("application/json")
    advisor.setHeaderAcceptLanguage(os.getenv("ADVISOR_ACCEPT_LANGUAGE", "en-US"))
    return advisor


@pytest.fixture(scope="session")
def locale_id():
    return _get_env_int("ADVISOR_LOCALE_ID", 3477)


@pytest.fixture(scope="session")
def plan_locale_id():
    return _get_env_int("ADVISOR_PLAN_LOCALE_ID", 5959)
