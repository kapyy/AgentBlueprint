import signal
import sys
import threading
import time

from factory.factory_server import FunctionServerThread
from function.core_server import CoreServerThread


def signal_handler(sig, frame):
    print("Signal received, stopping servers...")
    coreThread.stop()  # Graceful shutdown with a timeout of 0 seconds
    funcThread.stop()
    sys.exit(0)

signal.signal(signal.SIGINT, signal_handler)
signal.signal(signal.SIGTERM, signal_handler)

if __name__== '__main__':
    # threading.Thread(target=FactoryServerStart).start()
    coreThread = CoreServerThread()
    funcThread = FunctionServerThread()

    coreThread.start()
    funcThread.start()
    try:
        while True:
            time.sleep(86400)  # Main thread can perform other tasks here
    except KeyboardInterrupt:
        print("KeyboardInterrupt received, stopping server...")
        coreThread.stop()
        funcThread.stop()
        coreThread.join()
        funcThread.join()