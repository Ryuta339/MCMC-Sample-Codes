import numpy as np
from matplotlib import pyplot as plt

N = 10000//10

raw_data = []
for _ in range (N):
    raw_data.append ([float(c) for c in input().split()])

data = np.array (raw_data)

fig = plt.figure (figsize=(8,3))

ax = fig.add_subplot (1,3,1)
ax.scatter (data[:,0], data[:,1])
ax = fig.add_subplot (1,3,2)
ax.scatter (data[:,0], data[:,2])
ax = fig.add_subplot (1,3,3)
ax.scatter (data[:,1], data[:,2])

fig.show ()
plt.show ()
