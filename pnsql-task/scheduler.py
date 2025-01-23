import sys
import time
import datetime
import json
import traceback
from conf.config import config
from model.taskdb import TaskDB
from util.command import safe_execute_local_cmd
from factory.task_split import split_task
from util.log import logger



class Scheduler(object):
    def __init__(self):
        conn_setting = dict(config['database'])
        self.db = TaskDB(conn_setting)
        self._logger = logger(__name__)

    def split_tasks(self):
        tasks = self.db.get_tosplit_tasks()
        for task in tasks:
            try:
                if not task:
                    continue
                subtasks = split_task(task, task['type'])
                self.db.add_subtask(subtasks)
                self.db.update_task_status(task['id'], 'todo')
            except Exception as e:
                errmsg = "split task error <{}>, {}".format(task['id'], traceback.format_exc())
                self._logger.error(errmsg)

    def execute_subtasks(self):
        sub_ids = self.db.get_approve_subtasks()
        for sub_id in sub_ids:
            if sub_id['status'] =='tosplit':
                if sub_id['type'] == 'SubTaskDone':
                    self.db.update_subtask_status(sub_id['id'], 'done')
                    continue
                self.db.update_subtask_status(sub_id['id'], 'todo')

        subtasks = self.db.get_todo_subtasks()
        for subtask in subtasks:
            parameter = json.loads(subtask['parameter'])
            # if 'trigger_time' in paras and paras['trigger_time']:
            #     trigger_time = datetime.datetime.strptime(paras['trigger_time'], "%Y-%m-%d %H:%M:%S")
            #     if trigger_time > datetime.datetime.now():
            #         continue
            #
            self.db.update_subtask_status(subtask['id'], 'doing')
            try:
                cmd_array = ['python', '-u', 'worker.py', '--jobid={}'.format(subtask['id'])]
                print(cmd_array)
                # safe_execute_local_cmd(cmd_array)
            except Exception as e:
                errmsg = "worker crashed, subtask <{}>, {}".format(subtask['id'], traceback.format_exc())
                self._logger.error(errmsg)

    def run(self):
        while True:
            self.split_tasks()
            self.execute_subtasks()
            time.sleep(1)