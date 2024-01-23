from urllib.parse import urlparse

from bson.objectid import ObjectId
from pymongo import MongoClient

from config import config


class MongodbClient(object):
    def __init__(self):
        self.mongodb = MongoClient(host=config.CONST_MONGODB_CONFIG.get('host'),
                                   port=config.CONST_MONGODB_CONFIG.get('port'),
                                   username=config.CONST_MONGODB_CONFIG.get('username'),
                                   password=config.CONST_MONGODB_CONFIG.get('password'),
                                   authSource=config.CONST_MONGODB_CONFIG.get('database'),
                                   unicode_decode_error_handler='ignore',
                                   )
        self.db = self.mongodb[config.CONST_MONGODB_CONFIG.get('database')]
        self.collection = self.db[config.CONST_MONGODB_CONFIG.get('collection')]
        self.coll_password = self.db["password"]

    def save_password_to_db(self):
        cursor = self.collection.find({"flag": 0, "request_parameters": {"$ne": {}}})
        cursor.batch_size(1)
        for record in cursor:
            url = record.get('url')
            url_parse = urlparse(url)
            site = url_parse.netloc
            from_ip = record.get('origin')
            request_body = record.get('request_body')
            request_header = record.get('request_header')
            header = record.get('header')
            body = record.get('body')
            date_start = record.get('date_start')

            self.collection.update_one({"_id": ObjectId(record.get("_id"))},
                                       {"$set": {"flag": 1}})

            request_parameters = record.get('request_parameters')
            keys = request_parameters.keys()
            intersection = get_intersection(keys, config.CONST_KEYWORD)
            if len(intersection) >= 2:
                ret = dict()
                for i in intersection:
                    t = dict()
                    t[i] = record.get('request_parameters').get(i)[0]
                    ret.update(t)

                value = dict(
                    site=site,
                    url=url,
                    from_ip=from_ip,
                    data=ret,
                    request_parameters=request_parameters,
                    request_header=request_header,
                    request_body=request_body,
                    header=header,
                    body=body,
                    date_start=date_start,
                    status=0,
                )
                print("URL: {}, DATA: {}".format(url, ret))

                self.coll_password.update_one({"site": site, "data": ret},
                                              {"$set": value},
                                              upsert=True)
                # self.coll_password.insert(value)

        cursor.close()

    def clean_password(self):
        self.coll_password.delete_many({})


def get_intersection(a, b):
    """return intersection of two lists"""
    return list(set(a).intersection(b))
