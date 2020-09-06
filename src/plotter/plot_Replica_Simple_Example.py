import numpy as np
from matplotlib import pyplot as plt

N = 1000000
raw_data = []
for _ in range (N):
    raw_data.append ([float(c) for c in input().split()])
data = np.array (raw_data)

fig = plt.figure (figsize=(10,7))
hor = np.arange (N)+1

for i in range (3):
    ax = fig.add_subplot (2,3,i+1)
    ax.plot (hor, data[:,i], 'x')

    ax = fig.add_subplot (2,3,i+4)
    ax.hist (data[500000:,i], bins=60)

fig.show ()
plt.show ()
