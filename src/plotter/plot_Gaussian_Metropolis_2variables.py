import numpy as np
from matplotlib import pyplot as plt

N = 10000//10
fig = plt.figure ()

num_list = []
for i in range (N):
    num_list.append (list (map (float, input().split())))

data = np.array (num_list)

ax = fig.add_subplot (1,1,1)
ax.scatter (data[:,0], data[:,1], marker='+')

fig.show ()
plt.show ()
