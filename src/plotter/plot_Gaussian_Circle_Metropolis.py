import numpy as np
from matplotlib import pyplot as plt

N = 1000000
fig = plt.figure (figsize=(8,3))

num_list = []
for i in range (N):
    num_list.append (list (map (float, input().split())))

data = np.array (num_list)

x = np.linspace (-3,3,200);
p = np.exp (-x**2/2)/np.sqrt(2*np.pi)*(x>=0) + 2*np.sqrt(1-(x*(x**2<=1))**2)/np.pi*(-1<=x)*(x<0)

ax = fig.add_subplot(1,4,1)
ax.hist (data[:1000,0], bins=61, density=True)
ax.plot(x,p)

ax = fig.add_subplot(1,4,2)
ax.hist (data[:10000,0], bins=61, density=True)
ax.plot(x,p)

ax = fig.add_subplot(1,4,3)
ax.hist (data[:100000,0], bins=61, density=True)
ax.plot(x,p)

ax = fig.add_subplot(1,4,4)
ax.hist (data[:,0], bins=61, density=True)
ax.plot(x,p)

plt.show()
