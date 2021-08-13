import requests

if __name__ == '__main__':
    res = requests.post('http://localhost:9200/_xpack/license/start_basic?acknowledge=true')
    print(res.json())
