import json

class TaskFactory(object):
    def __init__(self, task):
        self.task = task
        if isinstance(task['parameter'], str):
            self.task['parameter'] = json.loads(task['parameter'])
        self.subtasks = [
            {
                'taskid': self.task['id'],
                'title': '提交工单',
                'owner': self.task['owner'],
                'approve': self.task['owner'],
                'status': 'done',
                'parameter': self.task['parameter'],
                'task_describe': '工单提交完毕',
                'type': '',
            },
        ]












