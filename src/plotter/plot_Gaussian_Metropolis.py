import numpy as np
from matplotlib import pyplot as plt

N = 1000000
fig = plt.figure (figsize=(8,5))

num_list = []
for i in range (N):
    num_list.append (list (map (float, input().split())))

data = np.array (num_list)

x = np.linspace (-3,3,200);
p = np.exp (-x**2/2)/np.sqrt(2*np.pi)

ax = fig.add_subplot(2,4,1)
ax.hist (data[:1000,0], bins=61, density=True)
ax.plot(x,p)

ax = fig.add_subplot(2,4,2)
ax.hist (data[:10000,0], bins=61, density=True)
ax.plot(x,p)

ax = fig.add_subplot(2,4,3)
ax.hist (data[:100000,0], bins=61, density=True)
ax.plot(x,p)

ax = fig.add_subplot(2,4,4)
ax.hist (data[:,0], bins=61, density=True)
ax.plot(x,p)

ns = np.arange(N) + 1
mn = np.cumsum (data[:,0]) / ns
sg = np.cumsum (data[:,0]**2) / ns

ax = fig.add_subplot (2,1,2)
ax.plot (ns, mn, '-')
ax.plot (ns, sg, '--')
ax.set_xscale ('log')

plt.show()
