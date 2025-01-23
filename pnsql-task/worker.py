import argparse
import sys
import traceback
from job_command import JobCommand
from model.taskdb import TaskDB
from util.log import logger
from conf.config import config
from util.command import safe_execute_local_cmd, simple_execute_ssh_cmd
from util.daemonize import daemonize


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

def exec_cmd(command, _logger):
    if command.config["exec_type"] == "ssh":
        cmd = "nohup {} > /dev/null 2>&1 &".format(' '.join(command.command))
        _logger.info(cmd)
        status, stdout, stderr = simple_execute_ssh_cmd(cmd, command.config["script_host"], int(config["ssh"]["port"]),
                                                        config["ssh"]["user"], config["ssh"]["passwd"])
        _logger.info("{} {} {}".format(status, stdout, stderr))
    elif command.config["exec_type"] == "bash":
        _logger.info(command.command)
        stdout = safe_execute_local_cmd(command.command)
        _logger.info(stdout)
    else:
        raise ValueError("unknown exec_type")


class Worker(object):
    def __init__(self, jobid):
        self.logger = logger(__name__, __file__)
        conn_setting = dict(config['database'])
        self.db = TaskDB(conn_setting)
        self.job = self.db.get_subtask_by_id(jobid)

    def run(self):
        command = JobCommand(self.db.get_subtask_config(self.job['type']), self.job)
        self.logger.info(command.command)
        exec_cmd(command, self.logger)


if __name__ == "__main__":
    options = parse_args(sys.argv[1:])
    daemonize()
    worker = Worker(options.jobid)
    try:
        worker.run()
    except Exception as e:
        worker.logger.error(traceback.format_exc())