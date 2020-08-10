import numpy as np
from matplotlib import pyplot as plt

s = input().split()
niter, wstep = [int(c) for c in s]
K = niter // wstep

num_list = []
for _ in range (K):
    num_list.append (list (map (float, input().split())))

data = np.array (num_list)

x = np.arange(K)

fig = plt.figure (figsize=(6,6))

ax = fig.add_subplot (2,1,1)
ax.plot (x, 2*np.ones_like(x)/3, '--')
ax.plot (x, data[:,0])

ax = fig.add_subplot (2,1,2)
ax.plot (x, 0.1305*np.ones_like(x), '--')
ax.plot (x, data[:,1])
ax.plot (x, data[:,2])
ax.set_ylim ([0.1, 0.16])

fig.show ()
plt.show ()
