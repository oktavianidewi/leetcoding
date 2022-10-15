# Python3 program for the above approach
import sys

# Function to count minimum
# squares that sum to n
def numSquares(n) :

	# Creating visited vector
	# of size n + 1
	visited = [0]*(n + 1)
	
	# Queue of pair to store node
	# and number of steps
	q = []
	
	# Initially ans variable is
	# initialized with inf
	ans = sys.maxsize
	
	# Push starting node with 0
	# 0 indicate current number
	# of step to reach n
	q.append([n, 0])
	
	# Mark starting node visited
	visited[n] = 1
	while(len(q) > 0) :
		
		p = q[0]
		q.pop(0)
	
		# If node reaches its destination
		# 0 update it with answer
		if(p[0] == 0) :
			ans = min(ans, p[1])
	
		# Loop for all possible path from
		# 1 to i*i <= current node(p.first)
		i = 1
		while i * i <= p[0] :
		
			# If we are standing at some node
			# then next node it can jump to will
			# be current node-
			# (some square less than or equal n)
			path = p[0] - i * i
		
			# Check if it is valid and
			# not visited yet
			if path >= 0 and (visited[path] == 0 or path == 0) :
			
				# Mark visited
				visited[path] = 1
				
				# Push it it Queue
				q.append([path,p[1] + 1])
			
			i += 1
	
	# Return ans to calling function
	return ans

print(numSquares(12))

# This code is contributed by divyeshrabadiya07