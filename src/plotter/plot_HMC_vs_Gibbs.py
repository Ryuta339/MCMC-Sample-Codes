import numpy as np
from matplotlib import pyplot as plt

N = int(1000000//10)
raw_data = []
for _ in range (N):
    raw_data.append ([float(c) for c in input().split()])
data1 = np.array (raw_data)


raw_data = []
for _ in range (N):
    raw_data.append ([float(c) for c in input().split()])
data2 = np.array (raw_data)

plt.hist (data1[:,0]*data1[:,1], range=(-15,5), bins=40, density=True, alpha=0.6)
plt.hist (data2[:,0]*data2[:,1], range=(-15,5), bins=40, density=True, alpha=0.6)
plt.show ()
