"""
@Time : 2021/1/26 上午11:24
@Author: machengcheng, xuxiaogang
@File : main.py
"""

from spider import query_train_list, get_stopovers, get_train_detail_info
from utils import save

if __name__ == '__main__':
    from_station = '北京'
    to_station = '上海'
    train_date = '2021-01-28'

    train_list = query_train_list(from_station, to_station, train_date)

    print(train_list)

    for train in train_list:
        train_num = train['TrainNo']

        stopovers = get_stopovers(from_station, to_station, train_date, train_num)
        train_detail_info = get_train_detail_info(from_station, to_station, train_date, train_num)

        save(train, stopovers, train_detail_info)
