import logging
from datetime import datetime
from conf.config import config
import logging
from datetime import datetime

def logger(log_name, log_level='INFO'):
    # 检查 log_level 是否有效
    if log_level and log_level.upper() not in ('DEBUG', 'INFO', 'WARNING', 'ERROR', 'CRITICAL'):
        raise ValueError("log_level must be one of ('DEBUG', 'INFO', 'WARNING', 'ERROR', 'CRITICAL')")

    # 日志格式
    formatter = logging.Formatter('%(asctime)s - %(levelname)s - %(message)s')

    # 日志文件路径
    log_file = '{}/{}.{}'.format(config["app"]["log_path"], log_name, datetime.today().strftime('%Y%m%d'))

    # 文件处理器
    handler = logging.FileHandler(log_file)
    handler.setFormatter(formatter)

    # 创建日志记录器
    logger = logging.getLogger(log_name)
    logger.setLevel(log_level.upper())  # 设置日志级别
    logger.addHandler(handler)

    return logger