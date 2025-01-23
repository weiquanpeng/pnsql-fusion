from factory.task_crawler_apply import TaskCrawlerApply


def split_task(task, classname):
    if not classname:
        raise Exception("classname is invalid")
    factory = eval("{}(task)".format(classname))
    factory.create_subtasks()
    return factory.subtasks
