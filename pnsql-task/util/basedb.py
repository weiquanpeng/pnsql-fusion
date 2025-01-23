import time
import pymysql
from warnings import filterwarnings
filterwarnings("ignore", category=pymysql.Warning)


class BaseDb(object):
    def __init__(self, conn_setting, transaction=False):
        if not (isinstance(conn_setting, dict) and {'user', 'passwd', 'db'} <= conn_setting.keys()):
            raise Exception("invalid conn_setting.")
        conn_setting['port'] = int(conn_setting['port'])
        conn_setting['autocommit'] = False if transaction else True
        self.conn_setting = conn_setting
        self.conn = None
        self.max_retries = 3
        self.delay = 10
        self.connect_to_db()

    def __del__(self):
        try:
            self.conn.close()
        except Exception:
            pass

    def connect_to_db(self):
        """尝试连接到数据库，支持重试机制"""
        retries = 0
        while retries < self.max_retries:
            try:
                self.conn = pymysql.connect(**self.conn_setting)
                return
            except Exception as e:
                retries += 1
                print(f"数据库连接失败: {e}，正在重试... ({retries}/{self.max_retries})")
                time.sleep(self.delay)
        raise Exception("无法连接到数据库，重试次数已达到上限。")

    def fetchone(self, sql, *parameters):
        self.conn.ping(reconnect=True)
        with self.conn.cursor(pymysql.cursors.DictCursor) as cursor:
            cursor.execute(sql, *parameters)
            return cursor.fetchone()

    def fetchall(self, sql, *parameters):
        self.conn.ping(reconnect=True)
        with self.conn.cursor(pymysql.cursors.DictCursor) as cursor:
            cursor.execute(sql, *parameters)
            return cursor.fetchall()

    def execute(self, sql, *parameters):
        retries = 0
        while retries < self.max_retries:
            try:
                with self.conn.cursor() as cursor:
                    cursor.execute(sql, *parameters)
                    return
            except Exception as e:
                retries += 1
                print(f"执行失败: {e}，尝试重新查询...")
                self.connect_to_db()  # 如果查询失败，重新连接
        raise Exception("执行失败，重试次数已达到上限。")