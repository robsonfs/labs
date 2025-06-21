import multiprocessing
import math
import time
import os

def cpu_stress():
    """
    Function that continuously performs heavy CPU calculations
    to stress a CPU core.
    """
    print(f"Starting stress on process {os.getpid()}")
    x = 0.0001
    while True:
        # Heavy floating point calculation to generate heat and load
        x = math.sqrt(math.exp(math.log(x + 1) * math.sin(x + 1)))

def main():
    num_cores = multiprocessing.cpu_count()
    print(f"Detected {num_cores} CPU cores. Spawning one process per core...")

    processes = []

    for i in range(num_cores):
        p = multiprocessing.Process(target=cpu_stress)
        processes.append(p)
        p.start()

    try:
        # Keep the main process alive to let workers run
        while True:
            time.sleep(10)
    except KeyboardInterrupt:
        print("\nTerminating stress test...")
        for p in processes:
            p.terminate()
        for p in processes:
            p.join()

if __name__ == "__main__":
    main()
