"""
@Time : 2021/1/26 上午11:31
@Author: machengcheng, xuxiaogang
@File : utils.py
"""
from spider import query_train_list, get_stopovers, get_train_detail_info
import json
import datetime

def crawl_and_save(from_station:str, to_station:str, date:datetime.datetime, crawl_date_num:int):
    """
    Save data.
    :param from_station: From station.
    :param to_station: To station.
    :param date: The begin date.
    :param crawl_date_num: The dates'num to be crawled.
    """
    for num in range(0, crawl_date_num):
        train_date = (date + datetime.timedelta(days=num)).strftime('%Y-%m-%d')
        train_list = []
        train_list1 = []
        
        train_list = query_train_list(from_station, to_station, train_date)

        for train in train_list:
            train_num = train['TrainNo']

            stopovers = get_stopovers(from_station, to_station, train_date, train_num)

            train_detail_info = get_train_detail_info(from_station, to_station, train_date, train_num)

            train_dict = {
                'train': train,
                'stopovers': stopovers,
                'train_detail_info': train_detail_info
            }
            train_list1.append(train_dict)
        
        test_dict = {
            'date': train_date,
            'from': from_station,
            'to': to_station,
            'train_list': train_list1
        }
        fileObject = open('./data/'+from_station+'+'+to_station+'+'+train_date+'.json', 'w')
        jsObj = json.dumps(test_dict, ensure_ascii=False, indent=4)
        fileObject.write(jsObj)
        fileObject.close()
        print(from_station+'+'+to_station+'+'+train_date+'.json'+' -----------------------')



