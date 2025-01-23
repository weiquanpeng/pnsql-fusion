from scheduler import Scheduler

if __name__ == "__main__":
    try:
        scheduler = Scheduler()
        scheduler.run()
    except Exception as e:
        print(e)