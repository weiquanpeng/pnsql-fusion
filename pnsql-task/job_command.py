import os
import json

JOB_PATH = "{}/job/".format(os.path.abspath(os.path.dirname(__file__)))
CRONJOB_PATH = "{}/cronjob/".format(os.path.abspath(os.path.dirname(__file__)))


class JobCommand(object):
    def __init__(self, subtask_config, job):
        self.job = job
        if isinstance(job['paras'], str):
            self.job['paras'] = json.loads(job['paras'])
        if subtask_config['exec_type'] == 'bash':
            if subtask_config['type'] == 'job':
                if not subtask_config['script'].startswith('/'):
                    subtask_config['script'] = JOB_PATH + subtask_config['script']
            elif subtask_config['type'] == 'cronjob':
                subtask_config['script'] = CRONJOB_PATH + subtask_config['script']
        self.config = subtask_config
        self.command = self.create_command()

    def create_command(self):
        options = self.create_options()
        command = [self.config["interpreter"], self.config["script"], ' '.join(options)]
        # command = "{} {} {}".format(self.config["interpreter"], self.config["script"], ' '.join(options))
        return command

    def create_options(self):
        options = ["--jobid={}".format(self.job["id"])]
        if 'paras' in self.config:
            for para in self.config['paras']:
                if para not in self.job["paras"]:
                    continue
                elif para in self.job["paras"] and not self.job["paras"][para]:
                    options.append("--{}".format(para))
                else:
                    options.append("--{opt}={val}".format(opt=para, val=self.job["paras"][para]))
        return options