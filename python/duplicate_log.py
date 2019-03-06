# coding=utf-8 
import logging
import logging.handlers

logger = logging.getLogger("logger")
handler = logging.StreamHandler()
handler.setLevel(logging.DEBUG)
formatter = logging.Formatter("%(asctime)s %(name)s %(levelname)s %(message)s")
handler.setFormatter(formatter)
logger.addHandler(handler)
#  logger.propagate = False
logger.debug('This is a customer debug message')
#  如果使用了logging里的方法，这些方案使用了默认配置的logger对象
#  logging.info('This is an customer info message')
logger.warning('This is a customer warning message')
logger.error('This is an customer error message')
logger.critical('This is a customer critical message')

