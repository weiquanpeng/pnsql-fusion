import argparse

import pandas as pd
import requests
import pymysql
import baostock as bs
import datetime
import threading
from urllib.parse import urlencode
import sys
import os

# project_root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
# sys.path.append(project_root)
import worker

# 数据库配置
DB_CONFIG = {
    'host': '127.0.0.1',
    'port': 3306,
    'user': 'pengweiquan',
    'password': '123123',
    'database': 'pwq',
    'charset': 'utf8mb4',
    'autocommit': True
}


def parse_args(args):
    parser = argparse.ArgumentParser(description='worker', add_help=False)
    parser.add_argument('--jobid', dest='jobid', type=str, help='jobid')
    parser.add_argument('--jobtype', dest='jobtype', type=str, help='jobtype')
    parser.add_argument('--start_time', dest='start_time', type=str, help='start_time')
    parser.add_argument('--stop_time', dest='stop_time', type=str, help='stop_time')
    parser.add_argument('-h', '--help', dest='help', action='store_true', help='help information', default=False)
    options = parser.parse_args(args)
    if options.help:
        parser.print_help()
        sys.exit(1)
    if not options.jobid:
        raise ValueError('lack of argument: jobid')
    return options

def get_stock_codes(date=None):
    bs.login()
    stock_df = bs.query_all_stock(date).get_data()
    if 0 == len(stock_df):
        if date is not None:
            print('当前选择日期为非交易日或尚无交易数据，请设置date为历史某交易日日期')
            sys.exit(0)
        delta = 1
        while 0 == len(stock_df):
            stock_df = bs.query_all_stock(datetime.date.today() - datetime.timedelta(days=delta)).get_data()
            delta += 1
    bs.logout()

    stock_df = stock_df[(stock_df['code'] >= 'sh.600000') & (stock_df['code'] < 'sz.399000')]
    return stock_df['code'].tolist()


def gen_secid(rawcode: str) -> str:
    if rawcode[:3] == '000':
        return f'1.{rawcode}'
    if rawcode[:3] == '399':
        return f'0.{rawcode}'
    if rawcode[0] != '6':
        return f'0.{rawcode}'
    return f'1.{rawcode}'


def get_k_history(code: str, beg: str, end: str, klt: int = 101, fqt: int = 1) -> pd.DataFrame:
    EastmoneyKlines = {
        'f51': '日期',
        'f52': '开盘',
        'f53': '收盘',
        'f54': '最高',
        'f55': '最低',
        'f56': '成交量',
        'f57': '成交额',
        'f58': '振幅',
        'f59': '涨跌幅',
        'f60': '涨跌额',
        'f61': '换手率',
    }
    EastmoneyHeaders = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko',
        'Accept': '*/*',
        'Accept-Language': 'zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2',
        'Referer': 'http://quote.eastmoney.com/center/gridlist.html',
    }
    fields = list(EastmoneyKlines.keys())
    columns = list(EastmoneyKlines.values())
    fields2 = ",".join(fields)
    secid = gen_secid(code)
    params = (
        ('fields1', 'f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13'),
        ('fields2', fields2),
        ('beg', beg),
        ('end', end),
        ('rtntype', '6'),
        ('secid', secid),
        ('klt', f'{klt}'),
        ('fqt', f'{fqt}'),
    )
    params = dict(params)
    base_url = 'https://push2his.eastmoney.com/api/qt/stock/kline/get'
    url = base_url + '?' + urlencode(params)
    json_response: dict = requests.get(url, headers=EastmoneyHeaders).json()

    data = json_response.get('data')
    if data is None:
        if secid[0] == '0':
            secid = f'1.{code}'
        else:
            secid = f'0.{code}'
        params['secid'] = secid
        url = base_url + '?' + urlencode(params)
        json_response: dict = requests.get(url, headers=EastmoneyHeaders).json()
        data = json_response.get('data')
    if data is None:
        print('股票代码:', code, '可能有误')
        return pd.DataFrame(columns=columns)

    klines = data['klines']
    rows = []
    for _kline in klines:
        kline = _kline.split(',')
        rows.append(kline)
    return rows


def hebing(p, y):
    averages = []
    len_list = 1
    while len_list < y:
        averages.append(0)
        len_list = len_list + 1
    for i in range(len(p) - (y - 1)):
        subset = p[i:i + y]
        subset = [float(x) for x in subset]
        average = sum(subset) / len(subset)
        average = round(average, 3)
        averages.append(format(average, '.3f'))
    return averages


def p_start(p_code):
    # 每个线程使用独立的数据库连接
    with pymysql.connect(**DB_CONFIG) as conn:
        df = get_k_history(p_code, start_date, end_date)

        p_list = [i[2] for i in df]
        x6 = hebing(p_list, 5)
        x18 = hebing(p_list, 10)
        x72 = hebing(p_list, 20)
        x96 = hebing(p_list, 30)
        x144 = hebing(p_list, 60)
        x288 = hebing(p_list, 250)

        c_list = [x + [y] for x, y in zip(df, x6)]
        c_list = [x + [y] for x, y in zip(c_list, x18)]
        c_list = [x + [y] for x, y in zip(c_list, x72)]
        c_list = [x + [y] for x, y in zip(c_list, x96)]
        c_list = [x + [y] for x, y in zip(c_list, x144)]
        c_list = [x + [y] for x, y in zip(c_list, x288)]

        p_list2 = [i[5] for i in df]
        m6 = hebing(p_list2, 5)
        c_list = [x + [y] for x, y in zip(c_list, m6)]

        with conn.cursor() as cursor:
            for i in c_list:
                insert_query = """
                INSERT INTO pwq_scecss (f0,f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13,f14,f15,f16,f17,f18) 
                VALUES (%s,%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
                """
                cursor.execute(insert_query, (p_code,) + tuple(i))
        print("股票 %s 已获取完成" % (p_code))


def execute_with_concurrency(stock2, max_concurrent):
    counter = 0
    lock = threading.Lock()

    def worker():
        nonlocal counter
        while True:
            with lock:
                if counter >= len(stock2):
                    break
                num = stock2[counter]
                counter += 1
            p_start(num)

    threads = []
    for _ in range(max_concurrent):
        thread = threading.Thread(target=worker)
        thread.start()
        threads.append(thread)

    for thread in threads:
        thread.join()


if __name__ == "__main__":
    options = worker.parse_args(sys.argv[1:])
    # 初始化数据库连接
    with pymysql.connect(**DB_CONFIG) as conn:
        with conn.cursor() as cursor:
            cursor.execute("TRUNCATE TABLE pwq_scecss")
        print("\n--------------------------------------初始化表完成-----------------------------------------------\n")

        stock_codes = get_stock_codes()
        pwq_s = ','.join([s.replace('.', '') for s in stock_codes])
        pd.set_option('expand_frame_repr', False)
        pd.set_option('display.max_rows', 5000)
        stock_code = pwq_s
        stock_list = stock_code.split(',')
        stock_list = [s[2:] for s in stock_list]
        print("--------------------------------------股票代码获取完毕-----------------------------------------------")

        # 开始日期
        start_date = options.start_time
        # 结束日期
        end_date = options.stop_time

        # 执行任务
        execute_with_concurrency(stock_list, max_concurrent=2)
