import numpy as np
from matplotlib import pyplot as plt

N = 100
raw_data = []
for _ in range (N):
    raw_data.append ([float(c) for c in input().split()])

data = np.array (raw_data)

plt.plot (np.arange(N), data)
plt.show ()
