"""
@Time : 2021/1/26 上午11:24
@Author: machengcheng, xuxiaogang
@File : main.py
"""


import requests
import configparser

conf = configparser.RawConfigParser()
conf.read('conf.ini')


def query_train_list(from_station: str, to_station: str, train_date: str) -> list:
    """
    Query train list.
    :param from_station: Start station.txt.
    :param to_station: End station.txt.
    :param train_date: Train date
    :return: Train list.
    """
    url = conf['url']['query_train_list'].format(
        **{'from_station': from_station, 'to_station': to_station, 'train_date': train_date})
    train_list = requests.get(url).json()['data']['Trains']
    return train_list


def get_stopovers(from_station: str, to_station: str, train_date: str, train_num: str) -> list:
    """
    Get stopovers.
    :param from_station: Start station.txt.
    :param to_station: End Station.
    :param train_date: Date.
    :param train_num: Train number.
    :return: The stopovers of station.txt.
    """
    pass
    # TODO: Completion code
    stopovers = list()
    return stopovers


def get_train_detail_info(from_station: str, to_station: str, train_date: str, train_num: str) -> dict:
    """
    Get train detail information.
    :param from_station: Start station.txt.
    :param to_station: End Station.
    :param train_date: Date.
    :param train_num: Train number.
    :return: The train detail information.
    """
    pass
    # TODO: Completion code
    train_detail_info = dict()
    return train_detail_info
