import json
import pymysql
import datetime
from util.basedb import BaseDb
from util.log import logger

_logger = logger(__name__)


class TaskDB(BaseDb):
    def __init__(self, conn_setting):
        BaseDb.__init__(self, conn_setting)

    def get_subtask_by_id(self, id):
        sql = "select * from process_subtask where id=%s"
        return self.fetchone(sql, id)
    def get_tosplit_tasks(self):
        sql = "SELECT * FROM task_config where status='tosplit'"
        return self.fetchall(sql)

    def get_task_config(self, name):
        sql = "SELECT * FROM task_config WHERE `type` = %s"
        return self.fetchone(sql, (name,))

    def update_task_status(self, id, status):
        sql = "UPDATE `task_config` SET `status`=%s WHERE id=%s"
        _logger.info(sql, status, id)
        self.execute(sql, (status, id))

    def get_todo_subtasks(self):
        """
        sql = "select * from (" \
              "select ps.* from `task_config` pt inner join `subtask_config` ps on pt.id=ps.`task_id` " \
              "where pt.status in ('todo','doing') and ps.status != 'done' group by ps.task_id) cs " \
              "where status='todo' and executor='system'"
        """
        sql = "SELECT * FROM subtask_config where `status`='todo' and `owner`='system'"
        return self.fetchall(sql)

    def update_subtask_status(self, id, status):
        sql = "UPDATE `subtask_config` SET `status`=%s WHERE id=%s"
        _logger.info(sql, status, id)
        self.execute(sql, (status, id))

    def get_approve_subtasks(self):
        sql = (
            "WITH ranked_subtasks AS ("
            "    SELECT s.id, s.created_at, s.updated_at, s.title, s.type, s.owner, s.approve, "
            "           s.task_id, s.status, s.parameter, s.task_describe, "
            "           ROW_NUMBER() OVER (PARTITION BY s.task_id ORDER BY s.id) AS row_num "
            "    FROM subtask_config s "
            "    WHERE s.task_id IN (SELECT t.id FROM task_config t WHERE t.status IN ('todo', 'doing')) "
            "      AND s.status!= 'done'"
            ") "
            "SELECT st.id, st.created_at, st.updated_at, st.title, st.type, st.owner, st.approve, "
            "       st.task_id, st.status, st.parameter, st.task_describe "
            "FROM ranked_subtasks st "
            "WHERE st.row_num = 1"
        )
        return self.fetchall(sql)

    def add_subtask(self, subtasks):
        sql = """
        INSERT INTO subtask_config (
            task_id, title, owner, approve, status, parameter, task_describe, type
        ) VALUES (
            %s, %s, %s, %s, %s, %s, %s, %s
        )
        """
        try:
            for subtask in subtasks:
                # 将 parameter 字典转换为 JSON 字符串
                parameter_json = json.dumps(subtask['parameter'], ensure_ascii=False)
                # 构造插入数据的元组
                data = (
                    subtask['taskid'],
                    subtask['title'],
                    subtask['owner'],
                    subtask['approve'],
                    subtask['status'],
                    parameter_json,
                    subtask['task_describe'],
                    subtask['type'],
                )
                # 执行插入操作
                self.execute(sql, data)
        except Exception as e:
            _logger.error("插入子任务数据失败: {}".format(e))
            raise

        def get_subtask_config(self, name):
            sql = "SELECT * FROM process_subtask_config where `name`=%s"
            return self.fetchone(sql, name)