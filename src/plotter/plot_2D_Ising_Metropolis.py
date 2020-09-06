import numpy as np
from matplotlib import pyplot as plt

N = 10000
raw_data = []
for i in range (N):
    raw_data.append ([float(c) for c in input().split()])

data = np.array (raw_data)

fig = plt.figure (figsize=(8,4))
ax = fig.add_subplot (1,2,1)
ax.plot (np.arange(N), data[:,0], 'x')
ax = fig.add_subplot (1,2,2)
ax.plot (np.arange(N), data[:,1])
plt.show ()
