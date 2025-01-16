from setuptools import setup, find_packages

setup(
    name = 'stormgeo.advisor-core',
    version = '0.1.0',
    author = 'climatempo',
    packages = find_packages(include=['src']),
    license = 'MIT',
    description = 'SDK to access the advisor core API.',
    long_description = open("README.md", "r", encoding="utf-8").read(),
    url = 'https://github.com/StormGeo/advisor-sdk',
    project_urls = {
        'Código fonte': 'https://github.com/StormGeo/advisor-sdk/tree/main/python-advisor-core'
    }
)