import os
import sys


def daemonize(stdin='/dev/null', stdout='/dev/null', stderr='/dev/null'):
    try:
        if os.fork() > 0:
            sys.exit(0)
    except Exception as e:
        sys.stderr.write("fork #1 failed : ({no}) {msg} \n".format(no=e.errorno, msg=e.strerror))
        sys.exit(1)
    # os.chdir('/')
    os.umask(0)
    os.setsid()

    try:
        if os.fork() > 0:
            sys.exit(0)
    except Exception as e:
        sys.stderr.write("fork #2 failed : ({no}) {msg} \n".format(no=e.errorno, msg=e.strerror))
        sys.exit(1)

    # 重定向标准IO
    sys.stdout.flush()
    sys.stderr.flush()

    stdin_new = open(stdin, 'r')
    stdout_new = open(stdout, 'w')
    stderr_new = open(stderr, 'w')

    os.dup2(stdin_new.fileno(), sys.stdin.fileno())
    os.dup2(stdout_new.fileno(), sys.stdout.fileno())
    os.dup2(stderr_new.fileno(), sys.stderr.fileno())