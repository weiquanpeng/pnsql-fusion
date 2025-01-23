from factory.task_factory import TaskFactory

class TaskCrawlerApply(TaskFactory):
    def create_subtasks(self):
        subtask = [
            {
                'taskid': 1,
                'title': 'owner 审批',
                'owner': 'admin',
                'approve': 'gaogao',
                'status': 'tosplit',
                'parameter': self.task['parameter'],
                'task_describe': '这是一个爬虫任务',
                'type': None,
            },
            {
                'taskid': 1,
                'title': 'leader 审批',
                'owner': 'admin',
                'approve': 'wanguangqiu,sikui',
                'status': 'tosplit',
                'parameter': self.task['parameter'],
                'task_describe': '这是一个爬虫任务',
                'type': None,
            },
            {
                'taskid': 1,
                'title': '任务执行',
                'owner': 'system',
                'approve': 'admin',
                'status': 'tosplit',
                'parameter': self.task['parameter'],
                'task_describe': '这是一个爬虫任务',
                'type': None,
            },
            {
                'taskid': 1,
                'title': '工单完成',
                'owner': 'done',
                'approve': 'admin',
                'status': 'tosplit',
                'parameter': self.task['parameter'],
                'task_describe': '这是一个爬虫任务',
                'type': None,
            }
        ]
        # 使用循环逐个追加
        for task in subtask:
            self.subtasks.append(task)
