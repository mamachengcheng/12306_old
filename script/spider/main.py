"""
@Time : 2021/1/26 上午11:24
@Author: machengcheng, xuxiaogang
@File : main.py
"""


from utils import crawl_and_save
import datetime

if __name__ == '__main__':
    
    from_station = '北京'
    to_station = '上海'
    date = datetime.datetime(2021, 2, 4)
    crawl_date_num = 30
    crawl_and_save(from_station, to_station, date, crawl_date_num)
    
    
